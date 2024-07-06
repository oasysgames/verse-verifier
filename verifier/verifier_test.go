package verifier

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
	"github.com/stretchr/testify/suite"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
)

type VerifierTestSuite struct {
	backend.BackendSuite

	verifier *Verifier
	task     verse.VerifiableVerse
	stopWork context.CancelFunc
	sigsCh   chan []*database.OptimismSignature
}

func TestVerifier(t *testing.T) {
	suite.Run(t, new(VerifierTestSuite))
}

type MockP2P struct {
	sigsCh chan []*database.OptimismSignature
}

func (m *MockP2P) PublishSignatures(ctx context.Context, sigs []*database.OptimismSignature) {
	m.sigsCh <- sigs
}

func (s *VerifierTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()

	s.sigsCh = make(chan []*database.OptimismSignature, 4)
	s.verifier = NewVerifier(&config.Verifier{
		Interval:            50 * time.Millisecond,
		Concurrency:         10,
		StateCollectLimit:   3,
		StateCollectTimeout: time.Second,
		Confirmations:       2,
	}, s.DB, &MockP2P{sigsCh: s.sigsCh}, s.SignableHub)

	s.task = verse.NewOPLegacy(s.DB, s.Hub, s.SCCAddr).WithVerifiable(s.Verse)
	s.verifier.AddTask(s.task)
}

func (s *VerifierTestSuite) TestVerify() {
	cases := []struct {
		batchRoot     string
		wantSignature string
		wantApproved  bool
	}{
		{
			"0x9ad778e5c9936769419b31119fb0bbc9d7b66c88ee10f0986ce46a6d302792b7",
			"0xa01df213459635dcd05e84b1828ba26b9469d52bf2860698437ac466d0e9afba5bda3efa378d9c36ca9eb7f4ce87f2aad73deeea357d7dba141d3469d095bb8c1c",
			true,
		},
		{
			"0x3b6af01f7666ff6990d8ccaa995f6efdae442ad24b5a354a70029ed8a2713357",
			"0x21c90d613eb6a8fbb43d858de6c6aa8c569e0c04e0e26af73f0a1043e533f26631e76beda58d5084bd93a5159a25cd6c80d5396916d1247644e63422e7cef85c1c",
			false,
		},
	}

	batchSize := 10

	s.startWorker()
	defer s.stopWork()

	// send transactions to verse-layer
	s.sendVerseTransactions(10 * len(cases))

	// emit and collect `StateBatchAppended` events
	for batchIndex, tt := range cases {
		_, err := s.task.EventDB().Save(
			s.task.RollupContract(),
			&scc.SccStateBatchAppended{
				BatchIndex:        big.NewInt(int64(batchIndex)),
				BatchRoot:         util.BytesToBytes32(common.FromHex(tt.batchRoot)),
				BatchSize:         big.NewInt(int64(batchSize)),
				PrevTotalElements: big.NewInt(int64(batchSize * batchIndex)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", batchSize))})
		s.NoError(err)
	}

	// subscribe new signature
	subscribes := s.waitPublished(len(cases))

	// assert
	for batchIndex, tt := range cases {
		bi64 := uint64(batchIndex)
		got0, _ := s.DB.OPSignature.Find(nil, nil, nil, &bi64, 1, 0)
		got1 := subscribes[batchIndex]

		s.Equal(tt.batchRoot, got0[0].RollupHash.Hex())
		s.Equal(tt.batchRoot, got1.RollupHash.Hex())

		s.Equal(tt.wantApproved, got0[0].Approved)
		s.Equal(tt.wantApproved, got1.Approved)

		s.Equal(tt.wantSignature, got0[0].Signature.Hex())
		s.Equal(tt.wantSignature, got1.Signature.Hex())
	}
}

func (s *VerifierTestSuite) TestDeleteInvalidSignature() {
	s.startWorker()
	defer s.stopWork()

	batches := s.Range(0, 10)
	batchSize := 5
	invalidBatch := 6

	// send transactions to verse-layer
	merkleRoots := make([][32]byte, len(batches))
	for batchIdx := range batches {
		elements := make([][32]byte, batchSize)
		for i, header := range s.sendVerseTransactions(batchSize) {
			elements[i] = header.Root
		}
		if merkleRoot, err := verse.CalcMerkleRoot(elements); s.NoError(err) {
			merkleRoots[batchIdx] = merkleRoot
		}
	}

	createds := make([]*database.OptimismSignature, len(batches))
	for batchIdx, merkleRoot := range merkleRoots {
		// save verify waiting state
		_, err := s.task.EventDB().Save(
			s.task.RollupContract(),
			&scc.SccStateBatchAppended{
				BatchIndex:        big.NewInt(int64(batchIdx)),
				BatchRoot:         merkleRoot,
				BatchSize:         big.NewInt(int64(batchSize)),
				PrevTotalElements: big.NewInt(int64(batchIdx * batchSize)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", batchIdx))})
		s.NoError(err)

		// run verification
		sigs := s.waitPublished(1)
		s.Len(sigs, 1)
		s.Equal(merkleRoot[:], sigs[0].RollupHash[:])
		createds[batchIdx] = sigs[0]
	}

	// increment `nextIndex`
	for batchIdx := range s.Range(0, invalidBatch) {
		s.TSCC.EmitStateBatchVerified(
			s.SignableHub.TransactOpts(context.Background()),
			big.NewInt(int64(batchIdx)),
			merkleRoots[batchIdx],
		)
		s.SignableHub.Commit()
	}
	// Confirm blocks
	for i := 0; i < s.verifier.cfg.Confirmations; i++ {
		s.Hub.Mining()
	}

	// run `deleteInvalidSignature`, but nothing happens
	s.Len(s.waitPublished(1), 0)

	signer := s.SignableHub.Signer()
	contract := s.task.RollupContract()
	gots, _ := s.DB.OPSignature.Find(nil, &signer, &contract, nil, 100, 0)
	// Make sure more than 3 depth older signature is cleaned up
	deletedBatches := invalidBatch - 3
	s.Equal(len(batches)-deletedBatches, len(gots))

	for batchIdx := range batches {
		if batchIdx < deletedBatches {
			continue
		}
		i := batchIdx - deletedBatches
		// should not be re-created
		s.Equal(createds[batchIdx].ID, gots[i].ID)
		s.Equal(createds[batchIdx].Signature, gots[i].Signature)
	}

	// update to invalid signature
	s.DB.OPSignature.Save(
		&createds[invalidBatch].ID,
		&createds[invalidBatch].PreviousID,
		createds[invalidBatch].Signer.Address,
		createds[invalidBatch].Contract.Address,
		createds[invalidBatch].RollupIndex,
		createds[invalidBatch].RollupHash,
		createds[invalidBatch].Approved,
		database.RandSignature())

	// run `deleteInvalidSignature`
	s.Len(s.waitPublished(len(batches)-invalidBatch), len(batches)-invalidBatch)

	gots, _ = s.DB.OPSignature.Find(nil, &signer, &contract, nil, 100, 0)
	s.Equal(len(batches)-deletedBatches, len(gots))

	for batchIdx := range batches {
		if batchIdx < deletedBatches {
			continue
		}
		i := batchIdx - deletedBatches
		if batchIdx < invalidBatch {
			s.Equal(createds[batchIdx].ID, gots[i].ID)
		} else {
			// should be re-created
			s.NotEqual(createds[batchIdx].ID, gots[i].ID)
		}
		s.Equal(createds[batchIdx].Signature, gots[i].Signature)
	}
}

func (s *VerifierTestSuite) TestStartVerifier() {
	// start verifier
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go s.verifier.startVerifier(ctx, s.task, 0)

	batches := s.Range(0, 3)
	batchSize := 5

	// send transactions to verse-layer
	merkleRoots := make([][32]byte, len(batches))
	for batchIdx := range batches {
		elements := make([][32]byte, batchSize)
		for i, header := range s.sendVerseTransactions(batchSize) {
			elements[i] = header.Root
		}
		if merkleRoot, err := verse.CalcMerkleRoot(elements); s.NoError(err) {
			merkleRoots[batchIdx] = merkleRoot
		}
	}
	for batchIdx, merkleRoot := range merkleRoots {
		_, err := s.TSCC.EmitStateBatchAppended(
			s.SignableHub.TransactOpts(context.Background()),
			big.NewInt(int64(batchIdx)),
			merkleRoot, big.NewInt(int64(batchSize)),
			big.NewInt(int64(batchIdx*batchSize)), []byte(fmt.Sprintf("test-%d", batchIdx)),
		)
		s.NoError(err)
		s.Mining()
	}

	// run verification
	sigs := <-s.sigsCh
	s.Len(sigs, len(merkleRoots))

	// assert
	for batchIndex, sig := range sigs {
		bi64 := uint64(batchIndex)
		got0, _ := s.DB.OPSignature.Find(nil, nil, nil, &bi64, 1, 0)

		s.Equal(sig.RollupHash.Hex(), got0[0].RollupHash.Hex())
		s.Equal(merkleRoots[batchIndex][:], got0[0].RollupHash[:])

		s.Equal(sig.Approved, got0[0].Approved)
		s.Equal(true, got0[0].Approved)

		s.Equal(sig.Signature.Hex(), got0[0].Signature.Hex())
	}

	// increment `nextIndex`
	nextIndex := 1
	s.TSCC.SetNextIndex(s.SignableHub.TransactOpts(ctx), big.NewInt(int64(nextIndex)))
	s.Hub.Commit()

	// confirm blocks
	s.Hub.Commit()
	s.Hub.Commit()

	// no prior signature than verified index should be sent
	for i := 0; i < 2; i++ {
		sigs = <-s.sigsCh
		s.Len(sigs, 2)
		for _, sig := range sigs {
			s.True(sig.RollupIndex >= uint64(nextIndex))
		}
	}
}

func (s *VerifierTestSuite) sendVerseTransactions(count int) (headers []*types.Header) {
	ctx := context.Background()
	to := common.HexToAddress("0x09ad74977844F513E61AdE2B50b0C06268A4f6d7")

	nonce, err := s.SignableVerse.PendingNonceAt(ctx, s.SignableVerse.Signer())
	s.NoError(err)

	for i := 0; i < count; i++ {
		gasPrice, err := s.SignableVerse.BaseGasPrice(ctx, nil)
		s.Nil(err)

		unsigned := types.NewTransaction(
			nonce+uint64(i), to, common.Big0, 21_000, gasPrice, nil)
		signed, err := s.SignableVerse.SignTx(unsigned)
		s.Nil(err)

		err = s.SignableVerse.SendTransaction(ctx, signed)
		s.NoError(err)

		h, err := s.SignableVerse.HeaderByHash(ctx, s.SignableVerse.Commit())
		s.NoError(err)

		headers = append(headers, h)
	}
	return headers
}

func (s *VerifierTestSuite) startWorker() {
	ctx := context.Background()
	ctx, s.stopWork = context.WithCancel(ctx)

	go func() {
		tick := time.NewTicker(s.verifier.cfg.Interval)
		defer tick.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				s.verifier.Work(ctx)
			}
		}
	}()
}

func (s *VerifierTestSuite) waitPublished(count int) []*database.OptimismSignature {
	ctx, candel := context.WithTimeout(context.Background(), time.Second/2)
	defer candel()

	sub := s.verifier.SubscribeNewSignature(ctx)
	defer sub.Cancel()

	var published []*database.OptimismSignature
	go func() {
		defer candel()

		for {
			select {
			case <-ctx.Done():
				return
			case sig := <-sub.Next():
				published = append(published, sig)
				if len(published) == count {
					return
				}
			}
		}
	}()

	<-ctx.Done()
	return published
}

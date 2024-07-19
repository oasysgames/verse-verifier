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
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
	"github.com/stretchr/testify/suite"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
)

type VerifierTestSuite struct {
	backend.BackendSuite

	verifier *Verifier
	task     verse.VerifiableVerse
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
		StateCollectLimit:   3,
		StateCollectTimeout: time.Second,
		Confirmations:       2,
		StartBlockOffset:    100,
	}, s.DB, &MockP2P{sigsCh: s.sigsCh}, s.SignableHub)

	s.task = verse.NewOPLegacy(s.DB, s.Hub, s.SCCAddr).WithVerifiable(s.Verse)
	s.verifier.AddTask(context.Background(), s.task, 0)
}

func (s *VerifierTestSuite) TestStartVerifier() {
	// start verifier by adding task
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s.verifier.AddTask(ctx, s.task, 0)

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
	for i := 0; i < 3; i++ {
		sigs = <-s.sigsCh
		if len(sigs) == len(batches) {
			// signatures before incrementing `nextIndex` are remained
			continue
		}
		for _, sig := range sigs {
			s.True(sig.RollupIndex >= uint64(nextIndex))
		}
	}
}

func (s *VerifierTestSuite) TestRetryBackoff() {
	verifier := &Verifier{
		cfg: &config.Verifier{
			MaxRetryBackoff: time.Minute,
			RetryTimeout:    time.Millisecond * 100,
		},
	}

	backoff := verifier.retryBackoff()

	wait := time.Millisecond * 10
	for i := 0; i < 10; i++ {
		delay, remain, attempts := backoff()

		s.Equal(i+1, attempts)
		s.Less(remain, verifier.cfg.RetryTimeout-wait*time.Duration(i))

		switch i {
		case 0:
			s.Equal(delay, time.Millisecond*100)
		case 1:
			s.Equal(delay, time.Millisecond*800)
		case 2:
			s.Equal(delay, time.Millisecond*6400)
		case 3:
			s.Equal(delay, time.Millisecond*51200)
		default: // i >= 4
			s.Equal(delay, time.Minute)
		}

		time.Sleep(wait)
	}

	_, remain, _ := backoff()
	s.Equal(remain, time.Duration(0))
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

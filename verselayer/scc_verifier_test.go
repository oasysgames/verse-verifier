package verselayer

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/stretchr/testify/suite"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	tscc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/scc"
)

type SccVerifierTestSuite struct {
	testhelper.Suite

	db    *database.Database
	hub   *backend.SignableBackend
	verse *backend.Backend

	scc     *tscc.Scc
	sccAddr common.Address

	verifier *SccVerifier
}

func TestSccVerifier(t *testing.T) {
	suite.Run(t, new(SccVerifierTestSuite))
}

func (s *SccVerifierTestSuite) SetupTest() {
	// setup test env
	s.db, _ = database.NewDatabase(&config.Database{Path: ":memory:"})
	s.hub = backend.NewSignableBackend(nil, nil, nil)
	s.verse = backend.NewBackend(nil, 0)

	// deploy `StateCommitmentChain` contract
	s.sccAddr, _, s.scc, _ = tscc.DeployScc(s.hub.TransactOpts(context.Background()), s.hub)
	s.hub.Mining()

	// setup verifier
	s.verifier = NewSccVerifier(&config.Verifier{
		Interval:            50 * time.Millisecond,
		Concurrency:         10,
		StateCollectLimit:   2,
		StateCollectTimeout: time.Second,
	}, s.db, s.hub)
	s.verifier.AddVerse(s.sccAddr, s.verse)
}

func (s *SccVerifierTestSuite) TestVerify() {
	cases := []struct {
		batchRoot     string
		wantSignature string
		wantApproved  bool
	}{
		{
			"0xa32f22db573ecdc5eafbb5d1cc99b51ebc603f26bb0becac52e46157eddbe005",
			"0x02141577e2d1ff2230c728436a61f2050a6cd7fb91c10c646b115447e6052e5355599c3f9225553c9bc06c5cfa4a0cbf1cd50f072e354a0a1d3549567f7e05591c",
			true,
		},
		{
			"0x3b6af01f7666ff6990d8ccaa995f6efdae442ad24b5a354a70029ed8a2713357",
			"0x7636fc9abc773fbebf17e87c3445f5a66a243a32595cf1f50d732e3db3e20d32082df7f820130eb14fe553a7223972ba3b68dfd4123bda67b28939fe51b3be181c",
			false,
		},
	}

	batchSize := 10

	// send transactions to verse-layer
	s.sendTransaction(10 * len(cases))

	// emit and collect `StateBatchAppended` events
	for i, tt := range cases {
		s.db.Optimism.SaveState(&scc.SccStateBatchAppended{
			Raw:               types.Log{Address: s.sccAddr},
			BatchIndex:        big.NewInt(int64(i)),
			BatchRoot:         util.BytesToBytes32(common.FromHex(tt.batchRoot)),
			BatchSize:         big.NewInt(int64(batchSize)),
			PrevTotalElements: big.NewInt(int64(batchSize * i)),
			ExtraData:         []byte(fmt.Sprintf("test-%d", batchSize))})
	}

	// subscribe new signature
	subscribes := s.startAndWait(s.verifier, len(cases))

	// assert
	for i, tt := range cases {
		index := uint64(i)
		got0, _ := s.db.Optimism.FindSignatures(nil, nil, nil, &index, 1, 0)
		got1 := subscribes[i]

		s.Equal(tt.batchRoot, got0[0].BatchRoot.Hex())
		s.Equal(tt.batchRoot, got1.BatchRoot.Hex())

		s.Equal(uint64(batchSize), got0[0].BatchSize)
		s.Equal(uint64(batchSize), got1.BatchSize)

		s.Equal(uint64(batchSize*i), got0[0].PrevTotalElements)
		s.Equal(uint64(batchSize*i), got1.PrevTotalElements)

		s.Equal(fmt.Sprintf("test-%d", batchSize), string(got0[0].ExtraData))
		s.Equal(fmt.Sprintf("test-%d", batchSize), string(got1.ExtraData))

		s.Equal(tt.wantApproved, got0[0].Approved)
		s.Equal(tt.wantApproved, got1.Approved)

		s.Equal(tt.wantSignature, got0[0].Signature.Hex())
		s.Equal(tt.wantSignature, got1.Signature.Hex())
	}
}

func (s *SccVerifierTestSuite) TestDeleteInvalidSignature() {
	batches := s.Range(0, 10)
	batchSize := 5
	invalidBatch := 6

	// send transactions to verse-layer
	merkleRoots := make([][32]byte, len(batches))
	for batchIdx := range batches {
		elements := make([][32]byte, batchSize)
		for i, header := range s.sendTransaction(batchSize) {
			elements[i] = header.Root
		}
		if merkleRoot, err := calcMerkleRoot(elements); s.NoError(err) {
			merkleRoots[batchIdx] = merkleRoot
		}
	}

	createds := make([]*database.OptimismSignature, len(batches))
	for batchIdx, merkleRoot := range merkleRoots {
		// save verify waiting state
		s.db.Optimism.SaveState(&scc.SccStateBatchAppended{
			Raw:               types.Log{Address: s.sccAddr},
			BatchIndex:        big.NewInt(int64(batchIdx)),
			BatchRoot:         merkleRoot,
			BatchSize:         big.NewInt(int64(batchSize)),
			PrevTotalElements: big.NewInt(int64(batchIdx * batchSize)),
			ExtraData:         []byte(fmt.Sprintf("test-%d", batchIdx))})

		// run verification
		sigs := s.startAndWait(s.verifier, 1)
		s.Len(sigs, 1)
		s.Equal(merkleRoot[:], sigs[0].BatchRoot[:])
		createds[batchIdx] = sigs[0]
	}

	// increment `nextIndex`
	for batchIdx := range s.Range(0, invalidBatch) {
		s.scc.EmitStateBatchVerified(
			s.hub.TransactOpts(context.Background()),
			big.NewInt(int64(batchIdx)),
			merkleRoots[batchIdx],
		)
		s.hub.Commit()
	}

	// run `deleteInvalidSignature`, but nothing happens
	s.Len(s.startAndWait(s.verifier, 1), 0)

	signer := s.hub.Signer()
	gots, _ := s.db.Optimism.FindSignatures(nil, &signer, &s.sccAddr, nil, 100, 0)
	s.Equal(len(batches), len(gots))

	for batchIdx := range batches {
		// should not be re-created
		s.Equal(createds[batchIdx].ID, gots[batchIdx].ID)
		s.Equal(createds[batchIdx].Signature, gots[batchIdx].Signature)
	}

	// update to invalid signature
	s.db.Optimism.SaveSignature(
		&createds[invalidBatch].ID,
		&createds[invalidBatch].PreviousID,
		createds[invalidBatch].Signer.Address,
		createds[invalidBatch].OptimismScc.Address,
		createds[invalidBatch].BatchIndex,
		createds[invalidBatch].BatchRoot,
		createds[invalidBatch].BatchSize,
		createds[invalidBatch].PrevTotalElements,
		createds[invalidBatch].ExtraData,
		createds[invalidBatch].Approved,
		database.RandSignature())

	// run `deleteInvalidSignature`
	s.Len(
		s.startAndWait(s.verifier, len(batches)-invalidBatch),
		len(batches)-invalidBatch,
	)

	gots, _ = s.db.Optimism.FindSignatures(nil, &signer, &s.sccAddr, nil, 100, 0)
	s.Equal(len(batches), len(gots))

	for batchIdx := range batches {
		if batchIdx < invalidBatch {
			s.Equal(createds[batchIdx].ID, gots[batchIdx].ID)
		} else {
			// should be re-created
			s.NotEqual(createds[batchIdx].ID, gots[batchIdx].ID)
		}
		s.Equal(createds[batchIdx].Signature, gots[batchIdx].Signature)
	}
}

func (s *SccVerifierTestSuite) TestCalcMerkleRoot() {
	cases := []struct {
		name string
		spec func()
	}{
		{
			"no elements",
			func() {
				_, err := calcMerkleRoot([][32]byte{})
				s.ErrorContains(err, "must provide at least one leaf hash")
			},
		},
		{
			"single element",
			func() {
				elements := [][32]byte{
					util.BytesToBytes32(
						common.FromHex(
							"0x56570de287d73cd1cb6092bb8fdee6173974955fdef345ae579ee9f475ea7432",
						),
					),
				}
				got, _ := calcMerkleRoot(elements)
				s.Equal(elements[0], got)
			},
		},
		{
			"more than one element",
			func() {
				wants := []string{
					"0x57d772147cdf27f5f67d679f0f3a513f8b87622ce598a3cf0b048ab178ddfc6e",
					"0x820919791e2ec4aea2fb218a7a3a5a89d06ba469585c824b60f0174ec13e1603",
					"0xe39e9f65a0fcee19f9b8aceadb3bbdbf7697be66b0632644e168d01dc103ddd6",
					"0x11f470d712bb3a84f0b01cb7c73493ec7d06eda480f567c99b9a6dc773679a72",
					"0x0faa9fa71909342540cabef2fdf911cf053141144b21d089641940533679cdd9",
					"0x0050d8ac9e23f46daf8be33332d201588cba3cee5c6141715756dc4b2c960ada",
					"0xfc61b646f502f97300b88afe15feaf046f90c8456f658273657d8a55e7fc79df",
					"0xa4329a43ffc1bc6195e1bddda04930ed1db6486df03a56a8df9a60bb2d5469e0",
					"0x9a68f697fd78c779e436dec655825d263066c5fee23f961fd15e9d14327ded6b",
					"0x437dc148af6b33ba532cf6e8d8c0c74ab680439cbd03f9000f7434fb217611b7",
					"0xeade7c5f57e013547c7cec95eff59e44616ab9bdadb73420545f741e4097f9c1",
					"0x7acdf7918c5b5dc8acac506737231e143f2dc6b8734ec02b3d92676852fd4880",
					"0x9b9b9244ced25fff4077e6bca56882d106981a5d949394ad509bb0b11e04d23a",
					"0xbba15d82445e21878b48a0e4b19854c4e0e75a68e644bdfb8ace0fc965264431",
				}
				sizes := []int{2, 3, 7, 9, 13, 63, 64, 123, 128, 129, 255, 1021, 1023, 1024}
				for i := range wants {
					want := util.BytesToBytes32(common.FromHex(wants[i]))
					size := sizes[i]

					s.Run(fmt.Sprintf("size %d", size), func() {
						elements := make([][32]byte, size)
						for i := range elements {
							bhash := common.FromHex(hexutil.EncodeBig(big.NewInt(int64(i))))
							elements[i] = util.BytesToBytes32(crypto.Keccak256(bhash))
						}
						got, _ := calcMerkleRoot(s.fillDefaultHashes(elements))
						s.Equal(want, got)
					})
				}
			},
		},
		{
			"odd number of elements",
			func() {
				elements := [][32]byte{
					util.BytesToBytes32(crypto.Keccak256(common.FromHex("0x12"))),
					util.BytesToBytes32(crypto.Keccak256(common.FromHex("0x34"))),
					util.BytesToBytes32(crypto.Keccak256(common.FromHex("0x56"))),
				}
				_, err := calcMerkleRoot(elements)
				s.Equal(nil, err)
			},
		},
	}

	for _, tt := range cases {
		s.Run(tt.name, tt.spec)
	}
}

func (s *SccVerifierTestSuite) sendTransaction(count int) (headers []*types.Header) {
	for i := 0; i < count; i++ {
		signedTx, _ := s.hub.SignTx(types.NewTransaction(
			uint64(i),
			common.HexToAddress("0x09ad74977844F513E61AdE2B50b0C06268A4f6d7"),
			common.Big0,
			uint64(21_000),
			big.NewInt(875_000_000),
			nil))

		s.verse.SendTransaction(context.Background(), signedTx)
		if h, err := s.verse.HeaderByHash(context.Background(), s.verse.Commit()); s.NoError(err) {
			headers = append(headers, h)
		}
	}
	return headers
}

func (s *SccVerifierTestSuite) fillDefaultHashes(elements [][32]byte) [][32]byte {
	fillhash := util.BytesToBytes32(
		crypto.Keccak256(common.FromHex("0x" + strings.Repeat("00", 32))),
	)

	filled := [][32]byte{}
	for i := 0; float64(i) < math.Pow(2, math.Ceil(math.Log2(float64(len(elements))))); i++ {
		if i < len(elements) {
			filled = append(filled, elements[i])
		} else {
			filled = append(filled, fillhash)
		}
	}

	return filled
}

func (s *SccVerifierTestSuite) startAndWait(
	verifier *SccVerifier,
	count int,
) []*database.OptimismSignature {
	ctx, candel := context.WithTimeout(context.Background(), time.Second/2)
	defer candel()

	sub := verifier.SubscribeNewSignature(ctx)
	defer sub.Cancel()

	published := []*database.OptimismSignature{}
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

	go verifier.Start(ctx)
	<-ctx.Done()

	return published
}

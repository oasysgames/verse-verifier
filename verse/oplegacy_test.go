package verse

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/stretchr/testify/suite"
)

type OPLegacyTestSuite struct {
	backend.BackendSuite

	verse        Verse
	verifiable   VerifiableVerse
	transactable TransactableVerse
}

func TestOPLegacy(t *testing.T) {
	suite.Run(t, new(OPLegacyTestSuite))
}

func (s *OPLegacyTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()

	s.verse = NewOPLegacy(s.DB, s.Hub, s.SCCAddr)
	s.verifiable = s.verse.WithVerifiable(s.Verse)
	s.transactable = s.verse.WithTransactable(s.SignableHub, s.SCCVAddr)
}

func (s *OPLegacyTestSuite) TestEventDB() {
	want := database.NewOPEventDB[database.OptimismState](s.DB)
	s.IsType(want, s.verse.EventDB())
	s.IsType(want, s.verifiable.EventDB())
	s.IsType(want, s.transactable.EventDB())
}

func (s *OPLegacyTestSuite) TestNextIndex() {
	s.TSCC.SetNextIndex(s.SignableHub.TransactOpts(context.Background()), big.NewInt(10))
	s.Mining()

	got0, _ := s.verse.NextIndex(&bind.CallOpts{})
	got1, _ := s.verifiable.NextIndex(&bind.CallOpts{})
	got2, _ := s.transactable.NextIndex(&bind.CallOpts{})
	s.Equal(uint64(10), got0.Uint64())
	s.Equal(uint64(10), got1.Uint64())
	s.Equal(uint64(10), got2.Uint64())
}

func (s *OPLegacyTestSuite) TestVerify() {
	ctx := context.Background()

	stateRoots := [][32]byte{}
	for range s.Range(0, 10) {
		nonce, err := s.Verse.PendingNonceAt(ctx, s.SignableVerse.Signer())
		s.Nil(err)

		gasPrice, err := s.SignableVerse.BaseGasPrice(ctx, nil)
		s.Nil(err)

		unsigned := types.NewTransaction(
			nonce, s.RandAddress(), common.Big1, 21_000, gasPrice, nil)

		_, err = s.SignableVerse.SendTxWithSign(ctx, unsigned)
		s.Nil(err)

		header := s.Verse.Blockchain().CurrentHeader()
		stateRoots = append(stateRoots, header.Root)
	}

	var (
		err   error
		event = &database.OptimismState{
			Contract:          database.OptimismContract{Address: s.RandAddress()},
			BatchIndex:        0,
			BatchSize:         uint64(len(stateRoots)),
			PrevTotalElements: 0,
			ExtraData:         nil,
		}
	)

	// if verification is successful
	event.BatchRoot, err = CalcMerkleRoot(stateRoots)
	s.Nil(err)

	approved, err := s.verifiable.Verify(log.New(), ctx, event, 3)
	s.True(approved)
	s.Nil(err)

	// if verification is failure
	event.BatchRoot = s.RandHash()
	approved, err = s.verifiable.Verify(log.New(), ctx, event, 3)
	s.False(approved)
	s.Nil(err)
}

func (s *OPLegacyTestSuite) TestTransact() {
	opts := s.SignableHub.TransactOpts(context.Background())

	// approve
	_, emitted := s.EmitStateBatchAppended(0)
	s.transactable.Transact(opts, 0, true, [][]byte{[]byte("test:approve")})
	s.Mining()

	assertLog, _ := s.TSCCV.AssertLogs(&bind.CallOpts{}, big.NewInt(0))
	s.Equal(s.SCCAddr, assertLog.StateCommitmentChain)
	s.Equal(emitted.BatchIndex.Uint64(), assertLog.BatchHeader.BatchIndex.Uint64())
	s.Equal(emitted.BatchRoot, assertLog.BatchHeader.BatchRoot)
	s.Equal(emitted.BatchSize.Uint64(), assertLog.BatchHeader.BatchSize.Uint64())
	s.Equal(emitted.PrevTotalElements.Uint64(), assertLog.BatchHeader.PrevTotalElements.Uint64())
	s.Equal(emitted.ExtraData, assertLog.BatchHeader.ExtraData)
	s.Equal([]byte("test:approve"), assertLog.Signatures)
	s.Equal(true, assertLog.Approve)

	// reject
	_, emitted = s.EmitStateBatchAppended(1)
	s.transactable.Transact(opts, 1, false, [][]byte{[]byte("test:reject")})
	s.Mining()

	assertLog, _ = s.TSCCV.AssertLogs(&bind.CallOpts{}, big.NewInt(1))
	s.Equal(s.SCCAddr, assertLog.StateCommitmentChain)
	s.Equal(emitted.BatchIndex.Uint64(), assertLog.BatchHeader.BatchIndex.Uint64())
	s.Equal(emitted.BatchRoot, assertLog.BatchHeader.BatchRoot)
	s.Equal(emitted.BatchSize.Uint64(), assertLog.BatchHeader.BatchSize.Uint64())
	s.Equal(emitted.PrevTotalElements.Uint64(), assertLog.BatchHeader.PrevTotalElements.Uint64())
	s.Equal(emitted.ExtraData, assertLog.BatchHeader.ExtraData)
	s.Equal([]byte("test:reject"), assertLog.Signatures)
	s.Equal(false, assertLog.Approve)
}

func (s *OPLegacyTestSuite) TestCalcMerkleRoot() {
	cases := []struct {
		name string
		spec func()
	}{
		{
			"no elements",
			func() {
				_, err := CalcMerkleRoot([][32]byte{})
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
				got, _ := CalcMerkleRoot(elements)
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
						got, _ := CalcMerkleRoot(fillDefaultHashes(elements))
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
				_, err := CalcMerkleRoot(elements)
				s.Equal(nil, err)
			},
		},
	}

	for _, tt := range cases {
		s.Run(tt.name, tt.spec)
	}
}

func fillDefaultHashes(elements [][32]byte) [][32]byte {
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

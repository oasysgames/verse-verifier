package verse

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	"github.com/stretchr/testify/suite"
)

type OPStackTestSuite struct {
	backend.BackendSuite

	verse        Verse
	verifiable   VerifiableVerse
	transactable TransactableVerse
}

func TestOPStack(t *testing.T) {
	suite.Run(t, new(OPStackTestSuite))
}

func (s *OPStackTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()

	s.verse = NewOPStack(s.DB, s.Hub, s.L2OOAddr)
	s.verifiable = s.verse.WithVerifiable(s.Verse)
	s.transactable = s.verse.WithTransactable(s.SignableHub, s.L2OOVAddr)
}

func (s *OPStackTestSuite) TestEventDB() {
	want := database.NewOPEventDB[database.OpstackProposal](s.DB)
	s.IsType(want, s.verse.EventDB())
	s.IsType(want, s.verifiable.EventDB())
	s.IsType(want, s.transactable.EventDB())
}

func (s *OPStackTestSuite) TestNextIndex() {
	s.TL2OO.SetNextVerifyIndex(s.SignableHub.TransactOpts(context.Background()), big.NewInt(10))
	s.Mining()

	got0, _ := s.verse.NextIndex(&bind.CallOpts{})
	got1, _ := s.verifiable.NextIndex(&bind.CallOpts{})
	got2, _ := s.transactable.NextIndex(&bind.CallOpts{})
	s.Equal(uint64(10), got0.Uint64())
	s.Equal(uint64(10), got1.Uint64())
	s.Equal(uint64(10), got2.Uint64())
}

func (s *OPStackTestSuite) TestVerify() {
	// address of the L2ToL1MessagePasser contract
	l2ToL1MessagePasser := common.HexToAddress("0x4200000000000000000000000000000000000016")

	ctx := context.Background()

	// send transaction to the L2ToL1MessagePasser
	nonce, err := s.Verse.PendingNonceAt(ctx, s.SignableVerse.Signer())
	s.Nil(err)

	unsigned := types.NewTransaction(nonce, l2ToL1MessagePasser,
		common.Big1, 21_000, big.NewInt(875_000_000), nil)

	_, err = s.SignableVerse.SendTxWithSign(ctx, unsigned)
	s.Nil(err)

	// create output root
	head := s.Verse.Blockchain().CurrentHeader()
	proof, err := s.Verse.GetProof(
		context.Background(), l2ToL1MessagePasser, []string{}, head.Number)
	s.Nil(err)

	event := &database.OpstackProposal{
		Contract:      database.OptimismContract{Address: s.RandAddress()},
		L2OutputIndex: 0,
		L2BlockNumber: head.Number.Uint64(),
		L1Timestamp:   uint64(time.Now().Unix()),
	}

	// if verification is successful
	outputV0 := &OpstackOutputV0{
		StateRoot:                head.Root,
		MessagePasserStorageRoot: proof.StorageHash,
		BlockHash:                head.Hash(),
	}
	event.OutputRoot = outputV0.OutputRoot()

	approved, err := s.verifiable.Verify(log.New(), ctx, event, 0)
	s.True(approved)
	s.Nil(err)

	// if verification is failure
	event.OutputRoot = s.RandHash()
	approved, err = s.verifiable.Verify(log.New(), ctx, event, 0)
	s.False(approved)
	s.Nil(err)
}

func (s *OPStackTestSuite) TestTransact() {
	opts := s.SignableHub.TransactOpts(context.Background())

	// approve
	_, emitted := s.EmitOutputProposed(0)
	s.transactable.Transact(opts, 0, true, [][]byte{[]byte("test:approve")})
	s.Mining()

	assertLog, _ := s.TL2OOV.AssertLogs(&bind.CallOpts{}, big.NewInt(0))
	s.Equal(s.L2OOAddr, assertLog.L2OutputOracle)
	s.Equal(emitted.L2OutputIndex.Uint64(), assertLog.L2OutputIndex.Uint64())
	s.Equal(emitted.OutputRoot, assertLog.L2Output.OutputRoot)
	s.Equal(emitted.L1Timestamp.Uint64(), assertLog.L2Output.Timestamp.Uint64())
	s.Equal(emitted.L2BlockNumber.Uint64(), assertLog.L2Output.L2BlockNumber.Uint64())
	s.Equal([]byte("test:approve"), assertLog.Signatures)
	s.Equal(true, assertLog.Approve)

	// reject
	_, emitted = s.EmitOutputProposed(1)
	s.transactable.Transact(opts, 1, false, [][]byte{[]byte("test:reject")})
	s.Mining()

	assertLog, _ = s.TL2OOV.AssertLogs(&bind.CallOpts{}, big.NewInt(1))
	s.Equal(s.L2OOAddr, assertLog.L2OutputOracle)
	s.Equal(emitted.L2OutputIndex.Uint64(), assertLog.L2OutputIndex.Uint64())
	s.Equal(emitted.OutputRoot, assertLog.L2Output.OutputRoot)
	s.Equal(emitted.L1Timestamp.Uint64(), assertLog.L2Output.Timestamp.Uint64())
	s.Equal(emitted.L2BlockNumber.Uint64(), assertLog.L2Output.L2BlockNumber.Uint64())
	s.Equal([]byte("test:reject"), assertLog.Signatures)
	s.Equal(false, assertLog.Approve)
}

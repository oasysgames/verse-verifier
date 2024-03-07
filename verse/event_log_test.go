package verse

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	"github.com/stretchr/testify/suite"
)

type EventLogTestSuite struct {
	backend.BackendSuite
}

func TestEventLog(t *testing.T) {
	suite.Run(t, new(EventLogTestSuite))
}

func (s *EventLogTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()
}

func (s *EventLogTestSuite) TestNewEventLogFilter() {
	txs := []*types.Transaction{}
	emit := func(tx *types.Transaction, _ any) {
		txs = append(txs, tx)
		s.Mining()
	}

	opts := s.SignableHub.TransactOpts(context.Background())

	// StateCommitmentChain
	emit(s.EmitStateBatchAppended(0))
	emit(s.TSCC.EmitStateBatchDeleted(opts, big.NewInt(0), s.RandHash()))
	emit(s.TSCC.EmitStateBatchVerified(opts, big.NewInt(0), s.RandHash()))

	// L2OutputOracle
	emit(s.EmitOutputProposed(0))
	emit(s.TL2OO.EmitOutputsDeleted(opts, big.NewInt(0), big.NewInt(0)))
	emit(s.TL2OO.EmitOutputVerified(opts, big.NewInt(0), s.RandHash(), big.NewInt(0)))

	wants := make([]*types.Log, len(txs))
	for i, tx := range txs {
		receipt, _ := s.Hub.TransactionReceipt(context.Background(), tx.Hash())
		wants[i] = receipt.Logs[0]
	}

	gots, _ := s.Hub.FilterLogs(context.Background(), NewEventLogFilter(0, 100))
	s.Len(gots, len(wants))
	for i, want := range wants {
		s.Equal(*want, gots[i])
	}

	scc, _ := abi.JSON(strings.NewReader(scc.SccABI))
	s.Equal(scc.Events["StateBatchAppended"].ID, gots[0].Topics[0])
	s.Equal(scc.Events["StateBatchDeleted"].ID, gots[1].Topics[0])
	s.Equal(scc.Events["StateBatchVerified"].ID, gots[2].Topics[0])

	l2oo, _ := abi.JSON(strings.NewReader(l2oo.OasysL2OutputOracleABI))
	s.Equal(l2oo.Events["OutputProposed"].ID, gots[3].Topics[0])
	s.Equal(l2oo.Events["OutputsDeleted"].ID, gots[4].Topics[0])
	s.Equal(l2oo.Events["OutputVerified"].ID, gots[5].Topics[0])
}

func (s *EventLogTestSuite) TestParseStateBatchAppendedEvent() {
	tx, want1 := s.EmitStateBatchAppended(10)
	receipt, _ := s.Hub.TransactionReceipt(context.Background(), tx.Hash())

	got, _ := ParseEventLog(receipt.Logs[0])
	gott0 := got.(*RollupedEvent)

	s.Equal("scc", gott0.Contract)
	s.Equal("StateBatchAppended", gott0.Event)
	s.Equal(receipt.Logs[0], gott0.Log)
	s.IsType(database.NewOPEventDB[database.OptimismState](s.DB), gott0.EventDB(s.DB))
	s.Equal(want1.BatchIndex.Uint64(), gott0.RollupIndex)

	gott1 := gott0.Parsed.(*scc.SccStateBatchAppended)
	s.Equal(want1.BatchIndex, gott1.BatchIndex)
	s.Equal(want1.BatchRoot, gott1.BatchRoot)
	s.Equal(want1.BatchSize, gott1.BatchSize)
	s.Equal(want1.PrevTotalElements, gott1.PrevTotalElements)
	s.Equal(want1.ExtraData, gott1.ExtraData)
	s.Equal(*receipt.Logs[0], gott1.Raw)
}

func (s *EventLogTestSuite) TestParseStateBatchDeletedEvent() {
	batchIndex := big.NewInt(10)
	batchRoot := s.RandHash()

	tx, _ := s.TSCC.EmitStateBatchDeleted(
		s.SignableHub.TransactOpts(context.Background()), batchIndex, batchRoot)
	s.Mining()
	receipt, _ := s.Hub.TransactionReceipt(context.Background(), tx.Hash())

	got, _ := ParseEventLog(receipt.Logs[0])
	gott0 := got.(*DeletedEvent)

	s.Equal("scc", gott0.Contract)
	s.Equal("StateBatchDeleted", gott0.Event)
	s.Equal(receipt.Logs[0], gott0.Log)
	s.IsType(database.NewOPEventDB[database.OptimismState](s.DB), gott0.EventDB(s.DB))
	s.Equal(batchIndex.Uint64(), gott0.RollupIndex)

	gott1 := gott0.Parsed.(*scc.SccStateBatchDeleted)
	s.Equal(batchIndex.Uint64(), gott1.BatchIndex.Uint64())
	s.Equal(batchRoot[:], gott1.BatchRoot[:])
	s.Equal(*receipt.Logs[0], gott1.Raw)
}

func (s *EventLogTestSuite) TestParseStateBatchVerifiedEvent() {
	batchIndex := big.NewInt(0)
	batchRoot := s.RandHash()

	tx, _ := s.TSCC.EmitStateBatchVerified(
		s.SignableHub.TransactOpts(context.Background()), batchIndex, batchRoot)
	s.Mining()
	receipt, _ := s.Hub.TransactionReceipt(context.Background(), tx.Hash())

	got, _ := ParseEventLog(receipt.Logs[0])
	gott0 := got.(*VerifiedEvent)

	s.Equal("scc", gott0.Contract)
	s.Equal("StateBatchVerified", gott0.Event)
	s.Equal(receipt.Logs[0], gott0.Log)
	s.IsType(database.NewOPEventDB[database.OptimismState](s.DB), gott0.EventDB(s.DB))
	s.Equal(batchIndex.Uint64(), gott0.RollupIndex)

	gott1 := gott0.Parsed.(*scc.SccStateBatchVerified)
	s.Equal(batchIndex.Uint64(), gott1.BatchIndex.Uint64())
	s.Equal(batchRoot[:], gott1.BatchRoot[:])
	s.Equal(*receipt.Logs[0], gott1.Raw)
}

func (s *EventLogTestSuite) TestParseOutputProposedEvent() {
	tx, want1 := s.EmitOutputProposed(10)
	receipt, _ := s.Hub.TransactionReceipt(context.Background(), tx.Hash())

	got, _ := ParseEventLog(receipt.Logs[0])
	gott0 := got.(*RollupedEvent)

	s.Equal("l2oo", gott0.Contract)
	s.Equal("OutputProposed", gott0.Event)
	s.Equal(receipt.Logs[0], gott0.Log)
	s.IsType(database.NewOPEventDB[database.OpstackProposal](s.DB), gott0.EventDB(s.DB))
	s.Equal(want1.L2OutputIndex.Uint64(), gott0.RollupIndex)

	gott1 := gott0.Parsed.(*l2oo.OasysL2OutputOracleOutputProposed)
	s.Equal(want1.OutputRoot, gott1.OutputRoot)
	s.Equal(want1.L2OutputIndex.Uint64(), gott1.L2OutputIndex.Uint64())
	s.Equal(want1.L2BlockNumber.Uint64(), gott1.L2BlockNumber.Uint64())
	s.Equal(want1.L1Timestamp.Uint64(), gott1.L1Timestamp.Uint64())
	s.Equal(*receipt.Logs[0], gott1.Raw)
}

func (s *EventLogTestSuite) TestParseOutputsDeletedEvent() {
	prevNextOutputIndex := big.NewInt(10)
	newNextOutputIndex := big.NewInt(5)

	tx, _ := s.TL2OO.EmitOutputsDeleted(
		s.SignableHub.TransactOpts(context.Background()), prevNextOutputIndex, newNextOutputIndex)
	s.Mining()
	receipt, _ := s.Hub.TransactionReceipt(context.Background(), tx.Hash())

	got, _ := ParseEventLog(receipt.Logs[0])
	gott0 := got.(*DeletedEvent)

	s.Equal("l2oo", gott0.Contract)
	s.Equal("OutputsDeleted", gott0.Event)
	s.Equal(receipt.Logs[0], gott0.Log)
	s.IsType(database.NewOPEventDB[database.OpstackProposal](s.DB), gott0.EventDB(s.DB))
	s.Equal(newNextOutputIndex.Uint64(), gott0.RollupIndex)

	gott1 := gott0.Parsed.(*l2oo.OasysL2OutputOracleOutputsDeleted)
	s.Equal(prevNextOutputIndex.Uint64(), gott1.PrevNextOutputIndex.Uint64())
	s.Equal(newNextOutputIndex.Uint64(), gott1.NewNextOutputIndex.Uint64())
	s.Equal(*receipt.Logs[0], gott1.Raw)
}

func (s *EventLogTestSuite) TestHandleOutputVerifiedEvent() {
	tx, want1 := s.EmitOutputProposed(10)
	receipt, _ := s.Hub.TransactionReceipt(context.Background(), tx.Hash())

	got, _ := ParseEventLog(receipt.Logs[0])
	gott0 := got.(*RollupedEvent)

	s.Equal("l2oo", gott0.Contract)
	s.Equal("OutputProposed", gott0.Event)
	s.Equal(receipt.Logs[0], gott0.Log)
	s.IsType(database.NewOPEventDB[database.OpstackProposal](s.DB), gott0.EventDB(s.DB))
	s.Equal(want1.L2OutputIndex.Uint64(), gott0.RollupIndex)

	gott1 := gott0.Parsed.(*l2oo.OasysL2OutputOracleOutputProposed)
	s.Equal(want1.OutputRoot, gott1.OutputRoot)
	s.Equal(want1.L2OutputIndex.Uint64(), gott1.L2OutputIndex.Uint64())
	s.Equal(want1.L2BlockNumber.Uint64(), gott1.L2BlockNumber.Uint64())
	s.Equal(want1.L1Timestamp.Uint64(), gott1.L1Timestamp.Uint64())
	s.Equal(*receipt.Logs[0], gott1.Raw)
}

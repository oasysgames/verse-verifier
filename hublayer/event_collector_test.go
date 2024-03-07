package hublayer

import (
	"context"
	"math/big"
	"testing"

	"github.com/oasysgames/oasys-optimism-verifier/database"
	tc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/scc"
	"github.com/stretchr/testify/suite"
)

type EventCollectorTestSuite struct {
	SccTestSuite
}

func TestEventCollector(t *testing.T) {
	suite.Run(t, new(EventCollectorTestSuite))
}

func (s *EventCollectorTestSuite) TestProcessStateBatchAppendedEvent() {
	// emit `StateBatchAppended` events
	var emits []*tc.SccStateBatchAppended
	for i := range s.Range(0, 10) {
		emits = append(emits, s.emitStateBatchAppendedEvent(i))
	}

	// collect `StateBatchAppended` events
	s.stateCollector.work(context.Background())

	// assert
	for i := range s.Range(0, 10) {
		got, _ := s.db.Optimism.FindState(s.sccAddr, uint64(i))
		s.Equal(s.sccAddr, got.OptimismScc.Address)
		s.Equal(emits[i].BatchIndex.Uint64(), got.BatchIndex)
		s.Equal(emits[i].BatchRoot[:], got.BatchRoot[:])
		s.Equal(emits[i].BatchSize.Uint64(), got.BatchSize)
		s.Equal(emits[i].PrevTotalElements.Uint64(), got.PrevTotalElements)
		s.Equal(emits[i].ExtraData, got.ExtraData)
	}
}

func (s *EventCollectorTestSuite) TestProcessStateBatchDeletedEvent() {
	ctx := context.Background()

	// emit `StateBatchAppended` events
	var emits []*tc.SccStateBatchAppended
	for i := range s.Range(0, 10) {
		emits = append(emits, s.emitStateBatchAppendedEvent(i))
	}

	// collect `StateBatchAppended` events
	s.stateCollector.work(ctx)

	// create signature records
	var creates []*database.OptimismSignature
	for i := range s.Range(0, 10) {
		sig, _ := s.db.OPSignature.Save(
			nil, nil,
			s.hub.Signer(),
			s.sccAddr,
			emits[i].BatchIndex.Uint64(),
			emits[i].BatchRoot,
			true,
			database.RandSignature(),
		)
		creates = append(creates, sig)
	}

	// emit `StateBatchDeleted` event
	s.scc.EmitStateBatchDeleted(
		s.hub.TransactOpts(ctx),
		emits[5].BatchIndex,
		emits[5].BatchRoot,
	)
	s.mining()

	// collect `StateBatchDeleted` events
	s.stateCollector.work(ctx)

	// assert
	for i := range s.Range(0, 10) {
		var want error
		if i >= 5 {
			want = database.ErrNotFound
		}
		_, err0 := s.db.Optimism.FindState(s.sccAddr, uint64(i))
		_, err1 := s.db.OPSignature.FindByID(creates[i].ID)
		s.Equal(want, err0)
		s.Equal(want, err1)
	}
}

func (s *EventCollectorTestSuite) TestProcessStateBatchVerifiedEvent() {
	// emit `EmitStateBatchVerified` events
	for index := range s.Range(0, 5) {
		s.scc.EmitStateBatchVerified(
			s.hub.TransactOpts(context.Background()),
			big.NewInt(int64(index)),
			s.RandHash(),
		)
		s.mining()
	}

	// collect `EmitStateBatchVerified` events
	s.stateCollector.work(context.Background())

	// assert
	scc, _ := s.db.OPContract.FindOrCreate(s.sccAddr)
	s.Equal(uint64(5), scc.NextIndex)
}

func (s *EventCollectorTestSuite) TestNoHandleOtherEvent() {
	ctx := context.Background()

	// emit `StateBatchAppended` and `Other` events
	for i := range s.Range(0, 10) {
		s.emitStateBatchAppendedEvent(i)
		s.scc.EmitOtherEvent(s.hub.TransactOpts(ctx), big.NewInt(11))
		s.mining()
	}

	// collect `StateBatchAppended` events
	s.stateCollector.work(ctx)

	// assert
	for i := range s.Range(0, 20) {
		var want error
		if i >= 10 {
			want = database.ErrNotFound
		}
		_, err := s.db.Optimism.FindState(s.sccAddr, uint64(i))
		s.ErrorIs(err, want)
	}
}

func (s *EventCollectorTestSuite) TestHandleReorganization() {
	ctx := context.Background()

	// emit `StateBatchAppended` events
	var emits []*tc.SccStateBatchAppended
	for i := range s.Range(0, 10) {
		emits = append(emits, s.emitStateBatchAppendedEvent(i))
	}

	// collect `StateBatchAppended` events
	s.stateCollector.work(ctx)

	// create signature records
	var creates []*database.OptimismSignature
	for i := range s.Range(0, 10) {
		sig, _ := s.db.OPSignature.Save(
			nil, nil,
			s.hub.Signer(),
			s.sccAddr,
			emits[i].BatchIndex.Uint64(),
			emits[i].BatchRoot,
			true,
			database.RandSignature(),
		)
		creates = append(creates, sig)
	}

	// simulate chain reorganization
	s.emitStateBatchAppendedEvent(4)
	s.stateCollector.work(ctx)

	// assert
	for i := range s.Range(0, 10) {
		_, err := s.db.Optimism.FindState(s.sccAddr, uint64(i))
		if i < 5 {
			s.NoError(err)
		} else {
			s.Error(err, database.ErrNotFound)
		}

		_, err = s.db.OPSignature.FindByID(creates[i].ID)
		if i < 4 {
			s.NoError(err)
		} else {
			s.Error(err, database.ErrNotFound)
		}
	}
}

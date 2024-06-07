package collector

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	tscc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/scc"
	"github.com/stretchr/testify/suite"
)

type EventCollectorTestSuite struct {
	backend.BackendSuite

	collector *EventCollector
	eventDB   database.IOPEventDB
}

func TestEventCollector(t *testing.T) {
	suite.Run(t, new(EventCollectorTestSuite))
}

func (s *EventCollectorTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()

	s.collector = NewEventCollector(&config.Verifier{
		Interval:         time.Millisecond,
		EventFilterLimit: 1000,
	}, s.DB, s.Hub, s.SignableHub.Signer())
	s.eventDB = database.NewOPEventDB[database.OptimismState](s.DB)
}

func (s *EventCollectorTestSuite) TestHandleRollupedEvent() {
	// emit `StateBatchAppended` events
	emits := make([]*tscc.SccStateBatchAppended, 10)
	for i := range s.Range(0, len(emits)) {
		_, emits[i] = s.EmitStateBatchAppended(i)
	}

	// collect `StateBatchAppended` events
	s.collector.Work(context.Background())

	// assert
	for i := range s.Range(0, len(emits)) {
		got, err := s.eventDB.FindByRollupIndex(s.SCCAddr, uint64(i))
		s.NoError(err)

		gott := got.(*database.OptimismState)
		s.Equal(s.SCCAddr, gott.Contract.Address)
		s.Equal(emits[i].BatchIndex.Uint64(), gott.BatchIndex)
		s.Equal(emits[i].BatchRoot[:], gott.BatchRoot[:])
		s.Equal(emits[i].BatchSize.Uint64(), gott.BatchSize)
		s.Equal(emits[i].PrevTotalElements.Uint64(), gott.PrevTotalElements)
		s.Equal(emits[i].ExtraData, gott.ExtraData)
	}
}

func (s *EventCollectorTestSuite) TestHandleDeletedEvent() {
	ctx := context.Background()

	// emit `StateBatchAppended` events
	emits := make([]*tscc.SccStateBatchAppended, 10)
	for i := range s.Range(0, len(emits)) {
		_, emits[i] = s.EmitStateBatchAppended(i)
	}

	// collect `StateBatchAppended` events
	s.collector.Work(ctx)

	// create signature records
	var creates []*database.OptimismSignature
	for i := range s.Range(0, len(emits)) {
		sig, err := s.DB.OPSignature.Save(
			nil, nil,
			s.SignableHub.Signer(),
			s.SCCAddr,
			emits[i].BatchIndex.Uint64(),
			emits[i].BatchRoot,
			true,
			database.RandSignature(),
		)
		s.NoError(err)
		creates = append(creates, sig)
	}

	// emit `StateBatchDeleted` event
	_, err := s.TSCC.EmitStateBatchDeleted(
		s.SignableHub.TransactOpts(ctx),
		emits[5].BatchIndex,
		emits[5].BatchRoot,
	)
	s.NoError(err)
	s.Mining()

	// collect `StateBatchDeleted` events
	s.collector.Work(ctx)

	// assert
	for i := range s.Range(0, 10) {
		var want error
		if i >= 5 {
			want = database.ErrNotFound
		}
		_, err0 := s.eventDB.FindByRollupIndex(s.SCCAddr, uint64(i))
		_, err1 := s.DB.OPSignature.FindByID(creates[i].ID)
		s.Equal(want, err0)
		s.Equal(want, err1)
	}
}

func (s *EventCollectorTestSuite) TestHandleVerifiedEvent() {
	// emit `EmitStateBatchVerified` events
	for index := range s.Range(0, 5) {
		_, err := s.TSCC.EmitStateBatchVerified(
			s.SignableHub.TransactOpts(context.Background()),
			big.NewInt(int64(index)),
			s.RandHash(),
		)
		s.NoError(err)
		s.Mining()
	}

	// collect `EmitStateBatchVerified` events
	s.collector.Work(context.Background())

	// assert
	scc, _ := s.DB.OPContract.FindOrCreate(s.SCCAddr)
	s.Equal(uint64(5), scc.NextIndex)
}

func (s *EventCollectorTestSuite) TestIgnoreOtherEvent() {
	ctx := context.Background()

	// emit `StateBatchAppended` and `Other` events
	for i := range s.Range(0, 10) {
		s.EmitStateBatchAppended(i)
		s.Mining()

		_, err := s.TSCC.EmitOtherEvent(s.SignableHub.TransactOpts(ctx), big.NewInt(11))
		s.NoError(err)
		s.Mining()
	}

	// collect `StateBatchAppended` events
	s.collector.Work(ctx)

	// assert
	for i := range s.Range(0, 20) {
		var want error
		if i >= 10 {
			want = database.ErrNotFound
		}
		_, err := s.eventDB.FindByRollupIndex(s.SCCAddr, uint64(i))
		s.ErrorIs(err, want)
	}
}

func (s *EventCollectorTestSuite) TestHandleReorganization() {
	ctx := context.Background()

	// emit `StateBatchAppended` events
	emits := make([]*tscc.SccStateBatchAppended, 10)
	for i := range s.Range(0, len(emits)) {
		_, emits[i] = s.EmitStateBatchAppended(i)
	}

	// collect `StateBatchAppended` events
	s.collector.Work(ctx)

	// create signature records
	var creates []*database.OptimismSignature
	for i := range s.Range(0, len(emits)) {
		sig, err := s.DB.OPSignature.Save(
			nil, nil,
			s.SignableHub.Signer(),
			s.SCCAddr,
			emits[i].BatchIndex.Uint64(),
			emits[i].BatchRoot,
			true,
			database.RandSignature(),
		)
		s.NoError(err)
		creates = append(creates, sig)
	}

	// simulate chain reorganization
	s.EmitStateBatchAppended(4)
	s.collector.Work(ctx)

	// assert
	for i := range s.Range(0, len(emits)) {
		_, err := s.eventDB.FindByRollupIndex(s.SCCAddr, uint64(i))
		if i < 5 {
			s.NoError(err)
		} else {
			s.Error(err, database.ErrNotFound)
		}

		_, err = s.DB.OPSignature.FindByID(creates[i].ID)
		if i < 4 {
			s.NoError(err)
		} else {
			s.Error(err, database.ErrNotFound)
		}
	}
}

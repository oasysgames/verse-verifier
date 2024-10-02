package verse

import (
	"context"
	"math/big"
	"testing"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	"github.com/stretchr/testify/suite"
)

type VersePoolTestSuite struct {
	backend.BackendSuite

	pool VersePool
	verse1,
	verse2 Verse
}

func TestVersePool(t *testing.T) {
	suite.Run(t, new(VersePoolTestSuite))
}

func (s *VersePoolTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()

	s.pool = NewVersePool(s.Hub)
	s.verse1 = NewOPLegacy(s.DB, s.Hub, 1, s.Hub.URL(), s.SCCAddr, s.SCCVAddr)
	s.verse2 = NewOPStack(s.DB, s.Hub, 2, s.Hub.URL(), s.L2OOAddr, s.L2OOVAddr)
}

func (s *VersePoolTestSuite) TestAdd_Get_Range() {
	// Get verse1
	got0, got1 := s.pool.Get(s.verse1.RollupContract())
	s.Nil(got0)
	s.False(got1)

	// Add verse1
	s.pool.Add(s.verse1, false)
	got0, _ = s.pool.Get(s.verse1.RollupContract())
	s.Equal(s.verse1, got0.Verse())
	s.False(got0.CanSubmit())

	// Range verse1
	rangeItems := make(map[uint64]*VersePoolItem)
	s.pool.Range(func(item *VersePoolItem) bool {
		rangeItems[item.Verse().ChainID()] = item
		return true
	})
	s.Len(rangeItems, 1)
	s.Equal(s.verse1, rangeItems[s.verse1.ChainID()].Verse())
	s.False(rangeItems[s.verse1.ChainID()].CanSubmit())

	// Get verse2
	got0, got1 = s.pool.Get(s.verse2.RollupContract())
	s.Nil(got0)
	s.False(got1)

	// Add verse1
	s.pool.Add(s.verse2, true)
	got0, _ = s.pool.Get(s.verse2.RollupContract())
	s.Equal(s.verse2, got0.Verse())
	s.True(got0.CanSubmit())

	// Range verse2
	rangeItems = make(map[uint64]*VersePoolItem)
	s.pool.Range(func(item *VersePoolItem) bool {
		rangeItems[item.Verse().ChainID()] = item
		return true
	})
	s.Len(rangeItems, 2)
	s.Equal(s.verse1, rangeItems[s.verse1.ChainID()].Verse())
	s.False(rangeItems[s.verse1.ChainID()].CanSubmit())
	s.Equal(s.verse2, rangeItems[s.verse2.ChainID()].Verse())
	s.True(rangeItems[s.verse2.ChainID()].CanSubmit())

}

func (s *VersePoolTestSuite) TestDelete() {
	s.pool.Add(s.verse1, false)
	s.pool.Delete(s.verse1.RollupContract())
	got0, got1 := s.pool.Get(s.verse1.RollupContract())
	s.Nil(got0)
	s.False(got1)
}

func (s *VersePoolTestSuite) TestNextIndex() {
	confirmation := 3
	ctx := context.Background()

	// If it does not exist in the pool, an error should be returned
	_, err := s.pool.NextIndex(ctx, s.verse1.RollupContract(), confirmation, true)
	s.ErrorContains(err, "not in the pool")

	// Add pool
	s.pool.Add(s.verse1, false)

	// 1st call
	got, _ := s.pool.NextIndex(ctx, s.verse1.RollupContract(), confirmation, true)
	s.Equal(uint64(0), got)

	// Update next index
	s.TSCC.SetNextIndex(s.SignableHub.TransactOpts(ctx), big.NewInt(1))

	// Since the block is not yet confirmed, the cache should be returned.
	for range s.Range(0, confirmation) {
		s.Hub.Mining()
		got, _ = s.pool.NextIndex(ctx, s.verse1.RollupContract(), confirmation, true)
		s.Equal(uint64(0), got)
	}

	// Since the block has been confirmed, a new value should be returned.
	s.Hub.Mining()
	got, _ = s.pool.NextIndex(ctx, s.verse1.RollupContract(), confirmation, true)
	s.Equal(uint64(1), got)

	// And the cache should also be updated.
	got, _ = s.pool.NextIndex(ctx, s.verse1.RollupContract(), confirmation, true)
	s.Equal(uint64(1), got)

	// Adjust latest block.
	s.Hub.Minings(10)
}

func (s *VersePoolTestSuite) TestEventEmittedBlock() {
	confirmation := 3
	ctx := context.Background()

	// If it does not exist in the pool, an error should be returned
	_, err := s.pool.NextIndex(ctx, s.verse1.RollupContract(), confirmation, true)
	s.ErrorContains(err, "not in the pool")

	// Add pool
	s.pool.Add(s.verse1, false)

	// Emit rollup event
	rollupIndex := 0
	s.EmitStateBatchAppended(rollupIndex)
	want0, _ := s.Hub.BlockNumber(ctx)

	// Since the block is not yet confirmed, the `ErrEventNotFound` should be returned.
	for range s.Range(0, confirmation-1) {
		s.Hub.Mining()
		_, err = s.pool.EventEmittedBlock(
			ctx, s.verse1.RollupContract(), uint64(rollupIndex), confirmation, true)
		s.ErrorIs(err, ErrEventNotFound)
	}

	// Since the block has been confirmed, emitted block should be returned.
	s.Hub.Mining()
	got0, _ := s.pool.EventEmittedBlock(
		ctx, s.verse1.RollupContract(), uint64(rollupIndex), confirmation, true)
	s.Equal(want0, got0)

	// Emit rollup event
	rollupIndex++
	s.EmitStateBatchAppended(1)
	want1, _ := s.Hub.BlockNumber(ctx)

	// Since the block is not yet confirmed, the cache should be returned.
	for range s.Range(0, confirmation-1) {
		s.Hub.Mining()
		_, err = s.pool.EventEmittedBlock(
			ctx, s.verse1.RollupContract(), uint64(rollupIndex), confirmation, true)
		s.ErrorIs(err, ErrEventNotFound)
	}

	// Since the block has been confirmed, emitted block should be returned.
	s.Hub.Mining()
	got1, _ := s.pool.EventEmittedBlock(
		ctx, s.verse1.RollupContract(), uint64(rollupIndex), confirmation, true)
	s.Equal(want1, got1)
}

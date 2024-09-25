package verifier

import (
	"context"
	"math/big"
	"testing"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	"github.com/stretchr/testify/suite"
)

type BlockRangeManagerTestSuite struct {
	backend.BackendSuite

	startOffSet       uint64
	maxRange          uint64
	blockRangeManager *eventFetchingBlockRangeManager
}

func TestBlockRangeManager(t *testing.T) {
	suite.Run(t, new(BlockRangeManagerTestSuite))
}

func (s *BlockRangeManagerTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()
	s.startOffSet = 10
	s.maxRange = 5
	s.blockRangeManager = NeweventFetchingBlockRangeManager(s.SignableHub, s.maxRange, s.startOffSet)
}

func (s *BlockRangeManagerTestSuite) TestGetBlockRange() {
	var (
		ctx = context.Background()
	)
	// Increment Block
	for i := 0; i < 10; i++ {
		s.Hub.Mining()
	}

	// get latest block
	latest, _ := s.Hub.Client().BlockByNumber(ctx, nil)

	// 1st call get block range
	start, end, skipFetchlog, err := s.blockRangeManager.GetBlockRange(ctx)
	s.NoError(err)
	expextedStart := latest.NumberU64() - s.startOffSet
	expectedEnd := latest.NumberU64() - s.startOffSet + s.maxRange
	s.Equal(expextedStart, start)
	s.Equal(expectedEnd, end)
	s.False(skipFetchlog)

	// 2nd call get block range
	expextedStart = start + s.maxRange + 1
	expectedEnd = latest.NumberU64()
	start, end, skipFetchlog, err = s.blockRangeManager.GetBlockRange(ctx)
	s.NoError(err)
	s.Equal(expextedStart, start)
	s.Equal(expectedEnd, end)
	s.False(skipFetchlog)

	// Restore next start
	s.blockRangeManager.RestoreNextStart()

	// 3rd call get block range
	start, end, skipFetchlog, err = s.blockRangeManager.GetBlockRange(ctx)
	s.NoError(err)
	s.Equal(expextedStart, start)
	s.Equal(expectedEnd, end)
	s.False(skipFetchlog)

	// 4th call get block range
	expextedStart = end + 1
	expectedEnd = latest.NumberU64()
	start, end, skipFetchlog, err = s.blockRangeManager.GetBlockRange(ctx)
	s.NoError(err)
	s.Equal(expextedStart, start)
	s.Equal(expectedEnd, end)
	s.True(skipFetchlog)
}

func (s *BlockRangeManagerTestSuite) TestCheckIfStartTooLarge() {
	var (
		ctx             = context.Background()
		nextRollupIndex = big.NewInt(0)
	)

	// get latest block
	latest, _ := s.Hub.Client().BlockByNumber(ctx, nil)

	// Set nextStart via calling GetBlockRange
	s.blockRangeManager.GetBlockRange(ctx)
	expextedNextStart := latest.NumberU64() + 1
	s.Equal(expextedNextStart, s.blockRangeManager.nextStart)

	// 1st call CheckIfStartTooLarge
	err := s.blockRangeManager.CheckIfStartTooLarge(nextRollupIndex, nextRollupIndex.Uint64()+1)
	s.ErrorIs(err, ErrStartBlockIsTooLarge)
	expextedNextStart = 0
	s.Equal(expextedNextStart, s.blockRangeManager.nextStart)

	// 2nd call CheckIfStartTooLarge
	err = s.blockRangeManager.CheckIfStartTooLarge(nextRollupIndex, nextRollupIndex.Uint64())
	s.NoError(err)
	s.True(s.blockRangeManager.startTooLargeCheckPassed)
}

package verifier

import (
	"testing"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/stretchr/testify/suite"
)

type BlockRangeManagerTestSuite struct {
	testhelper.Suite
}

func TestBlockRangeManager(t *testing.T) {
	suite.Run(t, new(BlockRangeManagerTestSuite))
}

func (s *BlockRangeManagerTestSuite) TestGetBlockRange() {
	maxRange := 5
	initialStart := uint64(1)
	latest := uint64(13)

	manager := newEventFetchingBlockRangeManager(maxRange, initialStart)

	// 1st call get block range
	start, end, skipFetchlog := manager.get(latest)
	s.Equal(uint64(1), start)
	s.Equal(uint64(5), end)
	s.False(skipFetchlog)

	// 2nd call get block range
	start, end, skipFetchlog = manager.get(latest)
	s.Equal(uint64(6), start)
	s.Equal(uint64(10), end)
	s.False(skipFetchlog)

	// Restore next start
	manager.restore()

	// 3rd call get block range
	start, end, skipFetchlog = manager.get(latest)
	s.Equal(uint64(6), start)
	s.Equal(uint64(10), end)
	s.False(skipFetchlog)

	// 4th call get block range
	start, end, skipFetchlog = manager.get(latest)
	s.Equal(uint64(11), start)
	s.Equal(uint64(13), end)
	s.False(skipFetchlog)

	// 4th call get block range
	start, end, skipFetchlog = manager.get(latest)
	s.Equal(uint64(14), start)
	s.Equal(uint64(13), end)
	s.True(skipFetchlog)
}

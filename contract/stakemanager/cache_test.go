package stakemanager

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/stretchr/testify/suite"
)

type CacheTestSuite struct {
	testhelper.Suite

	sm *testhelper.StakeManagerMock
	vs *Cache
}

func TestNewCache(t *testing.T) {
	suite.Run(t, new(CacheTestSuite))
}

func (s *CacheTestSuite) SetupTest() {
	s.sm = &testhelper.StakeManagerMock{}
	s.vs = NewCache(s.sm, time.Millisecond*5)

	for i := range s.Range(0, 1000) {
		s.sm.Owners = append(s.sm.Owners, s.RandAddress())
		s.sm.Operators = append(s.sm.Operators, s.RandAddress())
		s.sm.Stakes = append(s.sm.Stakes, big.NewInt(int64(i)))
		s.sm.Candidates = append(s.sm.Candidates, i%5 == 0)
	}
}

func (s *CacheTestSuite) TestTotalStake() {
	ctx := context.Background()

	s.Equal(big.NewInt(499500), s.vs.TotalStake(ctx))

	s.sm.Stakes[0] = new(big.Int).Add(s.sm.Stakes[0], common.Big1)
	s.Equal(big.NewInt(499500), s.vs.TotalStake(ctx))

	time.Sleep(time.Millisecond * 10)
	s.Equal(big.NewInt(499501), s.vs.TotalStake(ctx))
}

func (s *CacheTestSuite) TestStakeBySigner() {
	ctx := context.Background()

	for i, signer := range s.sm.Operators {
		s.Equal(s.sm.Stakes[i], s.vs.StakeBySigner(ctx, signer))
	}

	old := new(big.Int).Set(s.sm.Stakes[0])
	s.sm.Stakes[0] = new(big.Int).Add(s.sm.Stakes[0], common.Big1)
	s.Equal(old, s.vs.StakeBySigner(ctx, s.sm.Operators[0]))

	time.Sleep(time.Millisecond * 10)
	s.Equal(s.sm.Stakes[0], s.vs.StakeBySigner(ctx, s.sm.Operators[0]))
}

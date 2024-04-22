package stakemanager

import (
	"context"
	"math/big"
	"testing"

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
	s.vs = NewCache(s.sm)

	for i := range s.Range(0, 1000) {
		s.sm.Owners = append(s.sm.Owners, s.RandAddress())
		s.sm.Operators = append(s.sm.Operators, s.RandAddress())
		s.sm.Stakes = append(s.sm.Stakes, big.NewInt(int64(i)))
		s.sm.Candidates = append(s.sm.Candidates, i%5 == 0)
	}
}

func (s *CacheTestSuite) TestRefresh() {
	s.Nil(s.vs.Refresh(context.Background()))

	// assert `TotalStake()`
	s.Equal(big.NewInt(499500), s.vs.TotalStake())

	// assert `SignerStakes()`
	signerStakes := s.vs.SignerStakes()
	s.Len(signerStakes, len(s.sm.Operators))
	for i, signer := range s.sm.Operators {
		s.Equal(s.sm.Stakes[i], signerStakes[signer])
	}

	// assert `StakeBySigner(common.Address)`
	for i, signer := range s.sm.Operators {
		s.Equal(s.sm.Stakes[i], s.vs.StakeBySigner(signer))
	}

	// assert `Candidates()`
	candidates := s.vs.Candidates()
	s.Len(candidates, len(s.sm.Operators))
	for i, signer := range s.sm.Operators {
		s.Equal(s.sm.Candidates[i], candidates[signer])
	}

	// assert `Candidate(common.Address)`
	for i, signer := range s.sm.Operators {
		s.Equal(s.sm.Candidates[i], s.vs.Candidate(signer))
	}
}

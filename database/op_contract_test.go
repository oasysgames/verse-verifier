package database

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestOptimismContractDB(t *testing.T) {
	suite.Run(t, new(OptimismContractDBTestSuite))
}

type OptimismContractDBTestSuite struct {
	DatabaseTestSuite

	db *OptimismContractDB
}

func (s *OptimismContractDBTestSuite) SetupTest() {
	s.DatabaseTestSuite.SetupTest()
	s.db = s.DatabaseTestSuite.db.OPContract
}

func (s *OptimismContractDBTestSuite) TestFindOrCreate() {
	assert := func(got0, got1, got2 *OptimismContract) {
		var count int
		s.db.rawdb.Table("optimism_sccs").Select("COUNT(*)").Row().Scan(&count)
		s.Equal(3, count)

		s.Equal(uint64(1), got0.ID)
		s.Equal(uint64(2), got1.ID)
		s.Equal(uint64(3), got2.ID)
		s.Equal(s.ItoAddress(1), got0.Address)
		s.Equal(s.ItoAddress(2), got1.Address)
		s.Equal(s.ItoAddress(3), got2.Address)
	}

	addr0 := s.ItoAddress(1)
	addr1 := s.ItoAddress(2)
	addr2 := s.ItoAddress(3)

	got0, _ := s.db.FindOrCreate(addr0)
	got1, _ := s.db.FindOrCreate(addr1)
	got2, _ := s.db.FindOrCreate(addr2)
	assert(got0, got1, got2)

	got0, _ = s.db.FindOrCreate(addr0)
	got1, _ = s.db.FindOrCreate(addr1)
	got2, _ = s.db.FindOrCreate(addr2)
	assert(got0, got1, got2)

	// test upsert
	s.db.SaveNextIndex(addr0, 10)
	got0, _ = s.db.FindOrCreate(addr0)
	s.Equal(uint64(10), got0.NextIndex)
}

func (s *OptimismContractDBTestSuite) TestSaveNextIndex() {
	scc0 := s.createContract()
	scc1 := s.createContract()

	s.Equal(uint64(0), scc0.NextIndex)
	s.Equal(uint64(0), scc1.NextIndex)

	s.db.SaveNextIndex(scc0.Address, 5)
	s.db.SaveNextIndex(scc1.Address, 10)

	scc0, _ = s.db.FindOrCreate(scc0.Address)
	scc1, _ = s.db.FindOrCreate(scc1.Address)

	s.Equal(uint64(5), scc0.NextIndex)
	s.Equal(uint64(10), scc1.NextIndex)
}

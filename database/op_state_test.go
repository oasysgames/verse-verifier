package database

import (
	"math/big"
	"testing"

	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/stretchr/testify/suite"
)

func TestOptimismStateDB(t *testing.T) {
	suite.Run(t, new(OptimismStateDBTestSuite))
}

type OptimismStateDBTestSuite struct {
	DatabaseTestSuite

	db IOPEventDB
}

func (s *OptimismStateDBTestSuite) SetupTest() {
	s.DatabaseTestSuite.SetupTest()
	s.db = NewOPEventDB[OptimismState](s.DatabaseTestSuite.db)
}

func (s *OptimismStateDBTestSuite) TestFindByRollupIndex() {
	scc0 := s.createContract()
	scc1 := s.createContract()

	st0 := s.createState(scc0, 0)
	st1 := s.createState(scc1, 0)
	st2 := s.createState(scc1, 1)

	got0, _ := s.db.FindByRollupIndex(scc0.Address, 0)
	got1, _ := s.db.FindByRollupIndex(scc1.Address, 0)
	got2, _ := s.db.FindByRollupIndex(scc1.Address, 1)

	s.Equal(st0, got0)
	s.Equal(st1, got1)
	s.Equal(st2, got2)

	_, err := s.db.FindByRollupIndex(scc0.Address, 1)
	s.Error(err, ErrNotFound)

	_, err = s.db.FindByRollupIndex(scc1.Address, 2)
	s.Error(err, ErrNotFound)
}

func (s *OptimismStateDBTestSuite) TestFindForVerification() {
	assert := func(gots []OPEvent, wants []*OptimismState) {
		s.Len(gots, len(wants))
		for i, got := range gots {
			gott := got.(*OptimismState)
			s.Equal(wants[i].ID, gott.ID)
		}
	}

	// Create dummy records
	count := 10
	signer := s.createSigner()
	scc0, scc1 := s.createContract(), s.createContract()
	wants0, wants1 := make([]*OptimismState, count), make([]*OptimismState, count)
	for _, batchIndex := range s.Shuffle(s.Range(0, count)) {
		wants0[batchIndex] = s.createState(scc0, batchIndex)
		wants1[batchIndex] = s.createState(scc1, batchIndex)
	}

	gots, _ := s.db.FindForVerification(signer.Address, scc0.Address, 0, 100)
	assert(gots, wants0)

	gots, _ = s.db.FindForVerification(signer.Address, scc1.Address, 0, 100)
	assert(gots, wants1)

	// when limit is set
	gots, _ = s.db.FindForVerification(signer.Address, scc0.Address, 0, 2)
	assert(gots, wants0[:2])

	gots, _ = s.db.FindForVerification(signer.Address, scc1.Address, 0, 4)
	assert(gots, wants1[:4])

	// when `nextIndex` is set to query
	gots, _ = s.db.FindForVerification(signer.Address, scc0.Address, 2, 100)
	assert(gots, wants0[2:])

	gots, _ = s.db.FindForVerification(signer.Address, scc1.Address, 3, 100)
	assert(gots, wants1[3:])

	// when `nextIndex` is set to scc
	s.db.DB().OPContract.SaveNextIndex(scc0.Address, 3)
	s.db.DB().OPContract.SaveNextIndex(scc1.Address, 4)

	gots, _ = s.db.FindForVerification(signer.Address, scc0.Address, 0, 100)
	assert(gots, wants0[3:])

	gots, _ = s.db.FindForVerification(signer.Address, scc1.Address, 0, 100)
	assert(gots, wants1[4:])

	// when several signatures exist
	merge := func(a, b []*OptimismState) (m []*OptimismState) {
		m = append(m, a...)
		m = append(m, b...)
		return m
	}
	s.createSignature(signer, scc0, 6)
	s.createSignature(signer, scc1, 8)

	gots, _ = s.db.FindForVerification(signer.Address, scc0.Address, 0, 100)
	assert(gots, merge(wants0[3:6], wants0[7:]))

	gots, _ = s.db.FindForVerification(signer.Address, scc1.Address, 0, 100)
	assert(gots, merge(wants1[4:8], wants1[9:]))

	gots, _ = s.db.FindForVerification(signer.Address, scc0.Address, 8, 100)
	assert(gots, wants0[8:])

	gots, _ = s.db.FindForVerification(signer.Address, scc1.Address, 9, 100)
	assert(gots, wants1[9:])
}

func (s *OptimismStateDBTestSuite) TestSave() {
	scc_ := s.createContract()
	batchIndex := uint64(1)

	_, err := s.db.FindByRollupIndex(scc_.Address, batchIndex)
	s.Error(err, ErrNotFound)

	ev := &scc.SccStateBatchAppended{
		BatchIndex:        new(big.Int).SetUint64(batchIndex),
		BatchRoot:         s.RandHash(),
		BatchSize:         big.NewInt(12),
		PrevTotalElements: big.NewInt(3),
		ExtraData:         []byte("test"),
	}

	got, _ := s.db.Save(scc_.Address, ev)
	gott := got.(*OptimismState)
	s.Equal(uint64(1), gott.ID)
	s.Equal(*scc_, gott.Contract)
	s.Equal(ev.BatchIndex.Uint64(), gott.BatchIndex)
	s.Equal(ev.BatchRoot[:], gott.BatchRoot[:])
	s.Equal(ev.BatchSize.Uint64(), gott.BatchSize)
	s.Equal(ev.PrevTotalElements.Uint64(), gott.PrevTotalElements)
	s.Equal(ev.ExtraData, gott.ExtraData)

	found, _ := s.db.FindByRollupIndex(scc_.Address, batchIndex)
	s.Equal(gott, found)
}

func (s *OptimismStateDBTestSuite) TestDeletes() {
	assert := func(scc *OptimismContract, want []int) {
		var gots []int
		s.DatabaseTestSuite.db.rawdb.Model(&OptimismState{}).
			Where("optimism_scc_id = ?", scc.ID).
			Order("batch_index").
			Pluck("batch_index", &gots)
		s.Equal(want, gots)
	}

	scc0 := s.createContract()
	scc1 := s.createContract()
	for _, i := range s.Shuffle(s.Range(0, 10)) {
		s.createState(scc0, i)
		s.createState(scc1, i)
	}

	assert(scc0, s.Range(0, 10))
	assert(scc1, s.Range(0, 10))

	rows0, _ := s.db.Deletes(scc0.Address, 3)
	rows1, _ := s.db.Deletes(scc1.Address, 6)

	s.Equal(int64(7), rows0)
	s.Equal(int64(4), rows1)
	assert(scc0, s.Range(0, 3))
	assert(scc1, s.Range(0, 6))
}

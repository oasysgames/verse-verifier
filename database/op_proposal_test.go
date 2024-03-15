package database

import (
	"math/big"
	"testing"

	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/stretchr/testify/suite"
)

func TestOpstackProposalDB(t *testing.T) {
	suite.Run(t, new(OpstackProposalDBTestSuite))
}

type OpstackProposalDBTestSuite struct {
	DatabaseTestSuite

	db IOPEventDB
}

func (s *OpstackProposalDBTestSuite) SetupTest() {
	s.DatabaseTestSuite.SetupTest()
	s.db = NewOPEventDB[OpstackProposal](s.DatabaseTestSuite.db)
}

func (s *OpstackProposalDBTestSuite) TestFindByRollupIndex() {
	l2oo0 := s.createContract()
	l2oo1 := s.createContract()

	st0 := s.createProposal(l2oo0, 0)
	st1 := s.createProposal(l2oo1, 0)
	st2 := s.createProposal(l2oo1, 1)

	got0, _ := s.db.FindByRollupIndex(l2oo0.Address, 0)
	got1, _ := s.db.FindByRollupIndex(l2oo1.Address, 0)
	got2, _ := s.db.FindByRollupIndex(l2oo1.Address, 1)

	s.Equal(st0, got0)
	s.Equal(st1, got1)
	s.Equal(st2, got2)

	_, err := s.db.FindByRollupIndex(l2oo0.Address, 1)
	s.Error(err, ErrNotFound)

	_, err = s.db.FindByRollupIndex(l2oo1.Address, 2)
	s.Error(err, ErrNotFound)
}

func (s *OpstackProposalDBTestSuite) TestFindForVerification() {
	assert := func(gots []OPEvent, wants []*OpstackProposal) {
		s.Len(gots, len(wants))
		for i, got := range gots {
			gott := got.(*OpstackProposal)
			s.Equal(wants[i].ID, gott.ID)
		}
	}

	// Create dummy records
	count := 10
	signer := s.createSigner()
	l2oo0, l2oo1 := s.createContract(), s.createContract()
	wants0, wants1 := make([]*OpstackProposal, count), make([]*OpstackProposal, count)
	for _, outputIndex := range s.Shuffle(s.Range(0, count)) {
		wants0[outputIndex] = s.createProposal(l2oo0, outputIndex)
		wants1[outputIndex] = s.createProposal(l2oo1, outputIndex)
	}

	gots, _ := s.db.FindForVerification(signer.Address, l2oo0.Address, 0, 100)
	assert(gots, wants0)

	gots, _ = s.db.FindForVerification(signer.Address, l2oo1.Address, 0, 100)
	assert(gots, wants1)

	// when limit is set
	gots, _ = s.db.FindForVerification(signer.Address, l2oo0.Address, 0, 2)
	assert(gots, wants0[:2])

	gots, _ = s.db.FindForVerification(signer.Address, l2oo1.Address, 0, 4)
	assert(gots, wants1[:4])

	// when `nextIndex` is set to query
	gots, _ = s.db.FindForVerification(signer.Address, l2oo0.Address, 2, 100)
	assert(gots, wants0[2:])

	gots, _ = s.db.FindForVerification(signer.Address, l2oo1.Address, 3, 100)
	assert(gots, wants1[3:])

	// when `nextIndex` is set to scc
	s.db.DB().OPContract.SaveNextIndex(l2oo0.Address, 3)
	s.db.DB().OPContract.SaveNextIndex(l2oo1.Address, 4)

	gots, _ = s.db.FindForVerification(signer.Address, l2oo0.Address, 0, 100)
	assert(gots, wants0[3:])

	gots, _ = s.db.FindForVerification(signer.Address, l2oo1.Address, 0, 100)
	assert(gots, wants1[4:])

	// when several signatures exist
	merge := func(a, b []*OpstackProposal) (m []*OpstackProposal) {
		m = append(m, a...)
		m = append(m, b...)
		return m
	}
	s.createSignature(signer, l2oo0, 6)
	s.createSignature(signer, l2oo1, 8)

	gots, _ = s.db.FindForVerification(signer.Address, l2oo0.Address, 0, 100)
	assert(gots, merge(wants0[3:6], wants0[7:]))

	gots, _ = s.db.FindForVerification(signer.Address, l2oo1.Address, 0, 100)
	assert(gots, merge(wants1[4:8], wants1[9:]))

	gots, _ = s.db.FindForVerification(signer.Address, l2oo0.Address, 8, 100)
	assert(gots, wants0[8:])

	gots, _ = s.db.FindForVerification(signer.Address, l2oo1.Address, 9, 100)
	assert(gots, wants1[9:])
}

func (s *OpstackProposalDBTestSuite) TestSave() {
	l2oo_ := s.createContract()
	outputIndex := uint64(1)

	_, err := s.db.FindByRollupIndex(l2oo_.Address, outputIndex)
	s.Error(err, ErrNotFound)

	ev := &l2oo.OasysL2OutputOracleOutputProposed{
		L2OutputIndex: new(big.Int).SetUint64(outputIndex),
		OutputRoot:    s.RandHash(),
		L2BlockNumber: big.NewInt(12),
		L1Timestamp:   big.NewInt(1704067200),
	}

	got, _ := s.db.Save(l2oo_.Address, ev)
	gott := got.(*OpstackProposal)
	s.Equal(uint64(1), gott.ID)
	s.Equal(*l2oo_, gott.Contract)
	s.Equal(ev.L2OutputIndex.Uint64(), gott.L2OutputIndex)
	s.Equal(ev.OutputRoot[:], gott.OutputRoot[:])
	s.Equal(ev.L2BlockNumber.Uint64(), gott.L2BlockNumber)
	s.Equal(ev.L1Timestamp.Uint64(), gott.L1Timestamp)

	found, _ := s.db.FindByRollupIndex(l2oo_.Address, outputIndex)
	s.Equal(got, found)
}

func (s *OpstackProposalDBTestSuite) TestDeletes() {
	assert := func(scc *OptimismContract, want []int) {
		var gots []int
		s.DatabaseTestSuite.db.rawdb.Model(&OpstackProposal{}).
			Where("contract_id = ?", scc.ID).
			Order("l2_output_index").
			Pluck("l2_output_index", &gots)
		s.Equal(want, gots)
	}

	l2oo0 := s.createContract()
	l2oo1 := s.createContract()
	for _, i := range s.Shuffle(s.Range(0, 10)) {
		s.createProposal(l2oo0, i)
		s.createProposal(l2oo1, i)
	}

	assert(l2oo0, s.Range(0, 10))
	assert(l2oo1, s.Range(0, 10))

	rows0, _ := s.db.Deletes(l2oo0.Address, 3)
	rows1, _ := s.db.Deletes(l2oo1.Address, 6)

	s.Equal(int64(7), rows0)
	s.Equal(int64(4), rows1)
	assert(l2oo0, s.Range(0, 3))
	assert(l2oo1, s.Range(0, 6))
}

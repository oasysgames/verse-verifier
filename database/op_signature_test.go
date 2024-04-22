package database

import (
	"sort"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/suite"
)

func TestOptimismSignatureDB(t *testing.T) {
	suite.Run(t, new(OptimismSignatureDBTestSuite))
}

type OptimismSignatureDBTestSuite struct {
	DatabaseTestSuite

	db *OptimismSignatureDB
}

func (s *OptimismSignatureDBTestSuite) SetupTest() {
	s.DatabaseTestSuite.SetupTest()
	s.db = s.DatabaseTestSuite.db.OPSignature
}

func (s *OptimismSignatureDBTestSuite) TestFindByID() {
	signer0 := s.createSigner()
	signer1 := s.createSigner()

	contract0 := s.createContract()
	contract1 := s.createContract()

	sig0 := s.createSignature(signer0, contract0, 0)
	sig1 := s.createSignature(signer1, contract1, 0)
	sig2 := s.createSignature(signer1, contract1, 1)

	got0, _ := s.db.FindByID(sig0.ID)
	got1, _ := s.db.FindByID(sig1.ID)
	got2, _ := s.db.FindByID(sig2.ID)

	s.Equal(sig0, got0)
	s.Equal(sig1, got1)
	s.Equal(sig2, got2)

	_, err := s.db.FindByID("")
	s.Error(err, ErrNotFound)
}

func (s *OptimismSignatureDBTestSuite) TestFind() {
	concat := func(slices ...[]*OptimismSignature) (cpy []*OptimismSignature) {
		for _, s := range slices {
			cpy = append(cpy, s...)
		}
		return cpy
	}
	assert := func(gots, wants []*OptimismSignature) {
		s.Len(gots, len(wants))
		for i, want := range wants {
			s.Equal(want.ID, gots[i].ID, i)
			s.Equal(want.SignerID, gots[i].SignerID, i)
			s.Equal(want.ContractID, gots[i].ContractID, i)
		}
	}

	signer0 := s.createSigner()
	signer1 := s.createSigner()
	contract0 := s.createContract()
	contract1 := s.createContract()

	var creates0, creates1, creates2, creates3 []*OptimismSignature
	for _, index := range s.Range(0, 5) {
		creates0 = append(creates0, s.createSignature(signer0, contract0, index))
	}
	for _, index := range s.Range(5, 10) {
		creates1 = append(creates1, s.createSignature(signer1, contract0, index))
	}
	for _, index := range s.Range(10, 15) {
		creates2 = append(creates2, s.createSignature(signer0, contract1, index))
	}
	for _, index := range s.Range(15, 20) {
		creates3 = append(creates3, s.createSignature(signer1, contract1, index))
	}

	index := uint64(5)
	cases := []struct {
		name          string
		idAfter       *string
		signer        *common.Address
		contract      *common.Address
		index         *uint64
		limit, offset int
		wants         []*OptimismSignature
	}{
		{
			"all",
			nil, nil, nil, nil, 100, 0,
			concat(creates0, creates1, creates2, creates3),
		},
		{
			"idAfter",
			&creates2[1].ID, nil, nil, nil, 100, 0,
			concat(creates2[1:], creates3),
		},
		{
			"signer0",
			nil, &signer0.Address, nil, nil, 100, 0,
			concat(creates0, creates2),
		},
		{
			"contract0",
			nil, nil, &contract0.Address, nil, 100, 0,
			concat(creates0, creates1),
		},
		{
			"index=5",
			nil, nil, nil, &index, 100, 0,
			creates1[0:1],
		},
		{
			"limit",
			nil, nil, nil, nil, 8, 0,
			concat(creates0, creates1[:3]),
		},
		{
			"offset",
			nil, nil, nil, nil, 100, 13,
			concat(creates2[3:], creates3),
		},
		{
			"over offset",
			nil, nil, nil, nil, 100, 100,
			[]*OptimismSignature{},
		},
	}
	for _, c := range cases {
		s.Run(c.name, func() {
			gots, _ := s.db.Find(
				c.idAfter, c.signer, c.contract, c.index, c.limit, c.offset)
			assert(gots, c.wants)
		})
	}
}

func (s *OptimismSignatureDBTestSuite) TestFindLatestsPerSigners() {
	signer0 := s.createSigner()
	signer1 := s.createSigner()
	signer2 := s.createSigner()
	contracts := []*OptimismContract{s.createContract(), s.createContract(), s.createContract()}

	var want0, want1, want2 *OptimismSignature
	for i := range s.Shuffle(s.Range(0, len(contracts))) {
		for _, index := range s.Range(0, 5) {
			want0 = s.createSignature(signer0, contracts[i], index)
		}
		for _, index := range s.Range(0, 10) {
			want1 = s.createSignature(signer1, contracts[i], index)
		}
		for _, index := range s.Range(0, 15) {
			want2 = s.createSignature(signer2, contracts[i], index)
		}
	}

	gots, _ := s.db.FindLatestsPerSigners()
	s.Len(gots, 3)
	s.Equal(want0, gots[0])
	s.Equal(want1, gots[1])
	s.Equal(want2, gots[2])
}

func (s *OptimismSignatureDBTestSuite) TestFindLatestsBySigner() {
	signer0 := s.createSigner()
	signer1 := s.createSigner()
	contract := s.createContract()

	var wants0, wants1 []*OptimismSignature
	for _, index := range s.Range(0, 5) {
		wants0 = append(wants0, s.createSignature(signer0, contract, index))
	}
	for _, index := range s.Range(0, 10) {
		wants1 = append(wants1, s.createSignature(signer1, contract, index))
	}

	gots0, _ := s.db.FindLatestsBySigner(signer0.Address, 2, 0)
	s.Len(gots0, 2)
	s.Equal(wants0[len(wants0)-1].ID, gots0[0].ID)
	s.Equal(wants0[len(wants0)-2].ID, gots0[1].ID)

	gots1, _ := s.db.FindLatestsBySigner(signer1.Address, 2, 0)
	s.Len(gots1, 2)
	s.Equal(wants1[len(wants1)-1].ID, gots1[0].ID)
	s.Equal(wants1[len(wants1)-2].ID, gots1[1].ID)

	gots0, _ = s.db.FindLatestsBySigner(signer0.Address, 10, 2)
	s.Len(gots0, 3)
	s.Equal(wants0[len(wants0)-3].ID, gots0[0].ID)
	s.Equal(wants0[len(wants0)-4].ID, gots0[1].ID)
	s.Equal(wants0[len(wants0)-5].ID, gots0[2].ID)

	gots1, _ = s.db.FindLatestsBySigner(signer1.Address, 10, 2)
	s.Len(gots1, 8)
	s.Equal(wants1[len(wants1)-3].ID, gots1[0].ID)
	s.Equal(wants1[len(wants1)-4].ID, gots1[1].ID)
	s.Equal(wants1[len(wants1)-5].ID, gots1[2].ID)
	s.Equal(wants1[len(wants1)-6].ID, gots1[3].ID)
	s.Equal(wants1[len(wants1)-7].ID, gots1[4].ID)
	s.Equal(wants1[len(wants1)-8].ID, gots1[5].ID)
	s.Equal(wants1[len(wants1)-9].ID, gots1[6].ID)
	s.Equal(wants1[len(wants1)-10].ID, gots1[7].ID)
}

func (s *OptimismSignatureDBTestSuite) TestSave() {
	signer := s.createSigner()
	contract := s.createContract()
	rollupIndex := uint64(1)
	rollupHash := s.RandHash()
	approved := true
	signature := RandSignature()

	assert := func(got *OptimismSignature) {
		ulid.MustParse(got.ID)
		s.Equal(*signer, got.Signer)
		s.Equal(*contract, got.Contract)
		s.Equal(uint64(1), got.RollupIndex)
		s.Equal(rollupHash, got.RollupHash)
		s.Equal(approved, got.Approved)
		s.Equal(signature, got.Signature)
	}

	got0, _ := s.db.Save(nil, nil,
		signer.Address, contract.Address, rollupIndex, rollupHash,
		approved, signature)
	assert(got0)
	s.Equal("", got0.PreviousID)

	// overwrite test
	rollupHash = s.RandHash()
	approved = false
	signature = RandSignature()

	got0, _ = s.db.Save(nil, nil, signer.Address, contract.Address,
		rollupIndex, rollupHash, approved, signature)
	assert(got0)
	s.Equal("", got0.PreviousID)

	// add new
	got1, _ := s.db.Save(nil, nil, signer.Address, contract.Address,
		rollupIndex+1, rollupHash, approved, signature)
	s.Equal(got0.ID, got1.PreviousID)

	// other signer
	got2, _ := s.db.Save(nil, nil, s.RandAddress(), contract.Address,
		rollupIndex+2, rollupHash, approved, signature)
	s.Equal("", got2.PreviousID)

	// overtaking test
	id1, id2 := util.ULID(nil).String(), util.ULID(nil).String()
	_, err := s.db.Save(&id1, &id2, s.RandAddress(), contract.Address,
		rollupIndex+3, rollupHash, approved, signature)
	s.ErrorContains(err, "previous id is overtaking")
}

func (s *OptimismSignatureDBTestSuite) TestDeletes() {
	assert := func(signer *Signer, contract *OptimismContract, want []int) {
		var gots []int
		s.db.rawdb.Model(&OptimismSignature{}).
			Where("signer_id = ? AND optimism_scc_id = ?", signer.ID, contract.ID).
			Order("batch_index").
			Pluck("batch_index", &gots)
		s.Equal(want, gots)
	}

	signer0 := s.createSigner()
	signer1 := s.createSigner()
	contract0 := s.createContract()
	contract1 := s.createContract()
	for _, i := range s.Shuffle(s.Range(0, 10)) {
		s.createSignature(signer0, contract0, i)
		s.createSignature(signer1, contract1, i)
	}

	assert(signer0, contract0, s.Range(0, 10))
	assert(signer1, contract1, s.Range(0, 10))

	rows0, _ := s.db.Deletes(signer0.Address, contract0.Address, 3)
	rows1, _ := s.db.Deletes(signer1.Address, contract1.Address, 6)

	s.Equal(int64(7), rows0)
	s.Equal(int64(4), rows1)
	assert(signer0, contract0, s.Range(0, 3))
	assert(signer1, contract1, s.Range(0, 6))
}

func (s *OptimismSignatureDBTestSuite) TestSequentialFinder() {
	signer := s.createSigner()
	contract := s.createContract()

	var sigtree [5][]*OptimismSignature
	for _, depth := range s.Range(0, len(sigtree)) {
		sigtree[depth] = []*OptimismSignature{}

		var prevID string
		for i := 0; i <= depth; i++ {
			if depth > 0 && i < len(sigtree[depth-1]) {
				prevID = sigtree[depth-1][i].ID
			}

			sig := s.createSignature(signer, contract, (depth*10)+i)
			sig.PreviousID = prevID
			s.NoDBError(s.db.rawdb.Save(sig))

			sigtree[depth] = append(sigtree[depth], sig)

			// one signature at last depth
			if depth == len(sigtree)-1 {
				break
			}
		}
	}

	assert := func(gots, wants []*OptimismSignature) {
		s.Len(gots, len(wants))
		for i, want := range wants {
			s.Equal(want.ID, gots[i].ID, i)
			s.Equal(want.PreviousID, gots[i].PreviousID, i)
		}
	}

	finder := s.db.SequentialFinder("")
	gots0, _ := finder()
	gots1, _ := finder()
	gots2, _ := finder()
	gots3, _ := finder()
	gots4, _ := finder()
	gots5, _ := finder()

	assert(gots0, sigtree[0])
	assert(gots1, sigtree[1])
	assert(gots2, sigtree[2])
	assert(gots3, sigtree[3])
	assert(gots4, sigtree[4])
	assert(gots5, []*OptimismSignature{})
}

func (s *OptimismSignatureDBTestSuite) TestSort() {
	sigs := []*OptimismSignature{
		{Signer: Signer{Address: common.HexToAddress("0x2")}},
		{Signer: Signer{Address: common.HexToAddress("0x1")}},
		{Signer: Signer{Address: common.HexToAddress("0x0")}},
	}
	sort.Sort(OptimismSignatures(sigs))
	s.Equal(common.HexToAddress("0x0"), sigs[0].Signer.Address)
	s.Equal(common.HexToAddress("0x1"), sigs[1].Signer.Address)
	s.Equal(common.HexToAddress("0x2"), sigs[2].Signer.Address)
}

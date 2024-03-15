package submitter

import (
	"context"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
)

func (s *SubmitterTestSuite) TestSignatureIterator() {
	var signers [20]common.Address
	for i := range signers {
		signers[i] = s.RandAddress()
	}

	type signerGroup struct {
		signers []common.Address
		stakes  []int64
	}
	signerGroups := [4]*signerGroup{
		{signers: signers[0:5], stakes: []int64{1, 1, 1, 1, 1}},      // amount = 5
		{signers: signers[5:10], stakes: []int64{3, 3, 3, 3, 3}},     // amount = 15
		{signers: signers[15:20], stakes: []int64{10, 6, 20, 4, 30}}, // amount = 70
		{signers: signers[10:15], stakes: []int64{2, 2, 2, 2, 2}},    // amount = 10
	}

	type signatureGroup struct {
		sg         *signerGroup
		rollupHash common.Hash
		approved   bool
		signatures []*database.OptimismSignature
	}
	hash0, hash1 := s.RandHash(), s.RandHash()
	hash2, hash3 := s.RandHash(), s.RandHash()
	sigGroups := [3][]*signatureGroup{
		// rollupIndex = 0
		{
			{sg: signerGroups[1], rollupHash: hash0, approved: true},
			{sg: signerGroups[2], rollupHash: hash0, approved: false}, // want
			{sg: signerGroups[3], rollupHash: hash1, approved: true},
			{sg: signerGroups[0], rollupHash: hash1, approved: false},
		},
		// rollupIndex = 1
		{
			{sg: signerGroups[3], rollupHash: hash2, approved: true},
			{sg: signerGroups[0], rollupHash: hash2, approved: false},
			{sg: signerGroups[1], rollupHash: hash3, approved: true},
			{sg: signerGroups[2], rollupHash: hash3, approved: false}, // want
		},
		// rollupIndex = 2 (should return `*StakeAmountShortage`)
		{
			{sg: signerGroups[0], rollupHash: s.RandHash(), approved: true},
		},
	}

	// setup stakemanager
	sm := &testhelper.StakeManagerMock{}
	smcache := stakemanager.NewCache(sm)
	for _, group := range signerGroups {
		for i, signer := range group.signers {
			sm.Owners = append(sm.Owners, s.RandAddress())
			sm.Operators = append(sm.Operators, signer)
			sm.Stakes = append(sm.Stakes,
				new(big.Int).Mul(ethutil.TenMillionOAS, big.NewInt(group.stakes[i])))
			sm.Candidates = append(sm.Candidates, true)
		}
	}
	smcache.Refresh(context.Background())

	// save signatures
	for rollupIndex, c := range sigGroups {
		for _, group := range c {
			for _, signer := range group.sg.signers {
				sig, _ := s.DB.OPSignature.Save(
					nil, nil,
					signer,
					s.SCCAddr,
					uint64(rollupIndex),
					group.rollupHash,
					group.approved,
					database.RandSignature(),
				)
				group.signatures = append(group.signatures, sig)
			}
		}
	}

	// signers with stakes above the minimum
	highStakes := map[common.Address]bool{
		signerGroups[2].signers[4]: true, // amount = 30
		signerGroups[2].signers[2]: true, // amount = 20
		signerGroups[2].signers[0]: true, // amount = 10
	}
	// signatures of high stakes signers
	want0 := []*database.OptimismSignature{
		sigGroups[0][1].signatures[4], // amount = 30
		sigGroups[0][1].signatures[2], // amount = 20
		sigGroups[0][1].signatures[0], // amount = 10
	}
	want1 := []*database.OptimismSignature{
		sigGroups[1][3].signatures[4], // amount = 30
		sigGroups[1][3].signatures[2], // amount = 20
		sigGroups[1][3].signatures[0], // amount = 10
	}

	sort.Sort(database.OptimismSignatures(want0))
	sort.Sort(database.OptimismSignatures(want1))

	iter := &signatureIterator{
		db:           s.DB,
		stakemanager: smcache,
		contract:     s.SCCAddr,
		rollupIndex:  0,
	}

	// assert
	gots0, err0 := iter.next()
	gots1, err1 := iter.next()

	s.Nil(err0)
	s.Nil(err1)

	s.Len(gots0, len(highStakes))
	s.Len(gots1, len(highStakes))

	for i, want := range want0 {
		s.Equal(want.Signature, gots0[i].Signature)
		s.True(highStakes[gots0[i].Signer.Address])
	}
	for i, want := range want1 {
		s.Equal(want.Signature, gots1[i].Signature)
		s.True(highStakes[gots1[i].Signer.Address])
	}

	// should return `*StakeAmountShortage`
	for i := range sm.Operators {
		sm.Stakes[i] = ethutil.TenMillionOAS
	}
	smcache.Refresh(context.Background())
	_, err := iter.next()
	s.ErrorContains(err, "stake amount shortage")
}

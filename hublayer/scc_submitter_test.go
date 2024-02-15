package hublayer

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/stretchr/testify/suite"
)

type SccSubmitterTestSuite struct {
	SccTestSuite
}

func TestSccSubmitter(t *testing.T) {
	suite.Run(t, new(SccSubmitterTestSuite))
}

func (s *SccSubmitterTestSuite) TestWork() {
	var (
		indexes = 5

		signers    [20]common.Address
		signatures [][]*database.OptimismSignature
	)

	for i := range s.Range(0, len(signers)) {
		signers[i] = s.RandAddress()

		s.sm.Owners = append(s.sm.Owners, s.RandAddress())
		s.sm.Operators = append(s.sm.Operators, signers[i])
		s.sm.Stakes = append(
			s.sm.Stakes,
			new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10_000_000)),
		)
		s.sm.Candidates = append(s.sm.Candidates, true)
		s.sm.NewCursor = big.NewInt(0)
	}

	for i := range s.Range(0, indexes) {
		signatures = append(signatures, make([]*database.OptimismSignature, len(signers)))

		batchIndex := uint64(i)
		batchRoot := s.RandHash()
		batchSize := uint64(i)
		prevTotalElements := uint64(i + 1)
		extraData := []byte(fmt.Sprintf("%d", i))
		approved := i < indexes-1

		// create sample signatures
		for j := range s.Range(0, len(signers)) {
			sig, _ := s.db.Optimism.SaveSignature(
				nil, nil,
				signers[j],
				s.sccAddr,
				batchIndex,
				batchRoot,
				batchSize,
				prevTotalElements,
				extraData,
				approved,
				database.RandSignature(),
			)

			signatures[i][j] = sig
		}
		sort.Slice(signatures[i], func(x, y int) bool {
			a := signatures[i][x].Signer.Address.Hash().Big()
			b := signatures[i][y].Signer.Address.Hash().Big()
			return a.Cmp(b) == -1
		})

		// emit StateBatchAppended event to the test contract
		s.scc.EmitStateBatchAppended(
			s.hub.TransactOpts(context.Background()),
			new(big.Int).SetUint64(batchIndex),
			batchRoot,
			new(big.Int).SetUint64(batchSize),
			new(big.Int).SetUint64(prevTotalElements),
			extraData)
		s.mining()
	}

	s.sccSubmitter.stakemanager.Refresh(context.Background())

	for range s.Range(0, indexes/s.sccSubmitter.cfg.BatchSize+1) {
		go func() {
			time.Sleep(10 * time.Millisecond)
			s.hub.Commit()
		}()
		s.sccSubmitter.work(context.Background(), &submitTask{scc: s.sccAddr, hub: s.hub})
	}

	for i := range s.Range(0, indexes) {
		got, _ := s.sccv.AssertLogs(
			&bind.CallOpts{Context: context.Background()},
			big.NewInt(int64(i)),
		)

		s.Equal(s.sccAddr, got.StateCommitmentChain)

		s.Equal(uint64(i), got.BatchHeader.BatchIndex.Uint64())
		s.Equal(signatures[i][0].BatchRoot[:], got.BatchHeader.BatchRoot[:])
		s.Equal(uint64(i), got.BatchHeader.BatchSize.Uint64())
		s.Equal(uint64(i+1), got.BatchHeader.PrevTotalElements.Uint64())
		s.Equal([]byte(fmt.Sprintf("%d", i)), got.BatchHeader.ExtraData)

		s.Len(got.Signatures, len(signers)*65)
		for j, sig := range signatures[i] {
			start := j * 65
			end := start + 65
			s.Equal(sig.Signature[:], got.Signatures[start:end])
		}

		s.Equal(i < indexes-1, got.Approve)
	}
}

func (s *SccSubmitterTestSuite) TestFindSignatures() {
	type group struct {
		signers    []common.Address
		signatures []*database.OptimismSignature
		batchRoot  common.Hash
		approved   bool
		stake      int64
	}

	var (
		batchRoot0 = s.RandHash()
		batchRoot1 = s.RandHash()
		batchIndex = uint64(10)

		totalStake   = big.NewInt(0)
		signerStakes = map[common.Address]*big.Int{}

		groups = []*group{
			{batchRoot: batchRoot0, approved: true, stake: 1000},
			{batchRoot: batchRoot0, approved: false, stake: 2000},
			{batchRoot: batchRoot1, approved: true, stake: 10000}, // want
			{batchRoot: batchRoot1, approved: false, stake: 3000},
		}
		want = groups[2]
	)

	for _, group := range groups {
		totalStake.Add(totalStake, big.NewInt(group.stake))

		for i := range s.Range(0, 10) {
			group.signers = append(group.signers, s.RandAddress())
			signerStakes[group.signers[i]] = big.NewInt(group.stake / int64(10))

			// create sample signatures
			sig, _ := s.db.Optimism.SaveSignature(
				nil, nil,
				group.signers[i],
				s.sccAddr,
				batchIndex,
				group.batchRoot,
				0,
				0,
				[]byte(nil),
				group.approved,
				database.RandSignature(),
			)
			group.signatures = append(group.signatures, sig)
		}
	}

	sort.Slice(want.signatures, func(x, y int) bool {
		return bytes.Compare(
			want.signatures[x].Signer.Address[:],
			want.signatures[y].Signer.Address[:],
		) == -1
	})

	// assert
	gots, _ := s.sccSubmitter.findSignatures(
		s.sccAddr, batchIndex, common.Big0, totalStake, signerStakes)
	s.Len(gots, len(want.signatures))
	for i, want := range want.signatures {
		s.Equal(want.Signature, gots[i].Signature)
	}

	rows, err := s.sccSubmitter.findSignatures(
		s.sccAddr, batchIndex+1, common.Big0, totalStake, signerStakes)
	s.Len(rows, 0)
	s.NoError(err)

	// stake amount shortage
	rows, err = s.sccSubmitter.findSignatures(
		s.sccAddr, batchIndex, common.Big0, big.NewInt(20000), signerStakes)
	s.Len(rows, 0)
	s.NoError(err)
}

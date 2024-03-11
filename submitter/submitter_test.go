package submitter

import (
	"context"
	"math/big"
	"sort"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	tscc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
	"github.com/stretchr/testify/suite"
)

type SubmitterTestSuite struct {
	backend.BackendSuite

	submitter *Submitter
	task      verse.TransactableVerse
}

func TestSubmitter(t *testing.T) {
	suite.Run(t, new(SubmitterTestSuite))
}

func (s *SubmitterTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()

	// Setup `StakeManager` contract
	for i := range s.Range(0, 10) {
		s.StakeManager.Owners = append(s.StakeManager.Owners, s.RandAddress())
		s.StakeManager.Operators = append(s.StakeManager.Operators, s.RandAddress())
		s.StakeManager.Stakes = append(s.StakeManager.Stakes,
			new(big.Int).Add(ethutil.TenMillionOAS, big.NewInt(int64(10-i))))
		s.StakeManager.Candidates = append(s.StakeManager.Candidates, true)
	}

	// Setup submitter
	s.submitter = NewSubmitter(&config.Submitter{
		Interval:          0,
		Concurrency:       0,
		Confirmations:     0,
		GasMultiplier:     1.0,
		BatchSize:         20,
		MaxGas:            500_000_000,
		UseMulticall:      true, // TODO
		Multicall2Address: s.MulticallAddr.String(),
	}, s.DB, stakemanager.NewCache(s.StakeManager))

	s.task = verse.
		NewOPLegacy(s.DB, s.Hub, s.SCCAddr).
		WithTransactable(s.SignableHub, s.SCCVAddr)
}

func (s *SubmitterTestSuite) TestSubmit() {
	ctx := context.Background()
	batchIndexes := s.Range(0, 5)
	nextIndex := 2
	signers := s.StakeManager.Operators

	// save dummy signatures
	events := make([]*tscc.SccStateBatchAppended, len(batchIndexes))
	signatures := make([][]*database.OptimismSignature, len(batchIndexes))
	for i := range batchIndexes {
		_, events[i] = s.EmitStateBatchAppended(i)
		signatures[i] = make([]*database.OptimismSignature, len(signers))

		for j := range s.Range(0, len(signers)) {
			signatures[i][j], _ = s.DB.OPSignature.Save(
				nil, nil,
				signers[j],
				s.SCCAddr,
				events[i].BatchIndex.Uint64(),
				events[i].BatchRoot,
				i < len(batchIndexes)-1,
				database.RandSignature(),
			)
		}

		// no more signatures than the minimum stake should be sent
		sort.Slice(signatures[i], func(j, h int) bool {
			// sort by stake amount
			a := s.submitter.stakemanager.StakeBySigner(signatures[i][j].Signer.Address)
			b := s.submitter.stakemanager.StakeBySigner(signatures[i][h].Signer.Address)
			return a.Cmp(b) == 1 // order by desc
		})
		signatures[i] = signatures[i][:6]
		sort.Sort(database.OptimismSignatures(signatures[i]))
	}

	// set the `SCC.nextIndex`
	s.TSCC.SetNextIndex(s.SignableHub.TransactOpts(ctx), big.NewInt(int64(nextIndex)))
	s.Hub.Commit()

	// submitter do the work.
	s.submitter.stakemanager.Refresh(ctx)
	go s.submitter.work(ctx, s.task)
	time.Sleep(time.Second / 10)
	s.Hub.Commit()

	// assert multicall transaction
	mcallTx := s.Hub.Blockchain().CurrentBlock().Transactions()[0]
	sender, _ := s.Hub.TxSender(mcallTx)
	s.Equal(s.task.L1Signer().Signer(), sender)
	s.Equal(s.MulticallAddr, *mcallTx.To())

	mcallReceipt, _ := s.Hub.TransactionReceipt(context.Background(), mcallTx.Hash())
	s.Len(mcallReceipt.Logs, 6)
	s.Equal(s.SCCAddr, mcallReceipt.Logs[0].Address)
	s.Equal(s.SCCVAddr, mcallReceipt.Logs[1].Address)
	s.Equal(s.SCCAddr, mcallReceipt.Logs[2].Address)
	s.Equal(s.SCCVAddr, mcallReceipt.Logs[3].Address)
	s.Equal(s.SCCAddr, mcallReceipt.Logs[4].Address)
	s.Equal(s.SCCVAddr, mcallReceipt.Logs[5].Address)

	// assert call parameters
	length, _ := s.TSCCV.SccAssertLogsLen(&bind.CallOpts{Context: ctx})
	s.Equal(uint64(3), length.Uint64())

	for i := range batchIndexes {
		if i < nextIndex {
			got, err := s.TSCCV.AssertLogs(
				&bind.CallOpts{Context: ctx},
				big.NewInt(int64(i+nextIndex+1)))
			s.ErrorContains(err, "execution reverted")
			s.Equal(common.Address{}, got.StateCommitmentChain)
		} else {
			got, err := s.TSCCV.AssertLogs(
				&bind.CallOpts{Context: ctx},
				big.NewInt(int64(i-nextIndex)))
			s.NoError(err)
			s.Equal(s.SCCAddr, got.StateCommitmentChain)
			s.Equal(events[i].BatchIndex.Uint64(), got.BatchHeader.BatchIndex.Uint64())
			s.Equal(events[i].BatchRoot, got.BatchHeader.BatchRoot)
			s.Equal(events[i].BatchSize.Uint64(), got.BatchHeader.BatchSize.Uint64())
			s.Equal(events[i].PrevTotalElements.Uint64(), got.BatchHeader.PrevTotalElements.Uint64())
			s.Equal(events[i].ExtraData, got.BatchHeader.ExtraData)
			s.Equal(i < len(batchIndexes)-1, got.Approve)

			// no more signatures than the minimum stake should be sent
			s.Len(got.Signatures, len(signatures[i])*65)
			for j, sig := range signatures[i] {
				start := j * 65
				end := start + 65
				s.Equal(sig.Signature[:], got.Signatures[start:end])
			}
		}
	}
}

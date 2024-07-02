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
		Interval:         100 * time.Millisecond,
		Concurrency:      0,
		Confirmations:    2,
		GasMultiplier:    1.0,
		BatchSize:        20,
		MaxGas:           500_000_000,
		UseMulticall:     true, // TODO
		MulticallAddress: s.MulticallAddr.String(),
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

	// Confirm blocks
	for i := 0; i < s.submitter.cfg.Confirmations; i++ {
		s.Hub.Mining()
	}

	// submitter do the work.
	s.submitter.stakemanager.Refresh(ctx)
	go s.submitter.work(ctx, s.task, nil)
	time.Sleep(time.Second / 10)
	s.Hub.Commit()

	// assert multicall transaction
	currBlock, _ := s.Hub.Client().BlockByNumber(ctx, nil)
	mcallTx := currBlock.Transactions()[0]
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

func (s *SubmitterTestSuite) TestStartSubmit() {
	ctx, cancel := context.WithCancel(context.Background())
	batchIndexes := s.Range(0, 5)
	nextIndex := 2
	signers := s.StakeManager.Operators

	// Start submitter
	s.submitter.stakemanager.Refresh(ctx)
	go s.submitter.startSubmitter(ctx, s.task)
	// Dry run to cover no signature case
	// Manually confirmed by checking the logs
	time.Sleep(s.submitter.cfg.Interval)

	// Confirm the `stake amount shortage` case is covered
	// Manually confirmed by checking the logs
	_, event := s.EmitStateBatchAppended(0)
	s.DB.OPSignature.Save(
		nil, nil,
		signers[0],
		s.SCCAddr,
		event.BatchIndex.Uint64(),
		event.BatchRoot,
		true,
		database.RandSignature(),
	)
	// wait for the submitter to work
	time.Sleep(s.submitter.cfg.Interval * 2)

	// Confirm succcesfully tx submission case
	// set the `SCC.nextIndex`
	s.TSCC.SetNextIndex(s.SignableHub.TransactOpts(ctx), big.NewInt(int64(nextIndex)))
	s.Hub.Commit()
	// Confirm blocks
	for i := 0; i < s.submitter.cfg.Confirmations; i++ {
		s.Hub.Mining()
	}

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

	// submitter do the work.
	time.Sleep(s.submitter.cfg.Interval * 2)
	s.Hub.Commit()

	// assert multicall transaction
	currBlock, _ := s.Hub.Client().BlockByNumber(ctx, nil)
	mcallTx := currBlock.Transactions()[0]
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

	// Cancel will exit receipt waiting loop
	cancel()

	// Confirm old signatures are cleaned up
	time.Sleep(s.submitter.cfg.Interval)
	deleteIndex := uint64(1)
	rows, err := s.DB.OPSignature.Find(nil, nil, &s.SCCAddr, &deleteIndex, 1000, 0)
	s.NoError(err)
	s.True(len(rows) == 0)
}

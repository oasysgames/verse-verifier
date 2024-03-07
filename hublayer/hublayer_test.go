package hublayer

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/multicall2"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	tscc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/scc"
	tsccv "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/sccverifier"
)

type SccTestSuite struct {
	testhelper.Suite

	db    *database.Database
	hub   *backend.SignableBackend
	verse *backend.Backend

	sm *testhelper.StakeManagerMock

	mcall2     *multicall2.Multicall2
	mcall2Addr common.Address

	scc     *tscc.Scc
	sccAddr common.Address

	sccv     *tsccv.Sccverifier
	sccvAddr common.Address

	stateCollector *EventCollector
	sccSubmitter   *SccSubmitter
}

func (s *SccTestSuite) SetupTest() {
	ctx := context.Background()
	s.db, _ = database.NewDatabase(&config.Database{Path: ":memory:"})

	// setup test chain
	s.hub = backend.NewSignableBackend(nil, nil, nil)
	s.verse = backend.NewBackend(nil, 0)
	s.sm = &testhelper.StakeManagerMock{}

	// deploy `Multicall2` contract
	s.mcall2Addr, _, s.mcall2, _ = multicall2.DeployMulticall2(s.hub.TransactOpts(ctx), s.hub)
	s.hub.Mining()

	// deploy `StateCommitmentChain` contract
	s.sccAddr, _, s.scc, _ = tscc.DeployScc(s.hub.TransactOpts(ctx), s.hub)
	s.hub.Mining()

	// deploy `OasysStateCommitmentChainVerifier` contract
	s.sccvAddr, _, s.sccv, _ = tsccv.DeploySccverifier(s.hub.TransactOpts(ctx), s.hub)
	s.hub.Mining()

	// setup workers
	hubSigner := s.hub.Signer()
	s.stateCollector = NewEventCollector(&config.Verifier{
		Interval:         time.Millisecond,
		EventFilterLimit: 1000,
	}, s.db, s.hub, hubSigner)

	s.sccSubmitter = NewSccSubmitter(&config.Submitter{
		Interval:          0,
		Concurrency:       0,
		Confirmations:     0,
		GasMultiplier:     1.0,
		BatchSize:         2,
		MaxGas:            math.MaxInt,
		VerifierAddress:   s.sccvAddr.String(),
		Multicall2Address: s.mcall2Addr.String(),
	}, s.db, stakemanager.NewCache(s.sm))
	s.sccSubmitter.AddVerse(s.sccAddr, s.hub)
}

func (s *SccTestSuite) mining() {
	s.hub.Commit()
	header, _ := s.hub.HeaderByNumber(context.Background(), nil)
	s.db.Block.Save(header.Number.Uint64(), header.Hash())
}

func (s *SccTestSuite) emitStateBatchAppendedEvent(index int) *tscc.SccStateBatchAppended {
	i64 := int64(index)
	event := &tscc.SccStateBatchAppended{
		BatchIndex:        big.NewInt(i64),
		BatchRoot:         [32]byte(common.BigToHash(big.NewInt(i64))),
		BatchSize:         big.NewInt(10),
		PrevTotalElements: big.NewInt(i64 * 10),
		ExtraData:         []byte("extra data"),
	}
	s.scc.EmitStateBatchAppended(
		s.hub.TransactOpts(context.Background()), event.BatchIndex,
		event.BatchRoot, event.BatchSize, event.PrevTotalElements, event.ExtraData)
	s.mining()
	return event
}

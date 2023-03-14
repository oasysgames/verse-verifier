package hublayer

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	tmcall2 "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/multicall2"
	tscc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/scc"
	tsccv "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/sccverifier"
)

type SccTestSuite struct {
	testhelper.Suite

	db    *database.Database
	hub   *testhelper.TestBackend
	verse *testhelper.TestBackend

	sm *stakeManagerMock

	mcall2     *tmcall2.Multicall2
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
	s.hub = testhelper.NewTestBackend()
	s.verse = testhelper.NewTestBackend()
	s.sm = &stakeManagerMock{}

	// deploy `Multicall2` contract
	s.mcall2Addr, _, s.mcall2, _ = tmcall2.DeployMulticall2(s.hub.TransactOpts(ctx), s.hub)
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
	}, s.db, s.sm)
	s.sccSubmitter.AddVerse(s.sccAddr, s.hub)
}

func (s *SccTestSuite) mining() {
	s.hub.Commit()
	header, _ := s.hub.HeaderByNumber(context.Background(), nil)
	s.db.Block.SaveNewBlock(header.Number.Uint64(), header.Hash())
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

type stakeManagerMock struct {
	Owners     []common.Address
	Operators  []common.Address
	Stakes     []*big.Int
	Candidates []bool
	NewCursor  *big.Int
}

func (s *stakeManagerMock) GetTotalStake(
	callOpts *bind.CallOpts,
	epoch *big.Int,
) (*big.Int, error) {
	tot := new(big.Int)
	for _, stake := range s.Stakes {
		tot.Add(tot, stake)
	}
	return tot, nil
}

func (s *stakeManagerMock) GetValidators(
	callOpts *bind.CallOpts,
	epoch, cursol, howMany *big.Int,
) (struct {
	Owners     []common.Address
	Operators  []common.Address
	Stakes     []*big.Int
	Candidates []bool
	NewCursor  *big.Int
}, error) {
	length := big.NewInt(int64(len(s.Owners)))
	if new(big.Int).Add(cursol, howMany).Cmp(length) >= 0 {
		howMany = new(big.Int).Sub(length, cursol)
	}

	start := cursol.Uint64()
	end := start + howMany.Uint64()

	ret := struct {
		Owners     []common.Address
		Operators  []common.Address
		Stakes     []*big.Int
		Candidates []bool
		NewCursor  *big.Int
	}{
		Owners:     s.Owners[start:end],
		Operators:  s.Operators[start:end],
		Stakes:     s.Stakes[start:end],
		Candidates: s.Candidates[start:end],
		NewCursor:  new(big.Int).Add(cursol, howMany),
	}

	return ret, nil
}

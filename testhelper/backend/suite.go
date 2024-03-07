package backend

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/multicall2"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/scc"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/sccverifier"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	tscc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/scc"
	tsccv "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/sccverifier"
)

type BackendSuite struct {
	testhelper.Suite

	DB *database.Database

	Hub, Verse *Backend
	SignableHub,
	SignableVerse *SignableBackend

	StakeManager *testhelper.StakeManagerMock

	Multicall     *multicall2.Multicall2
	MulticallAddr common.Address

	SCC     *scc.Scc
	TSCC    *tscc.Scc
	SCCAddr common.Address

	SCCV     *sccverifier.Sccverifier
	TSCCV    *tsccv.Sccverifier
	SCCVAddr common.Address
}

func (b *BackendSuite) SetupTest() {
	ctx := context.Background()
	b.DB, _ = database.NewDatabase(&config.Database{Path: ":memory:"})
	b.StakeManager = &testhelper.StakeManagerMock{}

	// setup test chain
	b.Hub = NewBackend(nil, 0)
	b.Verse = NewBackend(nil, 0)
	b.SignableHub = NewSignableBackend(b.Hub, nil, nil)
	b.SignableVerse = NewSignableBackend(b.Verse, nil, nil)

	// deploy `Multicall2` contract
	b.MulticallAddr, _, b.Multicall, _ = multicall2.DeployMulticall2(b.SignableHub.TransactOpts(ctx), b.SignableHub)
	b.SignableHub.Mining()

	// deploy `StateCommitmentChain` contract
	b.SCCAddr, _, b.TSCC, _ = tscc.DeployScc(b.SignableHub.TransactOpts(ctx), b.SignableHub)
	b.SCC, _ = scc.NewScc(b.SCCAddr, b.SignableHub)
	b.SignableHub.Mining()

	// deploy `OasysRollupVerifier` contract
	b.SCCVAddr, _, b.TSCCV, _ = tsccv.DeploySccverifier(b.SignableHub.TransactOpts(ctx), b.SignableHub)
	b.SCCV, _ = sccverifier.NewSccverifier(b.SCCVAddr, b.SignableHub)
	b.SignableHub.Mining()
}

func (b *BackendSuite) Mining() {
	b.NotEmpty(b.SignableHub.Commit())
	header, err := b.SignableHub.HeaderByNumber(context.Background(), nil)
	b.NoError(err)
	b.DB.Block.SaveNewBlock(header.Number.Uint64(), header.Hash())
}

func (b *BackendSuite) EmitStateBatchAppended(index int) (
	*types.Transaction,
	*tscc.SccStateBatchAppended,
) {
	i64 := int64(index)
	event := &tscc.SccStateBatchAppended{
		BatchIndex:        big.NewInt(i64),
		BatchRoot:         [32]byte(common.BigToHash(big.NewInt(i64))),
		BatchSize:         big.NewInt(10),
		PrevTotalElements: big.NewInt(i64 * 10),
		ExtraData:         []byte(fmt.Sprintf("ExtraData-%d", index)),
	}
	tx, err := b.TSCC.EmitStateBatchAppended(
		b.SignableHub.TransactOpts(context.Background()), event.BatchIndex,
		event.BatchRoot, event.BatchSize, event.PrevTotalElements, event.ExtraData)
	b.NoError(err)
	b.Mining()
	return tx, event
}

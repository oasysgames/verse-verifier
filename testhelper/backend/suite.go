package backend

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2ooverifier"
	"github.com/oasysgames/oasys-optimism-verifier/contract/multicall2"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/contract/sccverifier"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	tl2oo "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/l2oo"
	tl2oov "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/l2ooverifier"
	tscc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/scc"
	tsccv "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/sccverifier"
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

	L2OO     *l2oo.OasysL2OutputOracle
	TL2OO    *tl2oo.L2oo
	L2OOAddr common.Address

	SCCV      *sccverifier.Sccverifier
	L2OOV     *l2ooverifier.OasysL2OutputOracleVerifier
	TSCCV     *tsccv.Sccverifier
	TL2OOV    *tl2oov.L2ooverifier
	SCCVAddr  common.Address
	L2OOVAddr common.Address
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

	// deploy `OasysL2OutputOracle` contract
	b.L2OOAddr, _, b.TL2OO, _ = tl2oo.DeployL2oo(b.SignableHub.TransactOpts(ctx), b.SignableHub)
	b.L2OO, _ = l2oo.NewOasysL2OutputOracle(b.L2OOAddr, b.SignableHub)
	b.SignableHub.Mining()

	// deploy `OasysRollupVerifier` contract
	b.SCCVAddr, _, b.TSCCV, _ = tsccv.DeploySccverifier(b.SignableHub.TransactOpts(ctx), b.SignableHub)
	b.SCCV, _ = sccverifier.NewSccverifier(b.SCCVAddr, b.SignableHub)
	b.SignableHub.Mining()

	// deploy `OasysL2OutputOracleVerifier` contract
	b.L2OOVAddr, _, b.TL2OOV, _ = tl2oov.DeployL2ooverifier(b.SignableHub.TransactOpts(ctx), b.SignableHub)
	b.L2OOV, _ = l2ooverifier.NewOasysL2OutputOracleVerifier(b.L2OOVAddr, b.SignableHub)
	b.SignableHub.Mining()
}

func (b *BackendSuite) Mining() {
	b.NotEmpty(b.SignableHub.Commit())
	header, err := b.SignableHub.HeaderByNumber(context.Background(), nil)
	b.NoError(err)
	b.DB.Block.Save(header.Number.Uint64(), header.Hash())
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

func (b *BackendSuite) EmitOutputProposed(index int) (
	*types.Transaction,
	*tl2oo.L2ooOutputProposed,
) {
	event := &tl2oo.L2ooOutputProposed{
		OutputRoot:    b.RandHash(),
		L2OutputIndex: big.NewInt(int64(index)),
		L2BlockNumber: big.NewInt(int64((index + 1) * 10)),
		L1Timestamp:   big.NewInt(time.Now().Unix()),
	}
	tx, err := b.TL2OO.EmitOutputProposed(b.SignableHub.TransactOpts(context.Background()),
		event.OutputRoot, event.L2OutputIndex, event.L2BlockNumber, event.L1Timestamp)
	b.NoError(err)
	b.Mining()
	return tx, event
}

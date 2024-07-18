package backend

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/account"
)

var (
	ether = big.NewInt(params.Ether)
)

var _ ethutil.Client = &Backend{} // type checking

func NewBackend(alloc types.GenesisAlloc, gasLimit uint64) *Backend {
	if alloc == nil {
		alloc = types.GenesisAlloc{
			account.DefaultAccount.Address: {Balance: new(big.Int).Mul(ether, ether)},
		}
	}
	if gasLimit == 0 {
		gasLimit = 500_000_000
	}
	sim := simulated.NewBackend(alloc, simulated.WithBlockGasLimit(gasLimit))
	return &Backend{sim}
}

type Backend struct {
	*simulated.Backend
}

func (b *Backend) FilterLogsWithRateThottling(ctx context.Context, q ethereum.FilterQuery) (logs []types.Log, err error) {
	return b.FilterLogs(ctx, q)
}

func (b *Backend) NewBatchHeaderClient() (ethutil.BatchHeaderClient, error) {
	return &BatchHeaderClient{b}, nil
}

func (b *Backend) URL() string {
	return "SimulatedBackend"
}

func (b *Backend) BlockNumber(ctx context.Context) (uint64, error) {
	return b.Blockchain().CurrentHeader().Number.Uint64(), nil
}

func (b *Backend) HeaderByNumber(
	ctx context.Context,
	number *big.Int,
) (*types.Header, error) {
	return b.Client().HeaderByNumber(ctx, number)
}

func (b *Backend) HeaderByHash(
	ctx context.Context,
	hash common.Hash,
) (*types.Header, error) {
	return b.Client().HeaderByHash(ctx, hash)
}

func (b *Backend) TransactionByHash(
	ctx context.Context,
	txHash common.Hash,
) (tx *types.Transaction, isPending bool, err error) {
	return b.Client().TransactionByHash(ctx, txHash)
}

func (b *Backend) TransactionReceipt(
	ctx context.Context,
	txHash common.Hash,
) (*types.Receipt, error) {
	return b.Client().TransactionReceipt(ctx, txHash)
}

func (b *Backend) PendingCodeAt(
	ctx context.Context,
	account common.Address,
) ([]byte, error) {
	return b.Client().PendingCodeAt(ctx, account)
}

func (b *Backend) PendingNonceAt(
	ctx context.Context,
	account common.Address,
) (uint64, error) {
	return b.Client().PendingNonceAt(ctx, account)
}

func (b *Backend) EstimateGas(
	ctx context.Context,
	call ethereum.CallMsg,
) (uint64, error) {
	return b.Client().EstimateGas(ctx, call)
}

func (b *Backend) FilterLogs(
	ctx context.Context,
	q ethereum.FilterQuery,
) ([]types.Log, error) {
	return b.Client().FilterLogs(ctx, q)
}

func (b *Backend) SubscribeFilterLogs(
	ctx context.Context,
	q ethereum.FilterQuery,
	ch chan<- types.Log,
) (ethereum.Subscription, error) {
	return b.Client().SubscribeFilterLogs(ctx, q, ch)
}

func (b *Backend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return b.Client().SendTransaction(ctx, tx)
}

func (b *Backend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return b.Client().SuggestGasPrice(ctx)
}

func (b *Backend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return b.Client().SuggestGasTipCap(ctx)
}

func (b *Backend) CodeAt(
	ctx context.Context,
	contract common.Address,
	blockNumber *big.Int,
) ([]byte, error) {
	return b.Client().CodeAt(ctx, contract, blockNumber)
}

func (b *Backend) CallContract(
	ctx context.Context,
	call ethereum.CallMsg,
	blockNumber *big.Int,
) ([]byte, error) {
	return b.Client().CallContract(ctx, call, blockNumber)
}

func (b *Backend) GetProof(
	ctx context.Context,
	address common.Address,
	storageKeys []string,
	blockNumber *big.Int,
) (*gethclient.AccountResult, error) {
	return b.GethClient().GetProof(ctx, address, storageKeys, blockNumber)
}

func (b *Backend) GetHeaderBatch(
	ctx context.Context,
	start uint64,
	size, limit int,
) ([]*types.Header, error) {
	headers := make([]*types.Header, size)
	for i := 0; i < size; i++ {
		h, err := b.Client().HeaderByNumber(ctx, new(big.Int).SetUint64(start+uint64(i)))
		if err != nil {
			return nil, err
		}

		headers[i] = h
	}

	return headers, nil
}

func (b *Backend) TxSender(tx *types.Transaction) (common.Address, error) {
	signer := types.MakeSigner(b.Blockchain().Config(), common.Big0, 0)
	return types.Sender(signer, tx)
}

func (b *Backend) Mining() *types.Header {
	b.Commit()
	return b.Blockchain().CurrentHeader()
}

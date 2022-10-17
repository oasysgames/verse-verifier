package ethutil

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type SignDataFn = func(data []byte) (sig []byte, err error)

type ReadOnlyClient interface {
	bind.ContractBackend
	bind.DeployBackend

	URL() string
	TransactionByHash(
		ctx context.Context,
		hash common.Hash,
	) (tx *types.Transaction, isPending bool, err error)
	GetHeaderBatch(ctx context.Context, start, size int) ([]*types.Header, error)
}

type WritableClient interface {
	ReadOnlyClient

	ChainID() *big.Int
	Signer() common.Address
	SignData(data []byte) (sig []byte, err error)
	SignTx(tx *types.Transaction) (*types.Transaction, error)
	TransactOpts(ctx context.Context) *bind.TransactOpts
}

type readOnlyClient struct {
	ethclient.Client

	url string
	rpc rpc.Client
}

func NewReadOnlyClient(url string) (ReadOnlyClient, error) {
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	return &readOnlyClient{
		Client: *ethclient.NewClient(rpcClient),
		url:    url,
		rpc:    *rpcClient,
	}, nil
}

func (c *readOnlyClient) URL() string {
	return c.url
}

func (c *readOnlyClient) TransactionByHash(
	ctx context.Context,
	hash common.Hash,
) (tx *types.Transaction, isPending bool, err error) {
	return c.Client.TransactionByHash(ctx, hash)
}

func (c *readOnlyClient) GetHeaderBatch(
	ctx context.Context,
	start, size int,
) ([]*types.Header, error) {
	headers := make([]*types.Header, size)
	for i := 0; i < size; i++ {
		h, err := c.HeaderByNumber(ctx, big.NewInt(int64(start+i)))
		if err != nil {
			return nil, err
		}

		headers[i] = h
	}

	return headers, nil
}

type writableClient struct {
	ReadOnlyClient

	chainId *big.Int
	wallet  accounts.Wallet
	signer  *accounts.Account
}

func NewWritableClient(
	chainId *big.Int,
	rpc string,
	wallet accounts.Wallet,
	signer *accounts.Account,
) (WritableClient, error) {
	rc, err := NewReadOnlyClient(rpc)
	if err != nil {
		return nil, err
	}

	return &writableClient{
		ReadOnlyClient: rc,
		chainId:        chainId,
		wallet:         wallet,
		signer:         signer,
	}, nil
}

func (c *writableClient) ChainID() *big.Int {
	return new(big.Int).Set(c.chainId)
}

// Return signer address.
func (c *writableClient) Signer() common.Address {
	return c.signer.Address
}

// Return transaction authorization data.
func (c *writableClient) TransactOpts(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		Context: ctx,
		From:    c.Signer(),
		Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return c.wallet.SignTx(*c.signer, tx, c.ChainID())
		},
	}
}

func (c *writableClient) SignData(data []byte) (sig []byte, err error) {
	_, msg := accounts.TextAndHash(crypto.Keccak256(data))
	return c.wallet.SignData(*c.signer, "", []byte(msg))
}

func (c *writableClient) SignTx(tx *types.Transaction) (*types.Transaction, error) {
	return c.wallet.SignTx(*c.signer, tx, c.ChainID())
}

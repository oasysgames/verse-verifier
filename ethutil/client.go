package ethutil

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
)

type SignDataFn = func(hash []byte) (sig []byte, err error)

type Client interface {
	bind.ContractBackend
	bind.DeployBackend

	URL() string
	BlockNumber(ctx context.Context) (uint64, error)
	TransactionByHash(
		ctx context.Context,
		hash common.Hash,
	) (tx *types.Transaction, isPending bool, err error)
	NewBatchHeaderClient() (BatchHeaderClient, error)
	GetProof(ctx context.Context, account common.Address, keys []string, blockNumber *big.Int) (*gethclient.AccountResult, error)
}

type SignableClient interface {
	Client

	ChainID() *big.Int
	Signer() common.Address
	SignData(data []byte) (sig []byte, err error)
	SignTx(tx *types.Transaction) (*types.Transaction, error)
	TransactOpts(ctx context.Context) *bind.TransactOpts
}

type client struct {
	*ethclient.Client

	url string
	rpc *rpc.Client
}

func NewClient(url string) (Client, error) {
	c, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	return &client{
		Client: ethclient.NewClient(c),
		url:    url,
		rpc:    c,
	}, nil
}

func (c *client) URL() string {
	return c.url
}

func (c *client) TransactionByHash(
	ctx context.Context,
	hash common.Hash,
) (tx *types.Transaction, isPending bool, err error) {
	return c.Client.TransactionByHash(ctx, hash)
}

func (c *client) NewBatchHeaderClient() (BatchHeaderClient, error) {
	client, err := w3.Dial(c.URL())
	if err != nil {
		return nil, err
	}
	return &BatchHeaderRPCClient{client: client}, nil
}

func (c *client) GetProof(ctx context.Context, account common.Address, keys []string, blockNumber *big.Int) (*gethclient.AccountResult, error) {
	return gethclient.New(c.rpc).GetProof(ctx, account, keys, blockNumber)
}

type signableClient struct {
	Client

	chainId *big.Int
	signer  Signer
}

func NewSignableClient(chainId *big.Int, rpc string, signer Signer) (SignableClient, error) {
	c, err := NewClient(rpc)
	if err != nil {
		return nil, err
	}

	return &signableClient{
		Client:  c,
		chainId: chainId,
		signer:  signer,
	}, nil
}

func (c *signableClient) ChainID() *big.Int {
	return new(big.Int).Set(c.chainId)
}

// Return signer address.
func (c *signableClient) Signer() common.Address {
	return c.signer.From()
}

// Return transaction authorization data.
func (c *signableClient) TransactOpts(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		Context: ctx,
		From:    c.signer.From(),
		Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return c.signer.SignTx(tx, c.ChainID())
		},
	}
}

func (c *signableClient) SignData(data []byte) (sig []byte, err error) {
	return c.signer.SignData(data)
}

func (c *signableClient) SignTx(tx *types.Transaction) (*types.Transaction, error) {
	return c.signer.SignTx(tx, c.ChainID())
}

type BatchHeaderRPCClient struct {
	client *w3.Client
}

func (c *BatchHeaderRPCClient) Get(
	ctx context.Context,
	start, end uint64,
) ([]*types.Header, error) {
	size := int(end - start + 1)

	headers := make([]*types.Header, size)
	calls := make([]w3types.Caller, size)
	for i := 0; i < size; i++ {
		headers[i] = &types.Header{}
		calls[i] = eth.
			HeaderByNumber(new(big.Int).SetUint64(start + uint64(i))).
			Returns(headers[i])
	}

	if err := c.client.CallCtx(ctx, calls...); err != nil {
		return nil, err
	}
	return headers, nil
}

func (c *BatchHeaderRPCClient) Close() error {
	return c.client.Close()
}

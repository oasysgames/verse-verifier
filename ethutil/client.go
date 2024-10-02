package ethutil

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"golang.org/x/sync/semaphore"
)

var (
	ErrTooManyRequests = errors.New("too many requests")
)

type SignDataFn = func(hash []byte) (sig []byte, err error)

type Client interface {
	bind.ContractBackend
	bind.DeployBackend

	URL() string
	BlockNumber(ctx context.Context) (uint64, error)
	HeaderWithCache(ctx context.Context) (*types.Header, error)
	TransactionByHash(
		ctx context.Context,
		hash common.Hash,
	) (tx *types.Transaction, isPending bool, err error)
	FilterLogsWithRateThottling(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
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

	url       string
	blockTime time.Duration
	rpc       *rpc.Client
	// used for api rate thottling.
	sem *semaphore.Weighted

	headerCache atomic.Pointer[types.Header]
}

func NewClient(url string, blockTime time.Duration) (Client, error) {
	c, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	// This is magic number, it should be updated based on the network.
	// semaphore is used for log filtering rate thottling for now.
	// ethclient.NewClient(c).SendTransaction
	const concurrency = 2
	return &client{
		Client:    ethclient.NewClient(c),
		url:       url,
		blockTime: blockTime,
		rpc:       c,
		sem:       semaphore.NewWeighted(concurrency),
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

func (c *client) FilterLogsWithRateThottling(ctx context.Context, q ethereum.FilterQuery) (logs []types.Log, err error) {
	if err = c.sem.Acquire(ctx, 1); err != nil {
		// continue even if we can't acquire semaphore.
		fmt.Printf("***** WARN *****\nfailed to acquire semaphore in filter: %v\n", err)
	} else {
		defer c.sem.Release(1)
	}

	logs, err = c.Client.FilterLogs(ctx, q)
	if err != nil && strings.Contains(err.Error(), "too many requests") {
		// sleep longer if the rate limit is reached.
		time.Sleep(3 * time.Second)
		err = fmt.Errorf("%w: %v", ErrTooManyRequests, err)
		return
	}

	// sleep if the filter range is big or not set.
	if q.ToBlock == nil || q.FromBlock == nil || 1024 <= q.ToBlock.Sub(q.ToBlock, q.FromBlock).Uint64() {
		time.Sleep(300 * time.Microsecond)
	}

	return
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

func (c *client) HeaderWithCache(ctx context.Context) (*types.Header, error) {
	if c.blockTime == 0 {
		return c.HeaderByNumber(ctx, nil)
	}

	cache := c.headerCache.Load()
	if cache != nil && time.Unix(int64(cache.Time), 0).Add(c.blockTime).After(time.Now()) {
		return cache, nil
	}

	if latest, err := c.HeaderByNumber(ctx, nil); err != nil {
		return nil, err
	} else {
		c.headerCache.Store(latest)
		return latest, nil
	}
}

type signableClient struct {
	Client

	chainId *big.Int
	signer  Signer
}

func NewSignableClient(chainId *big.Int, c Client, signer Signer) SignableClient {
	return &signableClient{
		Client:  c,
		chainId: chainId,
		signer:  signer,
	}
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

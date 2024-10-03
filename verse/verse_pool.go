package verse

import (
	"context"
	"fmt"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

const (
	// Since the current usage doesn't result in a high
	// cache hit rate, the cache size will be kept modest.
	emittedBlockCacheSize = 16

	// If the difference between the next index and the cached index is
	// greater than or equal to this value, update the FilterStartBlockCache.
	filterStartBlockCacheUpdateThreshold = 5
)

var (
	_ VersePool = &versePool{}
)

// VersePool is a pool of `Verse` instances that can be shared across the entire application.
// All methods are goroutine safe.
type VersePool interface {
	// Add a new Verse to the pool.
	Add(new Verse, canSubmit bool) bool

	// Get Verse from the pool.
	Get(contract common.Address) (*VersePoolItem, bool)

	// Removes the Verse from the pool. Does nothing if it does not exist.
	Delete(contract common.Address)

	// Get Verse from the pool one by one and pass it to the callback function.
	Range(func(item *VersePoolItem) bool)

	// Returns the next rollup index to be verified. Since the values are cached within
	// the pool, the cache will be returned if the L1 block has not changed since the last access.
	// If `confirmation` is greater than 1, call the method with the block number specified as
	// `latest - confirmation`. If the latest block is smaller than `confirmation` and `waits`
	// is true, it will wait for the number of confirmations to pass.
	// (Since the mainnet/testnet has grown sufficiently, there won't be any waiting.)
	NextIndex(
		ctx context.Context,
		contract common.Address,
		confirmation int,
		waits bool,
	) (uint64, error)

	// Returns the block number at which the event with the given rollup index was emitted on the L1.
	// Since the values are cached within the pool, the cache will be returned if the L1 block has
	// not changed since the last retrieval. If the confirmation is greater than 1, call the method
	// with the block number of 'latest - confirmation'. If the latest block is smaller than
	// `confirmation` and `waits` is true, it will wait for the number of confirmations to pass.
	// (Since the mainnet/testnet has grown sufficiently, there won't be any waiting.)
	EventEmittedBlock(
		ctx context.Context,
		contract common.Address,
		rollupIndex uint64,
		confirmation int,
		waits bool,
	) (uint64, error)
}

type VersePoolItem struct {
	verse                 Verse
	canSubmit             bool
	nextIndexCache        atomic.Pointer[nextIndexCache]
	emittedBlockCache     *lru.Cache[uint64, uint64]
	filterStartBlockCache atomic.Pointer[filterStartBlockCache]
}

type nextIndexCache struct {
	fetchedBlock,
	value uint64
}

type filterStartBlockCache struct {
	rollupIndex,
	startBlock uint64
}

func NewVersePool(l1Client ethutil.Client) VersePool {
	return &versePool{l1Client: l1Client}
}

type versePool struct {
	l1Client ethutil.Client
	verses   util.SyncMap[common.Address, *VersePoolItem]
}

func (pool *versePool) Add(new Verse, canSubmit bool) bool {
	cacheKey := new.RollupContract()

	old, _ := pool.verses.Load(cacheKey)
	if old != nil && old.Verse().URL() == new.URL() && old.canSubmit == canSubmit {
		return false
	}

	blkCache, _ := lru.New[uint64, uint64](emittedBlockCacheSize)
	pool.verses.Store(cacheKey, &VersePoolItem{
		verse:             new,
		canSubmit:         canSubmit,
		emittedBlockCache: blkCache,
	})
	return true
}

func (pool *versePool) Get(contract common.Address) (*VersePoolItem, bool) {
	return pool.verses.Load(contract)
}

func (pool *versePool) Delete(contract common.Address) {
	pool.verses.Delete(contract)
}

func (pool *versePool) Range(fn func(*VersePoolItem) bool) {
	pool.verses.Range(func(key common.Address, value *VersePoolItem) bool {
		return fn(value)
	})
}

func (pool *versePool) NextIndex(
	ctx context.Context,
	contract common.Address,
	confirmation int,
	waits bool,
) (uint64, error) {
	item, ok := pool.verses.Load(contract)
	if !ok {
		return 0, fmt.Errorf("not in the pool: %s", contract)
	}

	// Assume the fetched nextIndex is not reorged, as we confirm `confirmation` blocks
	confirmed, err := decideConfirmationBlockNumber(ctx, confirmation, pool.l1Client, waits)
	if err != nil {
		return 0, err
	}

	cache := item.nextIndexCache.Load()
	if cache != nil && cache.fetchedBlock >= confirmed {
		return cache.value, nil
	}

	opts := &bind.CallOpts{Context: ctx, BlockNumber: new(big.Int).SetUint64(confirmed)}
	ni, err := item.verse.NextIndex(opts)
	if err != nil {
		return 0, err
	}

	item.nextIndexCache.Store(&nextIndexCache{fetchedBlock: confirmed, value: ni})

	// To reduce the RPC load when retrieving the event emission block,
	// also obtain the block number where the next index event was emitted.
	fsbCache := item.filterStartBlockCache.Load()
	if fsbCache != nil {
		min, max := util.MinMax(ni, fsbCache.rollupIndex)
		if max-min >= filterStartBlockCacheUpdateThreshold {
			fsbCache = nil
		}
	}
	if fsbCache == nil && ni > 0 {
		// Get one previous event since the event matching the next index may not have been emitted yet.
		previ := ni - 1
		emitted, err := pool.EventEmittedBlock(ctx, contract, previ, confirmation, true)
		if err == nil {
			item.filterStartBlockCache.Store(&filterStartBlockCache{rollupIndex: previ, startBlock: emitted})
		}
	}

	return ni, nil
}

func (pool *versePool) EventEmittedBlock(
	ctx context.Context,
	contract common.Address,
	rollupIndex uint64,
	confirmation int,
	waits bool,
) (uint64, error) {
	item, ok := pool.verses.Load(contract)
	if !ok {
		return 0, fmt.Errorf("not in the pool: %s", contract)
	}
	if cache, ok := item.emittedBlockCache.Get(rollupIndex); ok {
		return cache, nil
	}

	confirmed, err := decideConfirmationBlockNumber(ctx, confirmation, pool.l1Client, waits)
	if err != nil {
		return 0, err
	}

	opts := &bind.FilterOpts{Context: ctx, End: &confirmed}

	// To reduce RPC load, if the target index is larger than the next index cache,
	// use the block where the event was emitted as the starting point for the search.
	fsbCache := item.filterStartBlockCache.Load()
	if fsbCache != nil && rollupIndex >= fsbCache.rollupIndex {
		opts.Start = fsbCache.startBlock
	}

	emittedBlock, err := item.verse.EventEmittedBlock(opts, rollupIndex)
	if err != nil {
		return 0, err
	}

	item.emittedBlockCache.Add(rollupIndex, emittedBlock)
	return emittedBlock, nil
}

func (item *VersePoolItem) Verse() Verse    { return item.verse }
func (item *VersePoolItem) CanSubmit() bool { return item.canSubmit }

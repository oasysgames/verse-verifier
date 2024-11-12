package stakemanager

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

var (
	totalStakeKey = common.HexToAddress("0xffffffffffffffffffffffffffffffffffffffff")
)

type IStakeManager interface {
	GetTotalStake(callOpts *bind.CallOpts, epoch *big.Int) (*big.Int, error)

	GetOperatorStakes(callOpts *bind.CallOpts, operator common.Address, epoch *big.Int) (*big.Int, error)
}

type Cache struct {
	sm      IStakeManager
	ttl     time.Duration
	mu      sync.Mutex
	entries map[common.Address]*cacheEntry
}

type cacheEntry struct {
	amount   *big.Int
	expireAt time.Time
}

func NewCache(sm IStakeManager, ttl time.Duration) *Cache {
	return &Cache{
		sm:      sm,
		ttl:     ttl,
		entries: make(map[common.Address]*cacheEntry),
	}
}

func (c *Cache) TotalStake(ctx context.Context) *big.Int {
	amount, _ := c.get(totalStakeKey, func() (*big.Int, error) {
		return c.sm.GetTotalStake(&bind.CallOpts{Context: ctx}, common.Big0)
	})
	return amount
}

func (c *Cache) TotalStakeWithError(ctx context.Context) (*big.Int, error) {
	return c.get(totalStakeKey, func() (*big.Int, error) {
		return c.sm.GetTotalStake(&bind.CallOpts{Context: ctx}, common.Big0)
	})
}

func (c *Cache) StakeBySigner(ctx context.Context, signer common.Address) *big.Int {
	amount, _ := c.get(signer, func() (*big.Int, error) {
		return c.sm.GetOperatorStakes(&bind.CallOpts{Context: ctx}, signer, common.Big0)
	})
	return amount
}

func (c *Cache) get(key common.Address, getAmount func() (*big.Int, error)) (*big.Int, error) {
	c.mu.Lock()
	if _, ok := c.entries[key]; !ok {
		c.entries[key] = &cacheEntry{amount: big.NewInt(0)}
	}
	c.mu.Unlock()

	var (
		cache = c.entries[key]
		err   error
	)
	if time.Now().After(cache.expireAt) {
		ttl := c.ttl
		if value, innerErr := getAmount(); innerErr == nil {
			cache.amount = value
		} else {
			log.Error("Failed to refresh", "err", innerErr)
			err = innerErr
			ttl /= 10 // prevent requests when RPC is down
		}
		cache.expireAt = time.Now().Add(ttl)
	}

	return new(big.Int).Set(cache.amount), err
}

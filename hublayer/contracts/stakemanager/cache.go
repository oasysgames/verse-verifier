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

type IStakeManager interface {
	GetTotalStake(callOpts *bind.CallOpts, epoch *big.Int) (*big.Int, error)

	GetValidators(callOpts *bind.CallOpts, epoch, cursol, howMany *big.Int) (struct {
		Owners     []common.Address
		Operators  []common.Address
		Stakes     []*big.Int
		Candidates []bool
		NewCursor  *big.Int
	}, error)
}

type Cache struct {
	sm           IStakeManager
	mu           sync.Mutex
	total        *big.Int
	signerStakes map[common.Address]*big.Int
}

func NewCache(sm IStakeManager) *Cache {
	return &Cache{
		sm:           sm,
		total:        big.NewInt(0),
		signerStakes: make(map[common.Address]*big.Int),
	}
}

func (c *Cache) RefreshLoop(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := c.Refresh(ctx); err != nil {
				log.Error("Failed to refresh", "err", err)
			} else {
				ticker.Reset(interval)
			}
		}
	}
}

func (c *Cache) Refresh(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	total, err := c.sm.GetTotalStake(&bind.CallOpts{Context: ctx}, common.Big0)
	if err != nil {
		return err
	}

	cursor, howMany := big.NewInt(0), big.NewInt(50)
	signerStakes := make(map[common.Address]*big.Int)
	for {
		result, err := c.sm.GetValidators(&bind.CallOpts{Context: ctx}, common.Big0, cursor, howMany)
		if err != nil {
			return err
		} else if len(result.Owners) == 0 {
			break
		}

		for i, operator := range result.Operators {
			signerStakes[operator] = result.Stakes[i]
		}
		cursor = result.NewCursor
	}

	c.total = total
	c.signerStakes = signerStakes
	return nil
}

func (c *Cache) TotalStake() *big.Int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.total
}

func (c Cache) SignerStakes() map[common.Address]*big.Int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.signerStakes
}

func (c Cache) StakeBySigner(signer common.Address) *big.Int {
	c.mu.Lock()
	defer c.mu.Unlock()

	b := c.signerStakes[signer]
	if b == nil {
		b = big.NewInt(0)
	}
	return b
}

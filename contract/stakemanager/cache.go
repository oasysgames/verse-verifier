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
	candidates   map[common.Address]bool
}

func NewCache(sm IStakeManager) *Cache {
	return &Cache{
		sm:           sm,
		total:        big.NewInt(0),
		signerStakes: make(map[common.Address]*big.Int),
	}
}

func (c *Cache) RefreshLoop(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(time.Second * 60 * 3)
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

func (c *Cache) Refresh(parent context.Context) error {
	ctx, cancel := context.WithTimeout(parent, time.Second*15)
	defer cancel()

	total, err := c.sm.GetTotalStake(&bind.CallOpts{Context: ctx}, common.Big0)
	if err != nil {
		return err
	}

	cursor, howMany := big.NewInt(0), big.NewInt(50)
	signerStakes := make(map[common.Address]*big.Int)
	candidates := make(map[common.Address]bool)
	for {
		result, err := c.sm.GetValidators(&bind.CallOpts{Context: ctx}, common.Big0, cursor, howMany)
		if err != nil {
			return err
		} else if len(result.Owners) == 0 {
			break
		}

		for i, operator := range result.Operators {
			signerStakes[operator] = result.Stakes[i]
			candidates[operator] = result.Candidates[i]
		}
		cursor = result.NewCursor
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.total = total
	c.signerStakes = signerStakes
	c.candidates = candidates
	return nil
}

func (c *Cache) TotalStake() *big.Int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return new(big.Int).Set(c.total)
}

func (c *Cache) SignerStakes() map[common.Address]*big.Int {
	c.mu.Lock()
	defer c.mu.Unlock()

	cpy := make(map[common.Address]*big.Int)
	for k, v := range c.signerStakes {
		cpy[k] = new(big.Int).Set(v)
	}
	return cpy
}

func (c *Cache) StakeBySigner(signer common.Address) *big.Int {
	c.mu.Lock()
	defer c.mu.Unlock()

	if b := c.signerStakes[signer]; b != nil {
		return new(big.Int).Set(b)
	}
	return big.NewInt(0)
}

func (c *Cache) Candidates() map[common.Address]bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	cpy := make(map[common.Address]bool)
	for k, v := range c.candidates {
		cpy[k] = v
	}
	return cpy
}

func (c *Cache) Candidate(signer common.Address) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.candidates[signer]
}

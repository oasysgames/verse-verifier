package hublayer

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

// Worker to collect new blocks.
type BlockCollector struct {
	db       *database.Database
	hub      ethutil.ReadOnlyClient
	interval time.Duration

	log log.Logger
}

func NewBlockCollector(
	db *database.Database,
	hub ethutil.ReadOnlyClient,
	interval time.Duration,
) *BlockCollector {
	return &BlockCollector{
		db:       db,
		hub:      hub,
		interval: interval,
		log:      log.New("worker", "block-collector"),
	}
}

func (w *BlockCollector) Start(
	ctx context.Context,
) {
	w.log.Info("Worker started", "interval", w.interval)

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	reachedHead := false

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Worker stopped")
			return
		case <-ticker.C:
			if w.work(ctx) && !reachedHead {
				reachedHead = true
				ticker.Reset(w.interval)
			}
		}
	}
}

func (w *BlockCollector) work(ctx context.Context) (reachedHead bool) {
	// get the highest local block
	var (
		number  = uint64(1)
		highest *common.Hash
	)
	if h, err := w.db.Block.FindHighest(); err == nil {
		number = h.Number + 1
		highest = &h.Hash
	} else if !errors.Is(err, database.ErrNotFound) {
		w.log.Error("Failed to find highest block", "err", err)
		return false
	}

	// fetch block header from hub-layer
	h, err := w.hub.HeaderByNumber(ctx, new(big.Int).SetUint64(number))
	if errors.Is(err, ethereum.NotFound) {
		w.log.Debug("Wait for new block", "number", number)
		return true
	} else if err != nil {
		w.log.Error("Failed to fetch block header from hub-layer", "number", number, "err", err)
		return false
	}

	// check if reorganization has occurred
	if highest != nil && *highest != h.ParentHash {
		w.log.Info("Reorganization detected", "number", h.Number, "hash", h.Hash())
		if err = w.deleteReorganizedBlocks(ctx); err != nil {
			w.log.Error("Failed to delete reorganized blocks", "err", err)
		}
		return false
	}

	if err := w.db.Block.SaveNewBlock(number, h.Hash()); err != nil {
		w.log.Error("Failed to save new block", "err", err)
		return false
	}

	w.log.Info("New block", "number", number, "hash", h.Hash())
	return false
}

func (w *BlockCollector) deleteReorganizedBlocks(ctx context.Context) error {
	return w.db.Transaction(func(tx *database.Database) error {
		highest, err := tx.Block.FindHighest()
		if err != nil {
			return err
		}

		// delete from the head
		for number := highest.Number; number > 0; number-- {
			local, err := tx.Block.Find(number)
			if err != nil && !errors.Is(err, database.ErrNotFound) {
				return err
			}

			remote, err := w.hub.HeaderByNumber(ctx, new(big.Int).SetUint64(number))
			if err != nil {
				return err
			}

			if local.Hash == remote.Hash() {
				w.log.Info("Reached reorganization starting block",
					"number", number, "hash", remote.Hash().String())
				break
			}

			if _, err := tx.Block.Delete(number); err != nil {
				return err
			}

			w.log.Info("Deleted reorganized block", "number", number)
		}

		return nil
	})
}

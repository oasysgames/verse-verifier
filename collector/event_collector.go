package collector

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
)

// Worker to collect events for OasysStateCommitmentChain.
type EventCollector struct {
	cfg    *config.Verifier
	db     *database.Database
	hub    ethutil.Client
	signer common.Address
	log    log.Logger
}

func NewEventCollector(
	cfg *config.Verifier,
	db *database.Database,
	hub ethutil.Client,
	signer common.Address,
) *EventCollector {
	return &EventCollector{
		cfg:    cfg,
		db:     db,
		hub:    hub,
		signer: signer,
		log:    log.New("worker", "event-collector"),
	}
}

func (w *EventCollector) Start(ctx context.Context) {
	w.log.Info("Event collector started",
		"interval", w.cfg.Interval, "event-filter-limit", w.cfg.EventFilterLimit)

	ticker := time.NewTicker(w.cfg.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Event collector stopped")
			return
		case <-ticker.C:
			w.work(ctx)
		}
	}
}

func (w *EventCollector) work(ctx context.Context) {
	for {
		// get new blocks from database
		blocks, err := w.db.Block.FindUncollecteds(w.cfg.EventFilterLimit)
		if err != nil && !errors.Is(err, database.ErrNotFound) {
			w.log.Error("Failed to find uncollected blocks", "err", err)
			return
		} else if len(blocks) == 0 {
			w.log.Debug("Wait for new block")
			return
		}

		// collect event logs from hub-layer
		start, end := blocks[0], blocks[len(blocks)-1]
		logs, err := w.hub.FilterLogs(ctx, verse.NewEventLogFilter(start.Number, end.Number))
		if err != nil {
			w.log.Error("Failed to fetch event logs from hub-layer",
				"start", start, "end", end, "err", err)
			return
		}
		if len(logs) == 0 {
			w.log.Debug("No event log", "start", start, "end", end)
		}

		if err = w.db.Transaction(func(tx *database.Database) error {
			for _, log := range logs {
				if err := w.processLog(tx, &log); err != nil {
					return err
				}
			}
			if err := tx.Block.SaveCollected(end.Number, end.Hash); err != nil {
				w.log.Error("Failed to save collected block", "number", end, "err", err)
				return err
			}
			return nil
		}); err != nil {
			return
		}
	}
}

// Handler for event log from rollup contracts.
func (w *EventCollector) processLog(tx *database.Database, log *types.Log) error {
	event, err := verse.ParseEventLog(log)
	if err != nil {
		w.log.Error("Failed to parse event log",
			"block", log.BlockNumber, "contract", log.Address.Hex(), "err", err)
		return err
	}

	err = nil
	switch t := event.(type) {
	case *verse.RollupedEvent:
		err = w.handleRollupedEvent(tx, t)
	case *verse.DeletedEvent:
		err = w.handleDeletedEvent(tx, t)
	case *verse.VerifiedEvent:
		err = w.handleVerifiedEvent(tx, t)
	default:
		err = fmt.Errorf("unknown event")
	}

	return err
}

// Handler for new rollup event.
func (w *EventCollector) handleRollupedEvent(txdb *database.Database, e *verse.RollupedEvent) error {
	eventDB := e.EventDB(txdb)

	log := e.Logger(w.log)
	log.Info("New rollup event")

	// delete the `OptimismState` records in consideration of chain reorganization
	rows, err := eventDB.Deletes(e.Log.Address, e.RollupIndex)
	if err != nil {
		log.Error("Failed to delete reorganized events", "err", err)
		return err
	} else if rows > 0 {
		log.Info("Deleted reorganized events", "rows", rows)
	}

	// delete the `OptimismSignature` records in consideration of chain reorganization
	rows, err = txdb.OPSignature.Deletes(w.signer, e.Log.Address, e.RollupIndex)
	if err != nil {
		log.Error("Failed to delete reorganized signatures", "err", err)
		return err
	} else if rows > 0 {
		log.Info("Deleted reorganized signatures", "rows", rows)
	}

	if _, err := eventDB.Save(e.Log.Address, e.Parsed); err != nil {
		log.Error("Failed to save rollup event", "err", err)
		return err
	}

	return nil
}

// Handler for rollup delete event.
func (w *EventCollector) handleDeletedEvent(txdb *database.Database, e *verse.DeletedEvent) error {
	eventDB := e.EventDB(txdb)

	log := e.Logger(w.log)
	log.Info("New rollup delete event")

	// delete `OptimismState` records after target batchIndex
	rows, err := eventDB.Deletes(e.Log.Address, e.RollupIndex)
	if err != nil {
		log.Error("Failed to delete events", "err", err)
		return err
	} else if rows > 0 {
		log.Info("Deleted events", "rows", rows)
	}

	// delete the `OptimismSignature` records in consideration of chain reorganization
	rows, err = txdb.OPSignature.Deletes(w.signer, e.Log.Address, e.RollupIndex)
	if err != nil {
		log.Error("Failed to delete signatures", "err", err)
		return err
	} else if rows > 0 {
		log.Info("Deleted signatures", "rows", rows)
	}

	return nil
}

// Handler for rollup verified event.
func (w *EventCollector) handleVerifiedEvent(txdb *database.Database, e *verse.VerifiedEvent) error {
	log := e.Logger(w.log)
	log.Info("New rollup verified event")

	err := txdb.OPContract.SaveNextIndex(e.Log.Address, e.RollupIndex+1)
	if err != nil {
		log.Error("Failed to save next index", "err", err)
		return err
	}

	return nil
}

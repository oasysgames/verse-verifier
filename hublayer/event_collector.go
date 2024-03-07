package hublayer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

const (
	verseBuildEvent         = "Build"
	stateBatchAppendedEvent = "StateBatchAppended"
	stateBatchDeletedEvent  = "StateBatchDeleted"
	stateBatchVerifiedEvent = "StateBatchVerified"
)

var (
	sccABI                  abi.ABI
	stateBatchAppendedTopic common.Hash
	stateBatchDeletedTopic  common.Hash
	stateBatchVerifiedTopic common.Hash
	filterTopics            [][]common.Hash
)

func init() {
	if parsed, err := abi.JSON(strings.NewReader(scc.SccABI)); err != nil {
		panic(err)
	} else {
		sccABI = parsed
	}

	stateBatchAppendedTopic = sccABI.Events[stateBatchAppendedEvent].ID
	stateBatchDeletedTopic = sccABI.Events[stateBatchDeletedEvent].ID
	stateBatchVerifiedTopic = sccABI.Events[stateBatchVerifiedEvent].ID
	filterTopics = [][]common.Hash{{
		stateBatchAppendedTopic,
		stateBatchDeletedTopic,
		stateBatchVerifiedTopic,
	}}
}

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
	w.log.Info("Worker started",
		"interval", w.cfg.Interval, "event-filter-limit", w.cfg.EventFilterLimit)

	ticker := time.NewTicker(w.cfg.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Worker stopped")
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
		filter := ethereum.FilterQuery{
			Topics:    filterTopics,
			FromBlock: new(big.Int).SetUint64(start.Number),
			ToBlock:   new(big.Int).SetUint64(end.Number),
		}
		logs, err := w.hub.FilterLogs(ctx, filter)
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
				if err := w.processLog(tx, log); err != nil {
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

// Parse event logs and save to database.
func (w *EventCollector) processLog(tx *database.Database, log types.Log) error {
	event, err := parseLog(log)
	if err != nil {
		w.log.Error("Failed to parse event log",
			"block", log.BlockNumber, "scc", log.Address.Hex(), "err", err)
		return err
	}

	switch t := event.(type) {
	case *scc.SccStateBatchAppended:
		return w.processStateBatchAppendedEvent(tx, t)
	case *scc.SccStateBatchDeleted:
		return w.processStateBatchDeletedEvent(tx, t)
	case *scc.SccStateBatchVerified:
		return w.processStateBatchVerifiedEvent(tx, t)
	}

	return nil
}

func (w *EventCollector) processStateBatchAppendedEvent(
	tx *database.Database,
	e *scc.SccStateBatchAppended,
) error {
	var (
		address    = e.Raw.Address
		batchIndex = e.BatchIndex.Uint64()
		logCtx     = []interface{}{
			"block", e.Raw.BlockNumber,
			"scc", address.Hex(),
			"index", batchIndex,
		}
	)
	w.log.Info("New SCC.StateBatchAppended event", logCtx...)

	// delete the `OptimismState` records in consideration of chain reorganization
	if rows, err := tx.Optimism.DeleteStates(address, batchIndex); err != nil {
		w.log.Error("Failed to delete reorganized states", append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted reorganized states", append(logCtx, "rows", rows)...)
	}

	// delete the `OptimismSignature` records in consideration of chain reorganization
	if rows, err := tx.OPSignature.Deletes(w.signer, address, batchIndex); err != nil {
		w.log.Error("Failed to delete reorganized signatures",
			append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted reorganized signatures", append(logCtx, "rows", rows)...)
	}

	// save new state
	if _, err := tx.Optimism.SaveState(e); err != nil {
		w.log.Error("Failed to save SCC.StateBatchAppended event", append(logCtx, "err", err)...)
		return err
	}
	return nil
}

func (w *EventCollector) processStateBatchDeletedEvent(
	tx *database.Database,
	e *scc.SccStateBatchDeleted,
) error {
	var (
		address    = e.Raw.Address
		batchIndex = e.BatchIndex.Uint64()
		logCtx     = []interface{}{
			"block", e.Raw.BlockNumber,
			"scc", address.Hex(),
			"index", batchIndex,
		}
	)
	w.log.Info("New SCC.StateBatchDeleted event", logCtx...)

	// delete `OptimismState` records after target batchIndex
	if rows, err := tx.Optimism.DeleteStates(address, batchIndex); err != nil {
		w.log.Error("Failed to delete states", append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted states", append(logCtx, "rows", rows)...)
	}

	// delete the `OptimismSignature` records in consideration of chain reorganization
	if rows, err := tx.OPSignature.Deletes(w.signer, address, batchIndex); err != nil {
		w.log.Error("Failed to delete reorganized signatures", append(logCtx, "err", err)...)
		return err
	} else if rows > 0 {
		w.log.Info("Deleted reorganized signatures", append(logCtx, "rows", rows)...)
	}

	return nil
}

func (w *EventCollector) processStateBatchVerifiedEvent(
	tx *database.Database,
	e *scc.SccStateBatchVerified,
) error {
	nextIndex := e.BatchIndex.Uint64() + 1

	logCtx := []interface{}{
		"block", e.Raw.BlockNumber,
		"scc", e.Raw.Address.Hex(),
		"next_index", nextIndex,
	}
	w.log.Info("New SCC.StateBatchVerified event", logCtx...)

	if err := tx.OPContract.SaveNextIndex(e.Raw.Address, nextIndex); err != nil {
		w.log.Error("Failed to save next index", append(logCtx, "err", err)...)
		return err
	}
	return nil
}

func parseLog(log types.Log) (interface{}, error) {
	var (
		event string
		out   interface{}
	)
	switch log.Topics[0] {
	case stateBatchAppendedTopic:
		event = stateBatchAppendedEvent
		out = &scc.SccStateBatchAppended{Raw: log}
	case stateBatchDeletedTopic:
		event = stateBatchDeletedEvent
		out = &scc.SccStateBatchDeleted{Raw: log}
	case stateBatchVerifiedTopic:
		event = stateBatchVerifiedEvent
		out = &scc.SccStateBatchVerified{Raw: log}
	default:
		return nil, fmt.Errorf("invalid log topic: %s", log.Topics[0].String())
	}

	if err := sccABI.UnpackIntoInterface(out, event, log.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack log data: %w", err)
	}

	var indexed abi.Arguments
	for _, arg := range sccABI.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, fmt.Errorf("failed to parse indexed log data: %w", err)
	}

	return out, nil
}

package verifier

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
)

// Worker to verify rollups.
type Verifier struct {
	// fields passed during construction
	cfg      *config.Verifier
	db       *database.Database
	l1Signer ethutil.SignableClient
	newSigP2P,
	unverifiedSigP2P P2P
	log log.Logger

	// internal fields
	l1Head atomic.Pointer[types.Header]
	tasks  sync.Map
}

type P2P interface {
	PublishSignatures(ctx context.Context, sigs []*database.OptimismSignature)
}

// Returns the new verifier.
func NewVerifier(
	cfg *config.Verifier,
	db *database.Database,
	p2p P2P,
	l1Signer ethutil.SignableClient,
) *Verifier {
	verifier := &Verifier{
		cfg:              cfg,
		db:               db,
		newSigP2P:        p2p,
		unverifiedSigP2P: p2p,
		l1Signer:         l1Signer,
		log:              log.New("worker", "verifier"),
	}

	go verifier.l1HeadUpdater()

	return verifier
}

func (w *Verifier) L1Signer() ethutil.SignableClient {
	return w.l1Signer
}

func (w *Verifier) HasTask(contract common.Address, l2RPC string) bool {
	val, ok := w.tasks.Load(contract)
	if !ok {
		return false
	}
	// If the L2 RPC is changed, replace the worker.
	return l2RPC == val.(verse.VerifiableVerse).L2Client().URL()
}

func (w *Verifier) AddTask(ctx context.Context, task verse.VerifiableVerse, chainId uint64) {
	_, exists := w.tasks.Load(task.RollupContract())
	w.tasks.Store(task.RollupContract(), task)
	if !exists {
		// Start the verifier by each contract.
		go w.startVerifier(ctx, task.RollupContract(), chainId)
	}
}

func (w *Verifier) GetTask(contract common.Address) (task verse.VerifiableVerse, found bool) {
	var val any
	val, found = w.tasks.Load(contract)
	if !found {
		return
	}
	task, found = val.(verse.VerifiableVerse)
	return
}

func (w *Verifier) RemoveTask(contract common.Address) {
	w.tasks.Delete(contract)
}

func (w *Verifier) startVerifier(ctx context.Context, contract common.Address, chainId uint64) {
	log := w.log.New("chain-id", chainId)

	// Create block range manager.
	task, _ := w.GetTask(contract)
	rangeMgr, err := w.getBlockRangeManager(log, ctx, task)
	if err != nil {
		return // canceled by parent context
	}

	// Publish all unverified signatures every 4 times.
	var counter int
	publishAllUnverifiedSigs := func() bool {
		counter++
		return counter%4 == 0
	}

	tick := time.NewTicker(w.cfg.Interval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Info("Verifying work stopped")
			return
		case <-tick.C:
			task, found := w.GetTask(contract)
			if !found {
				log.Info("Exit verifier as task is evicted")
				return
			}
			if err := w.work(ctx, task, chainId, rangeMgr, publishAllUnverifiedSigs()); err != nil {
				log.Error("Failed to run verification", "err", err)
			}
		}
	}
}

func (w *Verifier) work(
	parent context.Context,
	task verse.VerifiableVerse,
	chainId uint64,
	rangeMgr *eventFetchingBlockRangeManager,
	publishAllUnverifiedSigs bool,
) error {
	l1ctx, l1cancel := context.WithTimeout(parent, w.cfg.StateCollectTimeout)
	defer l1cancel()

	// Assume the fetched nextIndex is not reorged, as we confirm `w.cfg.Confirmations` blocks
	nextIndex, err := task.NextIndex(l1ctx, w.cfg.Confirmations, true)
	if err != nil {
		return fmt.Errorf("failed to call the NextIndex method: %w", err)
	}
	log := w.log.New("chain-id", chainId, "next-index", nextIndex)

	// Clean up old signatures
	if err = w.cleanOldSignatures(task.RollupContract(), nextIndex); err != nil {
		log.Warn("Failed to delete old signatures", "err", err)
	}

	// determine the start block number
	maxEnd, err := w.determineMaxEnd(log, l1ctx, task, nextIndex)
	if err != nil {
		return err
	}
	start, end, skipFetchlog := rangeMgr.get(maxEnd)
	log = log.New("max-end", maxEnd, "start", start, "end", end)

	if skipFetchlog && !publishAllUnverifiedSigs {
		log.Info("Skip fetching logs")
		return nil
	}

	// fetch event logs
	var logs []types.Log
	if !skipFetchlog {
		logs, err = w.l1Signer.FilterLogsWithRateThottling(
			l1ctx, verse.NewEventLogFilter(start, end, []common.Address{task.RollupContract()}))
		if err != nil {
			if errors.Is(err, ethutil.ErrTooManyRequests) {
				log.Warn("Rate limit exceeded", "err", err)
			}
			// Restore the next start block number if the fetching logs failed.
			rangeMgr.restore()
			return fmt.Errorf("failed to fetch(start: %d, end: %d) event logs from hub-layer: %w", start, end, err)
		}
	}
	log = log.New("count-logs", len(logs))

	log.Info("Start verification of all fetched logs")
	elapsed := time.Now()

	// verify the fetched logs
	var (
		opsigs []*database.OptimismSignature
		// flag at least one log verification failed.
		atLeastOneLogVerificationFailed bool
		// As the replica syncing is not real-time, the retry mechanism is required.
		backoffIncr, backoffDecr = w.retryBackoff()
	)
	for i := range logs {
		log := log.New("log-index", i)
		var (
			row *database.OptimismSignature
			err error
		)
		for {
			l2ctx, l2cancel := context.WithTimeout(parent, w.cfg.StateCollectTimeout*2)
			row, err = w.verifyAndSaveLog(l2ctx, &logs[i], task, nextIndex, log)
			l2cancel()

			// break if the verification is successful or skip
			if err == nil {
				break
			}
			// exit if context have been canceled
			if errors.Is(err, context.Canceled) {
				return err
			}

			// retry immediately if deadline exceeded
			// (In cases of high-load or high-latency of L2)
			if errors.Is(err, context.DeadlineExceeded) {
				log.Warn("Retry verification immediately")
				continue
			}

			// retry after the backoff if any errors
			// (In cases of maintenance of L2)
			delay, remain, attempts := backoffIncr()
			// give up if the retry time limit is exceeded
			if remain <= 0 {
				log.Error("Exceeded retry limit")
				break
			}

			log.Warn("Retry verification",
				"delay", delay, "remain", remain, "attempts", attempts, "err", err)
			select {
			case <-parent.Done():
				return parent.Err()
			case <-time.NewTimer(delay).C:
			}
		}

		// verification failed
		if err != nil {
			// skip the log if the verification failed
			log.Error("Failed to verify a log", "err", err)
			atLeastOneLogVerificationFailed = true
		}

		if row == nil {
			// skip if the row is nil
			// - when the event is not a rollup event
			// - when the event is already verified
			continue
		}

		backoffDecr()
		opsigs = append(opsigs, row)

		if i > 0 && i%50 == 0 {
			log.Info("Verification progress",
				"approved", row.Approved, "rollup-index", row.RollupIndex, "remain", len(logs)-i-1)
		}
	}
	if len(opsigs) > 0 {
		log.Info("Completed verification of all fetched logs",
			"count-new-sigs", len(opsigs), "elapsed", time.Since(elapsed))
	}

	// Will publish all unverified signatures if the flag is set.
	if publishAllUnverifiedSigs {
		contract := task.RollupContract()
		rows, err := w.db.OPSignature.FindUnverifiedBySigner(w.l1Signer.Signer(), nextIndex, &contract, database.FindUnverifiedBySignerLimit)
		if err != nil {
			log.Error("Failed to find unverified signatures", "err", err)
		} else if len(rows) > 0 {
			w.unverifiedSigP2P.PublishSignatures(parent, rows)
			log.Info("Published unverified signatures", "count-unverified-sigs", len(rows),
				"first-rollup-index", rows[0].RollupIndex, "last-rollup-index", rows[len(rows)-1].RollupIndex)
		}
	}

	if len(opsigs) > 0 {
		// publish all signatures at once
		w.newSigP2P.PublishSignatures(parent, opsigs)
		log.Info("Published new signatures", "count-new-sigs", len(opsigs),
			"first-rollup-index", opsigs[0].RollupIndex, "last-rollup-index", opsigs[len(opsigs)-1].RollupIndex)
	} else {
		log.Info("No signatures to publish")
	}

	if atLeastOneLogVerificationFailed {
		// Remove task if at least one log verification failed.
		// dinamic discovery on : The removed task will be added again in the next verse discovery
		// dinamic discovery off: restarting is required to add the removed task again
		w.RemoveTask(task.RollupContract())
	}

	return nil
}

func (w *Verifier) verifyAndSaveLog(ctx context.Context, log *types.Log, task verse.VerifiableVerse, nextIndex uint64, logger log.Logger) (*database.OptimismSignature, error) {
	event, err := verse.ParseEventLog(log)
	if err != nil {
		return nil, fmt.Errorf("failed to parse event log. block: %d contract: %s,: %w", log.BlockNumber, log.Address.Hex(), err)
	}

	// parse event log
	rollupEvent, ok := event.(*verse.RollupedEvent)
	if !ok {
		// skip `*verse.DeletedEvent` or `*verse.VerifiedEvent`
		return nil, nil
	}

	// cast to database event
	contract, err := w.db.OPContract.FindOrCreate(log.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to find or create contract(%s): %w", log.Address.Hex(), err)
	}
	dbEvent, err := rollupEvent.CastToDatabaseOPEvent(contract)
	if err != nil {
		return nil, fmt.Errorf("failed to cast to database. rollup-index: %d, event: %w", dbEvent.GetRollupIndex(), err)
	}

	if dbEvent.GetRollupIndex() < nextIndex {
		// skip old events
		return nil, nil
	}

	approved, err := task.Verify(logger, ctx, dbEvent, w.cfg.StateCollectLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to verification. rollup-index: %d, : %w", dbEvent.GetRollupIndex(), err)
	}

	msg := database.NewMessage(dbEvent, w.l1Signer.ChainID(), approved)
	sig, err := msg.Signature(w.l1Signer.SignData)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate signature. rollup-index: %d, : %w", dbEvent.GetRollupIndex(), err)
	}

	row, err := w.db.OPSignature.Save(
		nil, nil,
		w.l1Signer.Signer(),
		dbEvent.GetContract().Address,
		dbEvent.GetRollupIndex(),
		dbEvent.GetRollupHash(),
		approved,
		sig)
	if err != nil {
		return nil, fmt.Errorf("failed to save signature. rollup-index: %d, : %w", dbEvent.GetRollupIndex(), err)
	}

	return row, nil
}

func (w *Verifier) cleanOldSignatures(contract common.Address, nextIndex uint64) error {
	var verifiedIndex uint64
	if nextIndex == 0 {
		verifiedIndex = 0
	} else {
		verifiedIndex = nextIndex - 1
	}

	// Keep the last 3 signatures.
	if verifiedIndex < 3 {
		return nil
	}
	deleteIndex := uint64(verifiedIndex - 3)
	if _, err := w.db.OPSignature.DeleteOlds(contract, deleteIndex, database.DeleteOldsLimit); err != nil {
		return fmt.Errorf("failed to delete old signatures. deleteIndex: %d, : %w", deleteIndex, err)
	}
	return nil
}

func (w *Verifier) retryBackoff() (incr func() (delay, remain time.Duration, attempts int), decr func()) {
	started := time.Now()

	var counter, gauge int
	incr = func() (time.Duration, time.Duration, int) {
		// backoff delay: 0.1s, 0.8s, 6.4s, 51.2s, 409.6s(7m), 3276.8s(54m),
		delay := 100 << (3 * gauge) * time.Millisecond
		if delay <= 0 || delay > w.cfg.MaxRetryBackoff { // delay <= 0 is overflow
			delay = w.cfg.MaxRetryBackoff
		} else {
			gauge++
		}

		// The remaining time will not be replenished even if `decr` is done.
		remain := w.cfg.RetryTimeout - time.Since(started)
		if remain < 0 {
			remain = 0
		}

		counter++
		return delay, remain, counter
	}
	decr = func() {
		if gauge > 0 {
			gauge--
		}
	}
	return
}

// Updater for the L1 head.
// Don't need to safely exit this method when the process stops.
func (w *Verifier) l1HeadUpdater() {
	tick := util.NewTicker(w.cfg.Interval, 1)

	for range tick.C {
		ctx, cancel := context.WithTimeout(context.Background(), w.cfg.Interval)
		if h, err := w.l1Signer.HeaderByNumber(ctx, nil); err != nil {
			w.log.Warn("Failed to update L1 head", "err", err)
		} else {
			w.log.Debug("Updated L1 head", "number", h.Number, "hash", h.Hash())
			w.l1Head.Store(h)
		}
		cancel()
	}
}

// Fetch the NextIndex that should be verified next, and create a BlockRangeManager
// starting from the block number where the corresponding rollup event was emitted.
// Note: This method retries infinitely until it succeeds or is canceled.
func (w *Verifier) getBlockRangeManager(
	log log.Logger,
	ctx context.Context,
	task verse.VerifiableVerse,
) (*eventFetchingBlockRangeManager, error) {
	var (
		nextIndex    *uint64
		emittedBlock uint64
	)
	err := util.Retry(ctx, 0, w.cfg.Interval, func() error {
		if nextIndex == nil {
			ni, err := task.NextIndex(ctx, w.cfg.Confirmations, true)
			if err != nil {
				log.Info("Failed to fetch next index", "err", err)
				return err // retry
			}
			nextIndex = &ni

			// If NextIndex is 1 or greater, there is a possibility that verification has been
			// completed up to the latest rollup, so set the starting point to the previous event.
			if ni > 0 {
				*nextIndex--
			}
		}

		// Fetch the L1 block number where the event matching the nextIndex was emitted.
		var err error
		emittedBlock, err = task.GetEventEmittedBlock(
			ctx, *nextIndex, w.cfg.Confirmations, true)

		// If event does not exist, wait until it is rollup.
		if errors.Is(err, verse.ErrEventNotFound) {
			log.Warn("Event not found")
			time.Sleep(w.cfg.Interval * 4)
			return err // retry
		}

		if err != nil {
			log.Info("Failed to fetch event for next index", "next-index", nextIndex, "err", err)
		}
		return err // exit or retry
	})

	// Exit if canceled by caller
	if err != nil {
		return nil, err
	}
	log.Info("Initial block has been determined", "block", emittedBlock, "next-index", *nextIndex)
	return newEventFetchingBlockRangeManager(w.cfg.MaxLogFetchBlockRange, emittedBlock), nil
}

// Determine the upper limit of the end block.
// Note: This method retries infinitely until it succeeds or is canceled.
func (w *Verifier) determineMaxEnd(
	log log.Logger,
	ctx context.Context,
	task verse.VerifiableVerse,
	nextIndex uint64,
) (max uint64, err error) {
	if cache := w.l1Head.Load(); cache == nil {
		// If the L1 head is not fetched, nothing to do.
		err = errors.New("L1 head is not fetched")
		return
	} else {
		// Basically, upper limit is the head.
		max = cache.Number.Uint64()
		if max > uint64(w.cfg.Confirmations) {
			max -= uint64(w.cfg.Confirmations)
		}
	}

	// Fetch the L1 block number where the event matching the `nextIndex+w.cfg.MaxIndexDiff`
	// was emitted. It is to avoid excessively verifying new events, as the Submitter node
	// might not be subscribed to PubSub.
	maxIndex := nextIndex + uint64(w.cfg.MaxIndexDiff)
	err = util.Retry(ctx, 0, w.cfg.Interval, func() error {
		emittedBlock, inErr := task.GetEventEmittedBlock(ctx, maxIndex, w.cfg.Confirmations, true)
		// If it does not exist, verify up to the head.
		if errors.Is(inErr, verse.ErrEventNotFound) {
			return nil
		}
		// Retry if any errors
		if inErr != nil {
			return inErr
		}
		max = emittedBlock
		return nil
	})
	// Exit if canceled by caller
	if err != nil {
		return 0, err
	}
	return max, nil
}

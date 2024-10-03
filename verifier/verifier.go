package verifier

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
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
)

const (
	// Parameters for worker pool.
	maxIdleWorkerDuration      = time.Minute
	workerReleaseCheckInterval = time.Second
	workerReleaseCheckTimeout  = time.Duration(0) // no timeout
)

// Worker to verify rollups.
type Verifier struct {
	// fields passed during construction
	cfg        *config.Verifier
	db         *database.Database
	l1Signer   ethutil.SignableClient
	l2ClientFn L2ClientFn
	newSigP2P,
	unverifiedSigP2P P2P // Separated for testing.
	versepool verse.VersePool
	log       log.Logger

	// internal fields
	tasks util.SyncMap[common.Address, *taskT]
}

type P2P interface {
	PublishSignatures(ctx context.Context, sigs []*database.OptimismSignature) error
}

type L2ClientFn func(url string, blockTime time.Duration) (ethutil.Client, error)

type taskT struct {
	verse    verse.VerifiableVerse
	rangeMgr *eventFetchingBlockRangeManager
}

// Returns the new verifier.
func NewVerifier(
	cfg *config.Verifier,
	db *database.Database,
	p2p P2P,
	l1Signer ethutil.SignableClient,
	l2ClientFn L2ClientFn,
	versepool verse.VersePool,
) *Verifier {
	return &Verifier{
		cfg:              cfg,
		db:               db,
		newSigP2P:        p2p,
		unverifiedSigP2P: p2p,
		l1Signer:         l1Signer,
		l2ClientFn:       l2ClientFn,
		versepool:        versepool,
		log:              log.New("worker", "verifier"),
	}
}

func (w *Verifier) RemoveTask(contract common.Address) {
	w.tasks.Delete(contract)
}

func (w *Verifier) Start(ctx context.Context) {
	w.log.Info("Verifier started", "config", w.cfg)

	maxVerificationWorkers := w.cfg.MaxWorkers
	verificationInterval := w.cfg.Interval

	// Publish workers run expensive database queries, so the number of workers should be small.
	// P2P publishing is asynchronous and sufficiently fast.
	_, maxPublishWorkers := util.MinMax(maxVerificationWorkers/5, 2)
	publishInterval, _ := util.MinMax(verificationInterval*10, time.Minute)

	// Workers performing verification.
	go func() {
		// Manage running tasks to prevent dups.
		var running util.SyncMap[common.Address, time.Time]

		// Create woker pool.
		wp := util.NewWorkerPool(w.log, w.verify, maxVerificationWorkers,
			maxIdleWorkerDuration, workerReleaseCheckInterval, workerReleaseCheckTimeout)
		wp.Start()
		defer wp.Stop()

		// Every hour, remove Verse that were deleted from the pool from the local cache as well.
		cacheCleanupTick := time.NewTicker(time.Hour)
		defer cacheCleanupTick.Stop()

		workTick := time.NewTicker(verificationInterval)
		defer workTick.Stop()

		w.log.Info("Verification workers started",
			"max-workers", maxVerificationWorkers, "interval", verificationInterval)

		for {
			select {
			case <-ctx.Done():
				log.Info("Verification workers stopped")
				return
			case <-cacheCleanupTick.C:
				w.tasks.Range(func(cacheKey common.Address, task *taskT) bool {
					_, exists := w.versepool.Get(cacheKey)

					// Not delete if the last task has not completed.
					_, isRunning := running.Load(cacheKey)

					if !exists && !isRunning {
						task.verse.Logger(w.log).Info("Close connection and delete task cache")
						task.verse.L2Client().Close()
						w.tasks.Delete(cacheKey)
					}
					return true
				})
			case <-workTick.C:
				w.versepool.Range(func(item *verse.VersePoolItem) bool {
					log := item.Verse().Logger(w.log)

					// The RPC URL should not be used as a cache key. This is because the upgraded
					// Verse-Layer holds two contracts, v0 and v1, although they are the same URL.
					cacheKey := item.Verse().RollupContract()

					// Skip if previous task is running
					if started, isRunning := running.Load(cacheKey); isRunning {
						log.Info("Skip duplicate verification", "elapsed", time.Since(started))
						return true
					}
					running.Store(cacheKey, time.Now())

					// Since worker run asynchronously, should not call `running.Delete()` here.
					release := func() { running.Delete(cacheKey) }

					// If the cache does not exist or the RPC URL has changed, open a new connection.
					var task *taskT
					if cache, ok := w.tasks.Load(cacheKey); ok && cache.verse.L2Client().URL() == item.Verse().URL() {
						task = cache
					} else {
						if l2Client, err := w.l2ClientFn(item.Verse().URL(), 0); err != nil {
							log.Error("Failed to construct verse-layer client", "err", err)
						} else if rangeMgr, err := w.getBlockRangeManager(log, ctx, item.Verse()); err != nil {
							log.Error("Failed to construct block range manager", "err", err)
						} else {
							task = &taskT{
								verse:    item.Verse().WithVerifiable(l2Client),
								rangeMgr: rangeMgr,
							}
							w.tasks.Store(cacheKey, task)
						}
					}

					if task != nil {
						wp.Work(ctx, task, func(context.Context, *taskT) { release() })
					} else {
						release()
					}
					return true
				})
			}
		}
	}()

	// Workers performing publish unverified signatures.
	go func() {
		// Manage running tasks to prevent dups.
		var running util.SyncMap[common.Address, time.Time]

		// Create woker pool.
		wp := util.NewWorkerPool(w.log, w.publish, maxPublishWorkers,
			maxIdleWorkerDuration, workerReleaseCheckInterval, workerReleaseCheckTimeout)
		wp.Start()
		defer wp.Stop()

		tick := time.NewTicker(publishInterval)
		defer tick.Stop()

		w.log.Info("Publish workers started",
			"max-workers", maxPublishWorkers, "interval", publishInterval)

		for {
			select {
			case <-ctx.Done():
				log.Info("Publish workers stopped")
				return
			case <-tick.C:
				w.versepool.Range(func(item *verse.VersePoolItem) bool {
					log := item.Verse().Logger(w.log)
					cacheKey := item.Verse().RollupContract()

					// Skip if previous task is running
					if started, isRunning := running.Load(cacheKey); isRunning {
						log.Info("Skip duplicate publish", "elapsed", time.Since(started))
						return true
					}
					running.Store(cacheKey, time.Now())
					release := func() { running.Delete(cacheKey) }

					if cache, ok := w.tasks.Load(cacheKey); ok {
						wp.Work(ctx, cache, func(context.Context, *taskT) { release() })
					} else {
						// Task generation is left to the verification worker
						log.Warn("Task not found")
						release()
					}
					return true
				})
			}
		}
	}()

	<-ctx.Done()
	w.log.Info("Verifier stopped", "config", w.cfg)
}

// Fetch and verify rollup events from the Hub-Layer.
func (w *Verifier) verify(parent context.Context, task *taskT) {
	log := task.verse.Logger(w.log)

	l1ctx, l1cancel := context.WithTimeout(parent, w.cfg.StateCollectTimeout)
	defer l1cancel()

	nextIndex, err := w.versepool.NextIndex(l1ctx, task.verse.RollupContract(), w.cfg.Confirmations, true)
	if err != nil {
		w.log.Error("Failed to fetch next index", "err", err)
		return
	}
	log = log.New("next-index", nextIndex)

	// Clean up old signatures
	if err = w.cleanOldSignatures(task.verse.RollupContract(), nextIndex); err != nil {
		log.Warn("Failed to delete old signatures", "err", err)
	}

	// determine the start block number
	maxEnd, err := w.determineMaxEnd(l1ctx, task.verse, nextIndex)
	if err != nil {
		w.log.Error("Failed to determine the maximum end block", "err", err)
		return
	}
	start, end, skipFetchlog := task.rangeMgr.get(maxEnd)
	log = log.New("max-end", maxEnd, "start", start, "end", end)

	if skipFetchlog {
		log.Info("Skip fetching logs")
		return
	}

	// fetch event logs
	var logs []types.Log
	if !skipFetchlog {
		logs, err = w.l1Signer.FilterLogsWithRateThottling(
			l1ctx, verse.NewEventLogFilter(start, end, []common.Address{task.verse.RollupContract()}))
		if err != nil {
			if errors.Is(err, ethutil.ErrTooManyRequests) {
				log.Warn("Rate limit exceeded", "err", err)
			}
			// Restore the next start block number if the fetching logs failed.
			task.rangeMgr.restore()
			log.Error("Failed to fetch event logs from hub-layer", "err", err)
			return
		}
	}

	if len(logs) == 0 {
		log.Info("Skip verify")
		return
	}

	log = log.New("count-logs", len(logs))
	log.Info("Start verification of all fetched logs")

	// verify the fetched logs
	var (
		opsigs  []*database.OptimismSignature
		elapsed = time.Now()
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
			row, err = w.verifyAndSaveLog(l2ctx, &logs[i], task.verse, nextIndex, log)
			l2cancel()

			// break if the verification is successful or skip
			if err == nil {
				break
			}
			// exit if context have been canceled
			if errors.Is(err, context.Canceled) {
				return
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
				return
			case <-time.NewTimer(delay).C:
			}
		}

		// verification failed
		if err != nil {
			// skip the log if the verification failed
			log.Error("Failed to verify a log", "err", err)
			atLeastOneLogVerificationFailed = true
		}

		backoffDecr()

		if row == nil {
			// skip if the row is nil
			// - when the event is not a rollup event
			// - when the event is already verified
			continue
		}

		opsigs = append(opsigs, row)

		if i > 0 && i%50 == 0 {
			log.Info("Verification progress",
				"approved", row.Approved, "rollup-index", row.RollupIndex, "remain", len(logs)-i-1)
		}
	}

	log.Info("Completed verification of all fetched logs",
		"count-sigs", len(opsigs), "elapsed", time.Since(elapsed))

	if len(opsigs) > 0 {
		// publish all signatures at once
		if err := w.newSigP2P.PublishSignatures(parent, opsigs); err != nil {
			log.Error("Failed to publish new signatures", "err", err)
		} else {
			log.Info("Published new signatures", "count-sigs", len(opsigs),
				"first-rollup-index", opsigs[0].RollupIndex, "last-rollup-index", opsigs[len(opsigs)-1].RollupIndex)
		}
	} else {
		log.Info("No signatures to publish")
	}

	if atLeastOneLogVerificationFailed {
		// Remove task if at least one log verification failed.
		// dinamic discovery on : The removed task will be added again in the next verse discovery
		// dinamic discovery off: restarting is required to add the removed task again
		w.RemoveTask(task.verse.RollupContract())
	}
}

// Publish all unverified signatures.
func (w *Verifier) publish(parent context.Context, task *taskT) {
	log := task.verse.Logger(w.log)

	contract := task.verse.RollupContract()
	nextIndex, err := w.versepool.NextIndex(parent, contract, w.cfg.Confirmations, true)
	if err != nil {
		log.Error("Failed to fetch next index", "err", err)
		return
	}

	rows, err := w.db.OPSignature.FindUnverifiedBySigner(
		w.l1Signer.Signer(), nextIndex, &contract, database.FindUnverifiedBySignerLimit)
	if err != nil {
		log.Error("Failed to find unverified signatures", "err", err)
		return
	}
	if len(rows) > 0 {
		if err := w.unverifiedSigP2P.PublishSignatures(parent, rows); err != nil {
			log.Error("Failed to publish unverified signatures", "err", err)
		} else {
			log.Info("Published unverified signatures", "count", len(rows),
				"first-rollup-index", rows[0].RollupIndex, "last-rollup-index", rows[len(rows)-1].RollupIndex)
		}
	}
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

// Fetch the NextIndex that should be verified next, and create a BlockRangeManager
// starting from the block number where the corresponding rollup event was emitted.
// Note: This method retries infinitely until it succeeds or is canceled.
func (w *Verifier) getBlockRangeManager(
	log log.Logger,
	ctx context.Context,
	task verse.Verse,
) (*eventFetchingBlockRangeManager, error) {
	nextIndex, err := w.versepool.NextIndex(ctx, task.RollupContract(), w.cfg.Confirmations, true)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch next index: %w", err)
	}

	// If NextIndex is 1 or greater, there is a possibility that verification has been
	// completed up to the latest rollup, so set the starting point to the previous event.
	if nextIndex > 0 {
		nextIndex--
	}
	log = log.New("next-index", nextIndex)

	// Fetch the L1 block number where the event matching the next index was emitted.
	tick := time.NewTicker(w.cfg.Interval)
	defer tick.Stop()

	for {
		emittedBlock, err := w.versepool.EventEmittedBlock(
			ctx, task.RollupContract(), nextIndex, w.cfg.Confirmations, true)
		if err == nil {
			log.Info("Initial block has been determined", "block", emittedBlock)
			return newEventFetchingBlockRangeManager(w.cfg.MaxLogFetchBlockRange, emittedBlock), nil
		}
		if errors.Is(err, verse.ErrEventNotFound) {
			log.Warn("Event not found, wait until it is rollup")
		} else {
			return nil, fmt.Errorf("failed to fetch the event(index=%d) emitted block: %w", nextIndex, err)
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-tick.C:
		}
	}
}

// Determine the upper limit of the end block.
// Note: This method retries infinitely until it succeeds or is canceled.
func (w *Verifier) determineMaxEnd(
	ctx context.Context,
	task verse.VerifiableVerse,
	nextIndex uint64,
) (uint64, error) {
	// Fetch the L1 block number where the event matching the `nextIndex+w.cfg.MaxIndexDiff`
	// was emitted. It is to avoid excessively verifying new events, as the Submitter node
	// might not be subscribed to PubSub.
	emittedBlock, err := w.versepool.EventEmittedBlock(
		ctx, task.RollupContract(), nextIndex+uint64(w.cfg.MaxIndexDiff), w.cfg.Confirmations, true)
	if err == nil {
		return emittedBlock, nil
	}

	// If it does not exist or any errors occurs, verify up to the head.
	header, err := w.l1Signer.HeaderWithCache(ctx)
	if err != nil {
		// If the L1 head is not fetched, nothing to do.
		return 0, fmt.Errorf("failed to fetch the L1 head: %w", err)
	}

	max := header.Number.Uint64()
	if max > uint64(w.cfg.Confirmations) {
		max -= uint64(w.cfg.Confirmations)
	}
	return max, nil
}

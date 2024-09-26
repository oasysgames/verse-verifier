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
	p2p      P2P
	log      log.Logger

	// internal fields
	l1Latest atomic.Pointer[types.Header]
	tasks    sync.Map
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
		cfg:      cfg,
		db:       db,
		p2p:      p2p,
		l1Signer: l1Signer,
		log:      log.New("worker", "verifier"),
	}

	go verifier.l1LatestUpdater()

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
	var (
		tick                     = time.NewTicker(w.cfg.Interval)
		counter                  int
		blockRangeManager        = NeweventFetchingBlockRangeManager(w.l1Signer, w.cfg.MaxLogFetchBlockRange, w.cfg.StartBlockOffset)
		publishAllUnverifiedSigs = func() bool {
			counter++
			// Publish all unverified signatures every 4 times.
			return counter%4 == 0
		}
	)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Verifying work stopped", "chainId", chainId)
			return
		case <-tick.C:
			task, found := w.GetTask(contract)
			if !found {
				w.log.Info("exit verifier as task is evicted", "chainId", chainId)
				return
			}
			if err := w.work(ctx, task, chainId, blockRangeManager, publishAllUnverifiedSigs()); err != nil {
				if errors.Is(err, ErrStartBlockIsTooLarge) {
					w.log.Warn("Walk back the start block number to ensure the next index is correctly verified", "chainId", chainId, "err", err)
				} else {
					w.log.Error("Failed to run verification", "err", err)
				}
			}
		}
	}
}

func (w *Verifier) work(parent context.Context, task verse.VerifiableVerse, chainId uint64, blockRangeManager *eventFetchingBlockRangeManager, publishAllUnverifiedSigs bool) error {
	ctx, cancel := context.WithTimeout(parent, w.cfg.StateCollectTimeout)
	defer cancel()

	// Assume the fetched nextIndex is not reorged, as we confirm `w.cfg.Confirmations` blocks
	nextIndex, err := task.NextIndex(ctx, uint64(w.cfg.Confirmations), true)
	if err != nil {
		return fmt.Errorf("failed to call the NextIndex method: %w", err)
	}
	log := log.New("chainId", chainId, "next-index", nextIndex)

	// Clean up old signatures
	if err := w.cleanOldSignatures(task.RollupContract(), nextIndex.Uint64()); err != nil {
		log.Warn("Failed to delete old signatures", "err", err)
	}

	// determine the start block number
	start, end, skipFetchlog, err := blockRangeManager.GetBlockRange(ctx)
	if err != nil {
		return err
	}
	log = log.New("start", start, "end", end)

	if skipFetchlog && !publishAllUnverifiedSigs {
		log.Info("Skip fetching logs")
		return nil
	}

	// fetch event logs
	var logs []types.Log
	if !skipFetchlog {
		if logs, err = w.l1Signer.FilterLogsWithRateThottling(ctx, verse.NewEventLogFilter(start, end, []common.Address{task.RollupContract()})); err != nil {
			if errors.Is(err, ethutil.ErrTooManyRequests) {
				log.Warn("Rate limit exceeded", "err", err)
			}
			// Restore the next start block number if the fetching logs failed.
			blockRangeManager.RestoreNextStart()
			return fmt.Errorf("failed to fetch(start: %d, end: %d) event logs from hub-layer: %w", start, end, err)
		}
	}
	log = log.New("count-logs", len(logs))

	// verify the fetched logs
	var (
		opsigs []*database.OptimismSignature
		// flag at least one log verification failed.
		atLeastOneLogVerificationFailed bool
		// As the replica syncing is not real-time, the retry mechanism is required.
		retryBackoff = w.retryBackoff()
	)
	for i := range logs {
		var (
			row *database.OptimismSignature
			err error
		)
		for {
			row, err = w.verifyAndSaveLog(ctx, &logs[i], task, nextIndex.Uint64(), log)
			// break if the verification is successful or skip
			if err == nil {
				break
			}
			// exit if context have been canceled
			if errors.Is(err, context.Canceled) {
				return err
			}

			delay, remain, attempts := retryBackoff()
			// give up if the retry time limit is exceeded
			if remain <= 0 {
				break
			}
			// expand the deadline if the deadline is exceeded and retry immediately
			if errors.Is(err, context.DeadlineExceeded) {
				log.Info("expand the deadline", "log-index", i)
				cancel() // cancel previous context
				ctx, cancel = context.WithTimeout(parent, w.cfg.StateCollectTimeout*2)
				defer cancel()
				continue
			}

			// exponential backoff til max delay
			log.Warn("retry verification",
				"delay", delay, "remain", remain, "attempts", attempts, "err", err)
			time.Sleep(delay)
		}

		// verification failed
		if err != nil {
			// skip the log if the verification failed
			log.Error("Failed to verify a log", "log-index", i, "err", err)
			atLeastOneLogVerificationFailed = true
		}

		if row == nil {
			// skip if the row is nil
			// - when the event is not a rollup event
			// - when the event is already verified
			continue
		}

		opsigs = append(opsigs, row)
		log.Debug("Verification completed", "approved", row.Approved, "rollup-index", row.RollupIndex)

		// Make sure the first event rollup index is less than the next index, to prosess the next index correctly.
		if err := blockRangeManager.CheckIfStartTooLarge(nextIndex, row.RollupIndex); err != nil {
			return err
		}
	}
	if len(opsigs) > 0 {
		log.Info("Completed verification of all fetched logs", "count-newsigs", len(opsigs))
	}

	// Will publish all unverified signatures if the flag is set.
	if publishAllUnverifiedSigs {
		contract := task.RollupContract()
		rows, err := w.db.OPSignature.FindUnverifiedBySigner(w.l1Signer.Signer(), nextIndex.Uint64(), &contract, database.FindUnverifiedBySignerLimit)
		if err != nil {
			log.Error("Failed to find unverified signatures", "err", err)
		}
		opsigs = append(opsigs, rows...)
	}

	if len(opsigs) > 0 {
		// publish all signatures at once
		w.p2p.PublishSignatures(ctx, opsigs)
		log.Info("Published signatures", "count-sigs", len(opsigs), "first-rollup-index", opsigs[0].RollupIndex, "last-rollup-index", opsigs[len(opsigs)-1].RollupIndex)
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

func (w *Verifier) retryBackoff() func() (delay, remain time.Duration, attempts int) {
	started := time.Now()
	attempts := 0

	return func() (time.Duration, time.Duration, int) {
		// backoff delay: 0.1s, 0.8s, 6.4s, 51.2s, 409.6s(7m), 3276.8s(54m),
		delay := 100 << (3 * attempts) * time.Millisecond
		if delay <= 0 || delay > w.cfg.MaxRetryBackoff { // delay <= 0 is overflow
			delay = w.cfg.MaxRetryBackoff
		}

		remain := w.cfg.RetryTimeout - time.Since(started)
		if remain < 0 {
			remain = 0
		}

		attempts++
		return delay, remain, attempts
	}
}

// Updater for the `l1Latest` field.
func (w *Verifier) l1LatestUpdater() {
	tick := util.NewTicker(w.cfg.Interval, 1)
	defer tick.Stop()

	for range tick.C {
		ctx, cancel := context.WithTimeout(context.Background(), w.cfg.Interval/2)
		if h, err := w.l1Signer.HeaderByNumber(ctx, nil); err != nil {
			w.log.Warn("Failed to update L1 latest block", "err", err)
		} else {
			w.log.Info("Updated L1 latest block", "number", h.Number, "hash", h.Hash())
			w.l1Latest.Store(h)
		}
		cancel()
	}
}

package verifier

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
)

// Worker to verify rollups.
type Verifier struct {
	cfg      *config.Verifier
	db       *database.Database
	l1Signer ethutil.SignableClient
	p2p      P2P
	tasks    sync.Map
	log      log.Logger
	running  *sync.Map
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
	return &Verifier{
		cfg:      cfg,
		db:       db,
		p2p:      p2p,
		l1Signer: l1Signer,
		log:      log.New("worker", "verifier"),
		running:  &sync.Map{},
	}
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
		nextEventFetchStartBlock uint64
		counter                  int
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
			if err := w.work(ctx, task, chainId, &nextEventFetchStartBlock, publishAllUnverifiedSigs()); err != nil {
				w.log.Error("Failed to run verification", "err", err)
			}
		}
	}
}

func (w *Verifier) work(ctx context.Context, task verse.VerifiableVerse, chainId uint64, nextStart *uint64, publishAllUnverifiedSigs bool) error {
	// run verification tasks until time out
	var (
		cancel       context.CancelFunc
		ctxOrigin    = ctx
		setNextStart = func(endBlock uint64) {
			// Next start block is the current end block + 1
			endBlock += 1
			*nextStart = endBlock
		}
	)
	ctx, cancel = context.WithTimeout(ctx, w.cfg.StateCollectTimeout)
	defer cancel()

	// Assume the fetched nextIndex is not reorged, as we confirm `w.cfg.Confirmations` blocks
	nextIndex, err := task.NextIndexWithConfirm(&bind.CallOpts{Context: ctx}, uint64(w.cfg.Confirmations), true)
	if err != nil {
		return fmt.Errorf("failed to call the NextIndex method: %w", err)
	}
	log := log.New("chainId", chainId, "next-index", nextIndex)

	// Clean up old signatures
	if err := w.cleanOldSignatures(task.RollupContract(), nextIndex.Uint64()); err != nil {
		log.Warn("Failed to delete old signatures", "err", err)
	}

	// determine the start block number
	var (
		start        uint64
		skipFetchlog bool
		oneDayBlocks = uint64(5760)
	)
	end, err := w.l1Signer.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch the latest block number: %w", err)
	}
	if 0 < *nextStart {
		start = *nextStart
		if start > end {
			// Block number is not updated yet.
			skipFetchlog = true
		}
	} else {
		offset := w.cfg.StartBlockOffset
		if end < offset {
			start = 0
		} else {
			start = end - offset
		}
	}
	if start < end && oneDayBlocks < end-start {
		// If the range is too wide, divide it into one-day blocks.
		end = start + oneDayBlocks
	}
	log = log.New("start", start, "end", end)

	if skipFetchlog && !publishAllUnverifiedSigs {
		log.Info("Skip fetching logs")
		setNextStart(end)
		return nil
	}

	// fetch event logs
	var logs []types.Log
	if !skipFetchlog {
		if logs, err = w.l1Signer.FilterLogsWithRateThottling(ctx, verse.NewEventLogFilter(start, end, []common.Address{task.RollupContract()})); err != nil {
			if errors.Is(err, ethutil.ErrTooManyRequests) {
				log.Warn("Rate limit exceeded", "err", err)
			}
			return fmt.Errorf("failed to fetch(start: %d, end: %d) event logs from hub-layer: %w", start, end, err)
		}
	}
	log = log.New("count-logs", len(logs))

	// Only if succeed to fetch logs, update the next start block.
	setNextStart(end)

	// verify the fetched logs
	var (
		opsigs = []*database.OptimismSignature{}
		// flag at least one log verification failed
		atLeastOneLogVerificationFailed bool
	)
	for i := range logs {
		row, err := w.verifyAndSaveLog(ctx, &logs[i], task, nextIndex.Uint64(), log)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				// exit if context have been canceled
				return err
			} else if errors.Is(err, context.DeadlineExceeded) {
				// retry if the deadline is exceeded
				log.Warn("too much time spent on log iteration", "current-index", i)
				cancel()                                                     // cancel previous context
				ctx, cancel = context.WithTimeout(ctxOrigin, 30*time.Second) // expand the deadline
				defer cancel()
				row, err = w.verifyAndSaveLog(ctx, &logs[i], task, nextIndex.Uint64(), log)
				if err != nil {
					// give up if the retry fails
					log.Error("Failed to verification", "err", err, "rollup-index", nextIndex.Uint64())
					atLeastOneLogVerificationFailed = true
					continue
				}
			} else {
				// continue if other errors
				log.Error("Failed to verification", "err", err, "rollup-index", nextIndex.Uint64())
				atLeastOneLogVerificationFailed = true
				continue
			}
		}

		if row == nil {
			// skip if the row is nil
			// - when the event is not a rollup event
			// - when the event is already verified
			continue
		}

		opsigs = append(opsigs, row)
		log.Debug("Verification completed", "approved", row.Approved, "rollup-index", row.RollupIndex)
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
		// The removed task will be added again in the next verse discovery.
		// As the verse discovery interval is 1h, the faild log verification will be retried 1h later.
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

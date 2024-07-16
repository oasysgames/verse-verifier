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
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
)

// Worker to verify rollups.
type Verifier struct {
	cfg      *config.Verifier
	db       *database.Database
	l1Signer ethutil.SignableClient
	p2p      P2P
	topic    *util.Topic
	tasks    sync.Map
	log      log.Logger
	wg       *util.WorkerGroup
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
		topic:    util.NewTopic(),
		log:      log.New("worker", "verifier"),
		wg:       util.NewWorkerGroup(cfg.Concurrency),
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

func (w *Verifier) AddTask(task verse.VerifiableVerse) {
	task.Logger(w.log).Info("Add verifier task")
	w.tasks.Store(task.RollupContract(), task)
}

func (w *Verifier) RemoveTask(contract common.Address) {
	w.tasks.Delete(contract)
}

func (w *Verifier) SubscribeNewSignature(ctx context.Context) *SignatureSubscription {
	ch := make(chan *database.OptimismSignature)
	cancel := w.topic.Subscribe(ctx, func(ctx context.Context, data interface{}) {
		ch <- data.(*database.OptimismSignature)
	})
	return &SignatureSubscription{Cancel: cancel, ch: ch}
}

func (w *Verifier) AddVerse(ctx context.Context, v verse.VerifiableVerse, chainId uint64) {
	go w.startVerifier(ctx, v, chainId)
}

func (w *Verifier) startVerifier(ctx context.Context, v verse.VerifiableVerse, chainId uint64) {
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
			if err := w.work2(ctx, v, chainId, &nextEventFetchStartBlock, publishAllUnverifiedSigs()); err != nil {
				w.log.Error("Failed to run verification", "err", err)
			}
		}
	}
}

func (w *Verifier) Work(ctx context.Context) {
	w.tasks.Range(func(key, val interface{}) bool {
		workerID := key.(common.Address).Hex()
		task := val.(verse.VerifiableVerse)

		// deduplication
		if _, ok := w.running.Load(workerID); ok {
			return true
		}
		w.running.Store(workerID, 1)

		if !w.wg.Has(workerID) {
			worker := func(ctx context.Context, rname string, data interface{}) {
				defer w.running.Delete(rname)
				w.work(ctx, data.(verse.VerifiableVerse))
			}
			w.wg.AddWorker(ctx, workerID, worker)
		}

		w.wg.Enqueue(workerID, task)
		return true
	})
}

func (w *Verifier) work(ctx context.Context, task verse.VerifiableVerse) {
	log := task.Logger(w.log)

	// Assume the fetched nextIndex is not reorged, as we confirm `w.cfg.Confirmations` blocks
	nextIndex, err := task.NextIndexWithConfirm(&bind.CallOpts{Context: ctx}, uint64(w.cfg.Confirmations), true)
	if err != nil {
		log.Error("Failed to call the NextIndex method", "err", err)
		return
	}
	log = log.New("next-index", nextIndex)

	// Clean up old signatures
	if err := w.cleanOldSignatures(task.RollupContract(), nextIndex.Uint64()); err != nil {
		w.log.Warn("Failed to delete old signatures", "nextIndex", nextIndex, "err", err)
	}

	// verify the signature that match the nextIndex
	// and delete after signatures if there is a problem.
	// Prevent getting stuck indefinitely in the Verify waiting
	// event due to a bug in the signature creation process.
	w.deleteInvalidNextIndexSignature(task, nextIndex.Uint64())

	// run verification tasks until time out
	ctx, cancel := context.WithTimeout(ctx, w.cfg.StateCollectTimeout)
	defer cancel()

	for rollupIndex := nextIndex.Uint64(); ; rollupIndex++ {
		events, err := task.EventDB().FindForVerification(
			w.l1Signer.Signer(), task.RollupContract(), rollupIndex, 1)
		if err != nil {
			log.Error("Failed to find rollup events", "err", err)
			return
		} else if len(events) == 0 {
			log.Debug("Wait for new rollup event")
			return
		}

		log := log.New("rollup-index", events[0].GetRollupIndex())
		log.Info("Start verification")

		approved, err := task.Verify(log, ctx, events[0], w.cfg.StateCollectLimit)
		if err != nil {
			log.Error("Failed to verification", "err", err)
			return
		}

		msg := database.NewMessage(events[0], w.l1Signer.ChainID(), approved)
		sig, err := msg.Signature(w.l1Signer.SignData)
		if err != nil {
			log.Error("Failed to calculate signature", "err", err)
			return
		}

		row, err := w.db.OPSignature.Save(
			nil, nil,
			w.l1Signer.Signer(),
			events[0].GetContract().Address,
			events[0].GetRollupIndex(),
			events[0].GetRollupHash(),
			approved,
			sig)
		if err != nil {
			log.Error("Failed to save signature", "err", err)
			return
		}

		w.topic.Publish(row)
		log.Info("Verification completed", "approved", approved)
	}
}

func (w *Verifier) work2(ctx context.Context, task verse.VerifiableVerse, chainId uint64, nextStart *uint64, publishAllUnverifiedSigs bool) error {
	// run verification tasks until time out
	var (
		cancel    context.CancelFunc
		ctxOrigin = ctx
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
		// override the timeout to prevent the timeout error
		// Fetching logs with wide range may take a long time.
		var cancelLogTimeout context.CancelFunc
		ctx, cancelLogTimeout = context.WithTimeout(ctxOrigin, 300*time.Second)
		defer cancelLogTimeout()
	}

	if skipFetchlog && !publishAllUnverifiedSigs {
		log.Info("Skip fetching logs", "start", start, "end", end)
		return nil
	}

	// fetch event logs
	var logs []types.Log
	if !skipFetchlog {
		if logs, err = w.l1Signer.FilterLogs(ctx, verse.NewEventLogFilter(start, end, []common.Address{task.RollupContract()})); err != nil {
			return fmt.Errorf("failed to fetch(start: %d, end: %d) event logs from hub-layer: %w", start, end, err)
		}
	}

	// Next start block is the current end block + 1
	end += 1
	*nextStart = end

	// verify the fetched logs
	opsigs := []*database.OptimismSignature{}
	for i := range logs {
		event, err := verse.ParseEventLog(&logs[i])
		if err != nil {
			log.Warn("Failed to parse event log", "block", logs[i].BlockNumber, "contract", logs[i].Address.Hex(), "err", err)
			continue
		}

		// parse event log
		rollupEvent, ok := event.(*verse.RollupedEvent)
		if !ok {
			// skip `*verse.DeletedEvent` or `*verse.VerifiedEvent`
			continue
		}

		contract, err := w.db.OPContract.FindOrCreate(logs[i].Address)
		if err != nil {
			log.Warn("Failed to find or create contract", "err", err)
			continue
		}

		var dbEvent database.OPEvent
		switch t := rollupEvent.Parsed.(type) {
		case *scc.SccStateBatchAppended:
			var model database.OptimismState
			model.BatchIndex = t.BatchIndex.Uint64()
			model.PrevTotalElements = t.PrevTotalElements.Uint64()
			model.BatchSize = t.BatchSize.Uint64()
			model.BatchRoot = t.BatchRoot
			model.ContractID = contract.ID
			model.Contract = *contract
			dbEvent = &model
		case *l2oo.OasysL2OutputOracleOutputProposed:
			var model database.OpstackProposal
			model.L2OutputIndex = t.L2OutputIndex.Uint64()
			model.OutputRoot = t.OutputRoot
			model.L2BlockNumber = t.L2BlockNumber.Uint64()
			model.L1Timestamp = t.L1Timestamp.Uint64()
			model.ContractID = contract.ID
			model.Contract = *contract
			dbEvent = &model
		default:
			return fmt.Errorf("unsupported event type: %T", t)
		}

		if dbEvent.GetRollupIndex() < nextIndex.Uint64() {
			// skip old events
			continue
		}

		approved, err := task.Verify(log, ctx, dbEvent, w.cfg.StateCollectLimit)
		if err != nil {
			log.Error("Failed to verification", "err", err, "rollup-index", dbEvent.GetRollupIndex())
			continue
		}

		msg := database.NewMessage(dbEvent, w.l1Signer.ChainID(), approved)
		sig, err := msg.Signature(w.l1Signer.SignData)
		if err != nil {
			log.Error("Failed to calculate signature", "err", err, "rollup-index", dbEvent.GetRollupIndex())
			continue
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
			log.Error("Failed to save signature", "err", err, "rollup-index", dbEvent.GetRollupIndex())
			continue
		}

		opsigs = append(opsigs, row)
		log.Debug("Verification completed", "approved", approved, "rollup-index", dbEvent.GetRollupIndex())
	}
	if len(opsigs) > 0 {
		log.Info("Completed verification of all fetched logs", "count-logs", len(logs), "count-newsigs", len(opsigs), "start", start, "end", end)
	}

	// Will publish all unverified signatures if the flag is set.
	if publishAllUnverifiedSigs {
		contract := task.RollupContract()
		rows, err := w.db.OPSignature.FindUnverifiedBySigner(w.l1Signer.Signer(), nextIndex.Uint64(), &contract)
		if err != nil {
			log.Error("Failed to find unverified signatures", "err", err)
		}
		opsigs = append(opsigs, rows...)
	}

	if len(opsigs) > 0 {
		// publish all signatures at once
		w.p2p.PublishSignatures(ctx, opsigs)
	} else {
		log.Info("No signatures to publish", "count-logs", len(logs), "start", start, "end", end)
	}

	return nil
}

func (w *Verifier) deleteInvalidNextIndexSignature(task verse.VerifiableVerse, nextIndex uint64) {
	log := task.Logger(w.log).New("next-index", nextIndex)

	signer := w.l1Signer.Signer()
	contract := task.RollupContract()
	sigs, err := w.db.OPSignature.Find(nil, &signer, &contract, &nextIndex, 1, 0)
	if err != nil {
		log.Error("Unable to find signatures", "err", err)
		return
	} else if len(sigs) == 0 {
		log.Debug("No invalid signature")
		return
	}

	event, err := task.EventDB().FindByRollupIndex(contract, nextIndex)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Debug("No rollup event")
		} else {
			log.Error("Unable to find rollup event", "err", err)
		}
		return
	}

	err = database.NewMessage(event, w.l1Signer.ChainID(), true).
		VerifySigner(sigs[0].Signature[:], signer)
	if _, ok := err.(*ethutil.SignerMismatchError); ok {
		// possible reject signature
		err = database.NewMessage(event, w.l1Signer.ChainID(), false).
			VerifySigner(sigs[0].Signature[:], signer)
	}
	if err == nil {
		log.Debug("No invalid signature")
		return
	}

	log.Warn("Found invalid signature", "signature", sigs[0].Signature.Hex())

	rows, err := w.db.OPSignature.Deletes(signer, contract, nextIndex)
	if err != nil {
		log.Error("Failed to delete invalid signatures", "err", err)
	} else {
		log.Warn("Deleted invalid signatures", "rows", rows)
	}
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
	if _, err := w.db.OPSignature.DeleteOlds(contract, deleteIndex); err != nil {
		return fmt.Errorf("failed to delete old signatures. deleteIndex: %d, : %w", deleteIndex, err)
	}
	return nil
}

type SignatureSubscription struct {
	Cancel context.CancelFunc
	ch     chan *database.OptimismSignature
}

func (s *SignatureSubscription) Next() <-chan *database.OptimismSignature {
	return s.ch
}

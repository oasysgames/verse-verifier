package verifier

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
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
	topic    *util.Topic
	tasks    sync.Map
	log      log.Logger
}

// Returns the new verifier.
func NewVerifier(
	cfg *config.Verifier,
	db *database.Database,
	l1Signer ethutil.SignableClient,
) *Verifier {
	return &Verifier{
		cfg:      cfg,
		db:       db,
		l1Signer: l1Signer,
		topic:    util.NewTopic(),
		log:      log.New("worker", "verifier"),
	}
}

// Start verifier.
func (w *Verifier) Start(ctx context.Context) {
	w.log.Info(
		"Worker started",
		"signer", w.l1Signer.Signer(),
		"interval", w.cfg.Interval,
		"state-collect-limit", w.cfg.StateCollectLimit,
		"concurrency", w.cfg.Concurrency,
	)

	wg := util.NewWorkerGroup(w.cfg.Concurrency)
	running := &sync.Map{}

	tick := time.NewTicker(w.cfg.Interval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Worker stopped")
			return
		case <-tick.C:
			w.tasks.Range(func(key, val interface{}) bool {
				workerID := key.(common.Address).Hex()
				task := val.(verse.VerifiableVerse)

				// deduplication
				if _, ok := running.Load(workerID); ok {
					return true
				}
				running.Store(workerID, 1)

				if !wg.Has(workerID) {
					worker := func(ctx context.Context, rname string, data interface{}) {
						defer running.Delete(rname)
						w.work(ctx, data.(verse.VerifiableVerse))
					}
					wg.AddWorker(ctx, workerID, worker)
				}

				wg.Enqueue(workerID, task)
				return true
			})
		}
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

func (w *Verifier) work(ctx context.Context, task verse.VerifiableVerse) {
	log := task.Logger(w.log)

	// fetch the next index from hub-layer
	nextIndex, err := task.NextIndex(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Error("Failed to call the NextIndex method", "err", err)
		return
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
		log := log.New("rollup-index", rollupIndex)

		events, err := task.EventDB().FindForVerification(
			w.l1Signer.Signer(), task.RollupContract(), rollupIndex, 1)
		if err != nil {
			log.Error("Failed to find rollup events", "err", err)
			return
		} else if len(events) == 0 {
			log.Debug("Wait for new rollup event")
			return
		}

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

type SignatureSubscription struct {
	Cancel context.CancelFunc
	ch     chan *database.OptimismSignature
}

func (s *SignatureSubscription) Next() <-chan *database.OptimismSignature {
	return s.ch
}

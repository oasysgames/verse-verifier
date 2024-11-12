package submitter

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/multicall2"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
	"golang.org/x/net/context"
)

const (
	maxTxSize = 120 * 1024 // 120KB ()
	minTxGas  = 24871      // Multicall minimum gas

	// Parameters for worker pool.
	maxIdleWorkerDuration      = time.Minute
	workerReleaseCheckInterval = time.Second
	workerReleaseCheckTimeout  = time.Duration(0) // no timeout
)

var (
	ErrNoSignatures    = errors.New("no signatures")
	ErrAlreadyVerified = errors.New("already verified")
)

type Submitter struct {
	// fields passed during construction
	cfg          *config.Submitter
	db           *database.Database
	l1SignerFn   L1SignerFn
	stakemanager *stakemanager.Cache
	versepool    verse.VersePool
	log          log.Logger

	// internal fields
	tasks util.SyncMap[common.Address, *taskT]
}

type L1SignerFn func(chainID uint64) ethutil.SignableClient

type taskT struct {
	verse         verse.TransactableVerse
	verifiedIndex *uint64
}

func NewSubmitter(
	cfg *config.Submitter,
	db *database.Database,
	l1SignerFn L1SignerFn,
	stakemanager *stakemanager.Cache,
	versepool verse.VersePool,
) *Submitter {
	return &Submitter{
		cfg:          cfg,
		db:           db,
		l1SignerFn:   l1SignerFn,
		stakemanager: stakemanager,
		versepool:    versepool,
		log:          log.New("worker", "submitter"),
	}
}

func (w *Submitter) Start(ctx context.Context) {
	w.log.Info("Submitter started", "config", w.cfg)

	// Create woker pool.
	wp := util.NewWorkerPool(w.log, w.work, w.cfg.MaxWorkers,
		maxIdleWorkerDuration, workerReleaseCheckInterval, workerReleaseCheckTimeout)
	wp.Start()
	defer wp.Stop()

	// Manage running tasks to prevent dups.
	var running util.SyncMap[common.Address, time.Time]

	// Every hour, remove Verse that were deleted from the pool from the local cache as well.
	cacheCleanupTick := time.NewTicker(time.Hour)
	defer cacheCleanupTick.Stop()

	workTick := time.NewTicker(w.cfg.Interval)
	defer workTick.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Info("Submitter stopped")
			return
		case <-cacheCleanupTick.C:
			w.tasks.Range(func(cacheKey common.Address, task *taskT) bool {
				_, exists := w.versepool.Get(cacheKey)
				if !exists {
					task.verse.Logger(w.log).Info("Delete task cache")
					w.tasks.Delete(cacheKey)
				}
				return true
			})
		case <-workTick.C:
			w.versepool.Range(func(item *verse.VersePoolItem) bool {
				if !item.CanSubmit() {
					return true
				}

				log := w.log.New("chain-id", item.Verse().ChainID())

				// Task has internal state and should be cached
				cacheKey := item.Verse().RollupContract()

				// Skip if previous task is running
				if started, isRunning := running.Load(cacheKey); isRunning {
					log.Info("Skip", "elapsed", time.Since(started))
					return true
				}
				running.Store(cacheKey, time.Now())

				// Since worker run asynchronously, should not call `running.Delete()` here.
				release := func() { running.Delete(cacheKey) }

				// If the cache does not exist, create a new task.
				var task *taskT
				if cache, ok := w.tasks.Load(cacheKey); ok {
					task = cache
				} else {
					l1Signer := w.l1SignerFn(item.Verse().ChainID())
					if l1Signer == nil {
						log.Error("Submitter wallet was not found")
					} else {
						task = &taskT{
							verse: item.Verse().WithTransactable(
								l1Signer, item.Verse().VerifyContract()),
							verifiedIndex: nil, // initial value should be nil
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
}

func (w *Submitter) work(ctx context.Context, task *taskT) {
	nextIndex, err := w.submit(ctx, task)

	log := task.verse.Logger(w.log).New("next-index", nextIndex)
	if task.verifiedIndex != nil {
		log = log.New("verified-index", *task.verifiedIndex)
	}

	if errors.Is(err, verse.ErrNotSufficientConfirmations) {
		log.Info("Not enough confirmations")
	} else if errors.Is(err, ErrNoSignatures) {
		log.Info("No signatures to submit")
	} else if errors.Is(err, ErrAlreadyVerified) {
		// Skip if the nextIndex is already verified
		log.Info("Already verified the rollup index")
	} else if errors.Is(err, &StakeAmountShortage{}) {
		// Wait until enough signatures are collected
		var (
			shortageErr      *StakeAmountShortage
			required, actual *big.Int
		)
		if errors.As(err, &shortageErr) {
			required, actual = fromWei(shortageErr.required), fromWei(shortageErr.actual)
		}
		log.Info("Not enough signatures(stake amount shortage)",
			"nextIndex", nextIndex, "required", required, "actual", actual)
	} else if err != nil {
		log.Error("Failed to verify the rollup index", "err", err)
	} else {
		// Finally, succeeded to verify the corresponding rollup index, So move to the next index
		task.verifiedIndex = &nextIndex
		log.Info("Successfully verified the rollup index", "next-verified-index", *task.verifiedIndex)
		if err := w.cleanupOldSignatures(task.verse.RollupContract(), *task.verifiedIndex); err != nil {
			log.Warn("Failed to delete old signatures", "verified-index", *task.verifiedIndex, "err", err)
		}
	}
}

func (w *Submitter) cleanupOldSignatures(contract common.Address, verifiedIndex uint64) error {
	if verifiedIndex == 0 {
		return nil
	}
	// Just keep the last verified index
	deleteIndex := uint64(verifiedIndex - 1)
	if _, err := w.db.OPSignature.DeleteOlds(contract, deleteIndex, database.DeleteOldsLimit); err != nil {
		return fmt.Errorf("failed to delete old signatures. deleteIndex: %d, : %w", deleteIndex, err)
	}
	return nil
}

func (w *Submitter) RemoveTask(contract common.Address) {
	w.tasks.Delete(contract)
}

func (w *Submitter) submit(ctx context.Context, task *taskT) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()

	nextIndex, err := w.versepool.NextIndex(ctx, task.verse.RollupContract(), w.cfg.Confirmations, false)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch next index: %w", err)
	}
	log := task.verse.Logger(w.log).New("next-index", nextIndex)

	if task.verifiedIndex != nil {
		if *task.verifiedIndex == nextIndex {
			// Skip if the nextIndex is already verified
			return nextIndex, ErrAlreadyVerified
		} else if *task.verifiedIndex > nextIndex {
			// Continue as purhaps reorged
			log.Warn("Possible reorged. next index is smaller than the verified index",
				"verified-index", *task.verifiedIndex, "next-index", nextIndex)
		}
	}

	iter := &signatureIterator{
		db:           w.db,
		stakemanager: w.stakemanager,
		contract:     task.verse.RollupContract(),
		rollupIndex:  nextIndex,
	}

	var tx *types.Transaction
	if w.cfg.UseMulticall {
		tx, err = w.sendMulticallTx(log, ctx, task.verse, iter)
	} else {
		tx, err = w.sendNormalTx(log, ctx, task.verse, iter)
	}
	if err != nil {
		log.Debug(err.Error())
		return nextIndex, fmt.Errorf("failed to send transaction: %w", err)
	}

	if err = w.waitForReceipt(ctx, task.verse.L1Signer(), tx); err != nil {
		return nextIndex, fmt.Errorf("failed to wait for receipt: %w", err)
	}

	return nextIndex, nil
}

func (w *Submitter) sendNormalTx(
	log log.Logger,
	ctx context.Context,
	task verse.TransactableVerse,
	iter *signatureIterator,
) (*types.Transaction, error) {
	rows, err := iter.next(ctx)
	if err != nil {
		log.Error("Failed to find signatures", "err", err)
		return nil, err
	} else if len(rows) == 0 {
		log.Debug("No signatures")
		return nil, ErrNoSignatures
	}

	opts := task.L1Signer().TransactOpts(ctx)

	// call estimateGas
	opts.NoSend = true
	tx, err := task.Transact(opts, rows[0].RollupIndex, rows[0].Approved, extSignatureBytes(rows))
	if err != nil {
		log.Error("Failed to estimate gas", "err", err)
		return nil, err
	}

	// send transaction
	opts.NoSend = false
	opts.GasLimit = w.cfg.MultiplyGas(tx.Gas())
	if err := task.L1Signer().SendTransaction(ctx, tx); err != nil {
		log.Error("Failed to send verify transaction", "err", err)
		return nil, err
	}

	log.Info(
		"Sent transaction",
		"tx", tx.Hash().Hex(),
		"nonce", tx.Nonce(),
		"gas-limit", tx.Gas(),
		"gas-fee", tx.GasFeeCap(),
		"gas-tip", tx.GasTipCap(),
	)
	return tx, nil
}

func (w *Submitter) sendMulticallTx(
	log log.Logger,
	ctx context.Context,
	task verse.TransactableVerse,
	iter *signatureIterator,
) (*types.Transaction, error) {
	mcall, err := multicall2.NewMulticall2(
		common.HexToAddress(w.cfg.MulticallAddress), task.L1Signer())
	if err != nil {
		log.Error("Failed to construct the multicall contract", "err", err)
		return nil, err
	}

	opts := &bind.TransactOpts{
		Context:  ctx,
		NoSend:   true,
		Nonce:    common.Big1, // prevent `eth_getNonce`
		GasPrice: common.Big1, // prevent `eth_gasPrice`
		GasLimit: 21_000,      // prevent `eth_estimateGas`
		From:     task.L1Signer().Signer(),
		Signer: func(a common.Address, rawTx *types.Transaction) (*types.Transaction, error) {
			return rawTx, nil
		},
	}

	var (
		calls       []multicall2.Multicall2Call
		errShortage error
	)
	for i := 0; i < w.cfg.BatchSize; i++ {
		rows, err := iter.next(ctx)
		if _, ok := err.(*StakeAmountShortage); ok {
			errShortage = err
			break
		} else if err != nil {
			log.Debug("Failed to find signatures", "err", err)
			return nil, err
		} else if len(rows) == 0 {
			break
		}

		// build transaction (without sending).
		rawTx, err := task.Transact(opts, rows[0].RollupIndex, rows[0].Approved, extSignatureBytes(rows))
		if err != nil {
			log.Error("Failed to create verify transaction", "err", err)
			return nil, err
		}

		call := multicall2.Multicall2Call{
			Target:   task.VerifyContract(),
			CallData: rawTx.Data(),
		}
		rawTx, err = mcall.TryAggregate(opts, true, append(calls, call))
		if err != nil {
			log.Error("Failed to create multicall transaction", "err", err)
			return nil, err
		} else if len(rawTx.Data()) > maxTxSize {
			log.Warn("Oversized", "data-size", len(rawTx.Data()), "call-size", i+1)
			break
		}

		calls = append(calls, call)

		// if rejected, there is no need to approve any subsequent rollups.
		if !rows[0].Approved {
			break
		}
	}
	if len(calls) == 0 {
		if errShortage != nil {
			log.Debug("No calldata", "err", errShortage)
			return nil, errShortage
		}
		log.Debug("No calldata")
		return nil, ErrNoSignatures
	}

	// call estimateGas
	opts = task.L1Signer().TransactOpts(ctx)
	opts.NoSend = true
	tx, err := mcall.TryAggregate(opts, true, calls)
	if err != nil {
		log.Error("Failed to estimate gas", "err", err)
		return nil, err
	}

	// to fit max gas
	if tx.Gas() > w.cfg.MaxGas {
		gasPerCall := (tx.Gas() - minTxGas) / uint64(len(calls))
		end := uint64(len(calls))
		for ; end > 1 && end*gasPerCall > w.cfg.MaxGas; end-- {
		}
		calls = calls[:end]

		// re estimateGas
		tx, err = mcall.TryAggregate(opts, true, calls)
		if err != nil {
			log.Error("Failed to re-estimate gas", "err", err)
			return nil, err
		}
	}

	// send transaction
	opts.NoSend = false
	opts.GasLimit = w.cfg.MultiplyGas(tx.Gas())
	tx, err = mcall.TryAggregate(opts, true, calls)
	if err != nil {
		log.Error("Failed to send multicall verify transaction", "err", err)
		return nil, err
	}

	log.Info(
		"Sent transaction",
		"call-size", len(calls),
		"tx", tx.Hash().Hex(),
		"nonce", tx.Nonce(),
		"gas-limit", tx.Gas(),
		"gas-fee", tx.GasFeeCap(),
		"gas-tip", tx.GasTipCap(),
	)
	return tx, nil
}

func (w *Submitter) waitForReceipt(
	ctx context.Context,
	l1Client ethutil.SignableClient,
	tx *types.Transaction,
) error {
	// wait for block to be validated
	receipt, err := bind.WaitMined(ctx, l1Client, tx)
	if err != nil {
		return fmt.Errorf("failed to receive receipt. tx: %s, : %w", tx.Hash().Hex(), err)
	}
	if receipt.Status != 1 {
		return fmt.Errorf("transaction reverted. tx: %s", tx.Hash().Hex())
	}
	return nil
}

func fromWei(wei *big.Int) *big.Int {
	return new(big.Int).Div(wei, big.NewInt(params.Ether))
}

func extSignatureBytes(rows []*database.OptimismSignature) [][]byte {
	bytes := make([][]byte, len(rows))
	for i, row := range rows {
		bytes[i] = row.Signature[:]
	}
	return bytes
}

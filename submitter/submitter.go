package submitter

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
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
)

var (
	ErrNoSignatures    = errors.New("no signatures")
	ErrAlreadyVerified = errors.New("already verified")
)

type Submitter struct {
	cfg          *config.Submitter
	db           *database.Database
	stakemanager *stakemanager.Cache
	tasks        sync.Map
	log          log.Logger
}

func NewSubmitter(
	cfg *config.Submitter,
	db *database.Database,
	stakemanager *stakemanager.Cache,
) *Submitter {
	return &Submitter{
		cfg:          cfg,
		db:           db,
		stakemanager: stakemanager,
		log:          log.New("worker", "submitter"),
	}
}

// Deprecated
func (w *Submitter) Start(ctx context.Context) {
	w.log.Info("Submitter started",
		"interval", w.cfg.Interval,
		"concurrency", w.cfg.Concurrency,
		"confirmations", w.cfg.Confirmations,
		"gas-multiplier", w.cfg.GasMultiplier,
		"batch-size", w.cfg.BatchSize,
		"max-gas", w.cfg.MaxGas,
		"scc-verifier", w.cfg.SCCVerifierAddress,
		"l2oo-verifier", w.cfg.L2OOVerifierAddress,
		"use-multicall", w.cfg.UseMulticall,
		"multicall", w.cfg.MulticallAddress)
	w.workLoop(ctx)
}

func (w *Submitter) AddVerse(ctx context.Context, v verse.TransactableVerse, chainId uint64) {
	// Start submitting loop
	// 1. Request signatures every interval
	// 2. Submit verify tx if enough signatures are collected
	go w.startSubmitter(ctx, v, chainId)
}

func (w *Submitter) startSubmitter(ctx context.Context, v verse.TransactableVerse, chainId uint64) {
	var (
		tick          = time.NewTicker(w.cfg.Interval)
		duration      = w.cfg.Interval
		verifiedIndex *uint64
		resetDuration = func(target time.Duration) {
			if duration == target {
				return
			}
			duration = target
			tick.Reset(duration)
		}
	)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Submitting work stopped", "chainId", chainId)
			return
		case <-tick.C:
			nextIndex, err := w.work(ctx, v, verifiedIndex)
			if errors.Is(err, verse.ErrNotSufficientConfirmations) {
				w.log.Info("Not enough confirmations", "nextIndex", nextIndex, "chainId", chainId)
				continue
			} else if errors.Is(err, ErrNoSignatures) {
				w.log.Info("No signatures to submit", "nextIndex", nextIndex, "chainId", chainId)
				// Reset the ticker to the original interval
				resetDuration(w.cfg.Interval)
				continue
			} else if err != nil && strings.Contains(err.Error(), "stake amount shortage") {
				// Wait until enough signatures are collected
				w.log.Info("Waiting for enough signatures", "nextIndex", nextIndex, "chainId", chainId)
				// Reset the ticker to shorten the interval to be able to submit verify tx without waiting for the next interval
				resetDuration(w.cfg.Interval / 10)
				continue
			} else if err == nil {
				// Finally, succeeded to verify the corresponding rollup index, So move to the next index
				verifiedIndex = &nextIndex
				w.log.Info("Successfully verified the rollup index", "verifiedIndex", *verifiedIndex, "chainId", chainId)
				// Clean up old signatures
				if err := w.cleanOldSignatures(v.RollupContract(), *verifiedIndex); err != nil {
					w.log.Warn("Failed to delete old signatures", "verifiedIndex", *verifiedIndex, "chainId", chainId, "err", err)
				}
				// Reset the ticker to the original interval
				resetDuration(w.cfg.Interval)
				continue
			} else if errors.Is(err, ErrAlreadyVerified) {
				// Skip if the nextIndex is already verified
				w.log.Info("Already verified the rollup index", "nextIndex", nextIndex, "chainId", chainId)
				continue
			} else {
				w.log.Error("Failed to verify the rollup index", "nextIndex", nextIndex, "chainId", chainId, "err", err)
			}
		}
	}
}

func (w *Submitter) cleanOldSignatures(contract common.Address, verifiedIndex uint64) error {
	if verifiedIndex == 0 {
		return nil
	}
	// Just keep the last verified index
	deleteIndex := uint64(verifiedIndex - 1)
	if _, err := w.db.OPSignature.DeleteOlds(contract, deleteIndex); err != nil {
		return fmt.Errorf("failed to delete old signatures. deleteIndex: %d, : %w", deleteIndex, err)
	}
	return nil
}

// Deprecated
func (w *Submitter) workLoop(ctx context.Context) {
	wg := util.NewWorkerGroup(w.cfg.Concurrency)
	running := &sync.Map{}

	tick := time.NewTicker(w.cfg.Interval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Submitter stopped")
			return
		case <-tick.C:
			w.tasks.Range(func(key, val any) bool {
				workerID := key.(common.Address).Hex()
				task := val.(verse.TransactableVerse)

				// deduplication
				if _, ok := running.Load(workerID); ok {
					return true
				}
				running.Store(workerID, 1)

				if !wg.Has(workerID) {
					worker := func(ctx context.Context, rname string, data interface{}) {
						defer running.Delete(rname)
						w.work(ctx, data.(verse.TransactableVerse), nil)
					}
					wg.AddWorker(ctx, workerID, worker)
				}

				wg.Enqueue(workerID, task)
				return true
			})
		}
	}
}

func (w *Submitter) HasTask(contract common.Address) bool {
	_, ok := w.tasks.Load(contract)
	return ok
}

func (w *Submitter) AddTask(task verse.TransactableVerse) {
	task.Logger(w.log).Info("Add submitter task")
	w.tasks.Store(task.RollupContract(), task)
}

func (w *Submitter) RemoveTask(contract common.Address) {
	w.tasks.Delete(contract)
}

func (w *Submitter) work(ctx context.Context, task verse.TransactableVerse, verifiedIndex *uint64) (uint64, error) {
	log := task.Logger(w.log)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()

	// Assume the fetched nextIndex is not reorged, as we confirm `w.cfg.Confirmations` blocks
	nextIndex, err := task.NextIndexWithConfirm(&bind.CallOpts{Context: ctx}, uint64(w.cfg.Confirmations), false)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch next index: %w", err)
	}
	log = log.New("next-index", nextIndex)

	if verifiedIndex != nil {
		if *verifiedIndex == nextIndex.Uint64() {
			// Skip if the nextIndex is already verified
			return nextIndex.Uint64(), ErrAlreadyVerified
		} else if *verifiedIndex > nextIndex.Uint64() {
			// Continue as purhaps reorged
			log.Warn("Possible reorged. next index is smaller than the verified index", "verified-index", *verifiedIndex, "next-index", nextIndex.Uint64())
		}
	}

	iter := &signatureIterator{
		db:           w.db,
		stakemanager: w.stakemanager,
		contract:     task.RollupContract(),
		rollupIndex:  nextIndex.Uint64(),
	}

	var tx *types.Transaction
	if w.cfg.UseMulticall {
		tx, err = w.sendMulticallTx(log, ctx, task, iter)
	} else {
		tx, err = w.sendNormalTx(log, ctx, task, iter)
	}

	if err != nil {
		log.Debug(err.Error())
		return nextIndex.Uint64(), fmt.Errorf("failed to send transaction: %w", err)
	} else if tx != nil {
		// Skip waiting for confirmation. if sent tx is failed, it will be re-sent in the next interval.
		// w.waitForConfirmation(log.New("tx", tx.Hash()), ctx, task.L1Signer(), tx)
		w.waitForReceipt(log.New("tx", tx.Hash()), ctx, task.L1Signer(), tx)
	}

	return nextIndex.Uint64(), nil
}

func (w *Submitter) sendNormalTx(
	log log.Logger,
	ctx context.Context,
	task verse.TransactableVerse,
	iter *signatureIterator,
) (*types.Transaction, error) {
	rows, err := iter.next()
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
		rows, err := iter.next()
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

func (w *Submitter) waitForConfirmation(
	log log.Logger,
	ctx context.Context,
	l1Client ethutil.SignableClient,
	tx *types.Transaction,
) {
	// wait for block to be validated
	receipt, err := bind.WaitMined(ctx, l1Client, tx)
	if err != nil {
		log.Error("Failed to receive receipt", "err", err)
		return
	}
	if receipt.Status != 1 {
		log.Error("Transaction reverted")
		return
	}

	// wait for confirmations
	confirmed := map[common.Hash]bool{receipt.BlockHash: true}
	for {
		remaining := w.cfg.Confirmations - len(confirmed)
		if remaining <= 0 {
			log.Info("Transaction succeeded")
			return
		}

		log.Info("Wait for confirmation", "remaining", remaining)
		time.Sleep(time.Second)

		h, err := l1Client.HeaderByNumber(ctx, nil)
		if err != nil {
			log.Error("Failed to fetch block header", "err", err)
			continue
		}
		confirmed[h.Hash()] = true
	}
}

func (w *Submitter) waitForReceipt(
	log log.Logger,
	ctx context.Context,
	l1Client ethutil.SignableClient,
	tx *types.Transaction,
) {
	// wait for block to be validated
	receipt, err := bind.WaitMined(ctx, l1Client, tx)
	if err != nil {
		log.Error("Failed to receive receipt", "err", err)
		return
	}
	if receipt.Status != 1 {
		log.Error("Transaction reverted")
		return
	}
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

package submitter

import (
	"math/big"
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

func (w *Submitter) Start(ctx context.Context) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		w.workLoop(ctx)
	}()

	w.log.Info("Worker started",
		"interval", w.cfg.Interval,
		"concurrency", w.cfg.Concurrency,
		"confirmations", w.cfg.Confirmations,
		"gas-multiplier", w.cfg.GasMultiplier,
		"batch-size", w.cfg.BatchSize,
		"max-gas", w.cfg.MaxGas,
		"scc-verifier", w.cfg.SCCVerifierAddress,
		"l2oo-verifier", w.cfg.L2OOVerifierAddress,
		"use-multicall", w.cfg.UseMulticall,
		"multicall", w.cfg.Multicall2Address)

	wg.Wait()
	w.log.Info("Worker stopped")
}

func (w *Submitter) workLoop(ctx context.Context) {
	wg := util.NewWorkerGroup(w.cfg.Concurrency)
	running := &sync.Map{}

	tick := time.NewTicker(w.cfg.Interval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
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
						w.work(ctx, data.(verse.TransactableVerse))
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

func (w *Submitter) work(ctx context.Context, task verse.TransactableVerse) {
	log := task.Logger(w.log)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()

	// fetch the next index from hub-layer
	nextIndex, err := task.NextIndex(&bind.CallOpts{Context: ctx})
	if err != nil {
		log.Error("Failed to get next index", "err", err)
		return
	}
	log = log.New("next-index", nextIndex)

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
		log.Error(err.Error())
	} else if tx != nil {
		w.waitForCconfirmation(log.New("tx", tx.Hash()), ctx, task.L1Signer(), tx)
	}
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
		return nil, nil
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
		common.HexToAddress(w.cfg.Multicall2Address), task.L1Signer())
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

	var calls []multicall2.Multicall2Call
	for i := 0; i < w.cfg.BatchSize; i++ {
		rows, err := iter.next()
		if _, ok := err.(*StakeAmountShortage); ok {
			break
		} else if err != nil {
			log.Error("Failed to find signatures", "err", err)
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
		log.Info("No calldata")
		return nil, nil
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

func (w *Submitter) waitForCconfirmation(
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

package hublayer

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/scc"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/sccverifier"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"golang.org/x/net/context"
	"golang.org/x/sync/semaphore"
)

var (
	errNoSignature = errors.New("no signatures")
	minStake       = new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10_000_000))
)

type stakeManager interface {
	GetTotalStake(callOpts *bind.CallOpts, epoch *big.Int) (*big.Int, error)

	GetValidators(callOpts *bind.CallOpts, epoch, cursol, howMany *big.Int) (struct {
		Owners     []common.Address
		Operators  []common.Address
		Stakes     []*big.Int
		Candidates []bool
		NewCursor  *big.Int
	}, error)
}

type submitTask struct {
	scc common.Address
	hub ethutil.WritableClient
}

type SccSubmitter struct {
	db            *database.Database
	sm            stakeManager
	sccvAddr      common.Address
	interval      time.Duration
	concurrency   int
	confirmations int
	gasMultiplier float64

	sem    *semaphore.Weighted
	hubs   *sync.Map
	stakes *sync.Map
	log    log.Logger
}

func NewSccSubmitter(
	db *database.Database,
	sm stakeManager,
	sccvAddr common.Address,
	interval time.Duration,
	concurrency int,
	confirmations int,
	gasMultiplier float64,
) *SccSubmitter {
	return &SccSubmitter{
		db:            db,
		sm:            sm,
		sccvAddr:      sccvAddr,
		interval:      interval,
		concurrency:   concurrency,
		confirmations: confirmations,
		gasMultiplier: gasMultiplier,
		sem:           semaphore.NewWeighted(int64(concurrency)),
		hubs:          &sync.Map{},
		stakes:        &sync.Map{},
		log:           log.New("worker", "scc-submitter"),
	}
}

func (w *SccSubmitter) Start(ctx context.Context) {
	wg := &sync.WaitGroup{}
	queue := make(chan *submitTask)

	wg.Add(1)
	go func() {
		defer wg.Done()
		w.stakeRefreshLoop(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		w.workLoop(ctx, queue)
	}()

	w.log.Info("Worker started", "sccv", w.sccvAddr,
		"interval", w.interval, "concurrency", w.concurrency, "gas-multiplier", w.gasMultiplier)

	wg.Wait()
	w.log.Info("Worker stopped")
}

func (w *SccSubmitter) stakeRefreshLoop(ctx context.Context) {
	w.refreshStakes(ctx)

	tick := time.NewTicker(time.Hour)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			w.refreshStakes(ctx)
		}
	}
}

func (w *SccSubmitter) workLoop(ctx context.Context, queue chan<- *submitTask) {
	wg := util.NewWorkerGroup(w.concurrency)
	running := &sync.Map{}

	tick := time.NewTicker(w.interval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			w.hubs.Range(func(key, value any) bool {
				scc, ok0 := key.(common.Address)
				hub, ok1 := value.(ethutil.WritableClient)
				if !(ok0 && ok1) {
					return true
				}

				// deduplication
				name := scc.Hex()
				if _, ok := running.Load(name); ok {
					return true
				}
				running.Store(name, 1)

				if !wg.Has(name) {
					handler := func(ctx context.Context, rname string, data interface{}) {
						defer running.Delete(rname)

						if task, ok := data.(*submitTask); ok {
							w.work(ctx, task)
						}
					}
					wg.AddWorker(ctx, name, handler)
				}

				wg.Enqueue(name, &submitTask{scc: scc, hub: hub})
				return true
			})
		}
	}
}

func (w *SccSubmitter) AddVerse(scc common.Address, hub ethutil.WritableClient) {
	if _, ok := w.hubs.Load(scc); !ok {
		w.hubs.Store(scc, hub)
	}
}

func (w *SccSubmitter) RemoveVerse(scc common.Address) {
	w.hubs.Delete(scc)
}

func (w *SccSubmitter) HasVerse(scc common.Address) bool {
	_, ok := w.hubs.Load(scc)
	return ok
}

func (w *SccSubmitter) work(ctx context.Context, task *submitTask) {
	logCtx := []interface{}{"scc", task.scc.Hex()}
	scc, err := scc.NewScc(task.scc, task.hub)
	if err != nil {
		log.Error("Failed to create OasysStateCommitmentChain contract",
			append(logCtx, "err", err)...)
		return
	}

	sccv, err := sccverifier.NewSccverifier(w.sccvAddr, task.hub)
	if err != nil {
		log.Error("Failed to create OasysStateCommitmentChainVerifier contract",
			append(logCtx, "err", err)...)
		return
	}

	// fetch the next index from hub-layer
	nextIndex, err := scc.NextIndex(&bind.CallOpts{Context: ctx})
	if err != nil {
		w.log.Error("Failed to call the SCC.nextIndex method", append(logCtx, "err", err)...)
		return
	}
	logCtx = append(logCtx, "index", nextIndex)

	rows, err := w.findSignatures(task.scc, nextIndex.Uint64(),
		minStake, w.getTotalStake(), w.getSignerStakes())
	if errors.Is(err, errNoSignature) {
		w.log.Info("No signatures", logCtx...)
		return
	} else if err != nil {
		w.log.Error("Failed to find signatures", append(logCtx, "err", err)...)
		return
	}

	// send transaction
	ctx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()

	tx, err := w.sendTransaction(ctx, logCtx, task, sccv, rows)
	if err != nil {
		return
	}

	w.waitForCconfirmation(ctx, append(logCtx, "tx", tx.Hash().String()), task, tx)
}

func (w *SccSubmitter) refreshStakes(ctx context.Context) {
	tot, err := w.fetchTotalStake(ctx)
	if err != nil {
		return
	}

	signerStakes, err := w.fetchSignerStakes(ctx)
	if err != nil {
		return
	}

	w.stakes.Store(common.Address{}, tot)
	for addr, stake := range signerStakes {
		w.stakes.Store(addr, stake)
	}
}

func (w *SccSubmitter) getTotalStake() *big.Int {
	if tot, ok := w.stakes.Load(common.Address{}); !ok {
		return big.NewInt(0)
	} else {
		return tot.(*big.Int)
	}
}

func (w SccSubmitter) getSignerStakes() map[common.Address]*big.Int {
	cpy := map[common.Address]*big.Int{}
	w.stakes.Range(func(key, value any) bool {
		addr := key.(common.Address)
		stake := value.(*big.Int)
		if addr != (common.Address{}) {
			cpy[addr] = stake
		}
		return true
	})
	return cpy
}

func (w *SccSubmitter) fetchTotalStake(ctx context.Context) (*big.Int, error) {
	tot, err := w.sm.GetTotalStake(&bind.CallOpts{Context: ctx}, common.Big0)
	if err != nil {
		w.log.Error("Failed to call StakeManager.totalStake method", "err", err)
		return nil, err
	}
	return tot, nil
}

func (w *SccSubmitter) fetchSignerStakes(ctx context.Context) (map[common.Address]*big.Int, error) {
	stakes := map[common.Address]*big.Int{}
	cursor := big.NewInt(0)
	howMany := big.NewInt(250)
	for {
		result, err := w.sm.GetValidators(
			&bind.CallOpts{Context: ctx},
			common.Big0,
			cursor,
			howMany,
		)
		if err != nil {
			w.log.Error("Failed to call StakeManager.getValidators method", "err", err)
			return nil, err
		} else if len(result.Owners) == 0 {
			break
		}

		for i, operator := range result.Operators {
			stakes[operator] = result.Stakes[i]
		}
		cursor = result.NewCursor
	}

	return stakes, nil
}

func (w *SccSubmitter) findSignatures(
	scc common.Address,
	nextIndex uint64,
	minStake *big.Int,
	totalStake *big.Int,
	signerStakes map[common.Address]*big.Int,
) ([]*database.OptimismSignature, error) {
	// find signatures from database
	rows, err := w.db.Optimism.FindSignatures(nil, nil, &scc, &nextIndex, 1000, 0)
	if err != nil {
		return nil, err
	} else if len(rows) == 0 {
		return nil, errNoSignature
	}

	// group by BatchRoot and Approved
	sigGroup := map[string][]*database.OptimismSignature{}
	stakeGroup := map[string]*big.Int{}
	for _, row := range rows {
		k := fmt.Sprintf("%s:%v", row.BatchRoot, row.Approved)
		if _, ok := sigGroup[k]; !ok {
			sigGroup[k] = []*database.OptimismSignature{}
			stakeGroup[k] = new(big.Int)
		}

		if stake, ok := signerStakes[row.Signer.Address]; ok && stake.Cmp(minStake) >= 0 {
			sigGroup[k] = append(sigGroup[k], row)
			stakeGroup[k].Add(stakeGroup[k], stake)
		}
	}

	// find the group key with the highest stake
	highestKey := ""
	highestStake := big.NewInt(0)
	for k, stake := range stakeGroup {
		if stake.Cmp(highestStake) == 1 {
			highestKey = k
			highestStake = stake
		}
	}

	// check over half
	required := new(big.Int).Mul(new(big.Int).Div(totalStake, big.NewInt(100)), big.NewInt(51))
	if highestStake.Cmp(required) == -1 {
		return nil, fmt.Errorf(
			"stake amount shortage, required: %s, actual: %s",
			fromWei(required).String(),
			fromWei(highestStake).String(),
		)
	}

	// sort by signer address
	rows = sigGroup[highestKey]
	sort.Slice(rows, func(i, j int) bool {
		return bytes.Compare(rows[i].Signer.Address[:], rows[j].Signer.Address[:]) == -1
	})

	return rows, nil
}

func (w *SccSubmitter) sendTransaction(
	ctx context.Context,
	logCtx []interface{},
	task *submitTask,
	sccv *sccverifier.Sccverifier,
	rows []*database.OptimismSignature,
) (*types.Transaction, error) {
	var method func(*bind.TransactOpts, common.Address, sccverifier.Lib_OVMCodecChainBatchHeader, [][]byte) (*types.Transaction, error)
	if rows[0].Approved {
		logCtx = append(logCtx, "method", "SCCVerifier.approve")
		method = sccv.Approve
	} else {
		logCtx = append(logCtx, "method", "SCCVerifier.reject")
		method = sccv.Reject
	}

	// create params
	header := sccverifier.Lib_OVMCodecChainBatchHeader{
		BatchIndex:        new(big.Int).SetUint64(rows[0].BatchIndex),
		BatchRoot:         rows[0].BatchRoot,
		BatchSize:         new(big.Int).SetUint64(rows[0].BatchSize),
		PrevTotalElements: new(big.Int).SetUint64(rows[0].PrevTotalElements),
		ExtraData:         rows[0].ExtraData,
	}

	signatures := make([][]byte, len(rows))
	for i, row := range rows {
		signatures[i] = row.Signature[:]
	}

	// estimate gas
	opts := task.hub.TransactOpts(ctx)
	opts.NoSend = true
	tx, err := method(opts, task.scc, header, signatures)
	if err != nil {
		w.log.Error("Failed to estimate gas", append(logCtx, "err", err)...)
		return nil, err
	}

	// send
	opts = task.hub.TransactOpts(ctx)
	opts.GasLimit = uint64(float64(tx.Gas()) * w.gasMultiplier)
	tx, err = method(opts, task.scc, header, signatures)
	if err != nil {
		log.Error("Failed to send transaction", append(logCtx, "err", err)...)
		return nil, err
	}

	w.log.Info(
		"Sent transaction",
		append(
			logCtx,
			"tx", tx.Hash().String(),
			"nonce", tx.Nonce(),
			"gas-limit", tx.Gas(),
			"gas-fee", tx.GasFeeCap(),
			"gas-tip", tx.GasTipCap(),
		)...)

	return tx, nil
}

func (w *SccSubmitter) waitForCconfirmation(
	ctx context.Context,
	logCtx []interface{},
	task *submitTask,
	tx *types.Transaction,
) {
	// wait for block to be validated
	receipt, err := bind.WaitMined(ctx, task.hub, tx)
	if err != nil {
		w.log.Error("Failed to receive receipt", append(logCtx, "err", err)...)
		return
	}
	if receipt.Status != 1 {
		w.log.Error("Transaction reverted", logCtx...)
		return
	}

	// wait for confirmations
	confirmed := map[common.Hash]bool{receipt.BlockHash: true}
	for {
		remaining := w.confirmations - len(confirmed)
		if remaining <= 0 {
			w.log.Info("Transaction succeeded", logCtx...)
			return
		}

		w.log.Info("Wait for confirmation", append(logCtx, "remaining", remaining)...)
		time.Sleep(time.Second)

		h, err := task.hub.HeaderByNumber(ctx, nil)
		if err != nil {
			w.log.Error("Failed to fetch block header", append(logCtx, "err", err)...)
			continue
		}
		confirmed[h.Hash()] = true
	}
}

func fromWei(wei *big.Int) *big.Int {
	return new(big.Int).Div(wei, big.NewInt(params.Ether))
}

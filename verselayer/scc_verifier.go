package verselayer

import (
	"bytes"
	"context"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

var (
	// See: https://github.com/oasysgames/oasys-optimism/blob/134491cc2cd9ec588bbaad7697beaf74deddece7/packages/contracts/contracts/libraries/utils/Lib_MerkleTree.sol#L29-L46
	merkleDefaultBytes [][32]byte
	merkleDefaultHexs  = []string{
		"0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563",
		"0x633dc4d7da7256660a892f8f1604a44b5432649cc8ec5cb3ced4c4e6ac94dd1d",
		"0x890740a8eb06ce9be422cb8da5cdafc2b58c0a5e24036c578de2a433c828ff7d",
		"0x3b8ec09e026fdc305365dfc94e189a81b38c7597b3d941c279f042e8206e0bd8",
		"0xecd50eee38e386bd62be9bedb990706951b65fe053bd9d8a521af753d139e2da",
		"0xdefff6d330bb5403f63b14f33b578274160de3a50df4efecf0e0db73bcdd3da5",
		"0x617bdd11f7c0a11f49db22f629387a12da7596f9d1704d7465177c63d88ec7d7",
		"0x292c23a9aa1d8bea7e2435e555a4a60e379a5a35f3f452bae60121073fb6eead",
		"0xe1cea92ed99acdcb045a6726b2f87107e8a61620a232cf4d7d5b5766b3952e10",
		"0x7ad66c0a68c72cb89e4fb4303841966e4062a76ab97451e3b9fb526a5ceb7f82",
		"0xe026cc5a4aed3c22a58cbd3d2ac754c9352c5436f638042dca99034e83636516",
		"0x3d04cffd8b46a874edf5cfae63077de85f849a660426697b06a829c70dd1409c",
		"0xad676aa337a485e4728a0b240d92b3ef7b3c372d06d189322bfd5f61f1e7203e",
		"0xa2fca4a49658f9fab7aa63289c91b7c7b6c832a6d0e69334ff5b0a3483d09dab",
		"0x4ebfd9cd7bca2505f7bef59cc1c12ecc708fff26ae4af19abe852afe9e20c862",
		"0x2def10d13dd169f550f578bda343d9717a138562e0093b380a1120789d53cf10",
	}
)

func init() {
	merkleDefaultBytes = make([][32]byte, len(merkleDefaultHexs))
	for i, hex := range merkleDefaultHexs {
		merkleDefaultBytes[i] = util.BytesToBytes32(common.FromHex(hex))
	}
}

type verifyTask struct {
	scc   common.Address
	verse ethutil.Client
}

// Worker to verify the events of OasysStateCommitmentChain.
type SccVerifier struct {
	cfg    *config.Verifier
	db     *database.Database
	signer ethutil.SignableClient

	verses *sync.Map
	topic  *util.Topic
	log    log.Logger
}

// Returns the new verifier.
func NewSccVerifier(
	cfg *config.Verifier,
	db *database.Database,
	signer ethutil.SignableClient,
) *SccVerifier {
	return &SccVerifier{
		cfg:    cfg,
		db:     db,
		signer: signer,
		verses: &sync.Map{},
		topic:  util.NewTopic(),
		log:    log.New("worker", "scc-verifier"),
	}
}

// Start verifier.
func (w *SccVerifier) Start(ctx context.Context) {
	w.log.Info(
		"Worker started",
		"signer", w.signer.Signer(),
		"interval", w.cfg.Interval,
		"stateroot-limit", w.cfg.StateCollectLimit,
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
			w.verses.Range(func(key, value interface{}) bool {
				scc, ok0 := key.(common.Address)
				verse, ok1 := value.(ethutil.Client)
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

						if task, ok := data.(*verifyTask); ok {
							w.work(ctx, task)
						}
					}
					wg.AddWorker(ctx, name, handler)
				}

				wg.Enqueue(name, &verifyTask{scc: scc, verse: verse})
				return true
			})
		}
	}
}

func (w *SccVerifier) Signer() ethutil.SignableClient {
	return w.signer
}

func (w *SccVerifier) AddVerse(scc common.Address, verse ethutil.Client) {
	w.verses.Store(scc, verse)
}

func (w *SccVerifier) RemoveVerse(scc common.Address) {
	w.verses.Delete(scc)
}

func (w *SccVerifier) HasVerse(rpc string, scc common.Address) bool {
	if value, ok := w.verses.Load(scc); !ok {
		return false
	} else {
		t, _ := value.(ethutil.Client)
		return t.URL() == rpc
	}
}

func (s *SccVerifier) SubscribeNewSignature(ctx context.Context) *SignatureSubscription {
	ch := make(chan *database.OptimismSignature)
	cancel := s.topic.Subscribe(ctx, func(ctx context.Context, data interface{}) {
		if t, ok := data.(*database.OptimismSignature); ok {
			ch <- t
		}
	})
	return &SignatureSubscription{Cancel: cancel, ch: ch}
}

func (w *SccVerifier) work(ctx context.Context, task *verifyTask) {
	logCtx := []interface{}{"scc", task.scc.Hex()}
	scc, err := scc.NewScc(task.scc, w.signer)
	if err != nil {
		log.Error("Failed to create OasysStateCommitmentChain contract",
			append(logCtx, "err", err)...)
		return
	}

	// fetch the next index from hub-layer
	nextIndex, err := scc.NextIndex(&bind.CallOpts{Context: ctx})
	if err != nil {
		w.log.Error("Failed to call the SCC.nextIndex method", append(logCtx, "err", err)...)
		return
	}

	// verify the signature that match the nextIndex
	// and delete after signatures if there is a problem.
	// Prevent getting stuck indefinitely in the Verify waiting
	// state due to a bug in the signature creation process.
	w.deleteInvalidSignature(task.scc, nextIndex.Uint64())

	// run verification tasks until time out
	ctx, cancel := context.WithTimeout(ctx, w.cfg.StateCollectTimeout)
	defer cancel()

	for i := nextIndex.Uint64(); ; i++ {
		states, err := w.db.Optimism.FindVerificationWaitingStates(
			w.signer.Signer(), task.scc, i, 1)
		if err != nil {
			w.log.Error("Failed to find states", append(logCtx, "err", err)...)
			return
		} else if len(states) == 0 {
			w.log.Debug("Wait for new state", logCtx...)
			return
		}

		state := states[0]
		logCtx := append(logCtx, "index", state.BatchIndex)

		w.log.Info("Start state verification", logCtx...)
		approved, sig, err := w.verifyState(ctx, task.verse, state)
		if err != nil {
			return
		}

		row, err := w.db.OPSignature.Save(
			nil, nil,
			w.signer.Signer(),
			state.OptimismScc.Address,
			state.BatchIndex,
			state.BatchRoot,
			approved,
			sig)
		if err != nil {
			w.log.Error("Failed to save signature", append(logCtx, "err", err)...)
			return
		}

		w.topic.Publish(row)
		w.log.Info("State verification completed", append(logCtx, "approved", approved)...)
	}
}

func (w *SccVerifier) verifyState(
	ctx context.Context,
	verse ethutil.Client,
	state *database.OptimismState,
) (bool, database.Signature, error) {
	logCtx := []interface{}{
		"scc", state.Contract.Address.Hex(),
		"index", state.BatchIndex,
	}

	// collect block headers from verse-layer
	var (
		start   = state.PrevTotalElements + 1
		end     = start + state.BatchSize - 1
		headers []*types.Header
		err     error
	)

	bc, err := verse.NewBatchHeaderClient()
	if err != nil {
		w.log.Error("Failed to construct batch client", append(logCtx, "err", err)...)
		return false, database.Signature{}, err
	}

	bi := ethutil.NewBatchHeaderIterator(bc, start, end, w.cfg.StateCollectLimit)
	defer bi.Close()

	st := time.Now()
	logCtx = append(logCtx, "start", start, "end", end, "batch-size", state.BatchSize)
	for {
		hs, err := bi.Next(ctx)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				w.log.Warn("Time up", logCtx...)
			} else {
				w.log.Error("Failed to collect state roots", append(logCtx, "err", err)...)
			}
			return false, database.Signature{}, err
		} else if len(hs) == 0 {
			break
		}
		headers = append(headers, hs...)
	}

	w.log.Info("Collected state roots", append(logCtx, "elapsed", time.Since(st))...)

	// calc and compare state root
	elements := make([][32]byte, len(headers))
	for i, header := range headers {
		elements[i] = header.Root
	}
	merkleRoot, err := calcMerkleRoot(elements)
	if err != nil {
		w.log.Error("Failed to calculate merkle root", append(logCtx, "err", err)...)
		return false, database.Signature{}, err
	}
	approved := bytes.Equal(state.BatchRoot[:], merkleRoot[:])

	// calc and save signature
	msg := ethutil.NewMessage(
		w.signer.ChainID(),
		state.Contract.Address,
		new(big.Int).SetUint64(state.BatchIndex),
		state.BatchRoot,
		approved,
	)
	if sig, err := msg.Signature(w.signer.SignData); err == nil {
		return approved, sig, nil
	} else {
		w.log.Error("Failed to calculate signature", append(logCtx, "err", err)...)
		return false, database.Signature{}, err
	}
}

func (w *SccVerifier) deleteInvalidSignature(scc common.Address, nextIndex uint64) {
	logCtx := []interface{}{"scc", scc.Hex(), "next-index", nextIndex}

	signer := w.signer.Signer()
	sigs, err := w.db.OPSignature.Find(nil, &signer, &scc, &nextIndex, 1, 0)
	if err != nil {
		w.log.Error("Unable to find signatures", append(logCtx, "err", err)...)
		return
	} else if len(sigs) == 0 {
		w.log.Debug("No invalid signature", logCtx...)
		return
	}

	msg := ethutil.NewMessage(
		w.signer.ChainID(),
		sigs[0].Contract.Address,
		new(big.Int).SetUint64(sigs[0].RollupIndex),
		sigs[0].RollupHash,
		sigs[0].Approved)
	if err := msg.VerifySigner(sigs[0].Signature[:], signer); err == nil {
		w.log.Debug("No invalid signature", logCtx...)
		return
	} else {
		w.log.Error("Unable to verify signature", append(logCtx, "err", err)...)
	}

	w.log.Warn("Found invalid signature", append(logCtx, "signature", sigs[0].Signature.Hex())...)

	if rows, err := w.db.OPSignature.Deletes(signer, scc, nextIndex); err != nil {
		w.log.Error("Unable to delete signatures", append(logCtx, "err", err)...)
	} else {
		w.log.Warn("Deleted invalid signature", append(logCtx, "delete-rows", rows)...)
	}
}

// Calculates a merkle root for a list of 32-byte leaf hashes.
// see: https://github.com/oasysgames/oasys-optimism/blob/134491cc2cd9ec588bbaad7697beaf74deddece7/packages/contracts/contracts/libraries/utils/Lib_MerkleTree.sol#L22
func calcMerkleRoot(elements [][32]byte) ([32]byte, error) {
	if len(elements) == 0 {
		return [32]byte{}, errors.New("must provide at least one leaf hash")
	}
	if len(elements) == 1 {
		return elements[0], nil
	}

	rowSize := len(elements)
	depth := 0

	for rowSize > 1 {
		halfRowSize := rowSize / 2
		rowSizeIsOdd := rowSize%2 == 1

		for i := 0; i < halfRowSize; i++ {
			leftSibling := elements[(2 * i)][:]
			rightSibling := elements[(2*i)+1][:]
			elements[i] = util.BytesToBytes32(
				crypto.Keccak256(bytes.Join([][]byte{leftSibling, rightSibling}, []byte(""))),
			)
		}

		if rowSizeIsOdd {
			leftSibling := elements[rowSize-1][:]
			rightSibling := merkleDefaultBytes[depth][:]
			elements[halfRowSize] = util.BytesToBytes32(
				crypto.Keccak256(bytes.Join([][]byte{leftSibling, rightSibling}, []byte(""))),
			)
		}

		rowSize = halfRowSize
		if rowSizeIsOdd {
			rowSize++
		}
		depth++
	}

	return elements[0], nil
}

type SignatureSubscription struct {
	Cancel context.CancelFunc
	ch     chan *database.OptimismSignature
}

func (s *SignatureSubscription) Next() <-chan *database.OptimismSignature {
	return s.ch
}

package verse

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/contract/sccverifier"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

var (
	_ Verse             = &oplegacy{}
	_ VerifiableVerse   = &verifiableOPLegacy{}
	_ TransactableVerse = &transactableOPLegacy{}

	NewOPLegacy = newVerseFactory(func(v Verse) Verse { return &oplegacy{v} })
)

type oplegacy struct {
	Verse
}

type verifiableOPLegacy struct {
	VerifiableVerse
}

type transactableOPLegacy struct {
	TransactableVerse
}

func (op *oplegacy) Logger(base log.Logger) log.Logger {
	return base.New("chain-id", fmt.Sprint(op.ChainID()), "op-version", "v0")
}

func (op *oplegacy) EventDB() database.IOPEventDB {
	return database.NewOPEventDB[database.OptimismState](op.DB())
}

func (op *oplegacy) NextIndex(opts *bind.CallOpts) (uint64, error) {
	sc, err := newSccContract(op)
	if err != nil {
		return 0, err
	}
	b, err := sc.NextIndex(opts)
	if err != nil {
		return 0, err
	}
	return b.Uint64(), nil
}

func (op *oplegacy) EventEmittedBlock(opts *bind.FilterOpts, rollupIndex uint64) (uint64, error) {
	sc, err := newSccContract(op)
	if err != nil {
		return 0, err
	}
	e, err := findStateBatchAppendedEvent(sc, opts, rollupIndex)
	if err != nil {
		return 0, err
	}
	return e.Raw.BlockNumber, nil
}

func (op *oplegacy) WithVerifiable(l2Client ethutil.Client) VerifiableVerse {
	return &verifiableOPLegacy{&verifiableVerse{op, l2Client}}
}

func (op *oplegacy) WithTransactable(
	l1Signer ethutil.SignableClient,
	rollupContract common.Address,
) TransactableVerse {
	return &transactableOPLegacy{&transactableVerse{op, l1Signer, rollupContract}}
}

func (op *verifiableOPLegacy) Verify(
	base log.Logger,
	ctx context.Context,
	event database.OPEvent,
	l2BatchSize int,
) (approved bool, err error) {
	row, ok := event.(*database.OptimismState)
	if !ok {
		return false, errors.New("not OptimismState event")
	}

	log := op.Logger(base).New("index", row.BatchIndex)

	// collect block headers from verse-layer
	var (
		start    = row.PrevTotalElements + 1
		end      = start + row.BatchSize - 1
		elements [][32]byte
	)

	bc, err := op.L2Client().NewBatchHeaderClient()
	if err != nil {
		log.Error("Failed to construct batch client", "err", err)
		return false, err
	}

	bi := ethutil.NewBatchHeaderIterator(bc, start, end, l2BatchSize)
	defer bi.Close()

	st := time.Now()
	log = log.New("start", start, "end", end, "batch-size", row.BatchSize)
	for {
		headers, err := bi.Next(ctx)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				log.Warn("Time up")
			} else {
				log.Warn("Failed to collect state roots", "err", err)
			}
			return false, err
		} else if len(headers) == 0 {
			break
		}
		for _, header := range headers {
			elements = append(elements, header.Root)
		}
	}

	log.Debug("Collected L2 states", "elapsed", time.Since(st))

	// calc and compare state root
	merkleRoot, err := CalcMerkleRoot(elements)
	if err != nil {
		log.Error("Failed to calculate merkle root", "err", err)
		return false, err
	}

	return bytes.Equal(row.BatchRoot[:], merkleRoot[:]), nil
}

func (op *transactableOPLegacy) Transact(
	opts *bind.TransactOpts,
	rollupIndex uint64,
	approved bool,
	signatures [][]byte,
) (*types.Transaction, error) {
	sc, err := newSccContract(op)
	if err != nil {
		return nil, err
	}

	vc, err := sccverifier.NewSccverifier(op.VerifyContract(), op.L1Signer())
	if err != nil {
		return nil, err
	}

	e, err := findStateBatchAppendedEvent(sc, &bind.FilterOpts{Context: opts.Context}, rollupIndex)
	if err != nil {
		return nil, err
	}

	method := vc.Approve
	if !approved {
		method = vc.Reject
	}

	return method(
		opts,
		op.RollupContract(),
		sccverifier.Lib_OVMCodecChainBatchHeader{
			BatchIndex:        e.BatchIndex,
			BatchRoot:         e.BatchRoot,
			BatchSize:         e.BatchSize,
			PrevTotalElements: e.PrevTotalElements,
			ExtraData:         e.ExtraData,
		},
		signatures,
	)
}

var (
	// See: https://github.com/oasysgames/oasys-optimism/blob/134491cc2cd9ec588bbaad7697beaf74deddece7/packages/contracts/contracts/libraries/utils/Lib_MerkleTree.sol#L29-L46
	merkleDefaults    []common.Hash
	merkleDefaultHexs = []string{
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
	merkleDefaults = make([]common.Hash, len(merkleDefaultHexs))
	for i, hex := range merkleDefaultHexs {
		merkleDefaults[i] = common.HexToHash(hex)
	}
}

// Calculates a merkle root for a list of 32-byte leaf hashes.
// see: https://github.com/oasysgames/oasys-optimism/blob/134491cc2cd9ec588bbaad7697beaf74deddece7/packages/contracts/contracts/libraries/utils/Lib_MerkleTree.sol#L22
func CalcMerkleRoot(elements [][32]byte) ([32]byte, error) {
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
			rightSibling := merkleDefaults[depth][:]
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

func newSccContract(op Verse) (*scc.Scc, error) {
	return scc.NewScc(op.RollupContract(), op.L1Client())
}

func findStateBatchAppendedEvent(
	scc *scc.Scc,
	opts *bind.FilterOpts,
	batchIndex uint64,
) (event *scc.SccStateBatchAppended, err error) {
	iter, err := scc.FilterStateBatchAppended(
		opts, []*big.Int{new(big.Int).SetUint64(batchIndex)})
	if err != nil {
		return nil, err
	}
	defer iter.Close()

	for {
		if iter.Next() {
			event = iter.Event // returns the last event
		} else if err := iter.Error(); err != nil {
			return nil, err
		} else {
			break
		}
	}

	if event == nil {
		err = ErrEventNotFound
	}
	return event, err
}

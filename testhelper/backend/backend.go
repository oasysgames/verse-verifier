package backend

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/account"
)

var _ ethutil.Client = &Backend{}

func NewBackend(alloc core.GenesisAlloc, gasLimit uint64) *Backend {
	if alloc == nil {
		alloc = core.GenesisAlloc{
			account.DefaultAccount.Address: {Balance: big.NewInt(params.Ether)},
		}
	}
	if gasLimit == 0 {
		gasLimit = 500_000_000
	}
	simulated := backends.NewSimulatedBackend(alloc, gasLimit)
	return &Backend{simulated}
}

type Backend struct {
	*backends.SimulatedBackend
}

func (b *Backend) URL() string {
	return "SimulatedBackend"
}

func (b *Backend) BlockNumber(ctx context.Context) (uint64, error) {
	return b.Blockchain().CurrentHeader().Number.Uint64(), nil
}

func (b *Backend) NewBatchHeaderClient() (ethutil.BatchHeaderClient, error) {
	return &BatchHeaderClient{b}, nil
}

func (b *Backend) GetProof(
	ctx context.Context,
	account common.Address,
	keys []string,
	blockNumber *big.Int,
) (*gethclient.AccountResult, error) {
	header, err := b.HeaderByNumber(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	state, err := b.Blockchain().StateAt(header.Root)
	if state == nil || err != nil {
		return nil, err
	}

	storageTrie := state.StorageTrie(account)
	storageHash := types.EmptyRootHash
	codeHash := state.GetCodeHash(account)
	storageProof := make([]gethclient.StorageResult, len(keys))

	// if we have a storageTrie, (which means the account exists), we can update the storagehash
	if storageTrie != nil {
		storageHash = storageTrie.Hash()
	} else {
		// no storageTrie means the account does not exist, so the codeHash is the hash of an empty bytearray.
		codeHash = crypto.Keccak256Hash(nil)
	}

	// create the proof for the storageKeys
	for i, key := range keys {
		if storageTrie != nil {
			proof, storageError := state.GetStorageProof(account, common.HexToHash(key))
			if storageError != nil {
				return nil, storageError
			}
			storageProof[i] = gethclient.StorageResult{
				Key:   key,
				Value: state.GetState(account, common.HexToHash(key)).Big(),
				Proof: toHexSlice(proof),
			}
		} else {
			storageProof[i] = gethclient.StorageResult{
				Key:   key,
				Value: new(big.Int),
				Proof: []string{},
			}
		}
	}

	// create the accountProof
	accountProof, proofErr := state.GetProof(account)
	if proofErr != nil {
		return nil, proofErr
	}

	return &gethclient.AccountResult{
		Address:      account,
		AccountProof: toHexSlice(accountProof),
		Balance:      state.GetBalance(account),
		CodeHash:     codeHash,
		Nonce:        state.GetNonce(account),
		StorageHash:  storageHash,
		StorageProof: storageProof,
	}, state.Error()
}

func (b *Backend) GetHeaderBatch(
	ctx context.Context,
	start uint64,
	size, limit int,
) ([]*types.Header, error) {
	headers := make([]*types.Header, size)
	for i := 0; i < size; i++ {
		h, err := b.HeaderByNumber(ctx, new(big.Int).SetUint64(start+uint64(i)))
		if err != nil {
			return nil, err
		}

		headers[i] = h
	}

	return headers, nil
}

func (b *Backend) TxSender(tx *types.Transaction) (common.Address, error) {
	signer := types.MakeSigner(b.Blockchain().Config(), common.Big0)
	return types.Sender(signer, tx)
}

func (b *Backend) Mining() *types.Header {
	b.Commit()
	return b.Blockchain().CurrentHeader()
}

func toHexSlice(b [][]byte) []string {
	r := make([]string, len(b))
	for i := range b {
		r[i] = hexutil.Encode(b[i])
	}
	return r
}

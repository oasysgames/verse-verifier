package testhelper

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"io/ioutil"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

var (
	chainId = big.NewInt(1337)

	ks      *keystore.KeyStore
	account accounts.Account
	alloc   map[common.Address]core.GenesisAccount
)

func init() {
	// create new wallet and key
	reader := bytes.NewReader([]byte(strings.Repeat("x", 40)))
	priv, _ := ecdsa.GenerateKey(crypto.S256(), reader)

	// setup keystore
	tempDir, _ := ioutil.TempDir(os.TempDir(), "keystore")
	ks = keystore.NewKeyStore(tempDir, keystore.StandardScryptN, keystore.StandardScryptP)

	// import private key into the keystore
	account, _ = ks.ImportECDSA(priv, "")
	ks.Unlock(account, "")

	// add balance to wallet
	alloc = map[common.Address]core.GenesisAccount{
		account.Address: {Balance: big.NewInt(params.Ether)},
	}
}

func NewTestBackend() *TestBackend {
	return &TestBackend{
		SimulatedBackend: backends.NewSimulatedBackend(alloc, 500_000_000),
		ks:               ks,
		account:          account,
	}
}

type TestBackend struct {
	*backends.SimulatedBackend

	ks      *keystore.KeyStore
	account accounts.Account
}

func (b *TestBackend) NewAccountBackend() *TestBackend {
	priv, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	account, _ := b.ks.ImportECDSA(priv, "")
	b.ks.Unlock(account, "")

	return &TestBackend{
		SimulatedBackend: b.SimulatedBackend,
		ks:               b.ks,
		account:          account,
	}
}

func (b *TestBackend) URL() string {
	return "SimulatedBackend"
}

func (b *TestBackend) TransactOpts(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		Context: ctx,
		From:    b.Signer(),
		Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return b.ks.SignTx(b.account, tx, b.ChainID())
		},
		GasPrice: big.NewInt(875_000_000),
		GasLimit: 500_000_000,
	}
}

func (b *TestBackend) BlockNumber(ctx context.Context) (uint64, error) {
	return b.Blockchain().CurrentHeader().Number.Uint64(), nil
}

func (b *TestBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	header, err := b.SimulatedBackend.HeaderByNumber(ctx, number)
	if err == nil && header == nil {
		return nil, ethereum.NotFound
	}
	return header, err
}

func (b *TestBackend) Mining() *types.Header {
	b.Commit()
	return b.Blockchain().CurrentHeader()
}

func (b *TestBackend) GetHeaderBatch(
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

func (b *TestBackend) ChainID() *big.Int {
	return chainId
}

func (b *TestBackend) Signer() common.Address {
	return b.account.Address
}

func (b *TestBackend) SignData(data []byte) (sig []byte, err error) {
	return b.ks.SignHash(b.account, crypto.Keccak256(data))
}

func (b *TestBackend) SignTx(tx *types.Transaction) (*types.Transaction, error) {
	return b.ks.SignTx(b.account, tx, b.ChainID())
}

func (b *TestBackend) NewBatchHeaderClient() (ethutil.BatchHeaderClient, error) {
	return &TestBatchHeaderClient{b}, nil
}

func (b *TestBackend) GetProof(
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

type TestBatchHeaderClient struct {
	ethutil.Client
}

func (c *TestBatchHeaderClient) Get(
	ctx context.Context,
	start, end uint64,
) ([]*types.Header, error) {
	size := int(end - start + 1)

	headers := make([]*types.Header, size)
	for i := 0; i < size; i++ {
		h, err := c.HeaderByNumber(ctx, new(big.Int).SetUint64(start+uint64(i)))
		if err != nil {
			return nil, err
		}
		headers[i] = h
	}

	return headers, nil
}

func (c *TestBatchHeaderClient) Close() error {
	return nil
}

func toHexSlice(b [][]byte) []string {
	r := make([]string, len(b))
	for i := range b {
		r[i] = hexutil.Encode(b[i])
	}
	return r
}

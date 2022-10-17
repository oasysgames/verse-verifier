package testhelper

import (
	"bytes"
	"context"
	"crypto/ecdsa"
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
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
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

	if err == nil && number != nil && number.Cmp(header.Number) == 1 {
		return nil, ethereum.NotFound
	}

	return header, err
}

func (s *TestBackend) Mining() *types.Header {
	s.Commit()
	return s.Blockchain().CurrentHeader()
}

func (s *TestBackend) GetHeaderBatch(
	ctx context.Context,
	start, size int,
) ([]*types.Header, error) {
	headers := make([]*types.Header, size)
	for i := 0; i < size; i++ {
		h, err := s.HeaderByNumber(ctx, big.NewInt(int64(start+i)))
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

func (b *TestBackend) SignData(hash []byte) (sig []byte, err error) {
	return b.ks.SignHash(b.account, crypto.Keccak256(hash))
}

func (b *TestBackend) SignTx(tx *types.Transaction) (*types.Transaction, error) {
	return b.ks.SignTx(b.account, tx, b.ChainID())
}

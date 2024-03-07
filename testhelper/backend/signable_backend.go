package backend

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/account"
)

var _ ethutil.SignableClient = &SignableBackend{}

func NewSignableBackend(
	b *Backend,
	ks *keystore.KeyStore,
	ac *accounts.Account,
) *SignableBackend {
	if b == nil {
		b = NewBackend(nil, 0)
	}
	if ks == nil {
		ks = account.DefaultKeyStore
	}
	if ac == nil {
		ac = account.DefaultAccount
	}
	return &SignableBackend{
		Backend: b,
		ks:      ks,
		account: ac,
	}
}

type SignableBackend struct {
	*Backend

	ks      *keystore.KeyStore
	account *accounts.Account
}

func (c *SignableBackend) ChainID() *big.Int {
	return big.NewInt(1337)
}

func (c *SignableBackend) Signer() common.Address {
	return c.account.Address
}

func (c *SignableBackend) SignData(data []byte) (sig []byte, err error) {
	return c.ks.SignHash(*c.account, crypto.Keccak256(data))
}

func (c *SignableBackend) SignTx(tx *types.Transaction) (*types.Transaction, error) {
	return c.ks.SignTx(*c.account, tx, c.ChainID())
}

func (c *SignableBackend) TransactOpts(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		Context: ctx,
		From:    c.Signer(),
		Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return c.ks.SignTx(*c.account, tx, c.ChainID())
		},
		GasPrice: big.NewInt(875_000_000),
		GasLimit: 500_000_000,
	}
}

func (b *SignableBackend) WithNewAccount() *SignableBackend {
	priv, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	account, _ := b.ks.ImportECDSA(priv, "")
	b.ks.Unlock(account, "")

	return &SignableBackend{
		Backend: b.Backend,
		ks:      b.ks,
		account: &account,
	}
}

func (c *SignableBackend) SendTxWithSign(
	ctx context.Context,
	unsigned *types.Transaction,
) (signed *types.Transaction, err error) {
	signed, err = c.SignTx(unsigned)
	if err != nil {
		return nil, err
	}
	if err = c.SendTransaction(context.Background(), signed); err != nil {
		return nil, err
	}
	c.Commit()
	return signed, nil
}

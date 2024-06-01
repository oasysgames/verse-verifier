package backend

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
	mrand "math/rand"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/account"
)

func NewSignableBackend(b *Backend, signer ethutil.Signer) *SignableBackend {
	if b == nil {
		b = NewBackend(nil, 0)
	}
	if signer == nil {
		// alternate the use of KeystoreSigner and PrivateKeySigner to test both signers
		if mrand.Intn(100)%2 == 0 {
			signer = ethutil.NewKeystoreSigner(account.DefaultWallet, account.DefaultAccount)
		} else {
			signer = ethutil.NewPrivateKeySigner(account.DefaultPrivateKey)
		}
	}
	return &SignableBackend{
		Backend: b,
		signer:  signer,
	}
}

type SignableBackend struct {
	*Backend

	signer ethutil.Signer
}

func (c *SignableBackend) ChainID() *big.Int {
	return big.NewInt(1337)
}

func (c *SignableBackend) Signer() common.Address {
	return c.signer.From()
}

func (c *SignableBackend) SignData(data []byte) (sig []byte, err error) {
	return c.signer.SignData(data)
}

func (c *SignableBackend) SignTx(tx *types.Transaction) (*types.Transaction, error) {
	return c.signer.SignTx(tx, c.ChainID())
}

func (c *SignableBackend) TransactOpts(ctx context.Context) *bind.TransactOpts {
	return &bind.TransactOpts{
		Context: ctx,
		From:    c.signer.From(),
		Signer: func(a common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return c.signer.SignTx(tx, c.ChainID())
		},
	}
}

func (b *SignableBackend) WithNewAccount() *SignableBackend {
	priv, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	return &SignableBackend{
		Backend: b.Backend,
		signer:  ethutil.NewPrivateKeySigner(priv),
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

// Returns the base fee of the block plus 1 gwei
func (c *SignableBackend) BaseGasPrice(ctx context.Context, number *big.Int) (*big.Int, error) {
	head, err := c.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	return new(big.Int).Add(head.BaseFee, big.NewInt(params.GWei)), nil
}

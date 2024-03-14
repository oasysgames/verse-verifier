package ethutil

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signer interface {
	From() common.Address
	SignData(data []byte) (sig []byte, err error)
	SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)
}

func NewKeystoreSigner(w accounts.Wallet, a *accounts.Account) Signer {
	return &keystoreSigner{wallet: w, account: a}
}

func NewPrivateKeySigner(key *ecdsa.PrivateKey) Signer {
	return &privateKeySigner{key: key}
}

type keystoreSigner struct {
	wallet  accounts.Wallet
	account *accounts.Account
}

func (s *keystoreSigner) From() common.Address {
	return s.account.Address
}

func (s *keystoreSigner) SignData(data []byte) (sig []byte, err error) {
	return s.wallet.SignData(*s.account, "", data)
}

func (s *keystoreSigner) SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return s.wallet.SignTx(*s.account, tx, chainID)
}

type privateKeySigner struct {
	key *ecdsa.PrivateKey
}

func (s *privateKeySigner) From() common.Address {
	return crypto.PubkeyToAddress(s.key.PublicKey)
}

func (s *privateKeySigner) SignData(data []byte) (sig []byte, err error) {
	return crypto.Sign(crypto.Keccak256(data), s.key)
}

func (s *privateKeySigner) SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return types.SignTx(tx, types.LatestSignerForChainID(chainID), s.key)
}

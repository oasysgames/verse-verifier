package account

import (
	"crypto/ecdsa"
	"crypto/rand"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	DefaultPrivateKey *ecdsa.PrivateKey
	DefaultKeyStore   *keystore.KeyStore
	DefaultAccount    *accounts.Account
	DefaultWallet     accounts.Wallet
)

var (
	privateKey = "7878787878787879118f980be9bfb497fd94f02a4adb77fdffea86d8da73e279"
)

func init() {
	// decode default private key
	if priv, err := crypto.HexToECDSA(privateKey); err != nil {
		panic(err)
	} else {
		DefaultPrivateKey = priv
	}

	// create default keystore
	if ks, err := NewKeyStore("default"); err != nil {
		panic(err)
	} else {
		DefaultKeyStore = ks
	}

	// create default account
	if ac, err := NewAccount(DefaultKeyStore, &privateKey); err != nil {
		panic(err)
	} else {
		DefaultAccount = ac
	}

	// kss := wallet.NewKeyStore(DefaultKeyStore)
LOOP:
	for _, wallet := range DefaultKeyStore.Wallets() {
		for _, account := range wallet.Accounts() {
			if account.Address == DefaultAccount.Address {
				DefaultWallet = wallet
				break LOOP
			}
		}
	}
	if DefaultWallet == nil {
		panic("Wallet not found")
	}
}

func NewKeyStore(dirname string) (*keystore.KeyStore, error) {
	keydir, err := ioutil.TempDir(os.TempDir(), dirname)
	if err != nil {
		return nil, err
	}
	return keystore.NewKeyStore(keydir, keystore.StandardScryptN, keystore.StandardScryptP), nil
}

func NewAccount(ks *keystore.KeyStore, privKey *string) (*accounts.Account, error) {
	var (
		priv *ecdsa.PrivateKey
		err  error
	)
	if privKey != nil {
		priv, err = crypto.HexToECDSA(*privKey)
		if err != nil {
			return nil, err
		}
	} else {
		priv, err = ecdsa.GenerateKey(crypto.S256(), rand.Reader)
		if err != nil {
			return nil, err
		}
	}

	account, err := ks.ImportECDSA(priv, "")
	if err != nil {
		return nil, err
	}
	ks.Unlock(account, "")

	return &account, nil
}

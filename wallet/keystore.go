package wallet

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

type KeyStore struct {
	*keystore.KeyStore
}

func NewKeyStore(dir string) *KeyStore {
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)

	return &KeyStore{
		KeyStore: ks,
	}
}

func (k *KeyStore) FindWallet(address common.Address) (accounts.Wallet, *accounts.Account, error) {
	for _, wallet := range k.Wallets() {
		for _, account := range wallet.Accounts() {
			if account.Address == address {
				return wallet, &account, nil
			}
		}
	}

	return nil, nil, errors.New("not found")
}

func (k *KeyStore) IsLocked(wallet accounts.Wallet) (bool, error) {
	status, err := wallet.Status()
	if err != nil {
		return false, err
	}

	return strings.ToLower(status) != "unlocked", nil
}

func (k *KeyStore) WaitForUnlock(ctx context.Context, wallet accounts.Wallet) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("cancelled")
		case <-ticker.C:
			locked, err := k.IsLocked(wallet)
			if err != nil {
				return err
			}
			if !locked {
				return nil
			}
		}
	}
}

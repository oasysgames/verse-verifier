package ethutil

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Ecrecover(hash []byte, sig []byte) (common.Address, error) {
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes long")
	}
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}

	// Transform V from 0/1 to 27/28 according to the yellow paper
	cpy := make([]byte, len(sig))
	copy(cpy, sig)
	cpy[crypto.RecoveryIDOffset] -= 27

	if pub, err := crypto.SigToPub(hash, cpy); err == nil {
		return crypto.PubkeyToAddress(*pub), nil
	} else {
		return common.Address{}, err
	}
}

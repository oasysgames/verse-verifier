package verselayer

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

type SccMessage struct {
	AbiPacked []byte
	Eip712Msg string
}

func NewSccMessage(
	hubChainID *big.Int,
	scc common.Address,
	batchIndex *big.Int,
	batchRoot [32]byte,
	approved bool,
) *SccMessage {
	b := common.Big0
	if approved {
		b = common.Big1
	}

	abiPacked := bytes.Join([][]byte{
		common.LeftPadBytes(hubChainID.Bytes(), 32),
		scc[:],
		common.LeftPadBytes(batchIndex.Bytes(), 32),
		batchRoot[:],
		b.Bytes(),
	}, nil)
	_, msg := accounts.TextAndHash(crypto.Keccak256(abiPacked))

	return &SccMessage{
		AbiPacked: abiPacked,
		Eip712Msg: msg,
	}
}

func (m *SccMessage) Signature(signDataFn ethutil.SignDataFn) ([65]byte, error) {
	var sig [65]byte
	signed, err := signDataFn([]byte(m.Eip712Msg))
	if err != nil {
		return sig, err
	}
	copy(sig[:], signed)

	// Transform V from 0/1 to 27/28 according to the yellow paper
	sig[crypto.RecoveryIDOffset] += 27
	return sig, nil
}

package ethutil

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignerMismatchError struct {
	Actual, Recoverd common.Address
}

func (e *SignerMismatchError) Error() string {
	return fmt.Sprintf("signer mismatch: actual: %s, recoverd: %s", e.Actual, e.Recoverd)
}

type Message struct {
	AbiPacked []byte
	Eip712Msg string
}

func (m *Message) Signature(signDataFn SignDataFn) ([65]byte, error) {
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

func (m *Message) Ecrecover(signature []byte) (common.Address, error) {
	hash := crypto.Keccak256([]byte(m.Eip712Msg))
	return Ecrecover(hash, signature)
}

func (m *Message) VerifySigner(signature []byte, signer common.Address) error {
	if recoverd, err := m.Ecrecover(signature); err != nil {
		return err
	} else if !bytes.Equal(recoverd.Bytes(), signer.Bytes()) {
		return &SignerMismatchError{Actual: signer, Recoverd: recoverd}
	}
	return nil
}

func NewMessage(
	hubChainID *big.Int,
	contract common.Address,
	rollupIndex *big.Int,
	rollupHash [32]byte,
	approved bool,
) *Message {
	abiPacked := bytes.Join([][]byte{
		padUint256(hubChainID),
		contract[:],
		padUint256(rollupIndex),
		rollupHash[:],
		padBool(approved),
	}, nil)
	_, msg := accounts.TextAndHash(crypto.Keccak256(abiPacked))

	return &Message{AbiPacked: abiPacked, Eip712Msg: msg}
}

// Deprecated: This is a signature with a bug in the boolean type abi-encode.
// It is retained for verification purposes because there are peers still
// sending signatures containing the bug.
func NewMessageWithApprovedBug(
	hubChainID *big.Int,
	scc common.Address,
	rollupIndex *big.Int,
	rollupHash [32]byte,
	approved bool,
) *Message {
	b := common.Big0
	if approved {
		b = common.Big1
	}

	abiPacked := bytes.Join([][]byte{
		common.LeftPadBytes(hubChainID.Bytes(), 32),
		scc[:],
		common.LeftPadBytes(rollupIndex.Bytes(), 32),
		rollupHash[:],
		b.Bytes(),
	}, nil)
	_, msg := accounts.TextAndHash(crypto.Keccak256(abiPacked))

	return &Message{AbiPacked: abiPacked, Eip712Msg: msg}
}

func L2OORollupHashSource(outputRoot common.Hash, l1Timestamp, l2BlockNumber *big.Int) []byte {
	return bytes.Join([][]byte{
		outputRoot[:],
		padUint128(l1Timestamp),
		padUint128(l2BlockNumber),
	}, nil)
}

func padUint128(val *big.Int) []byte {
	return common.LeftPadBytes(val.Bytes(), 16)
}

func padUint256(val *big.Int) []byte {
	return common.LeftPadBytes(val.Bytes(), 32)
}

func padBool(val bool) []byte {
	if val {
		return []byte{1}
	}
	return []byte{0}
}

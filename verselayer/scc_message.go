package verselayer

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

type SignerMismatchError struct {
	Actual, Recoverd common.Address
}

func (e *SignerMismatchError) Error() string {
	return fmt.Sprintf("signer mismatch: actual: %s, recoverd: %s", e.Actual, e.Recoverd)
}

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
	_approved := []byte{0}
	if approved {
		_approved = []byte{1}
	}

	// See: https://github.com/oasysgames/oasys-optimism/blob/5186190c3250121179064b70d8e2fbd2d0a03ce3/packages/contracts/contracts/oasys/L1/rollup/OasysStateCommitmentChainVerifier.sol#L111-L119
	abiPacked := bytes.Join([][]byte{
		common.LeftPadBytes(hubChainID.Bytes(), 32),
		scc[:],
		common.LeftPadBytes(batchIndex.Bytes(), 32),
		batchRoot[:],
		_approved,
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

func (m *SccMessage) Ecrecover(signature []byte) (common.Address, error) {
	hash := crypto.Keccak256([]byte(m.Eip712Msg))
	return ethutil.Ecrecover(hash, signature)
}

func (m *SccMessage) VerifySigner(signature []byte, signer common.Address) error {
	if recoverd, err := m.Ecrecover(signature); err != nil {
		return err
	} else if !bytes.Equal(recoverd.Bytes(), signer.Bytes()) {
		return &SignerMismatchError{Actual: signer, Recoverd: recoverd}
	}
	return nil
}

// Deprecated: This is a signature with a bug in the boolean type abi-encode.
// It is retained for verification purposes because there are peers still
// sending signatures containing the bug.
func NewSCCMessageWithApprovedBug(
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

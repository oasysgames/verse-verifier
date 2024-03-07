package ethutil

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/account"
	"github.com/stretchr/testify/suite"
)

type MessageTestSuite struct {
	suite.Suite

	account  *accounts.Account
	signData SignDataFn

	approveMsg, rejectMsg *Message
}

func TestMessage(t *testing.T) {
	suite.Run(t, new(MessageTestSuite))
}

func (s *MessageTestSuite) SetupSuite() {
	s.account = account.DefaultAccount
	s.signData = func(hash []byte) (sig []byte, err error) {
		return account.DefaultKeyStore.SignHash(*s.account, crypto.Keccak256(hash))
	}

	s.approveMsg = NewMessage(
		big.NewInt(1),
		common.HexToAddress("0x469b39F9425C26baF6E782C89C11425F93a02814"),
		big.NewInt(2),
		common.HexToHash("0x9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a"),
		true,
	)
	s.rejectMsg = NewMessage(
		big.NewInt(1),
		common.HexToAddress("0x469b39F9425C26baF6E782C89C11425F93a02814"),
		big.NewInt(2),
		common.HexToHash("0x9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a"),
		false,
	)
}

func (s *MessageTestSuite) TestNewApproveSccMessage() {
	wantAbiPacked, _ := hex.DecodeString(strings.Join([]string{
		"0000000000000000000000000000000000000000000000000000000000000001",
		"469b39F9425C26baF6E782C89C11425F93a02814",
		"0000000000000000000000000000000000000000000000000000000000000002",
		"9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a",
		"01",
	}, ""))

	hash := crypto.Keccak256(wantAbiPacked)
	wantEip712Msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), string(hash))

	s.Equal(wantAbiPacked, s.approveMsg.AbiPacked)
	s.Equal(wantEip712Msg, s.approveMsg.Eip712Msg)
}

func (s *MessageTestSuite) TestNewRejectSccMessage() {
	wantAbiPacked, _ := hex.DecodeString(strings.Join([]string{
		"0000000000000000000000000000000000000000000000000000000000000001",
		"469b39F9425C26baF6E782C89C11425F93a02814",
		"0000000000000000000000000000000000000000000000000000000000000002",
		"9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a",
		"00",
	}, ""))

	hash := crypto.Keccak256(wantAbiPacked)
	wantEip712Msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), string(hash))

	s.Equal(wantAbiPacked, s.rejectMsg.AbiPacked)
	s.Equal(wantEip712Msg, s.rejectMsg.Eip712Msg)
}

func (s *MessageTestSuite) TestSignature() {
	got1, _ := s.approveMsg.Signature(s.signData)
	got2, _ := s.rejectMsg.Signature(s.signData)

	s.Equal(hexutil.MustDecode(
		"0x1718cfc352e84bf50ced8b0aaf8a8955fb038389223b289cca33bdd1bd72b7d0"+
			"29b5f6ebf983f38ddc85086b58d48b16637b8bf8929230eec38ab05595504a5b1c"), got1[:])
	s.Equal(hexutil.MustDecode(
		"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729"+
			"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef1b"), got2[:])
}

func (s *MessageTestSuite) TestEcrecover() {
	got1, _ := s.approveMsg.Ecrecover(
		hexutil.MustDecode(
			"0x1718cfc352e84bf50ced8b0aaf8a8955fb038389223b289cca33bdd1bd72b7d0" +
				"29b5f6ebf983f38ddc85086b58d48b16637b8bf8929230eec38ab05595504a5b1c"))
	got2, _ := s.rejectMsg.Ecrecover(
		hexutil.MustDecode(
			"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729" +
				"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef1b"))
	got3, _ := s.rejectMsg.Ecrecover(
		hexutil.MustDecode(
			"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729" +
				"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef10"))

	s.Equal(s.account.Address, got1)
	s.Equal(s.account.Address, got2)
	s.NotEqual(s.account.Address, got3)
}

func (s *MessageTestSuite) TestVerifySigner() {
	got1 := s.approveMsg.VerifySigner(
		hexutil.MustDecode(
			"0x1718cfc352e84bf50ced8b0aaf8a8955fb038389223b289cca33bdd1bd72b7d0"+
				"29b5f6ebf983f38ddc85086b58d48b16637b8bf8929230eec38ab05595504a5b1c"), s.account.Address)
	got2 := s.rejectMsg.VerifySigner(
		hexutil.MustDecode(
			"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729"+
				"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef1b"), s.account.Address)
	got3 := s.rejectMsg.VerifySigner(
		hexutil.MustDecode(
			"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729"+
				"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef10"), s.account.Address)

	s.Nil(got1)
	s.Nil(got2)
	s.Error(got3)
}

func (s *MessageTestSuite) TestL2OORollupHashSource() {
	want := ("0000000000000000000000000000000000000000000000000000000000000001" +
		"00000000000000000000000000000002" +
		"00000000000000000000000000000003")
	got := L2OORollupHashSource(
		common.HexToHash("0x1"), big.NewInt(2), big.NewInt(3))
	s.Equal(want, hex.EncodeToString(got))
}

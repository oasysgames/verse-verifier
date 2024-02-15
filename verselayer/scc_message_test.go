package verselayer

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/stretchr/testify/suite"
)

type SccSigTestSuite struct {
	suite.Suite

	b          *testhelper.TestBackend
	approveMsg *SccMessage
	rejectMsg  *SccMessage
}

func TestSccSig(t *testing.T) {
	suite.Run(t, new(SccSigTestSuite))
}

func (s *SccSigTestSuite) SetupSuite() {
	s.b = testhelper.NewTestBackend()
	s.approveMsg = NewSccMessage(
		big.NewInt(1),
		common.HexToAddress("0x469b39F9425C26baF6E782C89C11425F93a02814"),
		big.NewInt(2),
		common.HexToHash("0x9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a"),
		true,
	)
	s.rejectMsg = NewSccMessage(
		big.NewInt(1),
		common.HexToAddress("0x469b39F9425C26baF6E782C89C11425F93a02814"),
		big.NewInt(2),
		common.HexToHash("0x9daca4c5cecc1ad42a57af6209e26bb49cca77a1642ce2385824bd7c2b5cba0a"),
		false,
	)
}

func (s *SccSigTestSuite) TestNewApproveSccMessage() {
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

func (s *SccSigTestSuite) TestNewRejectSccMessage() {
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

func (s *SccSigTestSuite) TestSignature() {
	got1, _ := s.approveMsg.Signature(s.b.SignData)
	got2, _ := s.rejectMsg.Signature(s.b.SignData)

	s.Equal(hexutil.MustDecode(
		"0x1718cfc352e84bf50ced8b0aaf8a8955fb038389223b289cca33bdd1bd72b7d0"+
			"29b5f6ebf983f38ddc85086b58d48b16637b8bf8929230eec38ab05595504a5b1c"), got1[:])
	s.Equal(hexutil.MustDecode(
		"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729"+
			"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef1b"), got2[:])
}

func (s *SccSigTestSuite) TestEcrecover() {
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

	s.Equal(s.b.Signer(), got1)
	s.Equal(s.b.Signer(), got2)
	s.NotEqual(s.b.Signer(), got3)
}

func (s *SccSigTestSuite) TestVerifySigner() {
	got1 := s.approveMsg.VerifySigner(
		hexutil.MustDecode(
			"0x1718cfc352e84bf50ced8b0aaf8a8955fb038389223b289cca33bdd1bd72b7d0"+
				"29b5f6ebf983f38ddc85086b58d48b16637b8bf8929230eec38ab05595504a5b1c"), s.b.Signer())
	got2 := s.rejectMsg.VerifySigner(
		hexutil.MustDecode(
			"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729"+
				"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef1b"), s.b.Signer())
	got3 := s.rejectMsg.VerifySigner(
		hexutil.MustDecode(
			"0x821d05b483cc69c0f50beb8828b597ea632a8ac0552d579996526665150c5729"+
				"0111f891cb9a4f82ab95667bb9d025dd7592b3f8d5a2217e3d173ca21cb374ef10"), s.b.Signer())

	s.Nil(got1)
	s.Nil(got2)
	s.Error(got3)
}

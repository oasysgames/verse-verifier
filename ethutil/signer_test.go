package ethutil

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/account"
	"github.com/stretchr/testify/suite"
)

type SignerTestSuite struct {
	testhelper.Suite

	ksSigner   Signer
	privSigner Signer
}

func TestSigner(t *testing.T) {
	suite.Run(t, new(SignerTestSuite))
}

func (s *SignerTestSuite) SetupTest() {
	s.ksSigner = NewKeystoreSigner(account.DefaultWallet, account.DefaultAccount)
	s.privSigner = NewPrivateKeySigner(account.DefaultPrivateKey)
}

func (s *SignerTestSuite) TestFrom() {
	s.Equal("0x26fE8b0aE2d2CD413B09935e927729C6Ef93EDDd", s.ksSigner.From().Hex())
	s.Equal("0x26fE8b0aE2d2CD413B09935e927729C6Ef93EDDd", s.privSigner.From().Hex())
}

func (s *SignerTestSuite) TestSignData() {
	data := []byte("hello world")

	want, err := hex.DecodeString("e4c0c5d033676523e17ff1148b31d4ee2cff44a606d1bec8b5099121412df9132fefd79c1cc752da4ef4d6501d814b49148cd9ca7774b29c48190e5c7b41da4f01")
	s.NoError(err)

	got0, _ := s.ksSigner.SignData(data)
	got1, _ := s.privSigner.SignData(data)

	s.Equal(want, got0)
	s.Equal(want, got1)
}

func (s *SignerTestSuite) TestSignTx() {
	chainID := big.NewInt(12345)
	to := s.RandAddress()

	got0, err := s.ksSigner.SignTx(types.NewTransaction(0, to, common.Big0, 21_000, common.Big1, []byte{}), chainID)
	s.NoError(err)

	got1, err := s.ksSigner.SignTx(types.NewTransaction(0, to, common.Big0, 21_000, common.Big1, []byte{}), chainID)
	s.NoError(err)

	s.Equal(chainID, got0.ChainId())
	s.Equal(chainID, got1.ChainId())
	s.Equal(got0.Hash(), got1.Hash())

	from0, _ := types.Sender(types.LatestSignerForChainID(chainID), got0)
	from1, _ := types.Sender(types.LatestSignerForChainID(chainID), got1)

	s.Equal("0x26fE8b0aE2d2CD413B09935e927729C6Ef93EDDd", from0.Hex())
	s.Equal("0x26fE8b0aE2d2CD413B09935e927729C6Ef93EDDd", from1.Hex())
}

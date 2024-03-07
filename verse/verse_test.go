package verse

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	"github.com/stretchr/testify/suite"
)

type VerseTestSuite struct {
	testhelper.Suite

	logger log.Logger
	db     *database.Database
	rollupContract,
	verifyContract common.Address
	l1Client,
	l2Client ethutil.Client
	l1Signer ethutil.SignableClient

	verse        Verse
	verifiable   VerifiableVerse
	transactable TransactableVerse
}

func TestVerse(t *testing.T) {
	suite.Run(t, new(VerseTestSuite))
}

func (s *VerseTestSuite) SetupTest() {
	s.db, _ = database.NewDatabase(&config.Database{Path: ":memory:"})
	s.rollupContract = common.HexToAddress("0x1")
	s.verifyContract = common.HexToAddress("0x2")
	s.l1Client = backend.NewBackend(nil, 0)
	s.l2Client = backend.NewBackend(nil, 0)
	s.l1Signer = backend.NewSignableBackend(nil, nil, nil)

	factory := newVerseFactory(func(v Verse) Verse { return v })

	s.verse = factory(s.db, s.l1Client, s.rollupContract)
	s.verifiable = s.verse.WithVerifiable(s.l2Client)
	s.transactable = s.verifiable.WithTransactable(s.l1Signer, s.verifyContract)
}

func (s *VerseTestSuite) TestLogger() {
	s.Equal(s.logger, s.verse.Logger(s.logger))
}

func (s *VerseTestSuite) TestDB() {
	s.Equal(s.db, s.verse.DB())
}

func (s *VerseTestSuite) TestL1Client() {
	s.Equal(s.l1Client, s.verse.L1Client())
}

func (s *VerseTestSuite) TestRollupContract() {
	s.Equal(s.rollupContract, s.verse.RollupContract())
}

func (s *VerseTestSuite) TestL2Client() {
	s.Equal(s.l2Client, s.verifiable.L2Client())
}

func (s *VerseTestSuite) TestL1Signer() {
	s.Equal(s.l1Signer, s.transactable.L1Signer())
}

func (s *VerseTestSuite) TestVerifyContract() {
	s.Equal(s.verifyContract, s.transactable.VerifyContract())
}

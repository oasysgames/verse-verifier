package database

import (
	"testing"

	"gorm.io/gorm"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/stretchr/testify/suite"
)

func TestBlockDatabase(t *testing.T) {
	suite.Run(t, new(BlockDatabaseTestSuite))
}

type BlockDatabaseTestSuite struct {
	testhelper.Suite

	db    *BlockDatabase
	rawdb *gorm.DB
}

func (s *BlockDatabaseTestSuite) SetupTest() {
	db, err := NewDatabase(":memory:")
	if err != nil {
		panic(err)
	}
	s.db = db.Block
	s.rawdb = db.db

	for _, number := range s.Shuffle(s.Range(0, 50)) {
		s.rawdb.Model(&Block{}).
			Create(&Block{Number: uint64(number + 1), Hash: s.ItoHash(number + 1)})
	}
}

func (s *BlockDatabaseTestSuite) TestFind() {
	got, _ := s.db.Find(10)
	s.Equal(uint64(10), got.Number)
	s.Equal(s.ItoHash(10), got.Hash)
	s.Equal(false, got.LogCollected)
}

func (s *BlockDatabaseTestSuite) TestFindHighest() {
	got, _ := s.db.FindHighest()
	s.Equal(uint64(50), got.Number)
	s.Equal(s.ItoHash(50), got.Hash)
	s.Equal(false, got.LogCollected)
}

func (s *BlockDatabaseTestSuite) TestFindUncollecteds() {
	assertGots := func(gots []*Block, expNumbers []int) {
		s.Equal(len(expNumbers), len(gots))

		for _, expNumber := range expNumbers {
			got := gots[0]
			gots = gots[1:]

			s.Equal(uint64(expNumber), got.Number)
			s.Equal(s.ItoHash(expNumber), got.Hash)
			s.Equal(false, got.LogCollected)
		}
	}

	// limit = 10
	gots, _ := s.db.FindUncollecteds(10)
	assertGots(gots, s.Range(1, 10+1))

	// limit = 100
	gots, _ = s.db.FindUncollecteds(100)
	assertGots(gots, s.Range(1, 50+1))

	s.rawdb.Model(&Block{}).
		Where("number <= 25").Update("log_collected", true)

	// limit = 10
	gots, _ = s.db.FindUncollecteds(10)
	assertGots(gots, s.Range(26, 35+1))

	// limit = 100
	gots, _ = s.db.FindUncollecteds(100)
	assertGots(gots, s.Range(26, 50+1))
}

func (s *BlockDatabaseTestSuite) TestSave() {
	number := uint64(100)

	s.db.SaveNewBlock(number, s.ItoHash(int(number)))

	got, _ := s.db.Find(number)
	s.Equal(number, got.Number)
	s.Equal(s.ItoHash(int(number)), got.Hash)
	s.Equal(false, got.LogCollected)
}

func (s *BlockDatabaseTestSuite) TestSaveLogCollected() {
	number := uint64(10)

	got, _ := s.db.Find(number)
	s.Equal(number, got.Number)
	s.Equal(s.ItoHash(int(number)), got.Hash)
	s.Equal(false, got.LogCollected)

	s.db.SaveLogCollected(got.Number)

	got, _ = s.db.Find(number)
	s.Equal(number, got.Number)
	s.Equal(s.ItoHash(int(number)), got.Hash)
	s.Equal(true, got.LogCollected)
}

func (s *BlockDatabaseTestSuite) TestDelete() {
	number := uint64(10)

	got, _ := s.db.Find(number)
	s.Equal(number, got.Number)

	s.db.Delete(number)

	_, err := s.db.Find(number)
	s.ErrorIs(err, ErrNotFound)
}

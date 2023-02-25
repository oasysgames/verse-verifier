package database

import (
	"math/rand"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"gorm.io/gorm"
)

type DatabaseTestSuite struct {
	testhelper.Suite

	db    *Database
	rawdb *gorm.DB
}

func (s *DatabaseTestSuite) SetupTest() {
	// Setup database
	db, err := NewDatabase(":memory:")
	if err != nil {
		panic(err)
	}
	s.db = db
	s.rawdb = db.db
}

func (s *DatabaseTestSuite) createSigner() *Signer {
	signer := &Signer{Address: s.RandAddress()}
	s.NoDBError(s.rawdb.Create(signer))
	return signer
}

func (s *DatabaseTestSuite) createSCC() *OptimismScc {
	scc := &OptimismScc{Address: s.RandAddress()}
	s.NoDBError(s.rawdb.Create(scc))
	return scc
}

func (s *DatabaseTestSuite) createState(scc *OptimismScc, index int) *OptimismState {
	state := &OptimismState{
		OptimismScc:       *scc,
		BatchIndex:        uint64(index),
		BatchRoot:         s.ItoHash(index),
		BatchSize:         uint64(rand.Intn(99)),
		PrevTotalElements: uint64(rand.Intn(99)),
		ExtraData:         s.RandBytes(),
	}
	s.NoDBError(s.rawdb.Create(state))
	return state
}

func (s *DatabaseTestSuite) createSignature(
	signer *Signer,
	scc *OptimismScc,
	index int,
) *OptimismSignature {
	sig := &OptimismSignature{
		ID:          util.ULID(nil).String(),
		PreviousID:  util.ULID(nil).String(),
		Signer:      *signer,
		OptimismScc: *scc,
		BatchIndex:  uint64(index),
		BatchRoot:   s.RandHash(),
		Signature:   RandSignature(),
	}
	s.NoDBError(s.rawdb.Create(sig))
	return sig
}

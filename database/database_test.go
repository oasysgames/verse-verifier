package database

import (
	"math/rand"
	"time"

	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

type DatabaseTestSuite struct {
	testhelper.Suite

	db *Database
}

func (s *DatabaseTestSuite) SetupTest() {
	// Setup database
	db, err := NewDatabase(&config.Database{Path: ":memory:"})
	if err != nil {
		panic(err)
	}
	s.db = db
}

func (s *DatabaseTestSuite) createSigner() *Signer {
	signer := &Signer{Address: s.RandAddress()}
	s.NoDBError(s.db.rawdb.Create(signer))
	return signer
}

func (s *DatabaseTestSuite) createContract() *OptimismContract {
	contract := &OptimismContract{Address: s.RandAddress()}
	s.NoDBError(s.db.rawdb.Create(contract))
	return contract
}

func (s *DatabaseTestSuite) createState(contract *OptimismContract, index int) *OptimismState {
	state := &OptimismState{
		Contract:          *contract,
		BatchIndex:        uint64(index),
		BatchRoot:         s.ItoHash(index),
		BatchSize:         uint64(rand.Intn(99)),
		PrevTotalElements: uint64(rand.Intn(99)),
		ExtraData:         s.RandBytes(),
	}
	s.NoDBError(s.db.rawdb.Create(state))
	return state
}

func (s *DatabaseTestSuite) createProposal(l2oo *OptimismContract, l2OutputIndex int) *OpstackProposal {
	proposal := &OpstackProposal{
		Contract:      *l2oo,
		L2OutputIndex: uint64(l2OutputIndex),
		OutputRoot:    s.RandHash(),
		L2BlockNumber: uint64(rand.Intn(99)),
		L1Timestamp:   uint64(time.Now().Unix()),
	}
	s.NoDBError(s.db.rawdb.Create(proposal))
	return proposal
}

func (s *DatabaseTestSuite) createSignature(
	signer *Signer,
	contract *OptimismContract,
	index int,
) *OptimismSignature {
	// avoiding duplication of ULID
	time.Sleep(time.Millisecond)

	sig := &OptimismSignature{
		ID:          util.ULID(nil).String(),
		PreviousID:  util.ULID(nil).String(),
		Signer:      *signer,
		Contract:    *contract,
		RollupIndex: uint64(index),
		RollupHash:  s.RandHash(),
		Signature:   RandSignature(),
	}
	s.NoDBError(s.db.rawdb.Create(sig))
	return sig
}

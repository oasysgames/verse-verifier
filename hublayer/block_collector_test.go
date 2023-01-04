package hublayer

import (
	"context"
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/stretchr/testify/suite"
)

type BlockCollectorTestSuite struct {
	testhelper.Suite

	db      *database.Database
	backend *testhelper.TestBackend
}

func TestBlockCollector(t *testing.T) {
	suite.Run(t, new(BlockCollectorTestSuite))
}

func (s *BlockCollectorTestSuite) SetupTest() {
	s.db, _ = database.NewDatabase(":memory:")
	s.backend = testhelper.NewTestBackend()
}

func (s *BlockCollectorTestSuite) TestCollectNewBlocks() {
	worker := NewBlockCollector(s.db, s.backend, 0)

	// mining 5 blocks
	var wants []*types.Header
	for range s.Range(0, 5) {
		wants = append(wants, s.backend.Mining())
	}

	// collect blocks
	for {
		if reached := worker.work(context.Background()); reached {
			break
		}
	}

	// assert
	gots, _ := s.db.Block.FindUncollecteds(20)
	s.Len(gots, len(wants))
	for i, want := range wants {
		s.Equal(want.Number.Uint64(), gots[i].Number)
		s.Equal(want.Hash(), gots[i].Hash)
	}
}

func (s *BlockCollectorTestSuite) TestHandleReorganization() {
	// mining 10 blocks
	var mined []*types.Header
	for range s.Range(0, 10) {
		mined = append(mined, s.backend.Mining())
	}

	// simulate chain reorganization
	rb := newReorgBackend(s.backend, mined)
	worker := NewBlockCollector(s.db, rb, 0)

	// collect blocks
	for {
		if reached := worker.work(context.Background()); reached {
			break
		}
	}

	// assert
	gots, _ := s.db.Block.FindUncollecteds(20)
	s.Len(gots, len(mined))
	for i, want := range mined {
		s.Equal(want.Number.Uint64(), gots[i].Number)
		s.Equal(want.Hash(), gots[i].Hash)
	}

	// reorg occurred
	rb.reorg()
	for {
		if reached := worker.work(context.Background()); reached {
			break
		}
	}

	// assert
	gots, _ = s.db.Block.FindUncollecteds(20)
	s.Len(gots, len(rb.reorged))
	for i, want := range mined[:len(mined)/2-1] {
		s.Equal(want.Number.Uint64(), gots[i].Number)
		s.Equal(want.Hash(), gots[i].Hash)
	}
	for i, want := range rb.reorged {
		s.Equal(want.Number.Uint64(), gots[i].Number)
		s.Equal(want.Hash(), gots[i].Hash)
	}
}

type reorgBackend struct {
	*testhelper.TestBackend

	mu *sync.Mutex
	do bool

	mined, reorged []*types.Header
}

func newReorgBackend(tb *testhelper.TestBackend, mined []*types.Header) *reorgBackend {
	reorged := make([]*types.Header, len(mined))
	for n := 1; n <= len(mined); n++ {
		cpy := types.CopyHeader(mined[n-1])

		if n == len(mined)/2 {
			// reorganization start block.
			cpy.Difficulty.Add(cpy.Difficulty, common.Big1)
		} else if n > len(mined)/2 {
			cpy.ParentHash = reorged[n-2].Hash()
		}

		reorged[n-1] = cpy
	}

	return &reorgBackend{
		TestBackend: tb,
		mu:          &sync.Mutex{},
		mined:       mined,
		reorged:     reorged,
	}
}

func (r *reorgBackend) reorg() {
	r.mu.Lock()
	defer r.mu.Unlock()

	h := r.TestBackend.Mining()
	h.ParentHash = r.reorged[len(r.reorged)-1].Hash()

	r.reorged = append(r.reorged, h)
	r.do = true
}

func (r *reorgBackend) HeaderByNumber(_ context.Context, b *big.Int) (*types.Header, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	n := int(b.Uint64())
	if r.do {
		if n <= len(r.reorged) {
			return r.reorged[n-1], nil
		}
	} else if n <= len(r.mined) {
		return r.mined[n-1], nil
	}

	return nil, ethereum.NotFound
}
package hublayer

import (
	"context"
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	"github.com/stretchr/testify/suite"
)

type BlockCollectorTestSuite struct {
	testhelper.Suite

	db      *database.Database
	backend *backend.Backend
}

func TestBlockCollector(t *testing.T) {
	suite.Run(t, new(BlockCollectorTestSuite))
}

func (s *BlockCollectorTestSuite) SetupTest() {
	s.db, _ = database.NewDatabase(&config.Database{Path: ":memory:"})
	s.backend = backend.NewBackend(nil, 0)
}

func (s *BlockCollectorTestSuite) TestCollectNewBlocks() {
	worker := NewBlockCollector(&config.Verifier{Interval: 0, BlockLimit: 2}, s.db, s.backend)

	// mining 5 blocks
	var wants []*types.Header
	for range s.Range(0, 5) {
		wants = append(wants, s.backend.Mining())
	}

	// collect blocks
	worker.work(context.Background())

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
	reorgedBlock := uint64(5)
	rb := newReorgBackend(s.backend, mined, reorgedBlock)
	worker := NewBlockCollector(&config.Verifier{Interval: 0, BlockLimit: 2}, s.db, rb)

	// collect blocks
	worker.work(context.Background())

	// assert
	gots, _ := s.db.Block.FindUncollecteds(20)
	s.Len(gots, len(mined))
	for i, want := range mined {
		s.Equal(want.Number.Uint64(), gots[i].Number)
		s.Equal(want.Hash(), gots[i].Hash)
	}

	// reorg occurred
	rb.reorg()
	worker.work(context.Background())

	// assert
	gots, _ = s.db.Block.FindUncollecteds(20)
	s.Equal(reorgedBlock-1, gots[len(gots)-1].Number)
	for i, want := range mined[:reorgedBlock-1] {
		s.Equal(want.Number.Uint64(), gots[i].Number)
		s.Equal(want.Hash(), gots[i].Hash)
	}

	// collect reorganized blocks
	worker.work(context.Background())

	// assert
	gots, _ = s.db.Block.FindUncollecteds(20)
	s.Len(gots, len(rb.reorged))
	for i, want := range mined[:reorgedBlock-1] {
		s.Equal(want.Number.Uint64(), gots[i].Number)
		s.Equal(want.Hash(), gots[i].Hash)
	}
	for i, want := range rb.reorged {
		s.Equal(want.Number.Uint64(), gots[i].Number)
		s.Equal(want.Hash(), gots[i].Hash)
	}
}

type reorgBackend struct {
	*backend.Backend

	mu *sync.Mutex
	do bool

	mined, reorged []*types.Header
}

func newReorgBackend(
	tb *backend.Backend,
	mined []*types.Header,
	reorgedBlock uint64,
) *reorgBackend {
	b := &reorgBackend{
		Backend: tb,
		mu:      &sync.Mutex{},
		mined:   mined,
		reorged: make([]*types.Header, len(mined)),
	}

	for i, src := range mined {
		cpy := types.CopyHeader(src)
		num := cpy.Number.Uint64()

		if num == reorgedBlock {
			// reorganization start block.
			cpy.Difficulty.Add(cpy.Difficulty, common.Big1)
		} else if num > reorgedBlock {
			cpy.ParentHash = b.reorged[i-1].Hash()
		}

		b.reorged[i] = cpy
	}

	return b
}

func (r *reorgBackend) reorg() {
	r.mu.Lock()
	defer r.mu.Unlock()

	h := r.Backend.Mining()
	h.ParentHash = r.reorged[len(r.reorged)-1].Hash()

	r.reorged = append(r.reorged, h)
	r.do = true
}

func (r *reorgBackend) NewBatchHeaderClient() (ethutil.BatchHeaderClient, error) {
	return &backend.BatchHeaderClient{Client: r}, nil
}

func (r *reorgBackend) HeaderByNumber(_ context.Context, b *big.Int) (*types.Header, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if b == nil {
		if r.do {
			return r.reorged[len(r.reorged)-1], nil
		}
		return r.mined[len(r.mined)-1], nil
	}

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

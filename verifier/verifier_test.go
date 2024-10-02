package verifier

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
	"github.com/stretchr/testify/suite"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
)

type VerifierTestSuite struct {
	backend.BackendSuite

	verifier *Verifier
	cfg      *config.Verifier
	task     verse.VerifiableVerse
	newSigP2P,
	unverifiedSigP2P *MockP2P
}

func TestVerifier(t *testing.T) {
	suite.Run(t, new(VerifierTestSuite))
}

type MockP2P struct {
	sigsCh chan []*database.OptimismSignature
}

func (m *MockP2P) PublishSignatures(ctx context.Context, sigs []*database.OptimismSignature) error {
	m.sigsCh <- sigs
	return nil
}

func (s *VerifierTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()

	s.cfg = &config.Verifier{
		Interval:              time.Millisecond * 50,
		StateCollectLimit:     3,
		StateCollectTimeout:   time.Second,
		Confirmations:         2,
		MaxLogFetchBlockRange: 5760,
		MaxIndexDiff:          3,
	}
	s.verifier = NewVerifier(s.cfg, s.DB, nil, s.SignableHub)
	s.task = verse.NewOPLegacy(s.DB, s.Hub, s.SCCAddr).WithVerifiable(s.Verse)

	s.newSigP2P = &MockP2P{sigsCh: make(chan []*database.OptimismSignature)}
	s.unverifiedSigP2P = &MockP2P{sigsCh: make(chan []*database.OptimismSignature)}
	s.verifier.newSigP2P = s.newSigP2P
	s.verifier.unverifiedSigP2P = s.unverifiedSigP2P
}

func (s *VerifierTestSuite) TestStartVerifier() {
	rollups := s.Range(0, 15) // index: 0~14
	nextIndex := 4
	emitInterval := 5

	// Set next index.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.TSCC.SetNextIndex(s.SignableHub.TransactOpts(ctx), big.NewInt(int64(nextIndex)))
	s.Hub.Mining()

	// Adjust latest block.
	s.Hub.MiningTo(100)

	// Send transactions to the Verse-Layer to generate a Merkle Root.
	merkleRoots := make([][32]byte, len(rollups))
	signatures := make([][65]byte, len(rollups))
	for batchIdx := range rollups {
		// Calc Merkle Root
		elements := make([][32]byte, 5)
		for i, header := range s.sendVerseTransactions(5) {
			elements[i] = header.Root
		}
		if merkleRoot, err := verse.CalcMerkleRoot(elements); s.NoError(err) {
			merkleRoots[batchIdx] = merkleRoot
		}

		// Calc signature of ethereum signed message
		ethmsg := ethutil.NewMessage(
			s.verifier.l1Signer.ChainID(),
			s.SCCAddr,
			big.NewInt(int64(batchIdx)),
			merkleRoots[batchIdx],
			true,
		)
		if sig, err := ethmsg.Signature(s.SignableHub.SignData); s.NoError(err) {
			signatures[batchIdx] = sig
		}
	}
	// Send generated Merkle Root to the Hub-Layer as a rollup event.
	for batchIdx, merkleRoot := range merkleRoots {
		for i := range s.Range(0, emitInterval) {
			if i+1 == emitInterval {
				_, err := s.TSCC.EmitStateBatchAppended(
					s.SignableHub.TransactOpts(context.Background()),
					big.NewInt(int64(batchIdx)),
					merkleRoot, big.NewInt(int64(5)),
					big.NewInt(int64(batchIdx*5)), []byte(fmt.Sprintf("test-%d", batchIdx)),
				)
				s.NoError(err)
			}
			s.Hub.Mining()
		}
	}

	assertP2P := func(batchIndex uint64, got *database.OptimismSignature) {
		s.Equal(batchIndex, got.RollupIndex)
		s.Equal(merkleRoots[batchIndex][:], got.RollupHash[:])
		s.Equal(true, got.Approved)
		s.Equal(signatures[batchIndex][:], got.Signature[:])
	}
	assertDB := func(batchIndex uint64) {
		got, _ := s.DB.OPSignature.Find(nil, nil, nil, &batchIndex, 1, 0)
		s.Equal(batchIndex, got[0].RollupIndex)
		s.Equal(merkleRoots[batchIndex][:], got[0].RollupHash[:])
		s.Equal(true, got[0].Approved)
		s.Equal(signatures[batchIndex][:], got[0].Signature[:])
	}

	s.Hub.Minings(10)

	// Run a task to check that the received unverified
	// signatures are greater than or equal to NextIndex.
	var (
		unverifiedWait sync.Mutex
		unverifiedSigs []*database.OptimismSignature
	)
	go func() {
		for {
			sigs := <-s.unverifiedSigP2P.sigsCh
			unverifiedWait.Unlock()
			unverifiedSigs = append(unverifiedSigs, sigs...)

			s.Greater(len(sigs), 0)
			for i, sig := range sigs {
				assertP2P(uint64(nextIndex+i), sig)
			}
		}
	}()

	// start verifier by adding task
	s.verifier.AddTask(ctx, s.task, 0)

	// wait for verification
	sigs := <-s.newSigP2P.sigsCh
	unverifiedWait.Lock()

	// No prior signature than verified index should be sent.
	expSigs := rollups[nextIndex : 1+nextIndex+s.cfg.MaxIndexDiff]
	s.Len(sigs, len(expSigs))

	// Since increasing the NextIndex causes signatures with indices less
	// than it to be deleted from the database, we should perform assertions beforehand.
	for _, batchIndex := range expSigs {
		assertDB(uint64(batchIndex))
	}

	// Additionally verified each time NextIndex is incremented.
	// Each time NextIndex is incremented, one event should be verified.
	for range s.Range(0, 2) {
		nextIndex++
		s.TSCC.SetNextIndex(s.SignableHub.TransactOpts(ctx), big.NewInt(int64(nextIndex)))
		s.Hub.Minings(1 + s.cfg.Confirmations)

		sigs = append(sigs, <-s.newSigP2P.sigsCh...)
		unverifiedWait.Lock()

		expSigs = append(expSigs, rollups[nextIndex+s.cfg.MaxIndexDiff])

		assertDB(uint64(expSigs[len(expSigs)-1]))
	}
	s.Len(sigs, len(expSigs))

	// Since `NextIndex + MaxIndexDiff` exceeds the maximum index,
	// it will be verified all the way up to L1Head.
	nextIndex = rollups[len(rollups)-s.cfg.MaxIndexDiff]
	s.TSCC.SetNextIndex(s.SignableHub.TransactOpts(ctx), big.NewInt(int64(nextIndex)))
	s.Hub.Minings(1 + s.cfg.Confirmations)

	sigs = append(sigs, <-s.newSigP2P.sigsCh...)
	unverifiedWait.Lock()

	expSigs = append(expSigs, rollups[nextIndex:]...)

	s.Len(sigs, len(expSigs))
	s.GreaterOrEqual(len(unverifiedSigs), len(expSigs))

	for _, batchIndex := range rollups[nextIndex:] {
		assertDB(uint64(batchIndex))
	}

	// Assertions of signatures received via P2P.
	for i, batchIndex := range expSigs {
		assertP2P(uint64(batchIndex), sigs[i])
	}

	// Check if the old signature has been deleted from the DB.
	rows, _ := s.DB.OPSignature.Find(nil, nil, nil, nil, 100, 0)
	s.Len(rows, len(rollups[nextIndex-1:]))
	for _, batchIndex := range rollups[nextIndex-1:] {
		assertDB(uint64(batchIndex))
	}
}

func (s *VerifierTestSuite) TestRetryBackoff() {
	verifier := &Verifier{
		cfg: &config.Verifier{
			MaxRetryBackoff: time.Minute,
			RetryTimeout:    time.Millisecond * 100,
		},
	}

	incr, decr := verifier.retryBackoff()

	wait := time.Millisecond * 10
	for i := range s.Range(0, 10) {
		delay, remain, attempts := incr()

		s.Equal(i+1, attempts)
		s.Less(remain, verifier.cfg.RetryTimeout-wait*time.Duration(i))

		switch i {
		case 0:
			s.Equal(delay, time.Millisecond*100)
		case 1:
			s.Equal(delay, time.Millisecond*800)
		case 2:
			s.Equal(delay, time.Millisecond*6400)
		case 3:
			s.Equal(delay, time.Millisecond*51200)
		default: // i >= 4
			s.Equal(delay, time.Minute)
		}

		time.Sleep(wait)
	}

	_, remain, _ := incr()
	s.Equal(remain, time.Duration(0))

	for _, decrs := range s.Range(1, 10) {
		for range s.Range(0, decrs) {
			decr()
		}
		delay, remain, _ := incr()

		s.Equal(remain, time.Duration(0))

		switch decrs {
		case 1:
			s.Equal(delay, time.Millisecond*51200)
		case 2:
			s.Equal(delay, time.Millisecond*6400)
		case 3:
			s.Equal(delay, time.Millisecond*800)
		default:
			s.Equal(delay, time.Millisecond*100)
		}

		// Advance to the limit
		for range s.Range(0, 10) {
			incr()
		}
	}
}

func (s *VerifierTestSuite) TestDetermineMaxEnd() {
	rollups := s.Range(0, 10) // index: 0~9
	nextIndex := 0
	emitInterval := 5

	// Set next index.
	ctx := context.Background()
	s.TSCC.SetNextIndex(s.SignableHub.TransactOpts(ctx), big.NewInt(int64(nextIndex)))
	s.Hub.Mining()

	// Mining blocks and emitting events.
	emittedBlocks := make([]uint64, len(rollups))
	for _, batchIdx := range rollups {
		for i := range s.Range(0, emitInterval) {
			if i+1 == emitInterval {
				_, err := s.TSCC.EmitStateBatchAppended(
					s.SignableHub.TransactOpts(ctx),
					big.NewInt(int64(batchIdx)),
					s.RandHash(),
					big.NewInt(int64(5)),
					big.NewInt(int64(batchIdx*5)),
					[]byte(fmt.Sprintf("test-%d", batchIdx)))
				s.NoError(err)

				emittedBlocks[batchIdx] = s.Hub.Mining().Number.Uint64()
			} else {
				s.Hub.Mining()
			}
		}
	}
	latest := s.Hub.Minings(10)[9]

	got, _ := s.verifier.determineMaxEnd(ctx, s.task, uint64(nextIndex))
	s.Equal(emittedBlocks[nextIndex+s.cfg.MaxIndexDiff], got)

	// Increase NextIndex to 1.
	nextIndex = 1
	got, _ = s.verifier.determineMaxEnd(ctx, s.task, uint64(nextIndex))
	s.Equal(emittedBlocks[nextIndex+s.cfg.MaxIndexDiff], got)

	// Increase NextIndex to 3.
	nextIndex = 3
	got, _ = s.verifier.determineMaxEnd(ctx, s.task, uint64(nextIndex))
	s.Equal(emittedBlocks[nextIndex+s.cfg.MaxIndexDiff], got)

	// Increase NextIndex to 6.
	// The Key point is that `6+MaxIndexDiff`` has not yet exceeded the maximum index.
	nextIndex = 6
	got, _ = s.verifier.determineMaxEnd(ctx, s.task, uint64(nextIndex))
	s.Equal(emittedBlocks[nextIndex+s.cfg.MaxIndexDiff], got)

	// Increase NextIndex to 7.
	// Since `7+MaxIndexDiff` exceeds the maximum index, `L1Head-Confirmations` should be returned.
	nextIndex = 7
	got, _ = s.verifier.determineMaxEnd(ctx, s.task, uint64(nextIndex))
	s.Equal(latest.Number.Uint64()-uint64(s.cfg.Confirmations), got)
}

func (s *VerifierTestSuite) sendVerseTransactions(count int) (headers []*types.Header) {
	ctx := context.Background()
	to := common.HexToAddress("0x09ad74977844F513E61AdE2B50b0C06268A4f6d7")

	nonce, err := s.SignableVerse.PendingNonceAt(ctx, s.SignableVerse.Signer())
	s.NoError(err)

	for i := 0; i < count; i++ {
		gasPrice, err := s.SignableVerse.BaseGasPrice(ctx, nil)
		s.Nil(err)

		unsigned := types.NewTransaction(
			nonce+uint64(i), to, common.Big0, 21_000, gasPrice, nil)
		signed, err := s.SignableVerse.SignTx(unsigned)
		s.Nil(err)

		err = s.SignableVerse.SendTransaction(ctx, signed)
		s.NoError(err)

		h, err := s.SignableVerse.HeaderByHash(ctx, s.SignableVerse.Commit())
		s.NoError(err)

		headers = append(headers, h)
	}
	return headers
}

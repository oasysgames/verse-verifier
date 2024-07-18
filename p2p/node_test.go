package p2p

import (
	"context"
	"math/big"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	pb "github.com/oasysgames/oasys-optimism-verifier/proto/p2p/v1/gen"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/stretchr/testify/suite"
)

func TestNode(t *testing.T) {
	suite.Run(t, new(NodeTestSuite))
}

type NodeTestSuite struct {
	testhelper.Suite

	baseTime     time.Time
	stakemanager *stakemanager.Cache
	b0, b1, b2   *backend.SignableBackend

	signer0,
	signer1,
	signer2,
	contract0,
	contract1,
	contract2 common.Address

	sigs map[common.Address]map[common.Address][]*database.OptimismSignature

	bootnode, node1, node2 *Node
}

func (s *NodeTestSuite) SetupTest() {
	s.baseTime = time.Now().UTC()
	s.b0 = backend.NewSignableBackend(nil, nil)
	s.b1 = s.b0.WithNewAccount()
	s.b2 = s.b0.WithNewAccount()
	s.signer0 = s.b0.Signer()
	s.signer1 = s.b1.Signer()
	s.signer2 = s.b2.Signer()
	s.contract0 = s.RandAddress()
	s.contract1 = s.RandAddress()
	s.contract2 = s.RandAddress()
	s.sigs = map[common.Address]map[common.Address][]*database.OptimismSignature{}

	backends := []*backend.SignableBackend{s.b0, s.b1, s.b2}
	signers := []common.Address{s.signer0, s.signer1}
	contracts := []common.Address{s.contract0, s.contract1}

	// setup stakemanager mock
	sm := &testhelper.StakeManagerMock{}
	s.stakemanager = stakemanager.NewCache(sm)
	for _, signer := range signers {
		sm.Owners = append(sm.Owners, s.RandAddress())
		sm.Operators = append(sm.Operators, signer)
		sm.Stakes = append(sm.Stakes, ethutil.TenMillionOAS)
		sm.Candidates = append(sm.Candidates, true)
	}
	s.stakemanager.Refresh(context.Background())

	// setup libp2p
	s.bootnode = s.newWorker([]string{})
	bootnodes := []string{s.bootnode.cfg.Listens[0] + "/p2p/" + s.bootnode.h.ID().String()}
	s.node1 = s.newWorker(bootnodes)
	s.node2 = s.newWorker(bootnodes)

	// create sample records
	for _, node := range []*Node{s.node1, s.node2} {
		node.db.Signer.FindOrCreate(s.signer0)
		node.db.Signer.FindOrCreate(s.signer1)
		node.db.Signer.FindOrCreate(s.signer2)
		node.db.OPContract.FindOrCreate(s.contract0)
		node.db.OPContract.FindOrCreate(s.contract1)
		node.db.OPContract.FindOrCreate(s.contract2)
	}

	sizes := [][]int{
		{
			50,  // signer0, contract0
			100, // signer0, contract1
		},
		{
			150, // signer1, contract0
			200, // signer1, contract1
		},
	}
	for i, signer := range signers {
		backend := backends[i]
		s.sigs[signer] = map[common.Address][]*database.OptimismSignature{}

		for j, contract := range contracts {
			s.sigs[signer][contract] = []*database.OptimismSignature{}

			for _, rollupIndex := range s.Range(0, sizes[i][j]) {
				rollupHash := s.genStateRoot(contract[:], rollupIndex)
				approved := true

				msg := ethutil.NewMessage(
					backend.ChainID(),
					contract,
					big.NewInt(int64(rollupIndex)),
					util.BytesToBytes32(rollupHash[:]),
					approved,
				)
				sigbin, _ := msg.Signature(backend.SignData)

				sig, _ := s.node1.db.OPSignature.Save(
					nil, nil,
					signer,
					contract,
					uint64(rollupIndex),
					rollupHash,
					approved,
					sigbin,
				)
				s.sigs[signer][contract] = append(s.sigs[signer][contract], sig)
			}
		}
	}
}

func (s *NodeTestSuite) TestHandleOptimismSignatureExchangeFromPubSub() {
	// succeed to save signature
	msg := toProtoBufSig(s.sigs[s.signer0][s.contract0][49])
	saved := s.node2.handleOptimismSignatureExchangeFromPubSub(context.Background(), s.node1.h.ID(), msg)
	s.True(saved)
	got, err := s.node2.db.OPSignature.FindByID(msg.Id)
	s.NoError(err)
	s.Equal(msg.Id, got.ID)

	// saving duplicate signature
	saved = s.node2.handleOptimismSignatureExchangeFromPubSub(context.Background(), s.node1.h.ID(), msg)
	s.False(saved)

	// saving too old signature (no pruneRollupIndexDepth)
	msg = toProtoBufSig(s.sigs[s.signer0][s.contract0][0])
	saved = s.node2.handleOptimismSignatureExchangeFromPubSub(context.Background(), s.node1.h.ID(), msg)
	s.True(saved)
}

func (s *NodeTestSuite) TestHandleOptimismSignatureExchangeRequests() {
	wantss := [][]struct {
		signer       common.Address
		contract     common.Address
		batchIndexes []int
	}{
		{
			{s.signer0, s.contract0, s.Range(0, 50)},
			{s.signer0, s.contract1, s.Range(0, 50)},
		},
		{
			{s.signer0, s.contract1, s.Range(50, 100)},
		},
		{
			{s.signer1, s.contract1, s.Range(100, 200)},
		},
	}

	// send message
	st, _ := s.node2.h.NewStream(context.Background(), s.node1.h.ID(), streamProtocol)
	writeStream(st, &pb.Stream{Body: &pb.Stream_OptimismSignatureExchange{
		OptimismSignatureExchange: &pb.OptimismSignatureExchange{
			Requests: []*pb.OptimismSignatureExchange_Request{
				{IdAfter: "", Signer: s.signer0[:]},
				{IdAfter: s.sigs[s.signer1][s.contract1][100].ID, Signer: s.signer1[:]},
				{IdAfter: "", Signer: s.signer2[:]},
			},
		},
	}})

	var reads []*pb.Stream
	for _, wants := range wantss {
		for range wants {
			// read message
			m, _ := readStream(st)
			reads = append(reads, m)

			// send receipt notify
			writeStream(st, &pb.Stream{Body: &pb.Stream_Misc{Misc: misc_SIGRECEIVED}})
		}
	}
	closeStream(st)

	// assert
	s.Len(reads, len(wantss)+1)
	s.NotNil(reads[len(reads)-1])
	for i, wants := range wantss {
		gots := reads[i].GetOptimismSignatureExchange().Responses

		for _, want := range wants {
			for _, bi := range want.batchIndexes {
				got := gots[0]
				gots = gots[1:]

				wantsig := s.sigs[want.signer][want.contract][bi]
				s.Equal(wantsig.ID, got.Id)
				s.Equal(wantsig.Signer.Address[:], got.Signer)
				s.Equal(wantsig.Contract.Address[:], got.Contract)
				s.Equal(wantsig.RollupIndex, got.RollupIndex)
				s.Equal(wantsig.RollupHash[:], got.RollupHash)
				s.Equal(wantsig.Approved, got.Approved)
				s.Equal(wantsig.Signature[:], got.Signature)
			}
		}

		s.Len(gots, 0)
	}
}

func (s *NodeTestSuite) TestHandleOptimismSignatureExchangeResponses() {
	cases := []struct {
		b    *backend.SignableBackend
		want *pb.OptimismSignature
	}{
		{
			s.b0,
			&pb.OptimismSignature{
				Id:          util.ULID(nil).String(),
				PreviousId:  s.sigs[s.signer0][s.contract0][49].ID,
				Contract:    s.contract0[:],
				RollupIndex: 1000,
				RollupHash:  s.RandHash().Bytes(),
				Approved:    true,
			},
		},
		{
			s.b1,
			&pb.OptimismSignature{
				Id:          util.ULID(nil).String(),
				PreviousId:  s.sigs[s.signer1][s.contract0][99].ID,
				Contract:    s.contract0[:],
				RollupIndex: 1000,
				RollupHash:  s.RandHash().Bytes(),
				Approved:    true,
			},
		},
		{
			s.b1,
			&pb.OptimismSignature{
				Id:          util.ULID(nil).String(),
				PreviousId:  s.sigs[s.signer1][s.contract0][99].ID,
				Contract:    s.contract0[:],
				RollupIndex: 1001,
				RollupHash:  s.RandHash().Bytes(),
				Approved:    true,
			},
		},
		// new signer
		{
			s.b2,
			&pb.OptimismSignature{
				Id:          util.ULID(nil).String(),
				PreviousId:  "",
				Contract:    s.contract0[:],
				RollupIndex: 0,
				RollupHash:  s.RandHash().Bytes(),
				Approved:    true,
			},
		},
		// overwrite
		{
			s.b0,
			&pb.OptimismSignature{
				Id:          util.ULID(nil).String(),
				PreviousId:  s.sigs[s.signer0][s.contract0][49].ID,
				Contract:    s.contract0[:],
				RollupIndex: s.sigs[s.signer0][s.contract0][1].RollupIndex,
				RollupHash:  s.RandHash().Bytes(),
				Approved:    true,
			},
		},
	}

	// set stream handler to node1
	var wg sync.WaitGroup
	wg.Add(1)
	s.node1.h.SetStreamHandler(streamProtocol, func(st network.Stream) {
		s.node1.handleOptimismSignatureExchangeResponses(context.Background(), st)
		wg.Done()
	})

	// send message from node2 to node1
	responses := []*pb.OptimismSignature{}
	for _, tt := range cases {
		m := ethutil.NewMessage(
			tt.b.ChainID(),
			common.BytesToAddress(tt.want.Contract),
			new(big.Int).SetUint64(tt.want.RollupIndex),
			common.BytesToHash(tt.want.RollupHash),
			tt.want.Approved,
		)
		sig, _ := m.Signature(tt.b.SignData)

		tt.want.Signer = tt.b.Signer().Bytes()
		tt.want.Signature = sig[:]
		responses = append(responses, tt.want)
	}
	st, _ := s.node2.h.NewStream(context.Background(), s.node1.h.ID(), streamProtocol)
	writeStream(st, &pb.Stream{Body: &pb.Stream_OptimismSignatureExchange{
		OptimismSignatureExchange: &pb.OptimismSignatureExchange{Responses: responses},
	}})
	closeStream(st)

	wg.Wait()

	// assert
	for _, tt := range cases {
		signer := common.BytesToAddress(tt.want.Signer)
		contract := common.BytesToAddress(tt.want.Contract)
		index := tt.want.RollupIndex
		got, _ := s.node1.db.OPSignature.Find(nil, &signer, &contract, &index, 1, 0)

		s.Equal(tt.want.Id, got[0].ID)
		s.Equal(tt.want.Signer, got[0].Signer.Address[:])
		s.Equal(tt.want.Contract, got[0].Contract.Address[:])
		s.Equal(tt.want.RollupIndex, got[0].RollupIndex)
		s.Equal(tt.want.RollupHash, got[0].RollupHash[:])
		s.Equal(tt.want.Approved, got[0].Approved)
		s.Equal(tt.want.Signature, got[0].Signature[:])
	}
}

func (s *NodeTestSuite) TestHandleFindCommonOptimismSignature() {
	want := s.sigs[s.signer0][s.contract0][0]

	// sent request from node2 to node1
	st, _ := s.node2.h.NewStream(context.Background(), s.node1.h.ID(), streamProtocol)
	writeStream(st, &pb.Stream{Body: &pb.Stream_FindCommonOptimismSignature{
		FindCommonOptimismSignature: &pb.FindCommonOptimismSignature{
			Locals: []*pb.FindCommonOptimismSignature_Local{
				{Id: util.ULID(nil).String(), PreviousId: util.ULID(nil).String()},
				{Id: util.ULID(nil).String(), PreviousId: util.ULID(nil).String()},
			},
		},
	}})
	writeStream(st, &pb.Stream{Body: &pb.Stream_FindCommonOptimismSignature{
		FindCommonOptimismSignature: &pb.FindCommonOptimismSignature{
			Locals: []*pb.FindCommonOptimismSignature_Local{
				{Id: util.ULID(nil).String(), PreviousId: util.ULID(nil).String()},
				{Id: want.ID, PreviousId: want.PreviousID},
				{Id: util.ULID(nil).String(), PreviousId: util.ULID(nil).String()},
			},
		},
	}})
	writeStream(st, eom)

	// read response from node1
	reads := s.readsStream(st)

	// assert
	s.Len(reads, 3)
	s.Nil(reads[0].GetFindCommonOptimismSignature().Found)
	s.Equal(reads[1].GetFindCommonOptimismSignature().Found.Id, want.ID)
	s.Equal(reads[1].GetFindCommonOptimismSignature().Found.PreviousId, want.PreviousID)
	s.NotNil(reads[2].GetEom())
}

func (s *NodeTestSuite) TestPublishLatestSignatures() {
	// Wait for pubsub to be ready
	time.Sleep(50 * time.Millisecond)

	var got struct {
		peer peer.ID
		sigs []*pb.OptimismSignature
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go func() {
		defer cancel()

		var m pb.PubSub
		peerID, err := subscribe(ctx, s.node2.sub, s.node2.h.ID(), &m)
		if err != nil {
			s.Fail(err.Error())
		}
		got.peer = peerID
		got.sigs = m.GetOptimismSignatureExchange().Latests
	}()

	// publish message
	s.node1.publishLatestSignatures(ctx)

	// wait for subscribed
	<-ctx.Done()

	// assert
	s.Equal(s.node1.h.ID(), got.peer)
	s.Len(got.sigs, 2)

	s.Equal(s.sigs[s.signer0][s.contract1][99].ID, got.sigs[0].Id)
	s.Equal(s.sigs[s.signer1][s.contract1][199].ID, got.sigs[1].Id)

	s.Equal(s.signer0[:], got.sigs[0].Signer)
	s.Equal(s.signer1[:], got.sigs[1].Signer)

	s.Equal(s.contract1[:], got.sigs[0].Contract)
	s.Equal(s.contract1[:], got.sigs[1].Contract)

	s.Equal(uint64(99), got.sigs[0].RollupIndex)
	s.Equal(uint64(199), got.sigs[1].RollupIndex)
}

func (s *NodeTestSuite) newWorker(bootnodes []string) *Node {
	// Setup database.
	db, _ := database.NewDatabase(&config.Database{Path: ":memory:"})

	// Setup libp2p.
	priv, _, _, _ := GenerateKeyPair()
	cfg := &config.P2P{
		Listens:         []string{"/ip4/127.0.0.1/tcp/" + s.findPort(5)},
		PublishInterval: 0,
		StreamTimeout:   3 * time.Second,
		ExperimentalLanDHT: struct {
			Loopback  bool
			Bootnodes []string
		}{Loopback: true, Bootnodes: bootnodes},
	}
	cfg.OutboundLimits.Concurrency = 10
	cfg.OutboundLimits.Throttling = 500
	cfg.InboundLimits.Concurrency = 10
	cfg.InboundLimits.Throttling = 1000
	cfg.InboundLimits.MaxSendTime = time.Second * 5
	host, dht, bwm, hpHelper, _ := NewHost(context.Background(), cfg, priv)

	worker, _ := NewNode(cfg, db, host, dht, bwm, hpHelper,
		s.b0.ChainID().Uint64(), []common.Address{}, s.stakemanager)
	host.SetStreamHandler(streamProtocol,
		worker.newStreamHandler(context.Background()))

	return worker
}

func (s *NodeTestSuite) findPort(maxAttempts int) string {
	for i := 0; i < maxAttempts; i++ {
		addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort("127.0.0.1", "0"))
		if err != nil {
			continue
		}

		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			l.Close()
			continue
		}

		port := l.Addr().(*net.TCPAddr).Port
		l.Close()

		return strconv.Itoa(port)
	}

	s.Fail("no port")
	return ""
}

func (s *NodeTestSuite) genStateRoot(contract []byte, batchIndex int) common.Hash {
	b := new(big.Int).SetBytes(contract)
	b.Add(b, big.NewInt(int64(batchIndex)))
	return common.BigToHash(b)
}

func (s *NodeTestSuite) readsStream(st network.Stream) []*pb.Stream {
	var recvs []*pb.Stream
	for {
		recv, err := readStream(st)
		if err != nil {
			s.Fail(err.Error())
		}

		recvs = append(recvs, recv)
		if recv.GetEom() != nil {
			break
		}
	}
	return recvs
}

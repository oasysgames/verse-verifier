package p2p

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/p2p/pb"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oasysgames/oasys-optimism-verifier/verselayer"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/suite"
)

func TestNode(t *testing.T) {
	suite.Run(t, new(NodeTestSuite))
}

type NodeTestSuite struct {
	testhelper.Suite

	baseTime time.Time
	b0       *testhelper.TestBackend
	b1       *testhelper.TestBackend
	b2       *testhelper.TestBackend
	signer0  common.Address
	signer1  common.Address
	signer2  common.Address
	scc0     common.Address
	scc1     common.Address
	scc2     common.Address
	sigs     map[common.Address]map[common.Address][]*database.OptimismSignature

	bootnode *Node
	node1    *Node
	node2    *Node
}

func (s *NodeTestSuite) SetupTest() {
	s.baseTime = time.Now().UTC()
	s.b0 = testhelper.NewTestBackend()
	s.b1 = s.b0.NewAccountBackend()
	s.b2 = s.b0.NewAccountBackend()
	s.signer0 = s.b0.Signer()
	s.signer1 = s.b1.Signer()
	s.signer2 = s.b2.Signer()
	s.scc0 = s.RandAddress()
	s.scc1 = s.RandAddress()
	s.scc2 = s.RandAddress()
	s.sigs = map[common.Address]map[common.Address][]*database.OptimismSignature{}

	s.bootnode = s.newWorker()
	s.node1 = s.newWorker()
	s.node2 = s.newWorker()

	// setup libp2p
	ctx := context.Background()
	bootstrapPeers := ConvertPeers([]string{
		maToP2P(s.bootnode.h.Addrs()[0], s.bootnode.h.ID()),
	})
	Bootstrap(ctx, s.node1.h, s.node1.dht)
	Bootstrap(ctx, s.node2.h, s.node2.dht)
	ConnectPeers(ctx, s.node1.h, bootstrapPeers)
	ConnectPeers(ctx, s.node2.h, bootstrapPeers)

	// create sample records
	for _, node := range []*Node{s.node1, s.node2} {
		node.db.Optimism.FindOrCreateSigner(s.signer0)
		node.db.Optimism.FindOrCreateSigner(s.signer1)
		node.db.Optimism.FindOrCreateSigner(s.signer2)
		node.db.Optimism.FindOrCreateSCC(s.scc0)
		node.db.Optimism.FindOrCreateSCC(s.scc1)
		node.db.Optimism.FindOrCreateSCC(s.scc2)
	}

	backends := []*testhelper.TestBackend{s.b0, s.b1, s.b2}
	signers := []common.Address{s.signer0, s.signer1}
	sccs := []common.Address{s.scc0, s.scc1}
	sizes := [][]int{{50, 100}, {150, 200}}

	for i, signer := range signers {
		backend := backends[i]
		s.sigs[signer] = map[common.Address][]*database.OptimismSignature{}

		for j, scc := range sccs {
			s.sigs[signer][scc] = []*database.OptimismSignature{}

			for _, index := range s.Range(0, sizes[i][j]) {
				batchRoot := s.genStateRoot(scc[:], index)
				approved := true

				msg := verselayer.NewSccMessage(
					backend.ChainID(),
					scc,
					big.NewInt(int64(index)),
					util.BytesToBytes32(batchRoot[:]),
					approved,
				)
				sigbin, _ := msg.Signature(backend.SignData)

				sig, _ := s.node1.db.Optimism.SaveSignature(
					nil, nil,
					signer,
					scc,
					uint64(index),
					batchRoot,
					uint64(index),
					uint64(index),
					[]byte(fmt.Sprintf("test-%d", index)),
					approved,
					sigbin,
				)
				s.sigs[signer][scc] = append(s.sigs[signer][scc], sig)
			}
		}
	}
}

func (s *NodeTestSuite) TestHandleOptimismSignatureExchangeFromPubSub() {
	type want struct {
		signer  common.Address
		idAfter string
	}
	type testcase struct {
		msg  *pb.OptimismSignature
		want want
	}
	cases := []*testcase{
		{
			toProtoBufSig(s.sigs[s.signer0][s.scc1][99]),
			want{s.signer0, s.sigs[s.signer0][s.scc1][99].ID},
		},
		{
			toProtoBufSig(s.sigs[s.signer1][s.scc1][199]),
			want{s.signer1, s.sigs[s.signer1][s.scc1][199].ID},
		},
		{
			&pb.OptimismSignature{
				Id:                util.ULID(nil).String(),
				PreviousId:        util.ULID(nil).String(),
				Signer:            s.signer2[:],
				Scc:               s.scc2[:],
				BatchIndex:        0,
				BatchRoot:         s.genStateRoot(s.scc2[:], 0).Bytes(),
				BatchSize:         0,
				PrevTotalElements: 0,
				ExtraData:         []byte(fmt.Sprintf("test-%d", 0)),
				Approved:          true,
			},
			want{s.signer2, ""},
		},
	}

	lastcase := cases[len(cases)-1]
	msg := verselayer.NewSccMessage(
		s.b2.ChainID(),
		common.BytesToAddress(lastcase.msg.Scc),
		new(big.Int).SetUint64(lastcase.msg.BatchIndex),
		util.BytesToBytes32(lastcase.msg.BatchRoot),
		lastcase.msg.Approved,
	)
	sigbin, _ := msg.Signature(s.b2.SignData)
	lastcase.msg.Signature = sigbin[:]

	// set assertion func to subscriber
	var (
		mu          = &sync.Mutex{}
		reads       = []*pb.Stream{}
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	)
	defer cancel()
	s.node2.h.SetStreamHandler(streamProtocol, func(st network.Stream) {
		mu.Lock()
		defer mu.Unlock()
		defer closeStream(st)

		for {
			m, _ := readStream(st)
			reads = append(reads, m)

			switch m.Body.(type) {
			case *pb.Stream_Eom:
				return
			case *pb.Stream_FindCommonOptimismSignature:
				writeStream(st, &pb.Stream{Body: &pb.Stream_FindCommonOptimismSignature{
					FindCommonOptimismSignature: &pb.FindCommonOptimismSignature{
						Found: nil,
					},
				}})
			}
		}
	})

	// publish message
	for _, tt := range cases {
		go s.node1.handleOptimismSignatureExchangeFromPubSub(ctx, s.node2.h.ID(), tt.msg)
		time.Sleep(time.Millisecond * 50)
	}
	<-ctx.Done()

	s.Len(reads, 12)

	// signer0
	gots0 := reads[0].GetFindCommonOptimismSignature().Locals
	s.Len(gots0, 100)
	s.Equal(gots0[0].Id, s.sigs[s.signer0][s.scc1][99].ID)
	s.Equal(gots0[99].Id, s.sigs[s.signer0][s.scc1][0].ID)

	gots1 := reads[1].GetFindCommonOptimismSignature().Locals
	s.Len(gots1, 50)
	s.Equal(gots1[0].Id, s.sigs[s.signer0][s.scc0][49].ID)
	s.Equal(gots1[49].Id, s.sigs[s.signer0][s.scc0][0].ID)

	gots2 := reads[2].GetOptimismSignatureExchange().Requests
	s.Len(gots2, 1)
	s.Equal(cases[0].want.signer[:], gots2[0].Signer)
	func() {
		gt := ulid.MustParse(gots2[0].IdAfter)
		wt := ulid.MustParse(cases[0].want.idAfter)
		s.Equal(wt.Time()-1000, gt.Time())
	}()

	gots3 := reads[3].GetEom()
	s.NotNil(gots3)

	// signer1
	gots4 := reads[4].GetFindCommonOptimismSignature().Locals
	s.Len(gots4, 100)
	s.Equal(gots4[0].Id, s.sigs[s.signer1][s.scc1][199].ID)
	s.Equal(gots4[99].Id, s.sigs[s.signer1][s.scc1][100].ID)

	gots5 := reads[5].GetFindCommonOptimismSignature().Locals
	s.Len(gots5, 100)
	s.Equal(gots5[0].Id, s.sigs[s.signer1][s.scc1][99].ID)
	s.Equal(gots5[99].Id, s.sigs[s.signer1][s.scc1][0].ID)

	gots6 := reads[6].GetFindCommonOptimismSignature().Locals
	s.Len(gots6, 100)
	s.Equal(gots6[0].Id, s.sigs[s.signer1][s.scc0][149].ID)
	s.Equal(gots6[99].Id, s.sigs[s.signer1][s.scc0][50].ID)

	gots7 := reads[7].GetFindCommonOptimismSignature().Locals
	s.Len(gots7, 50)
	s.Equal(gots7[0].Id, s.sigs[s.signer1][s.scc0][49].ID)
	s.Equal(gots7[49].Id, s.sigs[s.signer1][s.scc0][0].ID)

	gots8 := reads[8].GetOptimismSignatureExchange().Requests
	s.Len(gots8, 1)
	s.Equal(cases[1].want.signer[:], gots8[0].Signer)
	func() {
		gt := ulid.MustParse(gots8[0].IdAfter)
		wt := ulid.MustParse(cases[1].want.idAfter)
		s.Equal(wt.Time()-1000, gt.Time())
	}()

	gots9 := reads[9].GetEom()
	s.NotNil(gots9)

	// signer2
	gots10 := reads[10].GetOptimismSignatureExchange().Requests
	s.Len(gots10, 1)
	s.Equal(cases[2].want.signer[:], gots10[0].Signer)
	s.Equal("", gots10[0].IdAfter)

	gots11 := reads[11].GetEom()
	s.NotNil(gots11)
}

func (s *NodeTestSuite) TestHandleOptimismSignatureExchangeRequests() {
	wantss := [][]struct {
		signer       common.Address
		scc          common.Address
		batchIndexes []int
	}{
		{
			{s.signer0, s.scc0, s.Range(0, 50)},
			{s.signer0, s.scc1, s.Range(0, 100)},
		},
		{
			{s.signer1, s.scc1, s.Range(100, 200)},
		},
	}

	// send message
	st, _ := s.node2.h.NewStream(context.Background(), s.node1.h.ID(), streamProtocol)
	writeStream(st, &pb.Stream{Body: &pb.Stream_OptimismSignatureExchange{
		OptimismSignatureExchange: &pb.OptimismSignatureExchange{
			Requests: []*pb.OptimismSignatureExchange_Request{
				{IdAfter: "", Signer: s.signer0[:]},
				{IdAfter: s.sigs[s.signer1][s.scc1][100].ID, Signer: s.signer1[:]},
				{IdAfter: "", Signer: s.signer2[:]},
			},
		},
	}})
	writeStream(st, eom)

	// read messages
	reads := s.readsStream(st)
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

				wantsig := s.sigs[want.signer][want.scc][bi]
				s.Equal(wantsig.ID, got.Id)
				s.Equal(wantsig.Signer.Address[:], got.Signer)
				s.Equal(wantsig.OptimismScc.Address[:], got.Scc)
				s.Equal(wantsig.BatchIndex, got.BatchIndex)
				s.Equal(wantsig.BatchRoot[:], got.BatchRoot)
				s.Equal(wantsig.BatchSize, got.BatchSize)
				s.Equal(wantsig.PrevTotalElements, got.PrevTotalElements)
				s.Equal(wantsig.ExtraData, got.ExtraData)
				s.Equal(wantsig.Approved, got.Approved)
				s.Equal(wantsig.Signature[:], got.Signature)
			}
		}

		s.Len(gots, 0)
	}
}

func (s *NodeTestSuite) TestHandleOptimismSignatureExchangeResponses() {
	cases := []struct {
		b    *testhelper.TestBackend
		want *pb.OptimismSignature
	}{
		{
			s.b0,
			&pb.OptimismSignature{
				Id:                util.ULID(nil).String(),
				PreviousId:        s.sigs[s.signer0][s.scc0][49].ID,
				Scc:               s.scc0[:],
				BatchIndex:        1000,
				BatchRoot:         s.RandHash().Bytes(),
				BatchSize:         uint64(rand.Intn(100)),
				PrevTotalElements: uint64(rand.Intn(100)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", rand.Intn(100))),
				Approved:          true,
			},
		},
		{
			s.b1,
			&pb.OptimismSignature{
				Id:                util.ULID(nil).String(),
				PreviousId:        s.sigs[s.signer1][s.scc0][99].ID,
				Scc:               s.scc0[:],
				BatchIndex:        1000,
				BatchRoot:         s.RandHash().Bytes(),
				BatchSize:         uint64(rand.Intn(100)),
				PrevTotalElements: uint64(rand.Intn(100)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", rand.Intn(100))),
				Approved:          true,
			},
		},
		{
			s.b1,
			&pb.OptimismSignature{
				Id:                util.ULID(nil).String(),
				PreviousId:        s.sigs[s.signer1][s.scc0][99].ID,
				Scc:               s.scc0[:],
				BatchIndex:        1001,
				BatchRoot:         s.RandHash().Bytes(),
				BatchSize:         uint64(rand.Intn(100)),
				PrevTotalElements: uint64(rand.Intn(100)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", rand.Intn(100))),
				Approved:          true,
			},
		},
		// new signer
		{
			s.b2,
			&pb.OptimismSignature{
				Id:                util.ULID(nil).String(),
				PreviousId:        "",
				Scc:               s.scc0[:],
				BatchIndex:        0,
				BatchRoot:         s.RandHash().Bytes(),
				BatchSize:         uint64(rand.Intn(100)),
				PrevTotalElements: uint64(rand.Intn(100)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", rand.Intn(100))),
				Approved:          true,
			},
		},
		// overwrite
		{
			s.b0,
			&pb.OptimismSignature{
				Id:                util.ULID(nil).String(),
				PreviousId:        s.sigs[s.signer0][s.scc0][49].ID,
				Scc:               s.scc0[:],
				BatchIndex:        s.sigs[s.signer0][s.scc0][1].BatchIndex,
				BatchRoot:         s.RandHash().Bytes(),
				BatchSize:         uint64(rand.Intn(99999)),
				PrevTotalElements: uint64(rand.Intn(99999)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", rand.Intn(99999))),
				Approved:          true,
			},
		},
	}

	// send message
	responses := []*pb.OptimismSignature{}
	for _, tt := range cases {
		m := verselayer.NewSccMessage(
			tt.b.ChainID(),
			common.BytesToAddress(tt.want.Scc),
			new(big.Int).SetUint64(tt.want.BatchIndex),
			common.BytesToHash(tt.want.BatchRoot),
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
	writeStream(st, eom)

	// read messages
	reads := s.readsStream(st)
	closeStream(st)

	// assert
	s.Len(reads, 1)
	s.NotNil(reads[0].GetEom())

	for _, tt := range cases {
		signer := common.BytesToAddress(tt.want.Signer)
		scc := common.BytesToAddress(tt.want.Scc)
		index := tt.want.BatchIndex
		got, _ := s.node1.db.Optimism.FindSignatures(nil, &signer, &scc, &index, 1, 0)

		s.Equal(tt.want.Id, got[0].ID)
		s.Equal(tt.want.Signer, got[0].Signer.Address[:])
		s.Equal(tt.want.Scc, got[0].OptimismScc.Address[:])
		s.Equal(tt.want.BatchIndex, got[0].BatchIndex)
		s.Equal(tt.want.BatchRoot, got[0].BatchRoot[:])
		s.Equal(tt.want.BatchSize, got[0].BatchSize)
		s.Equal(tt.want.PrevTotalElements, got[0].PrevTotalElements)
		s.Equal(tt.want.ExtraData, got[0].ExtraData)
		s.Equal(tt.want.Approved, got[0].Approved)
		s.Equal(tt.want.Signature, got[0].Signature[:])
	}
}

func (s *NodeTestSuite) TestHandleFindCommonOptimismSignature() {
	want := s.sigs[s.signer0][s.scc0][0]

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

		peerID, m, err := subscribe(ctx, s.node2.sub, s.node2.h.ID())
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

	s.Equal(s.sigs[s.signer0][s.scc1][99].ID, got.sigs[0].Id)
	s.Equal(s.sigs[s.signer1][s.scc1][199].ID, got.sigs[1].Id)

	s.Equal(s.signer0[:], got.sigs[0].Signer)
	s.Equal(s.signer1[:], got.sigs[1].Signer)

	s.Equal(s.scc1[:], got.sigs[0].Scc)
	s.Equal(s.scc1[:], got.sigs[1].Scc)

	s.Equal(uint64(99), got.sigs[0].BatchIndex)
	s.Equal(uint64(199), got.sigs[1].BatchIndex)
}

func (s *NodeTestSuite) newWorker() *Node {
	// Setup database.
	db, _ := database.NewDatabase(&config.Database{Path: ":memory:"})

	// Setup libp2p.
	priv, _, _, _ := GenerateKeyPair()
	host, dht, bwm, _ := NewHost(context.Background(), "127.0.0.1", s.findPort(5), priv)

	worker, err := NewNode(
		&config.P2P{PublishInterval: 0, StreamTimeout: 3 * time.Second},
		db, host, dht, bwm, s.b0.ChainID().Uint64(), []common.Address{})
	if err != nil {
		s.Fail(err.Error())
	}

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

func (s *NodeTestSuite) genStateRoot(scc []byte, batchIndex int) common.Hash {
	b := new(big.Int).SetBytes(scc)
	b.Add(b, big.NewInt(int64(batchIndex)))
	return common.BigToHash(b)
}

func (s *NodeTestSuite) genSignature(signer, scc []byte, batchIndex int) database.Signature {
	b := new(big.Int).Xor(new(big.Int).SetBytes(signer), new(big.Int).SetBytes(scc))
	b.Add(b, big.NewInt(int64(batchIndex)))
	return database.BytesSignature(b.Bytes())
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

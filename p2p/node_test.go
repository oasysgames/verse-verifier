package p2p

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/p2p/pb"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/suite"
)

func TestNode(t *testing.T) {
	suite.Run(t, new(NodeTestSuite))
}

type NodeTestSuite struct {
	testhelper.Suite

	bootnode *Node
	node1    *Node
	node2    *Node

	baseTime time.Time
	signer0  common.Address
	signer1  common.Address
	signer2  common.Address
	scc0     common.Address
	scc1     common.Address
	scc2     common.Address
	sigs     map[common.Address]map[common.Address][]*database.OptimismSignature
}

func (s *NodeTestSuite) SetupTest() {
	s.bootnode = s.newWorker()
	s.node1 = s.newWorker()
	s.node2 = s.newWorker()

	s.baseTime = time.Now().UTC()
	s.signer0 = s.RandAddress()
	s.signer1 = s.RandAddress()
	s.signer2 = s.RandAddress()
	s.scc0 = s.RandAddress()
	s.scc1 = s.RandAddress()
	s.scc2 = s.RandAddress()
	s.sigs = map[common.Address]map[common.Address][]*database.OptimismSignature{}

	// setup libp2p
	bootstrapPeers := ConvertPeers([]string{
		maToP2P(s.bootnode.h.Addrs()[0], s.bootnode.h.ID()),
	})
	Bootstrap(context.Background(), s.node1.h, s.node1.dht, bootstrapPeers)
	Bootstrap(context.Background(), s.node2.h, s.node2.dht, bootstrapPeers)

	// create sample records
	for _, node := range []*Node{s.node1, s.node2} {
		node.db.Optimism.FindOrCreateSigner(s.signer0)
		node.db.Optimism.FindOrCreateSigner(s.signer1)
		node.db.Optimism.FindOrCreateSigner(s.signer2)
		node.db.Optimism.FindOrCreateSCC(s.scc0)
		node.db.Optimism.FindOrCreateSCC(s.scc1)
		node.db.Optimism.FindOrCreateSCC(s.scc2)
	}
	sizes := [][]int{{50, 100}, {150, 200}}
	for i, signer := range []common.Address{s.signer0, s.signer1} {
		s.sigs[signer] = map[common.Address][]*database.OptimismSignature{}

		for j, scc := range []common.Address{s.scc0, s.scc1} {
			s.sigs[signer][scc] = []*database.OptimismSignature{}

			for _, index := range s.Range(0, sizes[i][j]) {
				sig, _ := s.node1.db.Optimism.SaveSignature(
					nil, nil,
					signer,
					scc,
					uint64(index),
					s.genStateRoot(scc[:], index),
					uint64(index),
					uint64(index),
					[]byte(fmt.Sprintf("test-%d", index)),
					true,
					s.genSignature(signer[:], scc[:], index),
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
	cases := []struct {
		pubsub *pb.OptimismSignature
		want   want
	}{
		{
			&pb.OptimismSignature{
				Id:     ulid.Make().String(),
				Signer: s.signer0[:],
			},
			want{s.signer0, s.sigs[s.signer0][s.scc1][99].ID},
		},
		{
			&pb.OptimismSignature{
				Id:     ulid.Make().String(),
				Signer: s.signer1[:],
			},
			want{s.signer1, s.sigs[s.signer1][s.scc1][199].ID},
		},
		{
			&pb.OptimismSignature{
				Id:     ulid.Make().String(),
				Signer: s.signer2[:],
			},
			want{s.signer2, ""},
		},
	}

	// set assertion func to subscriber
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	s.node2.h.SetStreamHandler(streamProtocol, func(st network.Stream) {
		defer cancel()
		defer closeStream(st)

		reads := s.readsStream(st)

		s.Len(reads, 2)
		s.NotNil(reads[0].GetOptimismSignatureExchange())
		s.NotNil(reads[1].GetEom())

		gots := reads[0].GetOptimismSignatureExchange().Requests
		s.Len(gots, len(cases))
		for i, tt := range cases {
			s.Equal(tt.want.signer[:], gots[i].Signer, i)
			s.Equal(tt.want.idAfter, gots[i].IdAfter, i)
		}
	})

	// publish message
	m := &pb.OptimismSignatureExchange{}
	for _, tt := range cases {
		m.Latests = append(m.Latests, tt.pubsub)
	}
	s.node1.handleOptimismSignatureExchangeFromPubSub(ctx, s.node2.h.ID(), m)

	<-ctx.Done()
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

				s.Equal(s.sigs[want.signer][want.scc][bi].ID, got.Id)
				s.Equal(want.signer[:], got.Signer)
				s.Equal(want.scc[:], got.Scc)
				s.Equal(uint64(bi), got.BatchIndex)
				s.Equal(s.genStateRoot(want.scc[:], bi).Bytes(), got.BatchRoot)
				s.Equal(uint64(bi), got.BatchSize)
				s.Equal(uint64(bi), got.PrevTotalElements)
				s.Equal([]byte(fmt.Sprintf("test-%d", bi)), got.ExtraData)
				s.True(got.Approved)
				s.Equal(s.genSignature(want.signer[:], want.scc[:], bi).Bytes(), got.Signature)
			}
		}

		s.Len(gots, 0)
	}
}

func (s *NodeTestSuite) TestHandleOptimismSignatureExchangeResponses() {
	cases := []struct {
		want *pb.OptimismSignature
	}{
		{
			&pb.OptimismSignature{
				Id:                ulid.Make().String(),
				PreviousId:        s.sigs[s.signer0][s.scc0][49].ID,
				Signer:            s.signer0[:],
				Scc:               s.scc0[:],
				BatchIndex:        1000,
				BatchRoot:         s.RandHash().Bytes(),
				BatchSize:         uint64(rand.Intn(100)),
				PrevTotalElements: uint64(rand.Intn(100)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", rand.Intn(100))),
				Approved:          true,
				Signature:         database.RandSignature().Bytes(),
			},
		},
		{
			&pb.OptimismSignature{
				Id:                ulid.Make().String(),
				PreviousId:        s.sigs[s.signer1][s.scc0][99].ID,
				Signer:            s.signer1[:],
				Scc:               s.scc0[:],
				BatchIndex:        1000,
				BatchRoot:         s.RandHash().Bytes(),
				BatchSize:         uint64(rand.Intn(100)),
				PrevTotalElements: uint64(rand.Intn(100)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", rand.Intn(100))),
				Approved:          true,
				Signature:         database.RandSignature().Bytes(),
			},
		},
		{
			&pb.OptimismSignature{
				Id:                ulid.Make().String(),
				PreviousId:        s.sigs[s.signer1][s.scc0][99].ID,
				Signer:            s.signer1[:],
				Scc:               s.scc0[:],
				BatchIndex:        1001,
				BatchRoot:         s.RandHash().Bytes(),
				BatchSize:         uint64(rand.Intn(100)),
				PrevTotalElements: uint64(rand.Intn(100)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", rand.Intn(100))),
				Approved:          true,
				Signature:         database.RandSignature().Bytes(),
			},
		},
		// new signer
		{
			&pb.OptimismSignature{
				Id:                ulid.Make().String(),
				PreviousId:        "",
				Signer:            s.signer2[:],
				Scc:               s.scc0[:],
				BatchIndex:        0,
				BatchRoot:         s.RandHash().Bytes(),
				BatchSize:         uint64(rand.Intn(100)),
				PrevTotalElements: uint64(rand.Intn(100)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", rand.Intn(100))),
				Approved:          true,
				Signature:         database.RandSignature().Bytes(),
			},
		},
		// overwrite
		{
			&pb.OptimismSignature{
				Id:                ulid.Make().String(),
				PreviousId:        s.sigs[s.signer0][s.scc0][49].ID,
				Signer:            s.signer0[:],
				Scc:               s.scc0[:],
				BatchIndex:        s.sigs[s.signer0][s.scc0][1].BatchIndex,
				BatchRoot:         s.RandHash().Bytes(),
				BatchSize:         uint64(rand.Intn(99999)),
				PrevTotalElements: uint64(rand.Intn(99999)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", rand.Intn(99999))),
				Approved:          true,
				Signature:         database.RandSignature().Bytes(),
			},
		},
		// Previous ID does not exists
		{
			&pb.OptimismSignature{
				Id:                ulid.Make().String(),
				PreviousId:        ulid.Make().String(),
				Signer:            s.signer0[:],
				Scc:               s.scc0[:],
				BatchIndex:        1001,
				BatchRoot:         s.RandHash().Bytes(),
				BatchSize:         uint64(rand.Intn(100)),
				PrevTotalElements: uint64(rand.Intn(100)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", rand.Intn(100))),
				Approved:          true,
				Signature:         database.RandSignature().Bytes(),
			},
		},
	}

	// send message
	responses := []*pb.OptimismSignature{}
	for _, tt := range cases {
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

	for i, tt := range cases {
		signer := common.BytesToAddress(tt.want.Signer)
		scc := common.BytesToAddress(tt.want.Scc)
		index := tt.want.BatchIndex
		got, _ := s.node1.db.Optimism.FindSignatures(nil, &signer, &scc, &index, 1, 0)

		if i == len(cases)-1 {
			s.Len(got, 0)
		} else {
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
	db, _ := database.NewDatabase(":memory:")

	// Setup libp2p.
	priv, _, _, _ := GenerateKeyPair()
	host, dht, bwm, _ := NewHost(context.Background(), "127.0.0.1", s.findPort(5), priv)

	worker, err := NewNode(db, host, dht, bwm, 0)
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

package p2p

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestNode2(t *testing.T) {
	suite.Run(t, new(Node2TestSuite))
}

type Node2TestSuite struct {
	NodeTestSuite
	bootnode2, node21, node22 *Node2
}

func (s *Node2TestSuite) SetupTest() {
	s.NodeTestSuite.SetupTest()

	// setup libp2p
	s.bootnode2 = s.newWorker([]string{}, false)
	bootnodes := []string{s.bootnode.cfg.Listens[0] + "/p2p/" + s.bootnode.h.ID().String()}
	s.node21 = s.newWorker(bootnodes, true)
	s.node22 = s.newWorker(bootnodes, true)

	// // create sample records
	// for _, node := range []*Node2{s.node1, s.node2} {
	// 	node.db.Signer.FindOrCreate(s.signer0)
	// 	node.db.Signer.FindOrCreate(s.signer1)
	// 	node.db.Signer.FindOrCreate(s.signer2)
	// 	node.db.OPContract.FindOrCreate(s.contract0)
	// 	node.db.OPContract.FindOrCreate(s.contract1)
	// 	node.db.OPContract.FindOrCreate(s.contract2)
	// }

	// sizes := [][]int{
	// 	{
	// 		50,  // signer0, contract0
	// 		100, // signer0, contract1
	// 	},
	// 	{
	// 		150, // signer1, contract0
	// 		200, // signer1, contract1
	// 	},
	// }
	// for i, signer := range signers {
	// 	backend := backends[i]
	// 	s.sigs[signer] = map[common.Address][]*database.OptimismSignature{}

	// 	for j, contract := range contracts {
	// 		s.sigs[signer][contract] = []*database.OptimismSignature{}

	// 		for _, rollupIndex := range s.Range(0, sizes[i][j]) {
	// 			rollupHash := s.genStateRoot(contract[:], rollupIndex)
	// 			approved := true

	// 			msg := ethutil.NewMessage(
	// 				backend.ChainID(),
	// 				contract,
	// 				big.NewInt(int64(rollupIndex)),
	// 				util.BytesToBytes32(rollupHash[:]),
	// 				approved,
	// 			)
	// 			sigbin, _ := msg.Signature(backend.SignData)

	// 			sig, _ := s.node1.db.OPSignature.Save(
	// 				nil, nil,
	// 				signer,
	// 				contract,
	// 				uint64(rollupIndex),
	// 				rollupHash,
	// 				approved,
	// 				sigbin,
	// 			)
	// 			s.sigs[signer][contract] = append(s.sigs[signer][contract], sig)
	// 		}
	// 	}
	// }
}

func (s *Node2TestSuite) AfterTest() {
	s.bootnode2.Stop()
	s.node21.Stop()
	s.node22.Stop()
}

func (s *Node2TestSuite) newWorker(bootnodes []string, isHandleCommiterSubReq bool) *Node2 {
	node := s.NodeTestSuite.newWorker(bootnodes)
	node2, err := NewNode2(context.Background(), node, WithIsHandleCommiterSubReq(isHandleCommiterSubReq))
	s.Require().NoError(err)
	return node2
}

func (s *Node2TestSuite) TestSubscribeCommonTopicLoop() {
	var (
		ctx, cancel        = context.WithCancel(context.Background())
		chainId     uint64 = 124
		err         error
	)
	defer cancel()

	// Start subscribeCommonTopicLoop
	for _, node := range []*Node2{s.node21, s.node22} {
		go node.subscribeCommonTopicLoop(ctx)
		time.Sleep(time.Millisecond * 50)
	}

	// publish message
	err = s.node21.PublishSubmitterTopicSubReq(ctx, chainId, "", []byte{}, true)
	s.NoError(err)

	// check if message is received
	counter := 0
	for {
		time.Sleep(100 * time.Millisecond)
		if _, ok := s.node22.topicSubmitter[chainId]; ok {
			s.True(ok)
			ctx.Done()
			break
		}
		if counter > 10 {
			// give up
			s.Fail("failed to subscribe common topic")
			break
		}
		counter++
	}
}

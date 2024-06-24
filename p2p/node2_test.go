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
	node21, node22 *Node2
}

func (s *Node2TestSuite) SetupTest() {
	s.NodeTestSuite.SetupTest()
	s.node21 = s.newWorker(s.node1, true, true)
	s.node22 = s.newWorker(s.node2, true, true)
}

func (s *Node2TestSuite) newWorker(node *Node, isHandlingSignatureRequest, isHandlingPublishedSignatures bool) *Node2 {
	node2, err := NewNode2(context.Background(), node, WithIsHandlingSignatureRequest(isHandlingSignatureRequest), WithIsHandlingPublishedSignatures(isHandlingPublishedSignatures))
	s.Require().NoError(err)
	return node2
}

func (s *Node2TestSuite) AfterTest() {
	s.node21.Close()
	s.node22.Close()
}

func (s *Node2TestSuite) TestSubscribeSubmitterTopic() {
	var (
		ctx, cancel        = context.WithCancel(context.Background())
		chainId     uint64 = 124
	)
	defer cancel()

	// Start subscribeCommonTopicLoop
	for _, node := range []*Node2{s.node21, s.node22} {
		node.SubscribeSubmitterTopic(ctx, chainId)
		time.Sleep(time.Millisecond * 1000)
	}

	// publish message
	rollupIndex := uint64(9999)
	err := s.node21.PublishSignatureRequest(ctx, chainId, rollupIndex, 0, []byte{}, false)
	s.NoError(err)

	timer, cancancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancancel()

	select {
	case <-timer.Done():
		s.Fail("timeout")
	case m := <-s.node22.sigReqC:
		s.Equal(rollupIndex, m.RollupIndex)
	}
}

func (s *Node2TestSuite) TestHandleOptimismSignatureExchangeFromPubSub2() {
	// succeed to save signature
	msg := toProtoBufSig(s.sigs[s.signer0][s.contract0][49])
	saved := s.node22.handleOptimismSignatureExchangeFromPubSub(context.Background(), s.node1.h.ID(), msg)
	s.True(saved)
	got, err := s.node2.db.OPSignature.FindByID(msg.Id)
	s.NoError(err)
	s.Equal(msg.Id, got.ID)

	// saving duplicate signature
	saved = s.node22.handleOptimismSignatureExchangeFromPubSub(context.Background(), s.node1.h.ID(), msg)
	s.False(saved)

	// saving too old signature (no pruneRollupIndexDepth)
	msg = toProtoBufSig(s.sigs[s.signer0][s.contract0][0])
	saved = s.node22.handleOptimismSignatureExchangeFromPubSub(context.Background(), s.node1.h.ID(), msg)
	s.False(saved)
}

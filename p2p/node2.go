package p2p

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	meter "github.com/oasysgames/oasys-optimism-verifier/metrics"
	pbV1 "github.com/oasysgames/oasys-optimism-verifier/proto/p2p/v1/gen"
	pb "github.com/oasysgames/oasys-optimism-verifier/proto/p2p/v2/gen"
)

const (
	commonTopic          = "/oasys-optimism-verifier/common/2.0.0"
	PrefixSubmitterTopic = "/oasys-optimism-verifier/submitter/2.0.0/"
)

func submitterTopic(chainId uint64) string {
	return fmt.Sprintf("%s/%d", PrefixSubmitterTopic, chainId)
}

type Node2 struct {
	Node

	topicCommon    *ps.Topic
	subCommon      *ps.Subscription
	topicSubmitter map[uint64]*ps.Topic
	subSubmitter   map[uint64]*ps.Subscription

	isHandleSubmitterSubReq bool
	isHandleOptimismSigReq  bool
	isHandleOptimismSigPub  bool

	subscReqC chan *pb.ReqSubmitterTopicSub
	sigReqC   chan *pb.ReqOptimismSignature

	meterCommonTopicSubscribed,
	meterCommonTopicUnknownMsg,
	meterSubmitterTopicSubscribed,
	meterSubmitterTopicUnknownMsg meter.Counter
}

func NewNode2(
	ctx context.Context,
	node *Node,
	opts ...Option,
) (*Node2, error) {
	var (
		node2 = &Node2{
			Node:                          *node,
			topicSubmitter:                make(map[uint64]*ps.Topic),
			subSubmitter:                  make(map[uint64]*ps.Subscription),
			subscReqC:                     make(chan *pb.ReqSubmitterTopicSub, 4),
			sigReqC:                       make(chan *pb.ReqOptimismSignature, 4),
			meterCommonTopicSubscribed:    meter.GetOrRegisterCounter([]string{"p2p", "commonTopic", "subscribed"}, ""),
			meterCommonTopicUnknownMsg:    meter.GetOrRegisterCounter([]string{"p2p", "commonTopic", "unknown", "messages"}, ""),
			meterSubmitterTopicSubscribed: meter.GetOrRegisterCounter([]string{"p2p", "submitterTopic", "subscribed"}, ""),
			meterSubmitterTopicUnknownMsg: meter.GetOrRegisterCounter([]string{"p2p", "submitterTopic", "unknown", "messages"}, ""),
		}
	)

	for _, opt := range opts {
		if err := opt.Apply(node2); err != nil {
			return nil, fmt.Errorf("failed to apply Node2 option: %w", err)
		}
	}

	return node2, nil
}

func (w *Node2) Start(ctx context.Context) error {
	w.Node.Start(ctx)

	// NOTE: subscribe legacy pubsub topic for backward compatibility
	go w.subscribeLegacyPubSubTopic(ctx)

	if err := w.SubscribeCommonTopic(ctx); err != nil {
		return fmt.Errorf("failed to subscribe common topic: %w", err)
	}

	meterTicker := time.NewTicker(time.Second * 60)
	defer meterTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("P2P node2 stopped")
			return nil
		case <-meterTicker.C:
			nwstat := newNetworkStatus(w.h)
			w.meterTCPConnections.Set(float64(nwstat.connections.tcp))
			w.meterUDPConnections.Set(float64(nwstat.connections.udp))
			w.meterRelayConnections.Set(float64(nwstat.connections.relay))
			w.meterRelayHopStreams.Set(float64(nwstat.streams.hop))
			w.meterRelayStopStreams.Set(float64(nwstat.streams.stop))
			w.meterVerifierStreams.Set(float64(nwstat.streams.verifier))
			w.meterPeers.Set(float64(w.h.Peerstore().Peers().Len()))
			// TODO: add more metrics
		}
	}
}

func (w *Node2) Stop() {
	w.log.Info("Stopping node2")
	w.Node.Stop()
	// stop common topic
	w.topicCommon.Close()
	w.subCommon.Cancel()
	// stop submitter topics
	for chainId := range w.topicSubmitter {
		w.topicSubmitter[chainId].Close()
		w.subSubmitter[chainId].Cancel()
	}
}

func (w *Node2) SubscribeCommonTopic(ctx context.Context) (err error) {
	if w.subCommon != nil {
		// already subscribed
		return
	}
	if _, w.topicCommon, w.subCommon, _ = setupPubSub(ctx, w.h, commonTopic); err != nil {
		w.log.Error("Failed to setup common pubsub", "err", err)
	}
	go w.subscribeCommonTopicLoop(ctx)
	return nil
}

func (w *Node2) PublishSubmitterTopicSubReq(ctx context.Context, chainId uint64, rpc string, contract []byte, isLegacy bool) error {
	m := pb.MsgCommonTopic{
		Body: &pb.MsgCommonTopic_ReqSubmitterTopicSub{
			ReqSubmitterTopicSub: &pb.ReqSubmitterTopicSub{
				ChainId:  chainId,
				Rpc:      rpc,
				Contract: contract,
				IsLegacy: isLegacy,
			},
		},
	}
	if err := publish(ctx, w.topicCommon, &m); err != nil {
		return fmt.Errorf("failed to publish submitter topic. chainId: %d, err: %w", chainId, err)
	}
	return nil
}

func (w *Node2) SubscribeSubmitterTopic(ctx context.Context, chainId uint64) (err error) {
	if _, ok := w.topicSubmitter[chainId]; ok {
		// already subscribed
		return nil
	}
	if _, w.topicSubmitter[chainId], w.subSubmitter[chainId], err = setupPubSub(ctx, w.h, submitterTopic(chainId)); err != nil {
		return fmt.Errorf("failed to setup submitter pubsub. chainId: %d, err: %w", chainId, err)
	}
	go w.subscribeSubmitterTopicLoop(ctx)
	return nil
}

func (w *Node2) PublishSubmitterTopic(ctx context.Context, rollupIndex, highestVerifiedIndex uint64, contract []byte, isLegacy bool) error {
	m := pb.MsgSubmitterTopic{
		Body: &pb.MsgSubmitterTopic_ReqOptimismSignature{
			ReqOptimismSignature: &pb.ReqOptimismSignature{
				RollupIndex:          rollupIndex,
				HighestVerifiedIndex: highestVerifiedIndex,
				Contract:             contract,
				IsLegacy:             isLegacy,
			},
		},
	}
	if err := publish(ctx, w.topicCommon, &m); err != nil {
		return fmt.Errorf("failed to publish submitter topic. rollupIndex: %d, highestVerifiedIndex: %d, err: %w", rollupIndex, highestVerifiedIndex, err)
	}
	return nil
}

func (w *Node2) subscribeCommonTopicLoop(ctx context.Context) {
	w.log.Info("Start subscribing common topic")

	for {
		var msg pb.MsgCommonTopic
		peer, err := subscribe(ctx, w.subCommon, w.h.ID(), &msg)
		if errors.Is(err, context.Canceled) {
			w.log.Info("Common topic subscription stopped")
			return
		} else if errors.Is(err, errSelfMessage) {
			continue
		} else if err != nil {
			w.log.Error("Failed to subscribe common topic", "peer", peer, "err", err)
			continue
		}
		w.meterCommonTopicSubscribed.Incr()

		if m, ok := msg.TryGetReqSubmitterTopicSub(); ok && w.isHandleSubmitterSubReq {
			if err := w.handleSubmitterTopicSubscribeRequest(ctx, m); err != nil {
				w.log.Error("Failed to handle submitter topic subscribe request", "peer", peer, "err", err)
			}
			continue
		}

		// unsupported message
		w.log.Debug("Unsupported common topic message", "peer", peer, "err", err)
		w.meterCommonTopicUnknownMsg.Incr()
	}
}

func (w *Node2) handleSubmitterTopicSubscribeRequest(ctx context.Context, m *pb.ReqSubmitterTopicSub) (err error) {
	w.subscReqC <- m
	return w.SubscribeSubmitterTopic(ctx, m.ChainId)
}

func (w *Node2) subscribeSubmitterTopicLoop(ctx context.Context) {
	w.log.Info("Start subscribing submitter topic")

	for {
		var msg pb.MsgSubmitterTopic
		peer, err := subscribe(ctx, w.subCommon, w.h.ID(), &msg)
		if errors.Is(err, context.Canceled) {
			w.log.Info("Submitter topic subscription stopped")
			return
		} else if errors.Is(err, errSelfMessage) {
			continue
		} else if err != nil {
			w.log.Error("Failed to subscribe submitter topic", "peer", peer, "err", err)
			continue
		}
		w.meterSubmitterTopicSubscribed.Incr()

		if m, ok := msg.TryGetReqOptimismSignature(); ok && w.isHandleOptimismSigReq {
			if err := w.handleOptimismSignatureRequest(ctx, m); err != nil {
				w.log.Error("Failed to handle optimism signature request", "peer", peer, "err", err)
			}
			continue
		}

		if m, ok := msg.TryGetPubOptimismSignature(); ok && w.isHandleOptimismSigPub {
			if err := w.handleOptimismSignaturePublish(ctx, m); err != nil {
				w.log.Error("Failed to handle optimism signature publish", "peer", peer, "err", err)
			}
			continue
		}

		// unsupported message
		w.log.Debug("Unsupported submitter topic message", "peer", peer, "err", err)
		w.meterSubmitterTopicUnknownMsg.Incr()
	}
}

func (w *Node2) handleOptimismSignatureRequest(ctx context.Context, m *pb.ReqOptimismSignature) (err error) {
	w.sigReqC <- m
	return
}

func (w *Node2) handleOptimismSignaturePublish(ctx context.Context, m *pb.PubOptimismSignature) error {
	// Save signatures if it is new
	for _, sig := range m.Signatures {
		var (
			signer   = common.BytesToAddress(sig.Signer)
			contract = common.BytesToAddress(sig.Contract)
			logctx   = []interface{}{
				"signer", signer,
				"id", sig.Id, "previous-id", sig.PreviousId,
				"contract", contract,
				"index", sig.RollupIndex,
			}
		)
		// Cast v2 signature to v1 signature, v2 signature is the copy of v1 signature
		sigV1 := (*pbV1.OptimismSignature)(unsafe.Pointer(sig))
		if err := verifySignature(w.hubLayerChainID, sigV1); err != nil {
			w.log.Debug("Invalid signature", append(logctx, "err", err)...)
			return fmt.Errorf("invalid signature: %w", err)
		}
		if _, ok := w.ignoreSigners[signer]; ok {
			w.log.Info("Ignored", logctx...)
			continue
		}

		// deduplication
		if local, err := w.db.OPSignature.FindByID(sig.Id); err == nil && local.PreviousID == sig.PreviousId {
			continue
		}

		// local is newer
		if local, err := w.db.OPSignature.Find(nil, &signer, &contract, &sig.RollupIndex, 1, 0); err != nil {
			w.log.Debug("Failed to find local signature", append(logctx, "err", err)...)
			return fmt.Errorf("failed to find local signature: %w", err)
		} else if len(local) > 0 && strings.Compare(local[0].ID, sig.Id) == 1 {
			continue
		}

		if sig.PreviousId != "" {
			_, err := w.db.OPSignature.FindByID(sig.PreviousId)
			if errors.Is(err, database.ErrNotFound) {
				w.log.Warn("Previous ID does not exist", logctx...)
			} else if err != nil {
				w.log.Debug("Failed to find previous signature", append(logctx, "err", err)...)
				return fmt.Errorf("failed to find previous signature: %w", err)
			}
		}

		_, err := w.db.OPSignature.Save(
			&sig.Id, &sig.PreviousId,
			signer,
			contract,
			sig.RollupIndex,
			common.BytesToHash(sig.RollupHash),
			sig.Approved,
			database.BytesSignature(sig.Signature),
		)
		if err != nil {
			w.log.Debug("Failed to save signature", append(logctx, "err", err)...)
			return fmt.Errorf("failed to save signature: %w", err)
		}
		w.log.Debug("Received new signature", logctx...)
	}

	return nil
}

func (w *Node2) subscribeLegacyPubSubTopic(ctx context.Context) {
	for {
		var msg pbV1.PubSub
		peer, err := subscribe(ctx, w.sub, w.h.ID(), &msg)
		if errors.Is(err, context.Canceled) {
			// worker stopped
			return
		} else if errors.Is(err, errSelfMessage) {
			continue
		} else if err != nil {
			w.log.Error("Failed to subscribe legacy pubsub topic", "peer", peer, "err", err)
			continue
		}
		w.meterPubsubSubscribed.Incr()

		t := msg.GetOptimismSignatureExchange()
		if t == nil {
			w.log.Warn("Unsupported pubsub message", "peer", peer, "err", err)
			w.meterPubsubUnknownMsg.Incr()
			continue
		}

		// Save signatures if it is new
		for _, sig := range t.Latests {
			var (
				signer   = common.BytesToAddress(sig.Signer)
				contract = common.BytesToAddress(sig.Contract)
				logctx   = []interface{}{
					"signer", signer,
					"id", sig.Id, "previous-id", sig.PreviousId,
					"contract", contract,
					"index", sig.RollupIndex,
				}
			)
			if _, ok := w.ignoreSigners[signer]; ok {
				w.log.Info("Ignored", logctx...)
				continue
			}

			// deduplication
			if local, err := w.db.OPSignature.FindByID(sig.Id); err == nil && local.PreviousID == sig.PreviousId {
				continue
			}

			// local is newer
			if local, err := w.db.OPSignature.Find(nil, &signer, &contract, &sig.RollupIndex, 1, 0); err != nil {
				w.log.Debug("Failed to find local signature", append(logctx, "err", err)...)
				continue
			} else if len(local) > 0 && strings.Compare(local[0].ID, sig.Id) == 1 {
				continue
			}

			_, err := w.db.OPSignature.Save(
				&sig.Id, &sig.PreviousId,
				signer,
				contract,
				sig.RollupIndex,
				common.BytesToHash(sig.RollupHash),
				sig.Approved,
				database.BytesSignature(sig.Signature),
			)
			if err != nil {
				w.log.Warn("Failed to save signature", append(logctx, "err", err)...)
				continue
			}
			w.log.Debug("Received new signature", logctx...)
		}
	}
}

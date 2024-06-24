package p2p

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	meter "github.com/oasysgames/oasys-optimism-verifier/metrics"
	pbV1 "github.com/oasysgames/oasys-optimism-verifier/proto/p2p/v1/gen"
	pb "github.com/oasysgames/oasys-optimism-verifier/proto/p2p/v2/gen"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

const (
	PrefixSubmitterTopic = "/oasys-optimism-verifier/submitter/2.0.0"
)

func submitterTopic(chainId uint64) string {
	// NOTE: We are not guaranteed to have a unique topic for each chainId.
	// As such, we temporarily use the common topic for all chainIds.
	commonId := 0
	return fmt.Sprintf("%s/%d", PrefixSubmitterTopic, commonId)
}

type Node2 struct {
	Node

	topicSubmitter map[uint64]*ps.Topic
	subSubmitter   map[uint64]*ps.Subscription
	sigReqC        chan *pb.ReqOptimismSignature

	isHandlingSignatureRequest    bool
	isHandlingPublishedSignatures bool

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
			sigReqC:                       make(chan *pb.ReqOptimismSignature, 4),
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
	w.h.SetStreamHandler(streamProtocol, w.newStreamHandler(ctx))

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		w.subscribeLoop(ctx)
	}()

	w.showBootstrapLog()

	meterTicker := time.NewTicker(time.Second * 60)
	defer meterTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("P2P node2 metrics stopped")
			wg.Wait()
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
		}
	}
}

func (w *Node2) Close() {
	w.log.Info("Stopping node2")
	w.h.Close()
	w.topic.Close()
	w.sub.Cancel()
	// stop submitter topics
	for chainId := range w.topicSubmitter {
		w.topicSubmitter[chainId].Close()
		w.subSubmitter[chainId].Cancel()
	}
}

func (w *Node2) SubscribeSubmitterTopic(ctx context.Context, chainId uint64) (err error) {
	if _, ok := w.topicSubmitter[chainId]; ok {
		// already subscribed
		return nil
	}
	if _, w.topicSubmitter[chainId], w.subSubmitter[chainId], err = setupPubSub(ctx, w.h, submitterTopic(chainId)); err != nil {
		return fmt.Errorf("failed to setup submitter pubsub. chainId: %d, err: %w", chainId, err)
	}
	go w.subscribeSubmitterTopicLoop(ctx, chainId)
	return nil
}

func (w *Node2) PublishSignatureRequest(ctx context.Context, chainId uint64, rollupIndex, highestVerifiedIndex uint64, contract []byte, isLegacy bool) error {
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

	topic, ok := w.topicSubmitter[chainId]
	if !ok {
		return fmt.Errorf("submitter topic not found. chainId: %d", chainId)
	}
	if err := publish(ctx, topic, &m); err != nil {
		return fmt.Errorf("failed to publish submitter topic. rollupIndex: %d, highestVerifiedIndex: %d, err: %w", rollupIndex, highestVerifiedIndex, err)
	}
	return nil
}

func (w *Node2) subscribeSubmitterTopicLoop(ctx context.Context, chainId uint64) {
	w.log.Info("Start subscribing submitter topic")

	for {
		var msg pb.MsgSubmitterTopic
		peer, err := subscribe(ctx, w.subSubmitter[chainId], w.h.ID(), &msg)
		fmt.Println("PeerID: ", peer, "Error: ", err, "Msg", msg, w.cfg.ExperimentalLanDHT.Bootnodes)
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

		if m, ok := msg.TryGetReqOptimismSignature(); ok && w.isHandlingSignatureRequest {
			if err := w.handleOptimismSignatureRequest(ctx, m); err != nil {
				w.log.Error("Failed to handle optimism signature request", "peer", peer, "err", err)
			}
			continue
		}

		if m, ok := msg.TryGetPubOptimismSignature(); ok && w.isHandlingPublishedSignatures {
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
	w.log.Info("Received signature request", "rollupIndex", m.RollupIndex, "highestVerifiedIndex", m.HighestVerifiedIndex, "contract", m.Contract, "isLegacy", m.IsLegacy)
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

// copy from v1 Node
func (w *Node2) subscribeLoop(ctx context.Context) {
	type job struct {
		ctx    context.Context
		cancel context.CancelFunc
		peer   peer.ID
		remote *pbV1.OptimismSignature
		logctx []any
	}

	// Storing workers and jobs.
	workers := util.NewWorkerGroup(100)
	procs := &sync.Map{}

	for {
		var msg pbV1.PubSub
		peer, err := subscribe(ctx, w.sub, w.h.ID(), &msg)
		if errors.Is(err, context.Canceled) {
			// worker stopped
			return
		} else if errors.Is(err, errSelfMessage) {
			continue
		} else if err != nil {
			w.log.Error("Failed to subscribe", "peer", peer, "err", err)
			continue
		}
		w.meterPubsubSubscribed.Incr()

		t := msg.GetOptimismSignatureExchange()
		if t == nil {
			w.log.Warn("Unsupported pubsub message", "peer", peer, "err", err)
			w.meterPubsubUnknownMsg.Incr()
			continue
		}

		for _, remote := range t.Latests {
			signer := common.BytesToAddress(remote.Signer)
			if _, ok := w.ignoreSigners[signer]; ok {
				continue
			} else if w.stakemanager.StakeBySigner(signer).Cmp(ethutil.TenMillionOAS) == -1 {
				continue
			}

			// add new worker
			wname := signer.Hex()
			if !workers.Has(wname) {
				workers.AddWorker(ctx, wname, func(_ context.Context, rname string, data interface{}) {
					job := data.(*job)
					defer job.cancel()

					procs.Store(rname, job)
					defer procs.Delete(rname)

					w.handleOptimismSignatureExchangeFromPubSub(job.ctx, job.peer, job.remote)
					w.meterPubsubJobs.Decr()
				})
				w.meterPubsubWorkers.Incr()
			}

			if data, ok := procs.Load(wname); ok {
				proc := data.(*job)
				if peer == proc.peer {
					continue
				}
				if strings.Compare(remote.Id, proc.remote.Id) < 1 {
					w.log.Debug("Skipped old signature",
						append(proc.logctx, "skipped-peer", peer, "skipped-id", remote.Id)...)
					continue
				}

				w.log.Info("Worker canceled because newer signature were received",
					append(proc.logctx, "newer-peer", peer, "newer-id", remote.Id)...)
				proc.cancel()
			}

			ctx, cancel := context.WithCancel(ctx)
			job := &job{
				ctx: ctx, cancel: cancel,
				peer: peer, remote: remote,
				logctx: []any{"peer", peer, "signer", wname, "remote-id", remote.Id},
			}
			workers.Enqueue(wname, job)
			w.meterPubsubJobs.Incr()
		}
	}
}

func (w *Node2) handleOptimismSignatureExchangeFromPubSub(
	ctx context.Context,
	sender peer.ID,
	remote *pbV1.OptimismSignature,
) bool {
	signer := common.BytesToAddress(remote.Signer)
	logctx := []interface{}{
		"peer", sender,
		"signer", signer,
		"remote-latest-id", remote.Id,
		"remote-latest-previous-id", remote.PreviousId,
		"remote-latest-index", remote.RollupIndex,
	}

	if err := verifySignature(w.hubLayerChainID, remote); err != nil {
		w.log.Error("Invalid signature", append(logctx, "err", err)...)
		return false
	}

	local, err := w.db.OPSignature.FindLatestsBySigner(signer, 1, 0)
	if err != nil {
		w.log.Error("Failed to find the latest signature", append(logctx, "err", err)...)
		return false
	} else if len(local) > 0 && local[0].PreviousID == remote.PreviousId {
		// duplicated
		return false
	} else if len(local) > 0 && strings.Compare(local[0].ID, remote.Id) == 1 {
		// fully synchronized or less than local
		w.log.Debug("Skip already possess signature", append(logctx, "local-latest-id", local[0].ID)...)
		return false
	}

	// save signature
	if _, err := w.db.OPSignature.Save(
		&remote.Id, &remote.PreviousId,
		signer,
		common.BytesToAddress(remote.Contract),
		remote.RollupIndex,
		common.BytesToHash(remote.RollupHash),
		remote.Approved,
		database.BytesSignature(remote.Signature),
	); err != nil {
		w.log.Error("Failed to save signature", append(logctx, "err", err)...)
		return false
	}

	return true
}

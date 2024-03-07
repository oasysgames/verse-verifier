package p2p

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/metrics"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	ps "github.com/libp2p/go-libp2p-pubsub"
	msgio "github.com/libp2p/go-msgio"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	meter "github.com/oasysgames/oasys-optimism-verifier/metrics"
	"github.com/oasysgames/oasys-optimism-verifier/p2p/pb"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oklog/ulid/v2"
	"golang.org/x/sync/semaphore"
	"golang.org/x/time/rate"
	"google.golang.org/protobuf/proto"
)

const (
	pubsubTopic    = "/oasys-optimism-verifier/pubsub/1.0.0"
	streamProtocol = "/oasys-optimism-verifier/stream/1.0.0"
)

var (
	eom = &pb.Stream{Body: &pb.Stream_Eom{Eom: nil}}

	// miscellaneous messages
	misc_SIGRECEIVED = []byte("SIGNATURES_RECEIVED")
)

type Node struct {
	cfg             *config.P2P
	db              *database.Database
	h               host.Host
	dht             routing.Routing
	bwm             *metrics.BandwidthCounter
	hpHelper        HolePunchHelper
	hubLayerChainID *big.Int
	ignoreSigners   map[common.Address]int
	stakemanager    *stakemanager.Cache

	topic *ps.Topic
	sub   *ps.Subscription
	log   log.Logger

	outboundSem, inboundSem     *semaphore.Weighted
	outboundThrot, inboundThrot *rate.Limiter

	meterPubsubSubscribed,
	meterPubsubUnknownMsg,
	meterPubsubWorkers,
	meterStreamOpend,
	meterStreamHandled,
	meterStreamClosed,
	meterStreamWrites,
	meterStreamReads,
	meterStreamUnknownMsg,
	meterHolePunchSuccess,
	meterHolePunchErrs,
	meterStreamOpenErrs,
	meterStreamReadErrs,
	meterStreamWriteErrs meter.Counter

	meterPeers,
	meterTCPConnections,
	meterUDPConnections,
	meterRelayConnections,
	meterRelayHopStreams,
	meterRelayStopStreams,
	meterVerifierStreams,
	meterPubsubJobs meter.Gauge
}

func NewNode(
	cfg *config.P2P,
	db *database.Database,
	host host.Host,
	dht routing.Routing,
	bwm *metrics.BandwidthCounter,
	hpHelper HolePunchHelper,
	hubLayerChainID uint64,
	ignoreSigners []common.Address,
	stakemanager *stakemanager.Cache,
) (*Node, error) {
	_, topic, sub, err := setupPubSub(context.Background(), host, pubsubTopic)
	if err != nil {
		return nil, err
	}

	worker := &Node{
		cfg:             cfg,
		db:              db,
		h:               host,
		dht:             dht,
		bwm:             bwm,
		hpHelper:        hpHelper,
		hubLayerChainID: new(big.Int).SetUint64(hubLayerChainID),
		ignoreSigners:   map[common.Address]int{},
		stakemanager:    stakemanager,
		topic:           topic,
		sub:             sub,
		log:             log.New("worker", "p2p"),

		outboundSem: semaphore.NewWeighted(int64(cfg.OutboundLimits.Concurrency)),
		inboundSem:  semaphore.NewWeighted(int64(cfg.InboundLimits.Concurrency)),
		outboundThrot: rate.NewLimiter(
			rate.Limit(cfg.OutboundLimits.Throttling), cfg.OutboundLimits.Throttling),
		inboundThrot: rate.NewLimiter(
			rate.Limit(cfg.InboundLimits.Throttling), cfg.InboundLimits.Throttling),

		meterPubsubSubscribed: meter.GetOrRegisterCounter([]string{"p2p", "pubsub", "subscribed"}, ""),
		meterPubsubUnknownMsg: meter.GetOrRegisterCounter([]string{"p2p", "pubsub", "unknown", "messages"}, ""),
		meterPubsubWorkers:    meter.GetOrRegisterCounter([]string{"p2p", "pubsub", "workers"}, ""),
		meterPubsubJobs:       meter.GetOrRegisterGauge([]string{"p2p", "pubsub", "jobs"}, ""),
		meterStreamOpend:      meter.GetOrRegisterCounter([]string{"p2p", "stream", "opened"}, ""),
		meterStreamHandled:    meter.GetOrRegisterCounter([]string{"p2p", "stream", "handled"}, ""),
		meterStreamClosed:     meter.GetOrRegisterCounter([]string{"p2p", "stream", "closed"}, ""),
		meterStreamWrites:     meter.GetOrRegisterCounter([]string{"p2p", "stream", "writes"}, ""),
		meterStreamReads:      meter.GetOrRegisterCounter([]string{"p2p", "stream", "reads"}, ""),
		meterStreamUnknownMsg: meter.GetOrRegisterCounter([]string{"p2p", "stream", "unknown", "messages"}, ""),
		meterHolePunchSuccess: meter.GetOrRegisterCounter([]string{"p2p", "holepunch", "successes"}, ""),
		meterHolePunchErrs:    meter.GetOrRegisterCounter([]string{"p2p", "holepunch", "errors"}, ""),
		meterStreamOpenErrs:   meter.GetOrRegisterCounter([]string{"p2p", "stream", "open", "errors"}, ""),
		meterStreamReadErrs:   meter.GetOrRegisterCounter([]string{"p2p", "stream", "read", "errors"}, ""),
		meterStreamWriteErrs:  meter.GetOrRegisterCounter([]string{"p2p", "stream", "write", "errors"}, ""),
		meterPeers:            meter.GetOrRegisterGauge([]string{"p2p", "peers"}, ""),
		meterTCPConnections:   meter.GetOrRegisterGauge([]string{"p2p", "tcp", "connections"}, ""),
		meterUDPConnections:   meter.GetOrRegisterGauge([]string{"p2p", "udp", "connections"}, ""),
		meterRelayConnections: meter.GetOrRegisterGauge([]string{"p2p", "relay", "connections"}, ""),
		meterRelayHopStreams:  meter.GetOrRegisterGauge([]string{"p2p", "relayhop", "streams"}, ""),
		meterRelayStopStreams: meter.GetOrRegisterGauge([]string{"p2p", "relaystop", "streams"}, ""),
		meterVerifierStreams:  meter.GetOrRegisterGauge([]string{"p2p", "verifier", "streams"}, ""),
	}

	for _, addr := range ignoreSigners {
		worker.ignoreSigners[addr] = 1
	}

	return worker, nil
}

func (w *Node) Start(ctx context.Context) {
	defer w.topic.Close()
	defer w.sub.Cancel()
	w.h.SetStreamHandler(streamProtocol, w.newStreamHandler(ctx))

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		w.meterLoop(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		w.publishLoop(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		w.subscribeLoop(ctx)
	}()

	w.showBootstrapLog()
	wg.Wait()
	w.log.Info("Worker stopped")
}

func (w *Node) PeerID() peer.ID                  { return w.h.ID() }
func (w *Node) Host() host.Host                  { return w.h }
func (w *Node) Routing() routing.Routing         { return w.dht }
func (w *Node) HolePunchHelper() HolePunchHelper { return w.hpHelper }

func (w *Node) meterLoop(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 15)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
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

func (w *Node) publishLoop(ctx context.Context) {
	ticker := time.NewTicker(w.cfg.PublishInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.publishLatestSignatures(ctx)
		}
	}
}

func (w *Node) subscribeLoop(ctx context.Context) {
	type job struct {
		ctx    context.Context
		cancel context.CancelFunc
		peer   peer.ID
		remote *pb.OptimismSignature
		logctx []any
	}

	// Storing workers and jobs.
	workers := util.NewWorkerGroup(100)
	procs := &sync.Map{}

	for {
		peer, msg, err := subscribe(ctx, w.sub, w.h.ID())
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

func (w *Node) newStreamHandler(ctx context.Context) network.StreamHandler {
	return func(s network.Stream) {
		defer w.closeStream(s)

		w.meterStreamHandled.Incr()

		peer := s.Conn().RemotePeer()
		for {
			m, err := w.readStream(s)
			if t, ok := err.(*ReadWriteError); ok {
				w.log.Debug("Failed to read stream message", "peer", peer, "err", t)
				return
			} else if err != nil {
				w.log.Debug(err.Error(), "peer", peer)
				continue
			}

			var disconnect bool
			switch t := m.Body.(type) {
			case *pb.Stream_FindCommonOptimismSignature:
				// received FindCommonOptimismSignature request
				disconnect = w.handleFindCommonOptimismSignature(s, t.FindCommonOptimismSignature)
			case *pb.Stream_OptimismSignatureExchange:
				// received signature exchange request
				disconnect = w.handleOptimismSignatureExchangeRequest(ctx, s, t)
			case *pb.Stream_Eom:
				// received last message
				return
			default:
				w.log.Warn("Received an unknown message", "peer", peer)
				w.meterStreamUnknownMsg.Incr()
				return
			}

			if disconnect {
				return
			}
		}
	}
}

func (w *Node) handleOptimismSignatureExchangeFromPubSub(
	ctx context.Context,
	sender peer.ID,
	remote *pb.OptimismSignature,
) {
	signer := common.BytesToAddress(remote.Signer)
	logctx := []interface{}{
		"peer", sender,
		"signer", signer,
		"remote-latest-id", remote.Id,
		"remote-latest-previous-id", remote.PreviousId,
		"remote-latest-index", remote.BatchIndex,
	}

	if err := verifySignature(w.hubLayerChainID, remote); err != nil {
		w.log.Error("Invalid signature", append(logctx, "err", err)...)
		return
	}

	local, err := w.db.OPSignature.FindLatestsBySigner(signer, 1, 0)
	if err != nil {
		w.log.Error("Failed to find the latest signature", append(logctx, "err", err)...)
		return
	} else if len(local) > 0 && strings.Compare(local[0].ID, remote.Id) == 1 {
		// fully synchronized or less than local
		return
	}

	// open stream to peer
	var s network.Stream
	openStream := func() error {
		if ss, err := w.openStream(ctx, sender); err != nil {
			return err
		} else {
			s = ss
			return nil
		}
	}
	returned := make(chan any)
	defer func() { close(returned) }()
	go func() {
		select {
		case <-ctx.Done():
			// canceled because newer signature were received
		case <-returned:
		}
		if s != nil {
			w.closeStream(s)
		}
	}()

	var idAfter string
	if len(local) == 0 {
		w.log.Info("Request all signatures", logctx...)
	} else {
		if openStream() != nil {
			return
		}
		if found, err := w.findCommonLatestSignature(ctx, s, signer); err == nil {
			fsigner := common.BytesToAddress(found.Signer)
			if fsigner != signer {
				w.log.Error("Signer does not match", append(logctx, "found-signer", fsigner)...)
				return
			}

			idAfter = found.Id
			w.log.Info("Found common signature from peer",
				"signer", signer, "id", found.Id, "previous-id", found.PreviousId)
		} else if errors.Is(err, database.ErrNotFound) {
			if localID, err := ulid.ParseStrict(local[0].ID); err == nil {
				// Prevent out-of-sync by specifying the ID of 1 second ago
				ms := localID.Time() - 1000
				idAfter = ulid.MustNew(ms, ulid.DefaultEntropy()).String()
				logctx = append(logctx, "local-id", local[0].ID, "created-after", time.UnixMilli(int64(ms)))
			} else {
				w.log.Error("Failed to parse ULID", "local-id", local[0].ID, "err", err)
				return
			}
		} else {
			return
		}

		w.log.Info("Request signatures", append(logctx, "id-after", idAfter)...)
	}

	// send request to peer
	m := &pb.Stream{
		Body: &pb.Stream_OptimismSignatureExchange{
			OptimismSignatureExchange: &pb.OptimismSignatureExchange{
				Requests: []*pb.OptimismSignatureExchange_Request{
					{
						Signer:  remote.Signer,
						IdAfter: idAfter,
					},
				},
			},
		},
	}
	if s == nil && openStream() != nil {
		return
	}
	if err = w.writeStream(s, m); err != nil {
		w.log.Error("Failed to send signature request", "err", err)
		return
	}

	w.handleOptimismSignatureExchangeResponses(ctx, s)
}

func (w *Node) handleOptimismSignatureExchangeRequest(
	ctx context.Context,
	s network.Stream,
	request *pb.Stream_OptimismSignatureExchange,
) (disconnect bool) {
	peerID := s.Conn().RemotePeer()
	logctx := []interface{}{"peer", peerID}

	requests := request.OptimismSignatureExchange.GetRequests()
	if len(requests) == 0 {
		w.log.Warn("No requests", logctx...)
		return false
	}

	// number of signatures obtained from the database
	queryLimit := w.cfg.InboundLimits.Throttling / w.cfg.InboundLimits.Concurrency

	// sending time limit
	isTimeup, timePenalty := func() (func() bool, func()) {
		limit := time.Now().Add(w.cfg.InboundLimits.MaxSendTime)
		return func() bool {
				return time.Now().After(limit)
			}, func() {
				limit = limit.Add(-(w.cfg.InboundLimits.MaxSendTime / 3))
			}
	}()

	// By finely acquiring the semaphore, it prevents
	// other peers from being blocked for a long time.
	sem := util.NewReleaseGuardSemaphore(w.inboundSem)
	defer sem.ReleaseALL()

	for _, req := range requests {
		signer := common.BytesToAddress(req.Signer)
		if w.stakemanager.StakeBySigner(signer).Cmp(ethutil.TenMillionOAS) == -1 {
			continue
		}

		logctx := append(logctx, "signer", signer, "id-after", req.IdAfter)
		w.log.Info("Received signature request", logctx...)

		for offset := 0; ; offset += queryLimit {
			if isTimeup() {
				w.log.Warn("Time up", logctx...)
				return true
			} else if err := sem.Acquire(ctx, 1); err != nil {
				w.log.Error("Failed to acquire inbound semaphore", append(logctx, "err", err)...)
				return true
			}

			// get latest signatures for each requested signer
			sigs, err := w.db.OPSignature.Find(
				&req.IdAfter, &signer, nil, nil, queryLimit, offset)
			sem.ReleaseALL()
			if err != nil {
				w.log.Error("Failed to find requested signatures",
					append(logctx, "err", err)...)
				break
			}

			sigLen := len(sigs)
			if sigLen == 0 {
				break // reached the last
			}
			w.throttling(w.inboundThrot, sigLen,
				"in", "handleOptimismSignatureExchangeRequest", "peer", peerID)

			responses := make([]*pb.OptimismSignature, sigLen)
			for i, sig := range sigs {
				responses[i] = toProtoBufSig(sig)
			}
			m := &pb.Stream{Body: &pb.Stream_OptimismSignatureExchange{
				OptimismSignatureExchange: &pb.OptimismSignatureExchange{
					Responses: responses,
				},
			}}

			// send response to peer
			if err := w.writeStream(s, m); err != nil {
				w.log.Error("Failed to send signatures", append(logctx, "err", err)...)
				return true
			}
			w.log.Info("Sent signatures", append(logctx, "sents", sigLen)...)

			// wait for received notify
			if m, err = w.readStream(s); err == nil && bytes.Equal(m.GetMisc(), misc_SIGRECEIVED) {
				w.log.Info("Received notification of receipt", logctx...)
			} else {
				timePenalty()
			}
		}
	}

	return false
}

func (w *Node) handleOptimismSignatureExchangeResponses(ctx context.Context, s network.Stream) {
	peerID := s.Conn().RemotePeer()
	logctx := []interface{}{"peer", peerID}

	for {
		m, err := w.readStream(s)
		if err != nil {
			w.log.Debug("Failed to read stream message", append(logctx, "err", err)...)
			return
		}

		body := m.GetOptimismSignatureExchange()
		if body == nil {
			if m.GetEom() != nil {
				w.log.Warn("Received an unknown message", logctx...)
				w.meterStreamUnknownMsg.Incr()
			}
			return
		}

		responses := body.GetResponses()
		if len(responses) == 0 {
			return
		}

		for _, res := range responses {
			signer := common.BytesToAddress(res.Signer)
			scc := common.BytesToAddress(res.Scc)
			logctx := append(logctx,
				"signer", signer,
				"id", res.Id, "previous-id", res.PreviousId,
				"scc", scc,
				"index", res.BatchIndex)

			if err := verifySignature(w.hubLayerChainID, res); err != nil {
				w.log.Error("Invalid signature", append(logctx, "err", err)...)
				return
			}
			if _, ok := w.ignoreSigners[signer]; ok {
				w.log.Info("Ignored", logctx...)
				return
			}

			// deduplication
			if local, err := w.db.OPSignature.FindByID(res.Id); err == nil && local.PreviousID == res.PreviousId {
				continue
			}

			// local is newer
			if local, err := w.db.OPSignature.Find(nil, &signer, &scc, &res.BatchIndex, 1, 0); err != nil {
				w.log.Error("Failed to find local signature", append(logctx, "err", err)...)
				return
			} else if len(local) > 0 && strings.Compare(local[0].ID, res.Id) == 1 {
				continue
			}

			if res.PreviousId != "" {
				_, err := w.db.OPSignature.FindByID(res.PreviousId)
				if errors.Is(err, database.ErrNotFound) {
					w.log.Warn("Previous ID does not exist", logctx...)
				} else if err != nil {
					w.log.Error("Failed to find previous signature", append(logctx, "err", err)...)
					return
				}
			}

			_, err := w.db.OPSignature.Save(
				&res.Id, &res.PreviousId,
				signer,
				scc,
				res.BatchIndex,
				common.BytesToHash(res.BatchRoot),
				res.Approved,
				database.BytesSignature(res.Signature),
			)
			if err != nil {
				w.log.Error("Failed to save signature", append(logctx, "err", err)...)
				return
			}
			w.log.Info("Received new signature", logctx...)
		}

		// send received notify
		w.writeStream(s, &pb.Stream{Body: &pb.Stream_Misc{Misc: misc_SIGRECEIVED}})
	}
}

func (w *Node) handleFindCommonOptimismSignature(
	s network.Stream,
	recv *pb.FindCommonOptimismSignature,
) (disconnect bool) {
	remotes := recv.Locals
	if len(remotes) == 0 {
		return false
	}

	w.log.Info("Received FindCommonOptimismSignature request",
		"from", remotes[0].Id, "to", remotes[len(remotes)-1].Id)

	var found *pb.OptimismSignature
	for _, remote := range remotes {
		local, err := w.db.OPSignature.FindByID(remote.Id)
		if errors.Is(err, database.ErrNotFound) {
			continue
		}
		if err != nil {
			w.log.Error("Failed to find signature", "remote-id", remote.Id, "err", err)
			return true
		}
		if local.PreviousID == remote.PreviousId {
			found = toProtoBufSig(local)
			break
		}
	}

	m := &pb.Stream{
		Body: &pb.Stream_FindCommonOptimismSignature{
			FindCommonOptimismSignature: &pb.FindCommonOptimismSignature{Found: found},
		},
	}
	if err := w.writeStream(s, m); err != nil {
		w.log.Error("Failed to send FindCommonOptimismSignature response", "err", err)
		return true
	} else {
		if found == nil {
			w.log.Info("Sent FindCommonOptimismSignature response", "found", found != nil)
		} else {
			w.log.Info("Sent FindCommonOptimismSignature response",
				"found", found != nil, "id", found.Id, "previous-id", found.PreviousId)
		}
	}
	return false
}

// Find the latest signature of the same ID and PreviousID from peer
func (w *Node) findCommonLatestSignature(
	ctx context.Context,
	s network.Stream,
	signer common.Address,
) (*pb.OptimismSignature, error) {
	peerID := s.Conn().RemotePeer()
	logctx := []interface{}{"peer", peerID, "signer", signer}
	limit := w.cfg.OutboundLimits.Throttling / w.cfg.OutboundLimits.Concurrency

	sem := util.NewReleaseGuardSemaphore(w.outboundSem)
	defer sem.ReleaseALL()

	for offset := 0; ; offset += limit {
		if err := sem.Acquire(ctx, 1); err != nil {
			w.log.Error("Failed to acquire outbound semaphore", append(logctx, "err", err)...)
			return nil, err
		}

		// find local latest signatures (order by: id desc)
		sigs, err := w.db.OPSignature.FindLatestsBySigner(signer, limit, offset)
		sem.ReleaseALL()
		if err != nil {
			w.log.Error("Failed to find latest signatures", append(logctx, "err", err)...)
			return nil, err
		}

		sigLen := len(sigs)
		if sigLen == 0 {
			break // reached the last
		}
		w.throttling(w.outboundThrot, sigLen, "in", "findCommonLatestSignature", "peer", peerID)

		logctx = append(logctx, "from", sigs[0].ID, "to", sigs[sigLen-1].ID)

		// construct protobuf message
		locals := make([]*pb.FindCommonOptimismSignature_Local, sigLen)
		for i, sig := range sigs {
			locals[i] = &pb.FindCommonOptimismSignature_Local{
				Id:         sig.ID,
				PreviousId: sig.PreviousID,
			}
		}
		req := &pb.Stream{Body: &pb.Stream_FindCommonOptimismSignature{
			FindCommonOptimismSignature: &pb.FindCommonOptimismSignature{Locals: locals},
		}}

		// send request
		if err = w.writeStream(s, req); err != nil {
			w.log.Error(
				"Failed to send FindCommonOptimismSignature request",
				append(logctx, "err", err)...)
			return nil, err
		}
		w.log.Info("Sent FindCommonOptimismSignature request", logctx...)

		// read response
		res, err := w.readStream(s)
		if err != nil {
			w.log.Error("Failed to read stream message", append(logctx, "err", err)...)
			return nil, err
		}

		t := res.GetFindCommonOptimismSignature()
		if t == nil {
			w.log.Error("Unexpected response", logctx...)
			return nil, errors.New("unexpected response")
		}
		if t.Found != nil {
			// found!
			return t.Found, nil
		}
	}

	w.log.Warn("Common signature not found", "signer", signer)
	return nil, database.ErrNotFound
}

func (w *Node) publishLatestSignatures(ctx context.Context) {
	latests, err := w.db.OPSignature.FindLatestsPerSigners()
	if err != nil {
		w.log.Error("Failed to find latest signatures", "err", err)
		return
	}
	filterd := []*database.OptimismSignature{}
	for _, sig := range latests {
		if w.stakemanager.StakeBySigner(sig.Signer.Address).Cmp(ethutil.TenMillionOAS) >= 0 {
			filterd = append(filterd, sig)
		}
	}
	if len(filterd) > 0 {
		w.PublishSignatures(ctx, filterd)
	}
}

func (w *Node) PublishSignatures(ctx context.Context, rows []*database.OptimismSignature) {
	sigs := &pb.OptimismSignatureExchange{
		Latests: make([]*pb.OptimismSignature, len(rows)),
	}
	for i, row := range rows {
		sigs.Latests[i] = toProtoBufSig(row)
	}

	m := &pb.PubSub{Body: &pb.PubSub_OptimismSignatureExchange{
		OptimismSignatureExchange: sigs,
	}}
	if err := publish(ctx, w.topic, m); err != nil {
		w.log.Error("Failed to publish latest signatures", "err", err)
		return
	}

	w.log.Info("Publish latest signatures", "len", len(rows))
}

func (w *Node) openStream(ctx context.Context, peer peer.ID) (network.Stream, error) {
	// If holepunch is available, attempt a direct connection.
	if !HasDirectConnection(w.h, peer) && w.hpHelper.Available(w.h) {
		if err := <-w.hpHelper.HolePunch(ctx, w.h, peer, DefaultHolePunchTimeout); err != nil {
			if !errors.Is(err, ErrPeerNotSupportHolePunch) {
				w.meterHolePunchErrs.Incr()
			}
		} else {
			w.meterHolePunchSuccess.Incr()
		}
	}

	// Note: `WithUseTransient` is required to open a stream via circuit relay.
	s, err := w.h.NewStream(network.WithUseTransient(ctx, streamProtocol), peer, streamProtocol)
	if err != nil {
		w.log.Error("Failed to open stream", "peer", peer, "err", err)
		w.meterStreamOpenErrs.Incr()
		return nil, err
	}

	w.meterStreamOpend.Incr()
	return s, nil
}

func (w *Node) writeStream(s network.Stream, m *pb.Stream) error {
	if w.cfg.StreamTimeout > 0 {
		s.SetWriteDeadline(time.Now().Add(w.cfg.StreamTimeout))
		defer s.SetWriteDeadline(time.Time{})
	}

	err := writeStream(s, m)
	_, isRWErr := err.(*ReadWriteError)
	_, isEOM := m.Body.(*pb.Stream_Eom)
	if err == nil {
		w.meterStreamWrites.Incr()
	} else if isRWErr && !isEOM {
		w.meterStreamWriteErrs.Incr()
	}
	return err
}

func (w *Node) readStream(s network.Stream) (m *pb.Stream, err error) {
	if w.cfg.StreamTimeout > 0 {
		s.SetReadDeadline(time.Now().Add(w.cfg.StreamTimeout))
		defer s.SetReadDeadline(time.Time{})
	}

	m, err = readStream(s)
	if err == nil {
		w.meterStreamReads.Incr()
		return m, nil
	} else if _, ok := err.(*ReadWriteError); ok {
		w.meterStreamReadErrs.Incr()
	}
	return nil, err
}

func (w *Node) closeStream(s network.Stream) {
	closeStream(s)
	w.meterStreamClosed.Incr()
}

func (w *Node) showBootstrapLog() {
	listens := []string{}
	for _, ma := range w.h.Network().ListenAddresses() {
		listens = append(listens, ma.String())
	}
	w.log.Info("Listening on: " + strings.Join(listens, ","))
	w.log.Info("Appended announce addresses: " + strings.Join(w.cfg.AppendAnnounce, ","))
	w.log.Info("No announce addresses: " + strings.Join(w.cfg.NoAnnounce, ","))
	w.log.Info("Connection filter addresses: " + strings.Join(w.cfg.ConnectionFilter, ","))
	if w.cfg.Transports.TCP {
		w.log.Info("Enabled TCP transport")
	}
	if w.cfg.Transports.QUIC {
		w.log.Info("Enabled QUIC transport")
	}
	w.log.Info("Bootnodes: " + strings.Join(w.cfg.Bootnodes, ","))
	w.log.Info("Enabled NAT Travasal features",
		"upnp", w.cfg.NAT.UPnP, "autonat", w.cfg.NAT.AutoNAT, "holepunch", w.hpHelper.Enabled())
	if w.cfg.RelayService.Enable {
		w.log.Info("Enabled circuit relay service")
	}
	if w.cfg.RelayClient.Enable {
		w.log.Info("Enabled circuit relay client, relay nodes: " + strings.Join(w.cfg.RelayClient.RelayNodes, ","))
	}
	w.log.Info("Worker started", "id", w.h.ID(),
		"publish-interval", w.cfg.PublishInterval,
		"stream-timeout", w.cfg.StreamTimeout,
		"outbound-limits-concurrency", w.cfg.OutboundLimits.Concurrency,
		"outbound-limits-throttling", w.cfg.OutboundLimits.Throttling,
		"inbound-limits-concurrency", w.cfg.InboundLimits.Concurrency,
		"inbound-limits-maxsendtime", w.cfg.InboundLimits.MaxSendTime,
		"inbound-limits-throttling", w.cfg.InboundLimits.Throttling,
	)
}

func (w *Node) throttling(limiter *rate.Limiter, num int, logCtx ...any) {
	rsv := limiter.ReserveN(time.Now(), num)
	if !rsv.OK() {
		w.log.Error("num is greater than burst", logCtx...)
		return
	}

	sleep := rsv.Delay()
	if sleep > 0 {
		w.log.Warn("Throttling", append(logCtx, "sleep", sleep)...)
		time.Sleep(sleep)
	}
}

// Write protobuf message to libp2p stream.
func writeStream(s io.Writer, m *pb.Stream) error {
	data, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	data, err = compress(data)
	if err != nil {
		return err
	}

	// Note: Intentionally not closing with `Close()` as it would also close the stream.
	if err := msgio.NewWriter(s).WriteMsg(data); err != nil {
		return &ReadWriteError{err}
	}

	return nil
}

// Read protobuf message from libp2p stream.
// Note: Will wait forever, should cancel.
func readStream(s io.Reader) (*pb.Stream, error) {
	reader := msgio.NewReader(s)
	msg, err := reader.ReadMsg()
	if err != nil {
		return nil, &ReadWriteError{err}
	}

	// Note: Forgetting to call `ReleaseMsg()` can result
	// in high memory consumption within libp2p/go-buffer-pool.
	defer reader.ReleaseMsg(msg)

	data, err := decompress(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress stream message: %w", err)
	}

	var m pb.Stream
	if err := proto.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stream message: %w", err)
	}

	return &m, nil
}

// Send end-of-message and close libp2p stream.
func closeStream(s network.Stream) {
	s.SetWriteDeadline(time.Now().Add(time.Second / 2))
	defer s.SetWriteDeadline(time.Time{})

	writeStream(s, eom)
	s.Close()
}

// Publish new message.
func publish(ctx context.Context, topic *ps.Topic, m *pb.PubSub) error {
	data, err := proto.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal pubsub message: %w", err)
	}

	if data, err = compress(data); err != nil {
		return fmt.Errorf("failed to compress pubsub message: %w", err)
	}
	if err := topic.Publish(ctx, data); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

// Subscribe new message.
// Note: Will wait forever, should cancel.
func subscribe(
	ctx context.Context,
	sub *ps.Subscription,
	self peer.ID,
) (peer.ID, *pb.PubSub, error) {
	recv, err := sub.Next(ctx)
	if err != nil {
		return "", nil, fmt.Errorf("failed to subscribe pubsub message: %w", err)
	}

	if recv.ReceivedFrom == self || recv.GetFrom() == self {
		return "", nil, errSelfMessage
	}

	data, err := decompress(recv.Data)
	if err != nil {
		return "", nil, fmt.Errorf("failed to decompress pubsub message: %w", err)
	}

	var m pb.PubSub
	if err = proto.Unmarshal(data, &m); err != nil {
		return "", nil, fmt.Errorf("failed to unmarshal pubsub message: %w", err)
	}

	return recv.GetFrom(), &m, nil
}

func verifySignature(hubLayerChainID *big.Int, sig *pb.OptimismSignature) error {
	// verify ulid
	if id, err := ulid.ParseStrict(sig.Id); err != nil {
		return err
	} else if id.Time() > uint64(time.Now().UnixMilli()) {
		return fmt.Errorf("future ulid: %s, timestamp: %d", sig.Id, id.Time())
	}

	signer := common.BytesToAddress(sig.Signer)
	scc := common.BytesToAddress(sig.Scc)
	batchIndex := new(big.Int).SetUint64(sig.BatchIndex)
	batchRoot := common.BytesToHash(sig.BatchRoot)

	msg := ethutil.NewMessage(hubLayerChainID, scc, batchIndex, batchRoot, sig.Approved)
	err := msg.VerifySigner(sig.Signature, signer)

	// possibly an old signature with an approved bug
	if _, ok := err.(*ethutil.SignerMismatchError); ok {
		msg = ethutil.NewMessageWithApprovedBug(
			hubLayerChainID, scc, batchIndex, batchRoot, sig.Approved)
		err = msg.VerifySigner(sig.Signature, signer)
	}

	return err
}

func toProtoBufSig(row *database.OptimismSignature) *pb.OptimismSignature {
	sig := &pb.OptimismSignature{
		Id:         row.ID,
		PreviousId: row.PreviousID,
		Signer:     row.Signer.Address[:],
		Scc:        row.Contract.Address[:],
		BatchIndex: row.RollupIndex,
		BatchRoot:  row.RollupHash[:],
		Approved:   row.Approved,
		Signature:  row.Signature[:],
	}

	if row.BatchSize != nil {
		sig.BatchSize = *row.BatchSize
	}
	if row.PrevTotalElements != nil {
		sig.PrevTotalElements = *row.PrevTotalElements
	}
	if row.ExtraData != nil {
		sig.ExtraData = *row.ExtraData
	}

	return sig
}

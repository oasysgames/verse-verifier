package p2p

import (
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
	"github.com/multiformats/go-multiaddr"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	meter "github.com/oasysgames/oasys-optimism-verifier/metrics"
	"github.com/oasysgames/oasys-optimism-verifier/p2p/pb"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oasysgames/oasys-optimism-verifier/verselayer"
	"github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/proto"
)

const (
	pubsubTopic    = "/oasys-optimism-verifier/pubsub/1.0.0"
	streamProtocol = "/oasys-optimism-verifier/stream/1.0.0"
)

const (
	warnQueueLen = 30
)

var (
	eom = &pb.Stream{Body: &pb.Stream_Eom{Eom: nil}}
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

	topic *ps.Topic
	sub   *ps.Subscription
	log   log.Logger

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

	meterConnections,
	meterPeers,
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
		topic:           topic,
		sub:             sub,
		log:             log.New("worker", "p2p"),

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
		meterConnections:      meter.GetOrRegisterGauge([]string{"p2p", "connections"}, ""),
		meterPeers:            meter.GetOrRegisterGauge([]string{"p2p", "peers"}, ""),
	}
	worker.h.SetStreamHandler(streamProtocol, worker.handleStream)

	for _, addr := range ignoreSigners {
		worker.ignoreSigners[addr] = 1
	}

	return worker, nil
}

func (w *Node) Start(ctx context.Context) {
	defer w.topic.Close()
	defer w.sub.Cancel()

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

	showBootstrapLog(w.log, w.cfg, w.h)
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
			w.meterConnections.Set(float64(len(w.h.Network().Conns())))
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
		from   peer.ID
		remote *pb.OptimismSignature
	}

	wg := util.NewWorkerGroup(100) // each signer address
	running := &sync.Map{}         // stores IDs in process for each signer

	for {
		from, msg, err := subscribe(ctx, w.sub, w.h.ID())
		if errors.Is(err, context.Canceled) {
			// worker stopped
			return
		} else if errors.Is(err, errSelfMessage) {
			continue
		} else if err != nil {
			w.log.Error("Failed to subscribe", "peer", from, "err", err)
			continue
		}
		w.meterPubsubSubscribed.Incr()

		t := msg.GetOptimismSignatureExchange()
		if t == nil {
			w.log.Warn("Unsupported pubsub message", "peer", from, "err", err)
			w.meterPubsubUnknownMsg.Incr()
			continue
		}

		for _, remote := range t.Latests {
			wname := common.BytesToAddress(remote.Signer).Hex()

			// skip if older than the ID being processed
			if proc, ok := running.Load(wname); ok &&
				strings.Compare(remote.Id, proc.(string)) < 1 {
				w.log.Debug("Skip pubsub",
					"peer", from, "signer", wname,
					"processed-id", proc, "remote-id", remote.Id)
				continue
			}
			running.Store(wname, remote.Id)

			// add new worker
			if !wg.Has(wname) {
				handler := func(ctx context.Context, rname string, data interface{}) {
					defer running.Delete(rname)
					defer w.meterPubsubJobs.Decr()

					if t, ok := data.(job); ok {
						st := time.Now()
						w.handleOptimismSignatureExchangeFromPubSub(ctx, t.from, t.remote)
						w.log.Debug("Worked pubsub",
							"peer", from, "signer", rname,
							"elapsed", time.Since(st), "remote-id", t.remote.Id)
					}
				}
				wg.AddWorker(ctx, wname, handler)
				w.meterPubsubWorkers.Incr()
			}

			wg.Enqueue(wname, job{from: from, remote: remote})
			w.meterPubsubJobs.Incr()

			qlen := len(wg.Queue(wname))
			w.log.Debug("Enqueue pubsub",
				"peer", from, "signer", wname,
				"remote-id", remote.Id, "queue-len", qlen)
			if qlen >= warnQueueLen {
				w.log.Warn("Long queue", "signer", wname, "queue-len", qlen)
			}
		}
	}
}

func (w *Node) handleStream(s network.Stream) {
	defer w.closeStream(s)
	w.meterStreamHandled.Incr()

	peer := s.Conn().RemotePeer()
	for {
		m, err := w.readStreamWithTimeout(context.Background(), s, w.cfg.StreamTimeout)
		if t, ok := err.(*ReadWriteError); ok {
			w.log.Error("Failed to read stream message", "peer", peer, "err", t)
			return
		} else if err != nil {
			w.log.Error(err.Error(), "peer", peer)
			continue
		}

		switch t := m.Body.(type) {
		case *pb.Stream_OptimismSignatureExchange:
			// received signature exchange request or response
			w.handleOptimismSignatureExchangeFromStream(s, t.OptimismSignatureExchange)
		case *pb.Stream_FindCommonOptimismSignature:
			// received FindCommonOptimismSignature request
			w.handleFindCommonOptimismSignature(s, t.FindCommonOptimismSignature)
		case *pb.Stream_Eom:
			// received last message
			return
		default:
			w.log.Warn("Received an unknown message", "peer", peer)
			w.meterStreamUnknownMsg.Incr()
			return
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
		"remote-id", remote.Id,
		"remote-previous-id", remote.PreviousId,
		"index", remote.BatchIndex,
	}

	if ok, err := verifySignature(w.hubLayerChainID, remote); !ok || err != nil {
		w.log.Error("Invalid signature", append(logctx, "verify", ok, "err", err)...)
		return
	}
	if _, ok := w.ignoreSigners[signer]; ok {
		w.log.Info("Ignored", logctx...)
		return
	}

	local, err := w.db.Optimism.FindLatestSignaturesBySigner(signer, 1, 0)
	if err != nil {
		w.log.Error("Failed to find the latest signature", append(logctx, "err", err)...)
		return
	}

	// open stream to peer
	var s network.Stream
	defer func() {
		if s != nil {
			w.closeStream(s)
		}
	}()
	openStream := func() error {
		if ss, err := w.openStream(ctx, sender); err != nil {
			return err
		} else {
			s = ss
			return nil
		}
	}

	var idAfter string
	if len(local) == 0 {
		w.log.Info("Request all signatures", logctx...)
	} else if strings.Compare(local[0].ID, remote.Id) == 1 {
		// fully synchronized or less than local
		return
	} else {
		if openStream() != nil {
			return
		}
		if found, err := w.findCommonLatestSignature(s, signer); err == nil {
			fsigner := common.BytesToAddress(found.Signer)
			if fsigner != signer {
				w.log.Error("Signer does not match", append(logctx, "found-signer", fsigner)...)
				return
			}

			idAfter = found.Id
			w.log.Info("Found common signature from peer",
				"signer", signer, "id", found.Id, "previous-id", found.PreviousId)
		} else {
			if localID, err := ulid.ParseStrict(local[0].ID); err == nil {
				// Prevent out-of-sync by specifying the ID of 1 second ago
				ms := localID.Time() - 1000
				idAfter = ulid.MustNew(ms, ulid.DefaultEntropy()).String()
				logctx = append(logctx, "local-id", local[0].ID, "created-after", time.UnixMilli(int64(ms)))
			} else {
				w.log.Error("Failed to parse ULID", "local-id", local[0].ID, "err", err)
				return
			}
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
	if err := w.writeStream(s, eom); err != nil {
		w.log.Error("Failed to send end-of-message", "err", err)
		return
	}

	// wait for signature exchange response
	w.handleStream(s)
}

func (w *Node) handleOptimismSignatureExchangeFromStream(
	s network.Stream,
	recv *pb.OptimismSignatureExchange,
) {
	if len(recv.Requests) > 0 {
		// received signature exchange request
		for _, req := range recv.Requests {
			signer := common.BytesToAddress(req.Signer)
			logctx := []interface{}{"signer", signer, "id-after", req.IdAfter}
			w.log.Info("Received signature request", logctx...)

			limit, offset := 1000, 0
			for {
				// get latest signatures for each requested signer
				sigs, err := w.db.Optimism.FindSignatures(
					&req.IdAfter, &signer, nil, nil, limit, offset)
				offset += limit
				if err != nil {
					w.log.Error("Failed to find requested signatures",
						append(logctx, "err", err)...)
					break
				} else if len(sigs) == 0 {
					break
				}

				responses := make([]*pb.OptimismSignature, len(sigs))
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
					return
				}

				w.log.Info("Sent signatures", "len", len(responses))
			}
		}
	} else if len(recv.Responses) > 0 {
		// save received signatures
		for _, res := range recv.Responses {
			signer := common.BytesToAddress(res.Signer)
			scc := common.BytesToAddress(res.Scc)
			logctx := []interface{}{
				"signer", signer, "id", res.Id,
				"scc", scc.Hex(), "index", res.BatchIndex,
			}

			if ok, err := verifySignature(w.hubLayerChainID, res); !ok || err != nil {
				w.log.Error("Invalid signature",
					append(logctx, "verify", ok, "err", err)...)
				return
			}
			if _, ok := w.ignoreSigners[signer]; ok {
				w.log.Info("Ignored", logctx...)
				return
			}

			// deduplication
			if local, err := w.db.Optimism.FindSignatureByID(res.Id); err == nil && local.PreviousID == res.PreviousId {
				continue
			}

			if res.PreviousId != "" {
				_, err := w.db.Optimism.FindSignatureByID(res.PreviousId)
				if errors.Is(err, database.ErrNotFound) {
					w.log.Warn("Previous ID does not exist", logctx...)
					return
				} else if err != nil {
					w.log.Error("Failed to find previous signature", append(logctx, "err", err)...)
					return
				}
			}

			_, err := w.db.Optimism.SaveSignature(
				&res.Id, &res.PreviousId,
				signer,
				scc,
				res.BatchIndex,
				common.BytesToHash(res.BatchRoot),
				res.BatchSize,
				res.PrevTotalElements,
				res.ExtraData,
				res.Approved,
				database.BytesSignature(res.Signature),
			)
			if err != nil {
				w.log.Error("Failed to save signature", append(logctx, "err", err)...)
				return
			}
			w.log.Info("Received new signature", logctx...)
		}
	}
}

func (w *Node) handleFindCommonOptimismSignature(
	s network.Stream,
	recv *pb.FindCommonOptimismSignature,
) {
	remotes := recv.Locals
	if len(remotes) == 0 {
		return
	}

	w.log.Info("Received FindCommonOptimismSignature request",
		"from", remotes[0].Id, "to", remotes[len(remotes)-1].Id)

	var found *pb.OptimismSignature
	for _, remote := range remotes {
		local, err := w.db.Optimism.FindSignatureByID(remote.Id)
		if errors.Is(err, database.ErrNotFound) {
			continue
		}
		if err != nil {
			w.log.Error("Failed to find signature", "remote-id", remote.Id, "err", err)
			return
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
	if err := w.writeStream(s, m); err == nil {
		if found == nil {
			w.log.Info("Sent FindCommonOptimismSignature response", "found", found != nil)
		} else {
			w.log.Info("Sent FindCommonOptimismSignature response",
				"found", found != nil, "id", found.Id, "previous-id", found.PreviousId)
		}
	} else {
		w.log.Error("Failed to send FindCommonOptimismSignature response", "err", err)
	}
}

// Find the latest signature of the same ID and PreviousID from peer
func (w *Node) findCommonLatestSignature(
	s network.Stream,
	signer common.Address,
) (*pb.OptimismSignature, error) {
	limit, offset := 100, 0
	for {
		logctx := []interface{}{"signer", signer}

		// find local latest signatures (order by: id desc)
		sigs, err := w.db.Optimism.FindLatestSignaturesBySigner(signer, limit, offset)
		if err != nil {
			w.log.Error("Failed to find latest signatures", append(logctx, "err", err)...)
			return nil, err
		}
		if len(sigs) == 0 {
			// reached the last
			break
		}
		logctx = append(logctx, "from", sigs[0].ID, "to", sigs[len(sigs)-1].ID)

		// construct protobuf message
		locals := make([]*pb.FindCommonOptimismSignature_Local, len(sigs))
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
		res, err := w.readStreamWithTimeout(context.Background(), s, time.Second*5)
		if errors.Is(err, context.DeadlineExceeded) {
			w.log.Warn("Timeout or peer does not support FindCommonOptimismSignature", logctx...)
			return nil, err
		} else if err != nil {
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

		offset += limit
	}

	w.log.Warn("Common signature not found", "signer", signer)
	return nil, errors.New("not found")
}

func (w *Node) publishLatestSignatures(ctx context.Context) {
	latests, err := w.db.Optimism.FindLatestSignaturePerSigners()
	if err != nil {
		w.log.Error("Failed to find latest signatures", "err", err)
		return
	}
	if len(latests) == 0 {
		return
	}
	w.PublishSignatures(ctx, latests)
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
	// If holepunch and UDP/QUIC is enabled, attempt a direct connection.
	if w.hpHelper.Enabled() && CheckAddressesProtocols(w.h.Network().ListenAddresses(),
		[]int{multiaddr.P_UDP, multiaddr.P_QUIC}, []int{multiaddr.P_CIRCUIT}) {
		if err := <-w.hpHelper.HolePunch(ctx, w.h, peer, DefaultHolePunchTimeout); err == nil {
			if !errors.Is(err, ErrPeerNotSupportHolePunch) {
				w.meterHolePunchErrs.Incr()
			}
		} else {
			w.meterHolePunchSuccess.Incr()
		}
	}

	s, err := w.h.NewStream(ctx, peer, streamProtocol)
	if err != nil {
		w.log.Error("Failed to open stream", "peer", peer, "err", err)
		w.meterStreamOpenErrs.Incr()
		return nil, err
	}

	w.meterStreamOpend.Incr()
	return s, nil
}

func (w *Node) writeStream(s network.Stream, m *pb.Stream) error {
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

func (w *Node) readStreamWithTimeout(
	parent context.Context,
	s network.Stream,
	timeout time.Duration,
) (m *pb.Stream, err error) {
	ctx, cancel := context.WithTimeout(parent, timeout)
	defer cancel()

	go func() {
		defer cancel()
		m, err = readStream(s)
		if err == nil {
			w.meterStreamReads.Incr()
		} else if _, ok := err.(*ReadWriteError); ok {
			w.meterStreamReadErrs.Incr()
		}
	}()
	<-ctx.Done()

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return nil, context.DeadlineExceeded
	}
	return m, err
}

func (w *Node) closeStream(s network.Stream) {
	closeStream(s)
	w.meterStreamClosed.Incr()
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

func verifySignature(hubLayerChainID *big.Int, sig *pb.OptimismSignature) (bool, error) {
	// verify ulid
	if id, err := ulid.ParseStrict(sig.Id); err != nil {
		return false, err
	} else if id.Time() > uint64(time.Now().UnixMilli()) {
		return false, fmt.Errorf("future ulid: %s, timestamp: %d", sig.Id, id.Time())
	}

	msg := verselayer.NewSccMessage(
		hubLayerChainID,
		common.BytesToAddress(sig.Scc),
		new(big.Int).SetUint64(sig.BatchIndex),
		common.BytesToHash(sig.BatchRoot),
		sig.Approved)
	return msg.VerifySigner(sig.Signature, common.BytesToAddress(sig.Signer))
}

func toProtoBufSig(row *database.OptimismSignature) *pb.OptimismSignature {
	return &pb.OptimismSignature{
		Id:                row.ID,
		PreviousId:        row.PreviousID,
		Signer:            row.Signer.Address[:],
		Scc:               row.OptimismScc.Address[:],
		BatchIndex:        row.BatchIndex,
		BatchRoot:         row.BatchRoot[:],
		BatchSize:         row.BatchSize,
		PrevTotalElements: row.PrevTotalElements,
		ExtraData:         row.ExtraData,
		Approved:          row.Approved,
		Signature:         row.Signature[:],
	}
}

func showBootstrapLog(log log.Logger, cfg *config.P2P, h host.Host) {
	listens := []string{}
	for _, ma := range h.Network().ListenAddresses() {
		listens = append(listens, ma.String())
	}
	log.Info("Listening on: " + strings.Join(listens, ","))
	log.Info("Appended announce addresses: " + strings.Join(cfg.AppendAnnounce, ","))
	log.Info("No announce addresses: " + strings.Join(cfg.NoAnnounce, ","))
	log.Info("Connection filter addresses: " + strings.Join(cfg.ConnectionFilter, ","))
	if cfg.Transports.TCP {
		log.Info("Enabled TCP transport")
	}
	if cfg.Transports.QUIC {
		log.Info("Enabled QUIC transport")
	}
	log.Info("Bootnodes: " + strings.Join(cfg.Bootnodes, ","))
	nats := []string{}
	if cfg.NAT.UPnP {
		nats = append(nats, "upnp")
	}
	if cfg.NAT.AutoNAT {
		nats = append(nats, "autonat")
	}
	if cfg.NAT.HolePunch {
		nats = append(nats, "holepunch")
	}
	log.Info("Enabled NAT Travasal features: " + strings.Join(nats, ","))
	if cfg.RelayService.Enable {
		log.Info("Enabled circuit relay service")
	}
	if cfg.RelayClient.Enable {
		log.Info("Enabled circuit relay client, relay nodes: " + strings.Join(cfg.RelayClient.RelayNodes, ","))
	}
	log.Info("Worker started", "id", h.ID(),
		"publish-interval", cfg.PublishInterval,
		"stream-timeout", cfg.StreamTimeout,
	)
}

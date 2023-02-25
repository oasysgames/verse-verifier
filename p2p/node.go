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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/metrics"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	msgio "github.com/libp2p/go-msgio"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
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

var (
	eom = &pb.Stream{Body: &pb.Stream_Eom{Eom: nil}}

	errUnavailableStream = errors.New("unavailable stream")
	errSelfMessage       = errors.New("self message")
)

type Node struct {
	db              *database.Database
	h               host.Host
	dht             *kaddht.IpfsDHT
	bwm             *metrics.BandwidthCounter
	publishInterval time.Duration
	hubLayerChainID *big.Int

	topic *ps.Topic
	sub   *ps.Subscription
	log   log.Logger
}

func NewNode(
	db *database.Database,
	host host.Host,
	dht *kaddht.IpfsDHT,
	bwm *metrics.BandwidthCounter,
	publishInterval time.Duration,
	hubLayerChainID uint64,
) (*Node, error) {
	_, topic, sub, err := setupPubSub(context.Background(), host, pubsubTopic)
	if err != nil {
		return nil, err
	}

	worker := &Node{
		db:              db,
		h:               host,
		dht:             dht,
		bwm:             bwm,
		publishInterval: publishInterval,
		hubLayerChainID: new(big.Int).SetUint64(hubLayerChainID),
		topic:           topic,
		sub:             sub,
		log:             log.New("worker", "p2p"),
	}
	worker.h.SetStreamHandler(streamProtocol, worker.handleStream)

	return worker, nil
}

func (w *Node) Start(ctx context.Context) {
	defer w.topic.Close()
	defer w.sub.Cancel()

	wg := &sync.WaitGroup{}

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

	w.log.Info("Worker started", "id", w.h.ID())
	wg.Wait()
	w.log.Info("Worker stopped")
}

func (w *Node) publishLoop(ctx context.Context) {
	ticker := time.NewTicker(w.publishInterval)
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
		from peer.ID
		msg  *pb.PubSub
	}

	qw := util.NewQueueWorker(ctx, func(ctx context.Context, data interface{}) {
		if t, ok := data.(job); ok {
			w.handlePubSubMessage(ctx, t.from, t.msg)
		}
	})
	go qw.Start(ctx)

	for {
		from, msg, err := subscribe(ctx, w.sub, w.h.ID())
		if err == nil {
			qw.Enqueue(job{from: from, msg: msg})
		} else if errors.Is(err, context.Canceled) {
			return
		} else if errors.Is(err, errSelfMessage) {
			continue
		} else {
			w.log.Error(err.Error())
		}
	}
}

func (w *Node) handleStream(s network.Stream) {
	defer closeStream(s)

	for {
		m, err := readStream(s)
		if errors.Is(err, errUnavailableStream) {
			return
		} else if err != nil {
			w.log.Error(err.Error())
			continue
		}

		switch t := m.Body.(type) {
		case *pb.Stream_Eom:
			// received the last message
			return
		case *pb.Stream_OptimismSignatureExchange:
			// received signature exchange request or response
			w.handleOptimismSignatureExchangeFromStream(s, t.OptimismSignatureExchange)
		}
	}
}

func (w *Node) handlePubSubMessage(ctx context.Context, sender peer.ID, m *pb.PubSub) {
	switch t := m.Body.(type) {
	case *pb.PubSub_OptimismSignatureExchange:
		// received peer's latest signature list
		w.handleOptimismSignatureExchangeFromPubSub(ctx, sender, t.OptimismSignatureExchange)
	}
}

func (w *Node) handleOptimismSignatureExchangeFromPubSub(
	ctx context.Context,
	sender peer.ID,
	recv *pb.OptimismSignatureExchange,
) {
	// get latest signatures for each signer
	var reqs []*pb.OptimismSignatureExchange_Request
	for _, remote := range recv.Latests {
		signer := common.BytesToAddress(remote.Signer)

		if ok, err := verifySignature(w.hubLayerChainID, remote); !ok || err != nil {
			w.log.Error("Invalid signature",
				"signer", signer, "id", remote.Id,
				"scc", common.BytesToAddress(remote.Scc), "index", remote.BatchIndex,
				"verify", ok, "err", err)
			continue
		}

		local, err := w.db.Optimism.FindLatestSignatureBySigner(signer)

		var idAfter string
		if err != nil {
			if errors.Is(err, database.ErrNotFound) {
				w.log.Info("Request all signatures", "signer", signer.Hex())
			} else {
				w.log.Error("Failed to find the latest signature", "signer", signer.Hex(), "err", err)
				continue
			}
		} else if strings.Compare(local.ID, remote.Id) == 1 {
			// fully synchronized or less than local
			continue
		} else {
			if localID, err := ulid.ParseStrict(local.ID); err == nil {
				// Prevent out-of-sync by specifying the ID of 1 second ago
				ms := localID.Time() - 1000
				idAfter = ulid.MustNew(ms, ulid.DefaultEntropy()).String()
				w.log.Info("Request signatures", "signer", signer.Hex(),
					"remote-id", remote.Id, "local-id", local.ID,
					"created-after", time.UnixMilli(int64(ms)))
			} else {
				w.log.Error("Failed to parse ULID", "local-id", local.ID, "err", err)
				continue
			}
		}

		reqs = append(reqs, &pb.OptimismSignatureExchange_Request{
			Signer:  remote.Signer,
			IdAfter: idAfter,
		})
	}
	if len(reqs) == 0 {
		return
	}

	// open stream to peer
	s, err := w.h.NewStream(ctx, sender, streamProtocol)
	if err != nil {
		w.log.Error("Failed to open stream", "err", err)
		return
	}
	defer closeStream(s)

	// send request to peer
	m := &pb.Stream{
		Body: &pb.Stream_OptimismSignatureExchange{
			OptimismSignatureExchange: &pb.OptimismSignatureExchange{Requests: reqs},
		},
	}
	if err = writeStream(s, m); err != nil {
		w.log.Error("Failed to send signature request", "err", err)
		return
	}

	if err := writeStream(s, eom); err != nil {
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
					responses[i] = &pb.OptimismSignature{
						Id:                sig.ID,
						PreviousId:        sig.PreviousID,
						Signer:            sig.Signer.Address[:],
						Scc:               sig.OptimismScc.Address[:],
						BatchIndex:        sig.BatchIndex,
						BatchRoot:         sig.BatchRoot[:],
						BatchSize:         sig.BatchSize,
						PrevTotalElements: sig.PrevTotalElements,
						ExtraData:         sig.ExtraData,
						Approved:          sig.Approved,
						Signature:         sig.Signature[:],
					}
				}
				m := &pb.Stream{Body: &pb.Stream_OptimismSignatureExchange{
					OptimismSignatureExchange: &pb.OptimismSignatureExchange{
						Responses: responses,
					},
				}}
				// send response to peer
				if err := writeStream(s, m); err != nil {
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

			// deduplication
			if _, err := w.db.Optimism.FindSignatureByID(res.Id); err == nil {
				continue
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

func (w *Node) publishLatestSignatures(ctx context.Context) {
	latests, err := w.db.Optimism.FindLatestSignaturePerSigners()
	if err != nil {
		w.log.Error("Failed to find latest signatures", "err", err)
		return
	}
	w.PublishSignatures(ctx, latests)
}

func (w *Node) PublishSignatures(ctx context.Context, rows []*database.OptimismSignature) {
	sigs := &pb.OptimismSignatureExchange{
		Latests: make([]*pb.OptimismSignature, len(rows)),
	}
	for i, row := range rows {
		sigs.Latests[i] = &pb.OptimismSignature{
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

	m := &pb.PubSub{Body: &pb.PubSub_OptimismSignatureExchange{
		OptimismSignatureExchange: sigs,
	}}
	if err := publish(ctx, w.topic, m); err != nil {
		w.log.Error("Failed to publish latest signatures", "err", err)
		return
	}

	w.log.Info("Publish latest signatures", "len", len(rows))
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
	if err := msgio.NewWriter(s).WriteMsg(data); err != nil {
		return errUnavailableStream
	}

	return nil
}

// Read protobuf message from libp2p stream.
// Note: Will wait forever, should cancel.
func readStream(s io.Reader) (*pb.Stream, error) {
	data, err := msgio.NewReader(s).ReadMsg()
	if err != nil {
		log.Error("Failed to read stream message", "err", err)
		return nil, errUnavailableStream
	}

	data, err = decompress(data)
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

	// verify signer
	msg := verselayer.NewSccMessage(
		hubLayerChainID,
		common.BytesToAddress(sig.Scc),
		new(big.Int).SetUint64(sig.BatchIndex),
		common.BytesToHash(sig.BatchRoot),
		sig.Approved)
	hash := crypto.Keccak256([]byte(msg.Eip712Msg))
	if recoverd, err := ethutil.Ecrecover(hash, sig.Signature); err != nil {
		return false, err
	} else {
		return bytes.Equal(recoverd.Bytes(), sig.Signer), nil
	}
}

package p2p

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/holepunch"
	p2pping "github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

var (
	// Empirically, there is no need to set a long timeout because
	// when it succeeds, it succeeds immediately, but when it fails,
	// it does not succeed no matter how long you wait.
	DefaultHolePunchTimeout = time.Second * 5
)

var (
	// This node does not have the libp2p hole punching service.
	ErrHolePunchDisabled = errors.New("holepunch disabled")
	// An unexpected error occurred during hole punching.
	ErrHolePunchProtocol = errors.New("holepunch protocol error")
	// Peer does not support libp2p hole punching.
	ErrPeerNotSupportHolePunch = errors.New("peer does not support hole punch")
)

type HolePunchHelper interface {
	holepunch.EventTracer

	// Return whether hole punch service is enabled at this node.
	Enabled() bool

	// Returns whether the hole punching protocol is actually available on this node.
	Available(host host.Host) bool

	// Attempting direct connection with peer using libp2p hole punching.
	// If the direct connection is successful, the channel will return nil.
	// Note that even if an error is returned from the channel, there may
	// still be a successful circuit relay connection, so it's up to
	// the caller to decide whether to continue processing.
	HolePunch(
		ctx context.Context,
		host host.Host,
		remote peer.ID,
		timeout time.Duration,
	) <-chan error
}

func NewHolePunchHelper(enableHolePunchingProtocol bool) HolePunchHelper {
	return &holePunchHelper{
		enabled: enableHolePunchingProtocol,
		topic:   util.NewTopic(),
		log:     log.New("worker", "holepunch"),
	}
}

type holePunchHelper struct {
	enabled bool
	topic   *util.Topic
	log     log.Logger
}

func (ht *holePunchHelper) Trace(evt *holepunch.Event) { ht.topic.Publish(evt) }

func (ht *holePunchHelper) Enabled() bool { return ht.enabled }

func (ht *holePunchHelper) Available(host host.Host) bool {
	if !ht.enabled {
		return false
	}
	for _, p := range host.Mux().Protocols() {
		if p == holepunch.Protocol {
			return true
		}
	}
	return false
}

func (ht *holePunchHelper) HolePunch(
	ctx context.Context,
	host host.Host,
	remote peer.ID,
	timeout time.Duration,
) <-chan error {
	result := make(chan error)

	if HasDirectConnection(host, remote) {
		ht.log.Debug("Already have a direct connection", "remote", remote)
		go func() { result <- nil }()
		return result
	}

	var preCheck error
	if !ht.enabled {
		preCheck = ErrHolePunchDisabled
	} else if err := ht.checkSupportHolePunch(host, remote); err != nil {
		preCheck = err
	}
	if preCheck != nil {
		go func() { result <- preCheck }()
		return result
	}

	go func() {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		// subscribe hole punching events
		handler, hpResult := ht.eventSubscriber(remote)
		sub := ht.topic.Subscribe(ctx, handler)
		defer sub()

		// start hole punching by opening a ping stream
		if r, ok := <-p2pping.Ping(ctx, host, remote); !ok || r.Error != nil {
			ht.log.Debug("Ping error", "remote", remote, "err", r.Error)
			result <- errors.New("ping error")
			return
		}

		// rechecking as supported protocols have been fetched via circuit relay
		if err := ht.checkSupportHolePunch(host, remote); err != nil {
			result <- err
			return
		}

		ht.log.Debug("Waiting for hole punch", "remote", remote)
		select {
		case <-ctx.Done():
			ht.log.Debug("Deadline exceeded", "remote", remote)
			result <- context.DeadlineExceeded
		// case result <- <-hpResult:
		case r := <-hpResult:
			ht.log.Debug("Hole punch finished", "remote", remote, "result", r)
			result <- r
		}
	}()

	return result
}

func (ht *holePunchHelper) eventSubscriber(remote peer.ID) (func(context.Context, interface{}), <-chan error) {
	result := make(chan error)
	return func(_ context.Context, data interface{}) {
		evt := data.(*holepunch.Event)
		if evt.Remote != remote {
			return
		}

		switch evt.Type {
		case holepunch.ProtocolErrorEvtT:
			ht.log.Debug("Protocol error", "remote", evt.Remote)
			result <- ErrHolePunchProtocol
		case holepunch.DirectDialEvtT:
			if evt.Evt.(*holepunch.DirectDialEvt).Success {
				ht.log.Debug("Successfully direct connection without holepunch", "remote", evt.Remote)
				result <- nil
			}
		case holepunch.EndHolePunchEvtT:
			if evt.Evt.(*holepunch.EndHolePunchEvt).Success {
				ht.log.Debug("Successfully holepunch", "remote", evt.Remote)
				result <- nil
			}
		case holepunch.StartHolePunchEvtT:
		case holepunch.HolePunchAttemptEvtT:
		default:
			ht.log.Debug("Unknown event", "type", evt.Type)
		}
	}, result
}

func (ht *holePunchHelper) checkSupportHolePunch(host host.Host, remote peer.ID) error {
	protos, err := host.Peerstore().GetProtocols(remote)
	if err != nil {
		return err
	} else if len(protos) == 0 {
		return nil
	}

	for _, p := range protos {
		if p == holepunch.Protocol {
			return nil
		}
	}
	return ErrPeerNotSupportHolePunch
}

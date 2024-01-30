package p2p

import (
	"errors"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/host/autorelay"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
	"github.com/oasysgames/oasys-optimism-verifier/config"
)

// Return the options for libp2p to enable the use of Circuit Relay.
// The `Relay Service` node requires public connectivity and must not be behind a NAT.
func CircuitRelayOpts(cfg *config.P2P) (opts []libp2p.Option, err error) {
	rs, rc := cfg.RelayService, cfg.RelayClient
	if rs.Enable && rc.Enable {
		return nil, errors.New("relay service and relay client cannot be used together")
	} else if !rs.Enable && !rc.Enable {
		return opts, nil
	}

	opts = append(opts, libp2p.EnableRelay())
	if rs.Enable {
		// Enable Circuit Relay Service.
		opts = append(opts, circuitRelayServiceOpts(cfg))
	} else if rc.Enable {
		// Enable Circuit Relay Client.
		if opt, err := circuitRelayClientOpts(cfg); err != nil {
			return nil, err
		} else {
			opts = append(opts, opt)
		}
	}

	return opts, nil
}

func circuitRelayServiceOpts(cfg *config.P2P) libp2p.Option {
	rs := cfg.RelayService
	resources := relay.DefaultResources()

	if rs.DurationLimit != nil {
		resources.Limit.Duration = *rs.DurationLimit
	}
	if rs.DataLimit != nil {
		resources.Limit.Data = *rs.DataLimit
	}
	if rs.ReservationTTL != nil {
		resources.ReservationTTL = *rs.ReservationTTL
	}
	if rs.MaxReservations != nil {
		resources.MaxReservations = *rs.MaxReservations
	}
	if rs.MaxCircuits != nil {
		resources.MaxCircuits = *rs.MaxCircuits
	}
	if rs.BufferSize != nil {
		resources.BufferSize = *rs.BufferSize
	}
	if rs.MaxReservationsPerPeer != nil {
		resources.MaxReservationsPerPeer = *rs.MaxReservationsPerPeer
	}
	if rs.MaxReservationsPerIP != nil {
		resources.MaxReservationsPerIP = *rs.MaxReservationsPerIP
	}
	if rs.MaxReservationsPerASN != nil {
		resources.MaxReservationsPerASN = *rs.MaxReservationsPerASN
	}

	return libp2p.EnableRelayService(relay.WithResources(resources))
}

func circuitRelayClientOpts(cfg *config.P2P) (libp2p.Option, error) {
	rc := cfg.RelayClient
	if len(rc.RelayNodes) == 0 {
		return nil, errors.New("relay node required")
	}

	var relayNodes []peer.AddrInfo
	for _, s := range rc.RelayNodes {
		if addr, err := peer.AddrInfoFromString(s); err != nil {
			return nil, err
		} else {
			relayNodes = append(relayNodes, *addr)
		}
	}

	relayOpts := []autorelay.Option{
		autorelay.WithCircuitV1Support(),
		autorelay.WithStaticRelays(relayNodes),
	}

	return libp2p.EnableAutoRelay(relayOpts...), nil
}

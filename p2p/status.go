package p2p

import (
	"sort"
	"strings"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	relayproto "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/proto"
	ma "github.com/multiformats/go-multiaddr"
)

type peerStream struct {
	ID        string `json:"id"`
	Opened    string `json:"opened"`
	Direction string `json:"direction"`
	Protocol  string `json:"protocol"`
}

type peerConn struct {
	ID            string        `json:"id"`
	Opened        string        `json:"opened"`
	Direction     string        `json:"direction"`
	Peer          string        `json:"peer"`
	LocalAddress  string        `json:"local_address"`
	RemoteAddress string        `json:"remote_address"`
	Streams       []*peerStream `json:"streams"`
}

type peerStatus struct {
	ID        string   `json:"id"`
	Addresses []string `json:"addresses"`
	Protocols []string `json:"protocols"`
}

type HostStatus struct {
	*peerStatus
	Connections []*peerConn   `json:"connections"`
	Peers       []*peerStatus `json:"peers"`
}

func NewHostStatus(h host.Host) (*HostStatus, error) {
	s := &HostStatus{
		peerStatus: &peerStatus{
			Addresses: []string{},
			Protocols: []string{},
		},
		Connections: []*peerConn{},
		Peers:       []*peerStatus{},
	}

	// set `ID`
	s.ID = h.ID().String()

	// set `Addresses`
	addrInfo := host.InfoFromHost(h)
	maddrs, err := peer.AddrInfoToP2pAddrs(addrInfo)
	if err != nil {
		return nil, err
	}
	for _, ma := range maddrs {
		s.Addresses = append(s.Addresses, ma.String())
	}
	sort.Strings(s.Addresses)

	// set `Protocols`
	s.Protocols = h.Mux().Protocols()
	sort.Strings(s.Protocols)

	// set `Connections`
	for _, conn := range h.Network().Conns() {
		streams := []*peerStream{}
		for _, st := range conn.GetStreams() {
			stat := st.Stat()
			streams = append(streams, &peerStream{
				ID:        st.ID(),
				Opened:    stat.Opened.UTC().Format("2006-01-02T15:04:05+00:00"),
				Protocol:  string(st.Protocol()),
				Direction: stat.Direction.String(),
			})
		}

		stat := conn.Stat()
		s.Connections = append(s.Connections, &peerConn{
			ID:            conn.ID(),
			Opened:        stat.Opened.UTC().Format("2006-01-02T15:04:05+00:00"),
			Direction:     stat.Direction.String(),
			Peer:          conn.RemotePeer().String(),
			LocalAddress:  conn.LocalMultiaddr().String(),
			RemoteAddress: conn.RemoteMultiaddr().String(),
			Streams:       streams,
		})
	}

	// set `Peers`
	for _, id := range h.Peerstore().Peers() {
		if id == h.ID() {
			continue
		}
		pinfo := &peerStatus{ID: id.String()}
		s.Peers = append(s.Peers, pinfo)

		// set `Peers[].Addresses`
		addrInfo := h.Peerstore().PeerInfo(id)
		maddrs, err := peer.AddrInfoToP2pAddrs(&addrInfo)
		if err != nil {
			return nil, err
		}
		for _, ma := range maddrs {
			pinfo.Addresses = append(pinfo.Addresses, ma.String())
		}
		sort.Strings(pinfo.Addresses)

		// set `Peers[].Protocols`
		pinfo.Protocols, _ = h.Peerstore().GetProtocols(id)
		sort.Strings(pinfo.Protocols)
	}
	sort.Slice(s.Peers, func(i, j int) bool {
		return strings.Compare(s.Peers[i].ID, s.Peers[j].ID) == -1
	})

	return s, nil
}

type networkStat struct {
	connections struct {
		tcp, udp, relay int
	}
	streams struct {
		hop, stop, verifier int
	}
}

func newNetworkStatus(h host.Host) *networkStat {
	var s networkStat
	for _, conn := range h.Network().Conns() {
		local := []ma.Multiaddr{conn.LocalMultiaddr()}
		if CheckAddressesProtocols(local, []int{ma.P_TCP}, nil) {
			s.connections.tcp++
		} else if CheckAddressesProtocols(local, []int{ma.P_UDP}, nil) {
			s.connections.udp++
		}

		remote := []ma.Multiaddr{conn.RemoteMultiaddr()}
		if CheckAddressesProtocols(remote, []int{ma.P_CIRCUIT}, nil) {
			s.connections.relay++
		}

		for _, st := range conn.GetStreams() {
			switch st.Protocol() {
			case relayproto.ProtoIDv2Hop:
				s.streams.hop++
			case relayproto.ProtoIDv2Stop:
				s.streams.stop++
			case streamProtocol:
				s.streams.verifier++
			}
		}
	}
	return &s
}

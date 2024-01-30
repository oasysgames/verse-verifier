package p2p

import (
	"sort"
	"strings"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
)

type peerConn struct {
	Peer    string `json:"peer"`
	Address string `json:"address"`
	Streams int    `json:"streams"`
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
		s.Connections = append(s.Connections, &peerConn{
			Peer:    conn.RemotePeer().String(),
			Address: conn.RemoteMultiaddr().String(),
			Streams: len(conn.GetStreams()),
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

package p2p

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/metrics"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	"github.com/libp2p/go-libp2p-kad-dht/dual"
	"github.com/multiformats/go-multiaddr"
	"github.com/oasysgames/oasys-optimism-verifier/config"

	ds "github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/libp2p/go-libp2p/p2p/protocol/holepunch"
	quic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
)

// Create libp2p host.
func NewHost(
	ctx context.Context,
	cfg *config.P2P,
	priv crypto.PrivKey,
) (host.Host, routing.Routing, *metrics.BandwidthCounter, HolePunchHelper, error) {
	// Construct libp2p host.
	bwm := metrics.NewBandwidthCounter()
	opts := []libp2p.Option{
		libp2p.DefaultMuxers,
		libp2p.DefaultSecurity,
		libp2p.Identity(priv),
		libp2p.BandwidthReporter(bwm),
	}

	appends, hpHelper, err := userOptions(cfg)
	if err != nil {
		return nil, nil, nil, nil, err
	} else {
		opts = append(opts, appends...)
	}

	h, err := libp2p.New(opts...)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Add log handler for peer connection and disconnection.
	h.Network().Notify(&network.NotifyBundle{
		ConnectedF: func(n network.Network, c network.Conn) {
			log.Info("Connected new peer", "peer", maToP2P(c.RemoteMultiaddr(), c.RemotePeer()))
		},
		DisconnectedF: func(n network.Network, c network.Conn) {
			log.Info("Disconnected peer", "peer", maToP2P(c.RemoteMultiaddr(), c.RemotePeer()))
		},
	})

	// Construct libp2p DHT.
	dht, err := newRouting(ctx, h, append(cfg.Bootnodes, cfg.RelayClient.RelayNodes...))
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to construct DHT: %w", err)
	}

	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table every five minutes.
	if err = dht.Bootstrap(ctx); err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to bootstrap DHT: %w", err)
	}

	return rhost.Wrap(h, dht), dht, bwm, hpHelper, nil
}

func userOptions(cfg *config.P2P) (opts []libp2p.Option, hpHelper HolePunchHelper, err error) {
	// Construct listening addresses.
	listens := cfg.Listens
	if cfg.Listen != "" {
		s := strings.Split(cfg.Listen, ":")
		listens = append(listens, fmt.Sprintf("/ip4/%s/tcp/%s", s[0], s[1]))
		listens = append(listens, fmt.Sprintf("/ip4/%s/udp/%s/quic", s[0], s[1]))
	}
	if len(listens) == 0 {
		return nil, nil, errors.New("no listening address")
	} else if listenAddrs, err := convertMultiaddrs(listens); err != nil {
		return nil, nil, fmt.Errorf("failed to parse listening addrs: %w", err)
	} else {
		opts = append(opts, libp2p.ListenAddrs(listenAddrs...))
	}

	// Construct address factory.
	opt, err := AddrsFactoryOpt(cfg.AppendAnnounce, cfg.NoAnnounce)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to construct addrs factory: %w", err)
	}
	opts = append(opts, opt)

	// Construct connection filter.
	if len(cfg.ConnectionFilter) > 0 {
		opt, err := ConnectionFilterOpt(cfg.ConnectionFilter)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to construct connection filter: %w", err)
		}
		opts = append(opts, opt)
	}

	// Enable transport protocols.
	if cfg.Transports.TCP {
		opts = append(opts, libp2p.Transport(tcp.NewTCPTransport))
	}
	if cfg.Transports.QUIC {
		opts = append(opts, libp2p.Transport(quic.NewTransport))
	}

	// Enable NAT traversal using UPnP.
	if cfg.NAT.UPnP {
		opts = append(opts, libp2p.NATPortMap())
	}

	// Enable NAT traversal using Hole Punching.
	hpHelper = NewHolePunchHelper(cfg.NAT.HolePunch && cfg.RelayClient.Enable)
	if cfg.NAT.HolePunch {
		if cfg.RelayClient.Enable {
			opts = append(opts, libp2p.EnableHolePunching(holepunch.WithTracer(hpHelper)))
		} else {
			log.Error("Holepunch has been disabled. Please enable the relay client to use holepunch")
		}
	}

	// Enable NAT traversal using `Circuit Relay v2`.
	if relayOpts, err := CircuitRelayOpts(cfg); err != nil {
		return nil, nil, err
	} else {
		opts = append(opts, relayOpts...)
	}

	// Enable AutoNAT service.
	if cfg.NAT.AutoNAT {
		opts = append(opts, libp2p.EnableNATService())
	}

	return opts, hpHelper, nil
}

func newRouting(ctx context.Context, h host.Host, bootstrapPeers []string) (routing.Routing, error) {
	opts := []kaddht.Option{
		kaddht.Datastore(dssync.MutexWrap(ds.NewMapDatastore())),
		kaddht.BootstrapPeers(ConvertPeers(bootstrapPeers)...),
	}
	return dual.New(ctx, h, dual.DHTOption(opts...))
}

// Convert libp2p multi-address string to peer.AddrInfo.
func ConvertPeers(peers []string) []peer.AddrInfo {
	pinfos := make([]peer.AddrInfo, len(peers))
	for i, addr := range peers {
		maddr := multiaddr.StringCast(addr)
		p, err := peer.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			log.Error("Failed to convert peer info", "addr", addr, "err", err)
		}
		pinfos[i] = *p
	}
	return pinfos
}

func convertMultiaddrs(in []string) (out []multiaddr.Multiaddr, err error) {
	out = make([]multiaddr.Multiaddr, len(in))
	for i, addr := range in {
		out[i], err = multiaddr.NewMultiaddr(addr)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func maToP2P(ma multiaddr.Multiaddr, id peer.ID) string {
	return fmt.Sprintf("%s/p2p/%s", ma.String(), id)
}

func GenerateKeyPair() (crypto.PrivKey, crypto.PubKey, peer.ID, error) {
	priv, pub, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, nil, "", err
	}

	peerID, err := peer.IDFromPrivateKey(priv)
	if err != nil {
		return nil, nil, "", err
	}

	return priv, pub, peerID, nil
}

func EncodePrivateKey(priv crypto.PrivKey) (string, error) {
	b, err := crypto.MarshalPrivateKey(priv)
	if err != nil {
		return "", err
	}
	return crypto.ConfigEncodeKey(b), nil
}

func DecodePrivateKey(b64encodedKey string) (crypto.PrivKey, peer.ID, error) {
	dec, err := crypto.ConfigDecodeKey(b64encodedKey)
	if err != nil {
		return nil, "", err
	}

	priv, err := crypto.UnmarshalPrivateKey(dec)
	if err != nil {
		return nil, "", err
	}

	peerID, err := peer.IDFromPrivateKey(priv)
	if err != nil {
		return nil, "", err
	}

	return priv, peerID, nil
}

// Check if a direct connection exists with the specified remote peer.
func HasDirectConnection(host host.Host, remote peer.ID) bool {
	for _, conn := range host.Network().ConnsToPeer(remote) {
		_, err := conn.RemoteMultiaddr().ValueForProtocol(multiaddr.P_CIRCUIT)
		if err != nil {
			return true
		}
	}
	return false
}

// Check if the addresses support the protocol. It can also check the local node,
// but be aware that it will always return false for unconnected remote peer.
func CheckAddressesProtocols(addrs []multiaddr.Multiaddr, desireds, excludeds []int) bool {
LOOP:
	for _, ma := range addrs {
		for _, p := range desireds {
			if _, err := ma.ValueForProtocol(p); err != nil {
				continue LOOP
			}
		}
		for _, p := range excludeds {
			if _, err := ma.ValueForProtocol(p); err == nil {
				continue LOOP
			}
		}
		return true
	}
	return false
}

func setupPubSub(
	ctx context.Context,
	h host.Host,
	topic_ string,
) (*ps.PubSub, *ps.Topic, *ps.Subscription, error) {
	// Create pubsub object.
	ps, err := ps.NewGossipSub(ctx, h)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create pubsub object: %w", err)
	}

	// Join to pubsub topic.
	topic, err := ps.Join(topic_)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to join the topic: %w", err)
	}

	// Subscribe to new message.
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to subscribe the topic: %w", err)
	}

	return ps, topic, sub, nil
}

func compress(data []byte) ([]byte, error) {
	var (
		buf bytes.Buffer
		w   = gzip.NewWriter(&buf)
	)
	if _, err := w.Write(data); err != nil {
		w.Close()
		return nil, err
	}
	if err := w.Flush(); err != nil {
		w.Close()
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decompress(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	if err = r.Close(); err != nil {
		return nil, err
	}
	if b, err = ioutil.ReadAll(r); err != nil {
		return nil, err
	}
	return b, nil
}

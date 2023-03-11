package p2p

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rand"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/log"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/metrics"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/multiformats/go-multiaddr"

	ds "github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
)

// Create libp2p host.
func NewHost(
	ctx context.Context,
	address, port string,
	priv crypto.PrivKey,
) (host.Host, *kaddht.IpfsDHT, *metrics.BandwidthCounter, error) {
	maddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s", address, port))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to parse multiaddress string: %w", err)
	}

	bwm := metrics.NewBandwidthCounter()
	h, err := libp2p.New(
		libp2p.DefaultTransports,
		libp2p.DefaultMuxers,
		libp2p.DefaultSecurity,
		libp2p.ListenAddrs(maddr),
		libp2p.Identity(priv),
		libp2p.NATPortMap(),
		libp2p.BandwidthReporter(bwm),
	)
	if err != nil {
		return nil, nil, nil, err
	}

	h.Network().Notify(&network.NotifyBundle{
		ConnectedF: func(n network.Network, c network.Conn) {
			log.Info("Connected new peer", "peer", maToP2P(c.RemoteMultiaddr(), c.RemotePeer()))
		},
		DisconnectedF: func(n network.Network, c network.Conn) {
			log.Info("Disconnected peer", "peer", maToP2P(c.RemoteMultiaddr(), c.RemotePeer()))
		},
	})

	// Create DHT.
	dht, err := kaddht.New(ctx, h, kaddht.Mode(kaddht.ModeAutoServer),
		// in-memory thread-safe datastore.
		kaddht.Datastore(dssync.MutexWrap(ds.NewMapDatastore())))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create Kademlia DHT: %w", err)
	}

	// Create routed host.
	h = rhost.Wrap(h, dht)

	return h, dht, bwm, nil
}

// Setup Kademlia DHT.
func Bootstrap(ctx context.Context, h host.Host, dht *kaddht.IpfsDHT) {
	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table every five minutes.
	dht.Bootstrap(ctx)
}

// Connect to peers.
func ConnectPeers(ctx context.Context, h host.Host, peers []peer.AddrInfo) {
	connectedPeers := map[peer.ID]int{}
	for _, p := range h.Network().Peers() {
		connectedPeers[p] = 1
	}

	for _, p := range peers {
		if _, ok := connectedPeers[p.ID]; ok {
			log.Debug("Peers are already connected", "id", p.ID)
			continue
		}

		go func(p peer.AddrInfo) {
			if err := h.Connect(ctx, p); err == nil {
				h.Peerstore().AddAddrs(p.ID, p.Addrs, peerstore.AddressTTL)
				log.Info("Connected to peer", "id", p.ID)
			} else {
				log.Error("Failed to connect to peer", "id", p.ID, "err", err)
			}
		}(p)
	}
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

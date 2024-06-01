package p2p

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/control"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	mafilter "github.com/whyrusleeping/multiaddr-filter"

	"github.com/multiformats/go-multiaddr"
)

func AddrsFactoryOpt(appendAnnounce, noAnnounce []string) (libp2p.Option, error) {
	appendAddrs, err := convertMultiaddrs(appendAnnounce)
	if err != nil {
		return nil, err
	}

	ipcidrMatches := multiaddr.NewFilters()
	exactMatches := map[string]bool{}
	for _, addr := range noAnnounce {
		if mask, err := mafilter.NewMask(addr); err == nil {
			ipcidrMatches.AddFilter(*mask, multiaddr.ActionDeny)
			continue
		}

		if maddr, err := multiaddr.NewMultiaddr(addr); err != nil {
			return nil, err
		} else {
			exactMatches[string(maddr.Bytes())] = true
		}
	}

	return libp2p.AddrsFactory(func(inputs []multiaddr.Multiaddr) (filterd []multiaddr.Multiaddr) {
		for _, maddr := range append(inputs, appendAddrs...) {
			if exactMatches[string(maddr.Bytes())] || ipcidrMatches.AddrBlocked(maddr) {
				continue
			}
			filterd = append(filterd, maddr)
		}
		return filterd
	}), nil
}

func ConnectionFilterOpt(filters []string) (libp2p.Option, error) {
	mafilters := multiaddr.NewFilters()
	for _, s := range filters {
		if mask, err := mafilter.NewMask(s); err != nil {
			return nil, fmt.Errorf("invalid formatted address filter: %s", s)
		} else {
			mafilters.AddFilter(*mask, multiaddr.ActionDeny)
		}
	}
	return libp2p.ConnectionGater((*multiAddrFilter)(mafilters)), nil
}

type multiAddrFilter multiaddr.Filters

func (f *multiAddrFilter) InterceptAddrDial(_ peer.ID, addr multiaddr.Multiaddr) (allow bool) {
	return !(*multiaddr.Filters)(f).AddrBlocked(addr)
}

func (f *multiAddrFilter) InterceptPeerDial(p peer.ID) (allow bool) {
	return true
}

func (f *multiAddrFilter) InterceptAccept(connAddr network.ConnMultiaddrs) (allow bool) {
	return !(*multiaddr.Filters)(f).AddrBlocked(connAddr.RemoteMultiaddr())
}

func (f *multiAddrFilter) InterceptSecured(_ network.Direction, _ peer.ID, connAddr network.ConnMultiaddrs) (allow bool) {
	return !(*multiaddr.Filters)(f).AddrBlocked(connAddr.RemoteMultiaddr())
}

func (f *multiAddrFilter) InterceptUpgraded(_ network.Conn) (allow bool, reason control.DisconnectReason) {
	return true, 0
}

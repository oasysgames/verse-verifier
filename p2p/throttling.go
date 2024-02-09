package p2p

import (
	"errors"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"golang.org/x/time/rate"
)

func newPeerThrottling(limit int) *peerThrottling {
	return &peerThrottling{
		limit:    limit,
		limiters: make(map[peer.ID]*rate.Limiter),
	}
}

type peerThrottling struct {
	limit    int
	limiters map[peer.ID]*rate.Limiter
	mu       sync.Mutex
}

func (th *peerThrottling) delay(peer peer.ID, num int, timeout time.Duration) (time.Duration, error) {
	rsv := th.limiter(peer).ReserveN(time.Now(), num)
	if !rsv.OK() {
		return 0, errors.New("num is greater than burst")
	}
	return rsv.Delay(), nil
}

func (th *peerThrottling) limiter(peer peer.ID) *rate.Limiter {
	th.mu.Lock()
	defer th.mu.Unlock()

	limiter, ok := th.limiters[peer]
	if !ok {
		limiter = rate.NewLimiter(rate.Limit(th.limit), th.limit)
		th.limiters[peer] = limiter
	}
	return limiter
}

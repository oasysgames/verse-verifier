package ipccmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	p2pping "github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/oasysgames/oasys-optimism-verifier/ipc"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

var PingCmd = &ping{handlerID: PING, attempts: 10}

type ping struct {
	handlerID int
	attempts  int
}

func (c *ping) Run(ctx context.Context, listen, peer string) {
	// attach to ipc
	cl, err := ipc.NewClient(listen, c.handlerID)
	if err != nil {
		util.Exit(1, "connection failure: %s\n", err)
	}
	defer cl.Close()

	// send ipc message
	if err = cl.Write([]byte(peer)); err != nil {
		util.Exit(1, "failed to write ipc message: %s\n", err)
	}

	// start read loop
	for {
		data, err := cl.Read()
		if errors.Is(err, io.EOF) {
			return
		} else if err != nil {
			util.Exit(1, "failed to read ipc message: %s\n", err)
		} else {
			fmt.Println(string(data))
		}
	}
}

func (c *ping) NewHandler(ctx context.Context, h host.Host) (handlerID int, handler ipc.Handler) {
	return c.handlerID, func(s *ipc.IPCServer, data []byte) {
		defer s.Write(ipc.EOM, nil)

		peerID, err := peer.Decode(string(data))
		if err != nil {
			s.Write(c.handlerID, []byte(fmt.Sprintf("failed to decode peer id: %s", err)))
			return
		}

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		pings := p2pping.Ping(ctx, h, peerID)
		for i := 0; i < c.attempts; i++ {
			r, ok := <-pings
			if !ok {
				return
			} else if r.Error != nil {
				s.Write(c.handlerID, []byte(fmt.Sprintf("error: %s", r.Error)))
				return
			}

			data = []byte(fmt.Sprintf("Pong received: time=%s", r.RTT))
			if err = s.Write(c.handlerID, data); err != nil {
				return
			}

			select {
			case <-ticker.C:
			case <-ctx.Done():
				return
			}
		}
	}
}

package ipccmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/oasysgames/oasys-optimism-verifier/ipc"
	"github.com/oasysgames/oasys-optimism-verifier/p2p"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

var StatusCmd = &status{handlerID: STATUS}

type status struct {
	handlerID int
}

func (c *status) NewHandler(h host.Host) (handlerID int, handler ipc.Handler) {
	type status struct {
		P2P *p2p.HostStatus `json:"p2p"`
	}

	return c.handlerID, func(s *ipc.IPCServer, _ []byte) {
		defer s.Write(ipc.EOM, nil)

		p2pStatus, err := p2p.NewHostStatus(h)
		if err != nil {
			s.Write(c.handlerID, []byte(fmt.Sprintf("failed to get p2p status: %s", err)))
			return
		}

		st := &status{P2P: p2pStatus}
		if data, err := json.Marshal(st); err != nil {
			s.Write(c.handlerID, []byte(fmt.Sprintf("failed to marshal status: %s", err)))
		} else {
			s.ChunkedWrite(c.handlerID, data)
		}
	}
}

func (c *status) Run(sockname string) {
	// attach to ipc
	cl, err := ipc.NewClient(sockname, c.handlerID)
	if err != nil {
		util.Exit(1, "connection failure: %s\n", err)
	}
	defer cl.Close()

	// send message
	if err = cl.Write(nil); err != nil {
		util.Exit(1, "failed to write ipc message: %s\n", err)
	}

	// read message
	var chunks [][]byte
	for {
		data, err := cl.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			util.Exit(1, "failed to read ipc message: %s\n", err)
		} else {
			chunks = append(chunks, data)
		}
	}

	fmt.Println(string(bytes.Join(chunks, nil)))
}

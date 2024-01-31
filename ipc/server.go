package ipc

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
	goipc "github.com/james-barrow/golang-ipc"
)

const (
	EOM = 65536
)

type Handler func(*IPCServer, []byte)

type IPCServer struct {
	sockname string

	s        *goipc.Server
	handlers *sync.Map
	log      log.Logger
}

func NewIPCServer(sockname string) (*IPCServer, error) {
	server, err := goipc.StartServer(sockname, nil)
	if err != nil {
		return nil, err
	}

	return &IPCServer{
		sockname: sockname,
		s:        server,
		handlers: &sync.Map{},
		log:      log.New("worker", "ipc"),
	}, nil
}

func (s *IPCServer) SetHandler(id int, handler Handler) {
	if _, ok := s.handlers.Load(id); !ok {
		s.handlers.Store(id, handler)
	}
}

func (s *IPCServer) Start(ctx context.Context) {
	go func() {
		for {
			func() {
				msg, err := s.s.Read()
				if err != nil {
					s.log.Error("Read error", "err", err)
					s.reConnect()
					return
				}

				if msg.MsgType == -1 {
					s.log.Debug("Status changed", "status", msg.Status)
					return
				}

				// message type -2 is an error, these won't automatically cause the recieve channel to close.
				if msg.MsgType == -2 {
					return
				}

				if handler, ok := s.handlers.Load(msg.MsgType); ok {
					handler.(Handler)(s, msg.Data)
				}
			}()
		}
	}()

	s.log.Info("Worker started", "sockname", s.sockname)
	<-ctx.Done()
	s.log.Info("Worker stopped")
}

func (s *IPCServer) Write(msgType int, message []byte) error {
	if err := s.s.Write(msgType, message); err != nil {
		s.log.Error("Failed to write ipc message", "err", err)
		return err
	}
	// If they do not sleep, clients will read messages in the wrong order.
	time.Sleep(time.Second / 4)
	return nil
}

func (s *IPCServer) reConnect() {
	s.s.Close()

	server, err := goipc.StartServer(s.sockname, nil)
	if err != nil {
		s.log.Error("Failed to re-connect", "err", err)
		return
	}

	s.s = server
}

package ipc

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
	goipc "github.com/james-barrow/golang-ipc"
)

const (
	EOM = 65536
)

const (
	chunkSize = 1024
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

func (s *IPCServer) Start() {
	s.log.Info("IPC server started", "sockname", s.sockname)

	for {
		msg, err := s.s.Read()
		if err != nil {
			s.log.Error("Read error", "err", err)
			s.reConnect()
			continue
		}

		if msg.MsgType == -1 {
			s.log.Debug("Status changed", "sockname", s.sockname, "status", msg.Status)
			if msg.Status == "Closed" || msg.Status == "Closing" {
				s.log.Info("IPC server closed", "sockname", s.sockname)
				return
			}
			if msg.Status == "Error" {
				s.log.Error("Error recieved", "sockname", s.sockname, "err", string(msg.Data))
				s.reConnect()
				continue
			}
			continue
		}

		// message type -2 is an error, these won't automatically cause the recieve channel to close.
		if msg.MsgType == -2 {
			s.log.Warn("Error recieved", "sockname", s.sockname, "err", string(msg.Data))
			continue
		}

		if handler, ok := s.handlers.Load(msg.MsgType); ok {
			handler.(Handler)(s, msg.Data)
		}
	}
}

func (s *IPCServer) Close() {
	s.s.Close()
}

func (s *IPCServer) Write(msgType int, message []byte) error {
	if err := s.s.Write(msgType, message); err != nil {
		s.log.Error("Failed to write ipc message", "err", err)
		return err
	}
	// If they do not sleep, clients will read messages in the wrong order.
	time.Sleep(time.Second / 10)
	return nil
}

func (s *IPCServer) ChunkedWrite(msgType int, message []byte) error {
	var chunks [][]byte
	for chunkSize < len(message) {
		message, chunks = message[chunkSize:],
			append(chunks, message[0:chunkSize:chunkSize])
	}
	chunks = append(chunks, message)

	for _, chunk := range chunks {
		if err := s.Write(msgType, chunk); err != nil {
			return err
		}
	}

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

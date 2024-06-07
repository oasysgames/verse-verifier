package ipc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type IPCServerTestSuite struct {
	suite.Suite
}

func TestIPCServer(t *testing.T) {
	suite.Run(t, new(IPCServerTestSuite))
}

func (s *IPCServerTestSuite) TestIPCServerClose() {
	var (
		sockname    = "ipc.sockname"
		ctx, cancel = context.WithTimeout(context.Background(), 200*time.Millisecond)
	)
	svr, err := NewIPCServer(sockname)
	s.NoError(err)

	go func() {
		svr.Start()
		cancel()
	}()

	err = svr.Close()
	s.NoError(err)

	<-ctx.Done()
}

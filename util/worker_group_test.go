package util

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/stretchr/testify/suite"
)

type WorkerGroupTestSuite struct {
	testhelper.Suite
}

func TestWorkerGroup(t *testing.T) {
	suite.Run(t, new(WorkerGroupTestSuite))
}

func (s *WorkerGroupTestSuite) TestWorkerGroup() {
	ctx := context.Background()
	msgs := strings.Split("hello, world!", "")

	workerSize := 1000
	taskSize := len(msgs) * workerSize
	semw := NewWorkerGroup(workerSize / 10)

	// add blocking worker
	semw.AddWorker(ctx, "blocking", func(ctx context.Context, _ string, data interface{}) {
		<-make(chan struct{})
	})
	semw.Enqueue("blocking", nil)

	// start workers
	wg := &sync.WaitGroup{}
	receivedCh := make(chan struct{ name, m string }, 999999)
	names := []string{}
	for i := 0; i < workerSize; i++ {
		name := fmt.Sprintf("worker-%d", i)
		names = append(names, name)

		semw.AddWorker(ctx, name, func(ctx context.Context, rname string, data interface{}) {
			defer wg.Done()

			s.Equal(rname, name)

			t, ok := data.(string)
			if !ok {
				s.Fail("invalid data")
			}
			receivedCh <- struct{ name, m string }{name, t}
		})
	}

	// enqueue
	wg.Add(taskSize)
	for _, m := range msgs {
		for _, name := range names {
			semw.Enqueue(name, m)
		}
	}
	wg.Wait()

	// assert
	received := map[string]string{}
	for range s.Range(0, taskSize) {
		r := <-receivedCh
		received[r.name] += r.m
	}
	s.Len(receivedCh, 0)

	s.Len(received, len(names))
	for _, name := range names {
		s.Equal("hello, world!", received[name])
	}
}

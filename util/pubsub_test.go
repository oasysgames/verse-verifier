package util

import (
	"context"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TopicTestSuite struct {
	suite.Suite
}

func TestTopic(t *testing.T) {
	suite.Run(t, new(TopicTestSuite))
}

func (s *TopicTestSuite) TestTopic() {
	subscribers := []*Handler{}
	cancels := map[*Handler]context.CancelFunc{}
	msgs := strings.Split("hello, world!", "")

	received := map[*Handler]string{}
	called := map[*Handler]int{}

	size := 1000
	topic := NewTopic()

	// blocking subscriber
	blocker := func(ctx context.Context, data interface{}) {
		<-make(chan struct{})
	}
	topic.Subscribe(context.Background(), blocker)

	// start subscribers
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for i := 0; i < size; i++ {
		var sub Handler
		sub = func(ctx context.Context, data interface{}) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			if t, ok := data.(string); ok {
				received[&sub] += t
				called[&sub] += 1
			}
		}

		subscribers = append(subscribers, &sub)
		cancels[&sub] = topic.Subscribe(context.Background(), sub)
	}
	s.Equal(topic.Len(), size+1)

	// publish messages
	wg.Add(size * len(msgs))
	for _, m := range msgs {
		topic.Publish(m)
	}
	wg.Wait()

	// stop subscribers
	wg.Add(len(cancels))
	for _, cancel := range cancels {
		go func(cancel context.CancelFunc) {
			cancel()
			wg.Done()
		}(cancel)
	}
	wg.Wait()
	s.Equal(topic.Len(), 1)

	// re-publish messages
	for _, m := range msgs {
		topic.Publish(m) // noop
	}

	// assert
	s.Len(received, size)
	s.Len(called, size)
	for _, sub := range subscribers {
		s.Equal("hello, world!", received[sub])
		s.Equal(len(msgs), called[sub])
	}
}

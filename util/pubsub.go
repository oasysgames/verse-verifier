package util

import (
	"context"
	"sync"
)

type Topic struct {
	workers *sync.Map
}

func NewTopic() *Topic {
	return &Topic{workers: &sync.Map{}}
}

func (t *Topic) Subscribe(workerCtx context.Context, handler Handler) (cancel context.CancelFunc) {
	w := NewQueueWorker(workerCtx, handler)
	go w.Start(workerCtx)

	t.workers.Store(w, w)
	return func() {
		t.workers.Delete(w)
	}
}

func (t *Topic) Publish(data interface{}) {
	t.workers.Range(func(key, value any) bool {
		value.(*JobQueueWorker).Enqueue(data)
		return true
	})
}

func (t *Topic) Len() int {
	n := 0
	t.workers.Range(func(key, value any) bool {
		n++
		return true
	})
	return n
}

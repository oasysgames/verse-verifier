package util

import (
	"context"
	"sync"
)

type Handler func(ctx context.Context, data interface{})

type JobQueueWorker struct {
	ctx     context.Context
	handler Handler
	queue   []interface{}
	notify  chan struct{}
	mu      *sync.Mutex
}

func NewQueueWorker(workerCtx context.Context, handler Handler) *JobQueueWorker {
	return &JobQueueWorker{
		ctx:     workerCtx,
		handler: handler,
		queue:   []interface{}{},
		notify:  make(chan struct{}),
		mu:      &sync.Mutex{},
	}
}

func (w *JobQueueWorker) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-w.notify:
			if data := w.shift(); data != nil {
				w.handler(ctx, data)
			}
		}
	}
}

func (w *JobQueueWorker) Enqueue(data interface{}) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.queue = append(w.queue, data)

	go func() {
		select {
		case <-w.ctx.Done():
			return
		case w.notify <- struct{}{}:
		}
	}()
}

func (w *JobQueueWorker) Queue() []interface{} {
	w.mu.Lock()
	defer w.mu.Unlock()

	cpy := make([]interface{}, len(w.queue))
	copy(cpy, w.queue)
	return cpy
}

func (w *JobQueueWorker) shift() interface{} {
	w.mu.Lock()
	defer w.mu.Unlock()

	if len(w.queue) == 0 {
		return nil
	}

	data := w.queue[0]
	w.queue = w.queue[1:]
	return data
}

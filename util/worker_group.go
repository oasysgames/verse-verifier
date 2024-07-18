package util

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

type WorkerGroup struct {
	sem     *semaphore.Weighted
	workers *sync.Map
}

func NewWorkerGroup(concurrency int) *WorkerGroup {
	return &WorkerGroup{
		sem:     semaphore.NewWeighted(int64(concurrency)),
		workers: &sync.Map{},
	}
}

func (sw *WorkerGroup) AddWorker(
	workerCtx context.Context,
	name string,
	handler func(ctx context.Context, name string, data interface{}),
) {
	if sw.Has(name) {
		return
	}

	w := NewQueueWorker(workerCtx, func(ctx context.Context, data interface{}) {
		if err := sw.sem.Acquire(ctx, 1); err != nil {
			return
		}
		defer sw.sem.Release(1)

		handler(ctx, name, data)
	})
	go w.Start(workerCtx)

	sw.workers.Store(name, w)
}

func (sw *WorkerGroup) RemoveWorker(name string) {
	if !sw.Has(name) {
		return
	}
	if w := sw.get(name); w != nil {
		w.Close()
		sw.workers.Delete(name)
	}
}

func (sw *WorkerGroup) Enqueue(name string, data interface{}) {
	if w := sw.get(name); w != nil {
		w.Enqueue(data)
	}
}

func (sw *WorkerGroup) Queue(name string) []interface{} {
	if w := sw.get(name); w != nil {
		return w.Queue()
	}
	return []interface{}{}
}

func (sw *WorkerGroup) Has(name string) bool {
	_, ok := sw.workers.Load(name)
	return ok
}

func (sw *WorkerGroup) get(name string) *JobQueueWorker {
	if v, ok := sw.workers.Load(name); !ok {
		return nil
	} else if t, ok := v.(*JobQueueWorker); !ok {
		return nil
	} else {
		return t
	}
}

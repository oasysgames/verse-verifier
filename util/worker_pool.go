package util

import (
	"context"
	"runtime"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

// WorkerPool is inspired by the worker pool implementation of fasthttp.
// https://github.com/valyala/fasthttp

// Returns a new worker pool.
// `maxWorkersCount` is the maximum number of goroutines that can be running concurrently within the pool.
// `maxIdleWorkerDuration(default=10s)` is the threshold for terminating idle workers.
// If there are no idle workers in the pool when a new job is executed via the `Work` method, the pool will
// recheck at intervals specified by `workerReleaseCheckInterval(default=0s)`. If zero is specified, the `Work` method
// will immediately return `false`. The loop that rechecks the pool will attempt retries up to the limit
// specified by `workerReleaseCheckTimeout(default=0s)`. If zero is specified, it will attempt retries indefinitely.
func NewWorkerPool[T any](
	log log.Logger,
	handler func(context.Context, T),
	maxWorkersCount int,
	maxIdleWorkerDuration,
	workerReleaseCheckInterval, workerReleaseCheckTimeout time.Duration,
) *WorkerPool[T] {
	return &WorkerPool[T]{
		logger:                     log,
		workerFuncFn:               handler,
		maxWorkersCount:            maxWorkersCount,
		maxIdleWorkerDuration:      maxIdleWorkerDuration,
		workerReleaseCheckInterval: workerReleaseCheckInterval,
		workerReleaseCheckTimeout:  workerReleaseCheckTimeout,
	}
}

type WorkerPool[T any] struct {
	logger          log.Logger
	workerFuncFn    func(context.Context, T)
	maxWorkersCount int
	maxIdleWorkerDuration,
	workerReleaseCheckInterval,
	workerReleaseCheckTimeout time.Duration

	workerChanPool sync.Pool
	stopCh         chan struct{}
	ready          []*workerChan[T]
	workersCount   int
	lock           sync.Mutex
	mustStop       bool
}

type workerPoolJob[T any] struct {
	ctx     context.Context
	data    T
	cleanup workerPoolJobCleanupFn[T]
}

type workerPoolJobCleanupFn[T any] func(context.Context, T)

type workerChan[T any] struct {
	lastUseTime time.Time
	ch          chan *workerPoolJob[T]
}

func (wp *WorkerPool[T]) Start() {
	if wp.stopCh != nil {
		return
	}
	wp.stopCh = make(chan struct{})
	stopCh := wp.stopCh
	wp.workerChanPool.New = func() any {
		return &workerChan[T]{
			ch: make(chan *workerPoolJob[T], workerChanCap),
		}
	}
	go func() {
		var scratch []*workerChan[T]
		for {
			wp.clean(&scratch)
			select {
			case <-stopCh:
				return
			default:
				time.Sleep(wp.getMaxIdle())
			}
		}
	}()
}

func (wp *WorkerPool[T]) Stop() {
	if wp.stopCh == nil {
		return
	}
	close(wp.stopCh)
	wp.stopCh = nil

	// Stop all the workers waiting for incoming connections.
	// Do not wait for busy workers - they will stop after
	// serving the connection and noticing wp.mustStop = true.
	wp.lock.Lock()
	ready := wp.ready
	for i := range ready {
		ready[i].ch <- nil
		ready[i] = nil
	}
	wp.ready = ready[:0]
	wp.mustStop = true
	wp.lock.Unlock()
}

func (wp *WorkerPool[T]) getMaxIdle() time.Duration {
	if wp.maxIdleWorkerDuration <= 0 {
		return 10 * time.Second
	}
	return wp.maxIdleWorkerDuration
}

var maxReleaseCheckLimit = time.Date(9999, 12, 31, 23, 59, 59, 999999999, time.UTC)

func (wp *WorkerPool[T]) getWorkerReleaseTimers() (interval *time.Ticker, timeout time.Time) {
	if wp.workerReleaseCheckInterval == 0 {
		return
	}
	interval = time.NewTicker(wp.workerReleaseCheckInterval)

	if wp.workerReleaseCheckTimeout == 0 {
		timeout = maxReleaseCheckLimit
	} else {
		timeout = time.Now().Add(wp.workerReleaseCheckTimeout)
	}
	return
}

func (wp *WorkerPool[T]) clean(scratch *[]*workerChan[T]) {
	maxIdleWorkerDuration := wp.getMaxIdle()

	// Clean least recently used workers if they didn't serve connections
	// for more than maxIdleWorkerDuration.
	criticalTime := time.Now().Add(-maxIdleWorkerDuration)

	wp.lock.Lock()
	ready := wp.ready
	n := len(ready)

	// Use binary-search algorithm to find out the index of the least recently worker which can be cleaned up.
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if criticalTime.After(wp.ready[mid].lastUseTime) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	i := r
	if i == -1 {
		wp.lock.Unlock()
		return
	}

	*scratch = append((*scratch)[:0], ready[:i+1]...)
	m := copy(ready, ready[i+1:])
	for i = m; i < n; i++ {
		ready[i] = nil
	}
	wp.ready = ready[:m]
	wp.lock.Unlock()

	// Notify obsolete workers to stop.
	// This notification must be outside the wp.lock, since ch.ch
	// may be blocking and may consume a lot of time if many workers
	// are located on non-local CPUs.
	tmp := *scratch
	for i := range tmp {
		tmp[i].ch <- nil
		tmp[i] = nil
	}
}

// Enqueue a new job to the pool. If there are no idle workers in the pool and
// `workerReleaseCheckInterval` is set to zero, it will immediately return false.
// Otherwise, the call will be blocked until it is canceled.  Jobs are executed
// asynchronously, so if you want to do something when the job finishes, pass `cleanup`.
// Note: `cleanup` will still be called even if there are no idle workers.
// Also, since workers are not released until `cleanup` is complete, long-running tasks are not recommended.
func (wp *WorkerPool[T]) Work(ctx context.Context, data T, cleanup workerPoolJobCleanupFn[T]) bool {
	waitTick, waitTimeout := wp.getWorkerReleaseTimers()
	if waitTick != nil {
		defer waitTick.Stop()
	}

	var ch *workerChan[T]
LOOP:
	for {
		ch = wp.getCh()
		if ch != nil || waitTick == nil {
			break
		}

		if time.Now().After(waitTimeout) {
			wp.logger.Warn("Exceeded maximum wait time")
			break
		}

		select {
		case <-ctx.Done():
			break LOOP
		case <-waitTick.C:
		}
	}

	if ch == nil {
		if cleanup != nil {
			cleanup(ctx, data)
		}
		return false
	}
	ch.ch <- &workerPoolJob[T]{ctx: ctx, data: data, cleanup: cleanup}
	return true
}

var workerChanCap = func() int {
	// Use blocking workerChan if GOMAXPROCS=1.
	// This immediately switches Serve to WorkerFunc, which results
	// in higher performance (under go1.5 at least).
	if runtime.GOMAXPROCS(0) == 1 {
		return 0
	}

	// Use non-blocking workerChan if GOMAXPROCS>1,
	// since otherwise the Serve caller (Acceptor) may lag accepting
	// new connections if WorkerFunc is CPU-bound.
	return 1
}()

func (wp *WorkerPool[T]) getCh() *workerChan[T] {
	var ch *workerChan[T]
	createWorker := false

	wp.lock.Lock()
	ready := wp.ready
	n := len(ready) - 1
	if n < 0 {
		if wp.workersCount < wp.maxWorkersCount {
			createWorker = true
			wp.workersCount++
		}
	} else {
		ch = ready[n]
		ready[n] = nil
		wp.ready = ready[:n]
	}
	wp.lock.Unlock()

	if ch == nil {
		if !createWorker {
			return nil
		}
		vch := wp.workerChanPool.Get()
		ch = vch.(*workerChan[T])
		go func() {
			wp.workerFunc(ch)
			wp.workerChanPool.Put(vch)
		}()
	}
	return ch
}

func (wp *WorkerPool[T]) release(ch *workerChan[T]) bool {
	ch.lastUseTime = time.Now()
	wp.lock.Lock()
	if wp.mustStop {
		wp.lock.Unlock()
		return false
	}
	wp.ready = append(wp.ready, ch)
	wp.lock.Unlock()
	return true
}

func (wp *WorkerPool[T]) workerFunc(ch *workerChan[T]) {
	var job *workerPoolJob[T]
	for job = range ch.ch {
		if job == nil {
			break
		}

		wp.workerFuncFn(job.ctx, job.data)
		if job.cleanup != nil {
			job.cleanup(job.ctx, job.data)
		}
		if !wp.release(ch) {
			break
		}
	}

	wp.lock.Lock()
	wp.workersCount--
	wp.lock.Unlock()
}

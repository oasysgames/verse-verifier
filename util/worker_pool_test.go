package util

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/assert"
)

func TestWorkerPool(t *testing.T) {
	assert := assert.New(t)

	maxWorkers := 10
	maxIdle := time.Millisecond * 10

	// Not waiting for worker release
	workerReleaseCheckInterval := time.Duration(0)
	workerReleaseCheckTimeout := time.Duration(0)

	var (
		mu               sync.Mutex
		calls, completed int
	)
	handler := func(ctx context.Context, job int) {
		mu.Lock()
		calls++
		mu.Unlock()

		<-ctx.Done()

		mu.Lock()
		completed++
		mu.Unlock()
	}

	wp := NewWorkerPool(log.Root(), handler, maxWorkers, maxIdle,
		workerReleaseCheckInterval, workerReleaseCheckTimeout)
	wp.Start()
	defer wp.Stop()

	// Check if the worker pool is empty
	assert.Equal(0, wp.workersCount)
	assert.Equal(0, len(wp.ready))

	// Run workers
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < maxWorkers*2; i++ {
		start := time.Now()
		ok := wp.Work(ctx, i, nil)
		elapsed := time.Since(start)

		if i < maxWorkers {
			assert.True(ok)
			assert.Equal(i+1, wp.workersCount)
		} else {
			assert.False(ok) // no idle workers
			assert.Equal(maxWorkers, wp.workersCount)
		}
		assert.Less(elapsed, time.Microsecond*100)

		assert.Equal(0, len(wp.ready))
	}

	// Wait for the worker to be executed by the worker pool
	time.Sleep(time.Millisecond)
	assert.Equal(maxWorkers, calls)

	// Wait for completes
	cancel()
	time.Sleep(time.Millisecond)
	assert.Equal(maxWorkers, completed)

	// Check if the idle worker has been released
	assert.Equal(maxWorkers, len(wp.ready))
	time.Sleep(maxIdle * 2)
	assert.Equal(0, len(wp.ready))

	// It should be run now that it has been released.
	assert.True(wp.Work(ctx, maxWorkers+1, nil))
}

func TestWorkerPool_WithWaitForRelease_WithoutTimeout(t *testing.T) {
	assert := assert.New(t)

	maxWorkers := 10
	maxIdle := time.Millisecond * 10

	// Wait for worker release without timeout
	workerReleaseCheckInterval := time.Millisecond
	workerReleaseCheckTimeout := time.Duration(0)

	var (
		mu               sync.Mutex
		calls, completed int
	)
	handler := func(ctx context.Context, job int) {
		mu.Lock()
		calls++
		mu.Unlock()

		time.Sleep(workerReleaseCheckInterval)

		mu.Lock()
		completed++
		mu.Unlock()
	}

	wp := NewWorkerPool(log.Root(), handler, maxWorkers, maxIdle,
		workerReleaseCheckInterval, workerReleaseCheckTimeout)
	wp.Start()
	defer wp.Stop()

	// Run workers
	ctx := context.Background()
	for i := 0; i < maxWorkers*2; i++ {
		start := time.Now()
		ok := wp.Work(ctx, i, nil)
		elapsed := time.Since(start)

		assert.True(ok)
		if i < maxWorkers {
			assert.Less(elapsed, time.Microsecond*100)
		} else {
			assert.Less(elapsed, workerReleaseCheckInterval*2)
		}
	}
}

func TestWorkerPool_WithWaitForRelease_WithTimeout(t *testing.T) {
	assert := assert.New(t)

	maxWorkers := 10
	maxIdle := time.Millisecond * 10

	// Wait for worker release with timeout
	workerReleaseCheckInterval := time.Millisecond
	workerReleaseCheckTimeout := workerReleaseCheckInterval * 2

	var (
		mu               sync.Mutex
		calls, completed int
	)
	handler := func(ctx context.Context, job int) {
		mu.Lock()
		calls++
		mu.Unlock()

		time.Sleep(time.Second)

		mu.Lock()
		completed++
		mu.Unlock()
	}

	wp := NewWorkerPool(log.Root(), handler, maxWorkers, maxIdle,
		workerReleaseCheckInterval, workerReleaseCheckTimeout)
	wp.Start()
	defer wp.Stop()

	// Run workers
	ctx := context.Background()
	for i := 0; i < maxWorkers*2; i++ {
		start := time.Now()
		ok := wp.Work(ctx, i, nil)
		elapsed := time.Since(start)

		if i < maxWorkers {
			assert.True(ok)
			assert.Less(elapsed, time.Microsecond*100)
		} else {
			assert.False(ok)
			assert.Greater(elapsed, workerReleaseCheckTimeout)
		}
	}
}

func TestWorkerPool_Cleanup(t *testing.T) {
	assert := assert.New(t)

	maxWorkers := 1
	maxIdle := time.Millisecond * 10
	workerReleaseCheckInterval := time.Millisecond
	workerReleaseCheckTimeout := workerReleaseCheckInterval * 2

	var handlerCalled, cleanupCalled atomic.Int32
	handler := func(ctx context.Context, job int32) {
		handlerCalled.Add(job)
	}
	cleanup := func(ctx context.Context, job int32) {
		cleanupCalled.Add(job)
		time.Sleep(time.Second)
	}

	wp := NewWorkerPool(log.Root(), handler, maxWorkers, maxIdle,
		workerReleaseCheckInterval, workerReleaseCheckTimeout)
	wp.Start()
	defer wp.Stop()

	assert.True(wp.Work(context.Background(), 1, cleanup))
	time.Sleep(time.Millisecond)

	// should be cleaned up even if there are no idle workers.
	assert.False(wp.Work(context.Background(), 2, cleanup))

	assert.Equal(int32(1), handlerCalled.Load())
	assert.Equal(int32(3), cleanupCalled.Load())
}

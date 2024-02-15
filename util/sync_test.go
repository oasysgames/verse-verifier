package util

import (
	"context"
	"errors"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"
)

func TestReleaseGuardSemaphore(t *testing.T) {
	parent := context.Background()
	gsem := NewReleaseGuardSemaphore(semaphore.NewWeighted(2))

	shouldErr := func() {
		ctx, cancel := context.WithTimeout(parent, time.Millisecond*5)
		defer cancel()
		if !errors.Is(gsem.Acquire(ctx, 1), context.DeadlineExceeded) {
			t.Error("Context.DeadlineExceeded should be returned")
		}
	}

	gsem.Acquire(parent, 1)
	gsem.Acquire(parent, 1)
	shouldErr()

	gsem.Release(1)
	gsem.Release(1)
	gsem.Acquire(parent, 2)
	shouldErr()

	gsem.ReleaseALL()
	gsem.Acquire(parent, 2)
	shouldErr()
}

package util

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

type ReleaseGuardSemaphore struct {
	sem *semaphore.Weighted
	mu  sync.Mutex
	cnt int64
}

func NewReleaseGuardSemaphore(sem *semaphore.Weighted) *ReleaseGuardSemaphore {
	return &ReleaseGuardSemaphore{sem: sem}
}

func (s *ReleaseGuardSemaphore) Acquire(ctx context.Context, n int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.sem.Acquire(ctx, n); err != nil {
		return err
	} else {
		s.cnt += n
		return nil
	}
}

func (s *ReleaseGuardSemaphore) Release(n int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if n > s.cnt {
		n = s.cnt
	}
	s.sem.Release(n)
	s.cnt -= n
}

func (s *ReleaseGuardSemaphore) ReleaseALL() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sem.Release(s.cnt)
	s.cnt = 0
}

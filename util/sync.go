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

type SyncMap[K comparable, V any] struct {
	in sync.Map
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.in.Store(key, value)
}

func (m *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	prev, loaded := m.in.Swap(key, value)
	if loaded {
		return prev.(V), true
	}
	return *new(V), false
}

func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	val, ok := m.in.Load(key)
	if ok {
		return val.(V), true
	}
	return *new(V), false
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.in.Delete(key)
}

func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.in.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

package util

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestSyncMap(t *testing.T) {
	var boolmap SyncMap[string, bool]
	got0, ok := boolmap.Load("key")
	assert.False(t, got0)
	assert.False(t, ok)

	boolmap.Store("key", true)
	got1, ok := boolmap.Load("key")
	assert.True(t, got1)
	assert.True(t, ok)
	boolmap.Range(func(key string, value bool) bool {
		assert.Equal(t, "key", key)
		assert.True(t, value)
		return true
	})

	boolmap.Delete("key")
	got2, ok := boolmap.Load("key")
	assert.False(t, got2)
	assert.False(t, ok)

	type item struct{ id int }
	var structmap SyncMap[string, *item]
	got3, ok := structmap.Load("key")
	assert.Nil(t, got3)
	assert.False(t, ok)

	structmap.Store("key1", &item{id: 100})
	got4, ok := structmap.Load("key1")
	assert.Equal(t, 100, got4.id)
	assert.True(t, ok)

	got5, loaded := structmap.Swap("key1", &item{id: 101})
	got6, ok := structmap.Load("key1")
	assert.Equal(t, 100, got5.id)
	assert.Equal(t, 101, got6.id)
	assert.True(t, loaded)
	assert.True(t, ok)

	got7, loaded := structmap.Swap("key2", &item{id: 102})
	assert.Nil(t, got7)
	assert.False(t, loaded)

	got8, ok := structmap.Load("key2")
	assert.Equal(t, 102, got8.id)
	assert.True(t, ok)

	total := 0
	structmap.Range(func(key string, value *item) bool {
		total++
		assert.Contains(t, []string{"key1", "key2"}, key)
		assert.Contains(t, []int{101, 102}, value.id)
		return true
	})
	assert.Equal(t, 2, total)

	structmap.Delete("key1")
	got9, ok := structmap.Load("key1")
	assert.Nil(t, got9)
	assert.False(t, ok)
}

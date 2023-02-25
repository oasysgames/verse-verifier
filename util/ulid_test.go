package util

import (
	"sync"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
)

func TestULID(t *testing.T) {
	proc := 10
	count := 10000

	wg := &sync.WaitGroup{}
	ch := make(chan ulid.ULID, proc*count)
	start := time.Now().UnixMilli()

	for i := 0; i < proc; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < count; j++ {
				id := ULID(nil)
				ch <- id
			}
		}()
	}

	wg.Wait()
	close(ch)
	end := time.Now().UnixMilli()

	m := &sync.Map{}
	for id := range ch {
		if _, loaded := m.LoadOrStore(id.String(), 0); loaded {
			t.Error("conflict")
		}

		ts := int64(id.Time())
		if ts < start || ts > end {
			t.Errorf("invalid time, id: %s, time: %d", id, ts)
		}
	}
}

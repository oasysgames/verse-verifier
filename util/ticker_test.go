package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTicker(t *testing.T) {
	d := time.Second / 2
	now := time.Now()

	times := []time.Time{}
	tick := NewTicker(d, 3)
	for i := 0; i < 4; i++ {
		times = append(times, <-tick.C)
	}

	assert.Less(t, times[0], now.Add(d))
	assert.Less(t, times[1], now.Add(d))
	assert.Less(t, times[2], now.Add(d))
	assert.Greater(t, times[3], now.Add(d))
}

package util

import "time"

type Ticker struct {
	*time.Ticker
	C <-chan time.Time
}

// A ticker that triggers immediately.
func NewTicker(d time.Duration, immediately int) *Ticker {
	ch := make(chan time.Time)
	parent := time.NewTicker(d)

	go func() {
		for i := 0; i < immediately; i++ {
			ch <- time.Now()
		}

		for t := range parent.C {
			ch <- t
		}
	}()

	return &Ticker{parent, ch}
}

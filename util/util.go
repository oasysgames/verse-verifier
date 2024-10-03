package util

import (
	"context"
	"fmt"
	"os"
	"sort"
	"time"
)

func Exit(code int, format string, args ...any) {
	fmt.Printf(format, args...)
	os.Exit(code)
}

func BytesToBytes32(s []byte) (a [32]byte) {
	var b32 [32]byte
	copy(b32[:], s)
	return b32
}

// Retry the given `fn` up to `attempts` times at the specified interval.
// If `ctx` is canceled, it will exit with an error regardless of the result.
// If `attempts“ is exceeded, the last error will be returned.
// If `attempts` is 0, it will retry indefinitely.
func Retry(ctx context.Context, attempts uint64, interval time.Duration, fn func() error) (err error) {
	condition := func() func() bool {
		if attempts == 0 {
			return func() bool { return true }
		}
		return func() bool {
			ok := attempts > 0
			if ok {
				attempts--
			}
			return ok
		}
	}()

	done := make(chan error)
	for condition() {
		go func() { done <- fn() }()

		select {
		case <-ctx.Done():
		case err = <-done:
		}

		// If canceled, prioritize that error.
		if ctx.Err() != nil {
			err = ctx.Err()
			break
		}
		// Success
		if err == nil {
			break
		}

		select {
		case <-ctx.Done():
		case <-time.NewTimer(interval).C:
		}
	}

	return err
}

// Return minimum and maximum values.
// If the `len(arr)` is zero, both default values are returned.
func MinMax[T int | int64 | uint | uint64 | time.Duration](arr ...T) (T, T) {
	if len(arr) == 0 {
		defv := new(T)
		return *defv, *defv
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return arr[0], arr[len(arr)-1]
}

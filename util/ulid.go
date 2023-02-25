package util

import (
	"crypto/rand"

	"github.com/oklog/ulid/v2"
)

var (
	defaultEntropy ulid.MonotonicReader
)

func init() {
	defaultEntropy = &ulid.LockedMonotonicReader{
		MonotonicReader: ulid.Monotonic(rand.Reader, 0),
	}
}

func ULID(entropy ulid.MonotonicReader) ulid.ULID {
	if entropy == nil {
		return ulid.MustNew(ulid.Now(), defaultEntropy)
	}
	return ulid.MustNew(ulid.Now(), entropy)
}

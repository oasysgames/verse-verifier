package util

import (
	"fmt"
	"os"
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

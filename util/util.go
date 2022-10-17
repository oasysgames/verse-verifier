package util

import (
	"fmt"
	"os"
)

func Exit(code int, format string, args ...any) {
	fmt.Printf(format, args...)
	os.Exit(code)
}

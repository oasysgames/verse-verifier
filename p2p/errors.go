package p2p

import "errors"

var (
	errSelfMessage = errors.New("self message")
)

type ReadWriteError struct {
	origin error
}

func (e *ReadWriteError) Error() string {
	return e.origin.Error()
}

func (e *ReadWriteError) String() string {
	return e.Error()
}

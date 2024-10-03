package util

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/stretchr/testify/suite"
)

type UtilTestSuite struct {
	testhelper.Suite
}

func TestUtil(t *testing.T) {
	suite.Run(t, new(UtilTestSuite))
}

func (s *UtilTestSuite) TestRetry() {
	parent := context.Background()

	// case: success
	var called int
	err := Retry(parent, 5, time.Millisecond, func() error {
		called++
		if called == 2 {
			return nil
		}
		return fmt.Errorf("called=%d", called)
	})
	s.Equal(2, called)
	s.NoError(err)

	// case: give up
	called = 0
	err = Retry(parent, 5, time.Millisecond, func() error {
		called++
		return fmt.Errorf("called=%d", called)
	})
	s.Equal(5, called)
	s.ErrorContains(err, "called=5")

	// case: infinite
	called = 0
	err = Retry(parent, 0, time.Millisecond, func() error {
		called++
		if called == 100 {
			return nil
		}
		return fmt.Errorf("called=%d", called)
	})
	s.Equal(100, called)
	s.NoError(err)

	// case: canceled
	called = 0
	ctx, cancel := context.WithTimeout(parent, time.Millisecond*10)
	defer cancel()

	st := time.Now()
	err = Retry(ctx, 0, time.Nanosecond, func() error {
		called++
		time.Sleep(time.Second) // heavy task
		return nil
	})

	s.Equal(1, called)
	s.Less(time.Since(st), time.Millisecond*15)
	s.ErrorIs(err, context.DeadlineExceeded)
}

func (s *UtilTestSuite) TestMinMax() {
	ints := []int{4, 5, 9, 2, 1, 6, 7, 3, 8}

	durations := make([]time.Duration, len(ints))
	for i, val := range ints {
		durations[i] = time.Duration(val)
	}

	min, max := MinMax(ints...)
	s.Equal(1, min)
	s.Equal(9, max)

	minD, maxD := MinMax(durations...)
	s.Equal(time.Duration(1), minD)
	s.Equal(time.Duration(9), maxD)
}

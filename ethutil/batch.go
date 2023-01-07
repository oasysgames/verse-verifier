package ethutil

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

type BatchHeaderClient interface {
	// Returns "start" to "end" block headers.
	// Returns an error if "end" is out of range
	Get(ctx context.Context, start, end uint64) ([]*types.Header, error)

	// Closes the client connection.
	Close() error
}

// A client that returns block headers from "start" to "end" like an iterator.
type BatchHeaderIterator struct {
	client     BatchHeaderClient
	start, end uint64
	limit      int
	call       int
}

// Returns the next block headers.
// Return empty slice if end is reached.
func (c *BatchHeaderIterator) Next(ctx context.Context) ([]*types.Header, error) {
	start := c.start + uint64(c.limit*c.call)
	if start > c.end {
		return []*types.Header{}, nil
	}

	end := start + uint64(c.limit-1)
	if end > c.end {
		end = c.end
	}

	if headers, err := c.client.Get(ctx, start, end); err != nil {
		return []*types.Header{}, nil
	} else {
		c.call += 1
		return headers, nil
	}
}

// Closes the client connection.
func (c *BatchHeaderIterator) Close() error {
	return c.client.Close()
}

func NewBatchHeaderIterator(
	client BatchHeaderClient,
	start, end uint64,
	limit int,
) *BatchHeaderIterator {
	return &BatchHeaderIterator{
		client: client,
		start:  start,
		end:    end,
		limit:  limit,
	}
}

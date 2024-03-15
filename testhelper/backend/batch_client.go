package backend

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

type BatchHeaderClient struct {
	ethutil.Client
}

func (c *BatchHeaderClient) Get(
	ctx context.Context,
	start, end uint64,
) ([]*types.Header, error) {
	size := int(end - start + 1)

	headers := make([]*types.Header, size)
	for i := 0; i < size; i++ {
		h, err := c.HeaderByNumber(ctx, new(big.Int).SetUint64(start+uint64(i)))
		if err != nil {
			return nil, err
		}
		headers[i] = h
	}

	return headers, nil
}

func (c *BatchHeaderClient) Close() error {
	return nil
}

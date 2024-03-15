package ethutil

import (
	"math/big"
)

var (
	OAS           = big.NewInt(1e18)
	TenMillionOAS = new(big.Int).Mul(OAS, big.NewInt(10_000_000))
)

package testhelper

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type StakeManagerMock struct {
	Owners     []common.Address
	Operators  []common.Address
	Stakes     []*big.Int
	Candidates []bool
}

func (b *StakeManagerMock) GetTotalStake(
	callOpts *bind.CallOpts,
	epoch *big.Int,
) (*big.Int, error) {
	tot := new(big.Int)
	for _, stake := range b.Stakes {
		tot.Add(tot, stake)
	}
	return tot, nil
}

func (b *StakeManagerMock) GetOperatorStakes(
	callOpts *bind.CallOpts,
	operator common.Address,
	epoch *big.Int,
) (*big.Int, error) {
	for i, addr := range b.Operators {
		if addr == operator {
			return b.Stakes[i], nil
		}
	}
	return big.NewInt(0), nil
}

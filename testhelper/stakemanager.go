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

func (b *StakeManagerMock) GetValidators(
	callOpts *bind.CallOpts,
	epoch, cursol, howMany *big.Int,
) (struct {
	Owners        []common.Address
	Operators     []common.Address
	Stakes        []*big.Int
	BlsPublicKeys [][]byte
	Candidates    []bool
	NewCursor     *big.Int
}, error) {
	length := big.NewInt(int64(len(b.Owners)))
	if new(big.Int).Add(cursol, howMany).Cmp(length) >= 0 {
		howMany = new(big.Int).Sub(length, cursol)
	}

	start := cursol.Uint64()
	end := start + howMany.Uint64()

	ret := struct {
		Owners        []common.Address
		Operators     []common.Address
		Stakes        []*big.Int
		BlsPublicKeys [][]byte
		Candidates    []bool
		NewCursor     *big.Int
	}{
		Owners:     b.Owners[start:end],
		Operators:  b.Operators[start:end],
		Stakes:     b.Stakes[start:end],
		Candidates: b.Candidates[start:end],
		NewCursor:  new(big.Int).Add(cursol, howMany),
	}

	return ret, nil
}

package submitter

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

type signatureIterator struct {
	db           *database.Database
	stakemanager *stakemanager.Cache
	contract     common.Address
	rollupIndex  uint64
}

func (si *signatureIterator) next() ([]*database.OptimismSignature, error) {
	rows, err := si.db.OPSignature.Find(nil, nil, &si.contract, &si.rollupIndex, 1000, 0)
	if err != nil {
		return nil, err
	}

	rows, err = filterSignatures(rows, ethutil.TenMillionOAS,
		si.stakemanager.TotalStake(), si.stakemanager.SignerStakes())
	if err != nil {
		return nil, err
	}

	si.rollupIndex++
	return rows, nil
}

func filterSignatures(
	rows []*database.OptimismSignature,
	minStake, totalStake *big.Int,
	signerStakes map[common.Address]*big.Int,
) (filterd []*database.OptimismSignature, err error) {
	// group by RollupHash and Approved
	type group struct {
		stake *big.Int
		rows  []*database.OptimismSignature
	}
	groups := map[string]*group{}

	for _, row := range rows {
		stake, ok := signerStakes[row.Signer.Address]
		if !ok || stake.Cmp(minStake) == -1 {
			continue
		}

		key := fmt.Sprintf("%s:%v", row.RollupHash, row.Approved)
		if _, ok := groups[key]; !ok {
			groups[key] = &group{stake: new(big.Int)}
		}

		groups[key].stake = new(big.Int).Add(groups[key].stake, stake)
		groups[key].rows = append(groups[key].rows, row)
	}
	if len(groups) == 0 {
		return nil, nil
	}

	// find the highest stake group
	var highest *group
	for key := range groups {
		if highest == nil || groups[key].stake.Cmp(highest.stake) == 1 {
			highest = groups[key]
		}
	}

	// check over half
	required := new(big.Int).Mul(new(big.Int).Div(totalStake, big.NewInt(100)), big.NewInt(51))
	if highest.stake.Cmp(required) == -1 {
		return nil, &StakeAmountShortage{required, highest.stake}
	}

	// sort by stake amount
	sort.Slice(highest.rows, func(i, j int) bool {
		a := signerStakes[highest.rows[i].Signer.Address]
		b := signerStakes[highest.rows[j].Signer.Address]
		return a.Cmp(b) == 1 // order by desc
	})

	// extract only amounts above the minimum stake
	exts := []*database.OptimismSignature{}
	amount := big.NewInt(0)
	for _, row := range highest.rows {
		exts = append(exts, row)
		amount.Add(amount, signerStakes[row.Signer.Address])
		if amount.Cmp(required) >= 0 {
			break
		}
	}

	// sort by signer address
	sort.Sort(database.OptimismSignatures(exts))
	return exts, nil
}

type StakeAmountShortage struct {
	required, actual *big.Int
}

func (err *StakeAmountShortage) Error() string {
	return fmt.Sprintf("stake amount shortage, required=%s actual=%s",
		fromWei(err.required), fromWei(err.actual))
}

func (err *StakeAmountShortage) Is(target error) bool {
	_, ok := target.(*StakeAmountShortage)
	return ok
}

package database

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
)

func (*OptimismState) idCol() string {
	return "optimism_states.id"
}

func (*OptimismState) contractCol() string {
	return "optimism_states.optimism_scc_id"
}

func (*OptimismState) rollupIndexCol() string {
	return "optimism_states.batch_index"
}

func (row *OptimismState) Logger(base log.Logger) log.Logger {
	return base.New("scc", row.Contract.Address, "rollup-index", row.BatchIndex)
}

func (row *OptimismState) GetContract() *OptimismContract {
	return &row.Contract
}

func (row *OptimismState) GetRollupIndex() uint64 {
	return row.BatchIndex
}

func (row *OptimismState) GetRollupHash() common.Hash {
	return row.BatchRoot
}

func (row *OptimismState) AssignEvent(contract *OptimismContract, e any) error {
	t, ok := e.(*scc.SccStateBatchAppended)
	if !ok {
		return errors.New("invalid event")
	}

	row.Contract = *contract
	row.BatchIndex = t.BatchIndex.Uint64()
	row.BatchRoot = t.BatchRoot
	row.BatchSize = t.BatchSize.Uint64()
	row.PrevTotalElements = t.PrevTotalElements.Uint64()
	row.ExtraData = t.ExtraData
	return nil
}

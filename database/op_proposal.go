package database

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

func (*OpstackProposal) idCol() string {
	return "opstack_proposals.id"
}

func (*OpstackProposal) contractCol() string {
	return "opstack_proposals.contract_id"
}

func (*OpstackProposal) rollupIndexCol() string {
	return "opstack_proposals.l2_output_index"
}

func (row *OpstackProposal) Logger(base log.Logger) log.Logger {
	return base.New("l2oo", row.Contract.Address, "rollup-index", row.L2OutputIndex)
}

func (row *OpstackProposal) GetContract() *OptimismContract {
	return &row.Contract
}

func (row *OpstackProposal) GetRollupIndex() uint64 {
	return row.L2OutputIndex
}

func (row *OpstackProposal) GetRollupHash() common.Hash {
	msg := ethutil.L2OORollupHashSource(
		row.OutputRoot,
		new(big.Int).SetUint64(row.L1Timestamp),
		new(big.Int).SetUint64(row.L2BlockNumber))
	return crypto.Keccak256Hash(msg)
}

func (row *OpstackProposal) AssignEvent(contract *OptimismContract, e any) error {
	t, ok := e.(*l2oo.OasysL2OutputOracleOutputProposed)
	if !ok {
		return errors.New("invalid event")
	}

	row.Contract = *contract
	row.L2OutputIndex = t.L2OutputIndex.Uint64()
	row.OutputRoot = t.OutputRoot
	row.L2BlockNumber = t.L2BlockNumber.Uint64()
	row.L1Timestamp = t.L1Timestamp.Uint64()
	return nil
}

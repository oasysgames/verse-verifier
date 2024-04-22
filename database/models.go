package database

import (
	"github.com/ethereum/go-ethereum/common"
)

// Model representing a block.
type Block struct {
	ID uint64 `gorm:"primarykey"`

	Number       uint64 `gorm:"uniqueIndex"`
	Hash         common.Hash
	LogCollected bool // Deprecated
}

// Model representing a rollup verifier.
type Signer struct {
	ID uint64 `gorm:"primarykey"`

	Address common.Address `gorm:"uniqueIndex"`
}

// Model representing a StateCommitmentChain or L2OutputOracle.
type OptimismContract struct {
	ID uint64 `gorm:"primarykey"`

	Address   common.Address `gorm:"uniqueIndex"`
	NextIndex uint64
}

func (OptimismContract) TableName() string {
	// Remnants of model name changed for OPStack support.
	return "optimism_sccs"
}

// Model representing rollup events from a StateCommitmentChain.
type OptimismState struct {
	ID uint64 `gorm:"primarykey"`

	ContractID uint64 `gorm:"column:optimism_scc_id;uniqueIndex:optimism_state_idx0,priority:1"`
	Contract   OptimismContract

	// Value of `StateBatchAppended` event.
	BatchIndex        uint64 `gorm:"uniqueIndex:optimism_state_idx0,priority:2;index:optimism_state_idx1"`
	BatchRoot         common.Hash
	BatchSize         uint64
	PrevTotalElements uint64
	ExtraData         []byte
}

// Model representing rollup events from a L2OutputOracle.
type OpstackProposal struct {
	ID uint64 `gorm:"primarykey"`

	ContractID uint64 `gorm:"uniqueIndex:opstack_proposal_idx0,priority:1"`
	Contract   OptimismContract

	// Value of `Types.OutputProposal` struct.
	L2OutputIndex uint64 `gorm:"uniqueIndex:opstack_proposal_idx0,priority:2;index:opstack_proposal_idx1"`
	OutputRoot    common.Hash
	L2BlockNumber uint64
	L1Timestamp   uint64
}

// Model representing signatures for rollups.
type OptimismSignature struct {
	ID         string `gorm:"primarykey;index:optimism_signature_idx3,priority:2"`
	PreviousID string

	SignerID uint64 `gorm:"uniqueIndex:optimism_signature_idx1,priority:1;index:optimism_signature_idx3,priority:1"`
	Signer   Signer

	ContractID uint64 `gorm:"column:optimism_scc_id;uniqueIndex:optimism_signature_idx1,priority:2"`
	Contract   OptimismContract

	RollupIndex uint64      `gorm:"column:batch_index;uniqueIndex:optimism_signature_idx1,priority:3;index:optimism_signature_idx2"`
	RollupHash  common.Hash `gorm:"column:batch_root"`

	BatchSize         *uint64 // Deprecated
	PrevTotalElements *uint64 // Deprecated
	ExtraData         *[]byte // Deprecated

	Approved  bool
	Signature Signature
}

func (OptimismSignature) TableName() string {
	// Remnants of model name changed for OPStack support.
	return "optimism_signatures"
}

// Model for storing miscellaneous data.
type Misc struct {
	ID    string `gorm:"primarykey"`
	Value []byte
}

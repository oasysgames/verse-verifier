package database

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"gorm.io/gorm"
)

type OPEvent interface {
	idCol() string
	contractCol() string
	rollupIndexCol() string

	Logger(base log.Logger) log.Logger
	GetContract() *OptimismContract
	GetRollupIndex() uint64
	GetRollupHash() common.Hash
	AssignEvent(contract *OptimismContract, e any) error
}

type OPEventConstraint[X any] interface {
	OPEvent
	*X
}

type OPEventDB[T any, PT OPEventConstraint[T]] struct {
	db *Database
	m  OPEvent
}

type IOPEventDB interface {
	DB() *Database
	FindByRollupIndex(contract common.Address, index uint64) (OPEvent, error)
	FindForVerification(
		signer common.Address,
		contract common.Address,
		nextIndex uint64,
		limit int,
	) ([]OPEvent, error)
	Save(contract common.Address, e any) (OPEvent, error)
	Deletes(contract common.Address, rollupIndex uint64) (int64, error)
}

func NewOPEventDB[T any, PT OPEventConstraint[T]](db *Database) IOPEventDB {
	return &OPEventDB[T, PT]{
		db: db,
		m:  PT(new(T)),
	}
}

func (db *OPEventDB[T, PT]) DB() *Database {
	return db.db
}

func (db *OPEventDB[T, PT]) FindByRollupIndex(contract common.Address, index uint64) (OPEvent, error) {
	_contract, err := db.db.OPContract.FindOrCreate(contract)
	if err != nil {
		return nil, err
	}

	row := PT(new(T))
	tx := db.db.rawdb.
		Joins("Contract").
		Where(fmt.Sprintf("%s = ?", db.m.contractCol()), _contract.ID).
		Where(fmt.Sprintf("%s = ?", db.m.rollupIndexCol()), index).
		First(row)

	if err := errconv(tx.Error); err != nil {
		return nil, err
	}
	return row, nil
}

// Return events waiting verification(order by BatchIndex).
func (db *OPEventDB[T, PT]) FindForVerification(
	signer common.Address,
	contract common.Address,
	nextIndex uint64,
	limit int,
) ([]OPEvent, error) {
	_signer, err := db.db.Signer.FindOrCreate(signer)
	if err != nil {
		return nil, err
	}

	_contract, err := db.db.OPContract.FindOrCreate(contract)
	if err != nil {
		return nil, err
	}

	if _contract.NextIndex > nextIndex {
		nextIndex = _contract.NextIndex
	}

	sub := db.db.rawdb.Model(&OptimismSignature{}).
		Select("batch_index").
		Where("optimism_scc_id = ? AND signer_id = ?", _contract.ID, _signer.ID).
		Where("batch_index >= ?", nextIndex)
	if sub.Error != nil {
		return nil, sub.Error
	}

	var rows []PT
	tx := db.db.rawdb.
		Joins("Contract").
		Where(fmt.Sprintf("%s = ?", db.m.contractCol()), _contract.ID).
		Where(fmt.Sprintf("%s >= ?", db.m.rollupIndexCol()), nextIndex).
		Where(fmt.Sprintf("%s NOT IN (?)", db.m.rollupIndexCol()), sub).
		Order(fmt.Sprintf("%s ASC", db.m.rollupIndexCol())).
		Limit(limit).
		Find(&rows)
	if tx.Error != nil {
		return nil, tx.Error
	}

	retRows := make([]OPEvent, len(rows))
	for i, row := range rows {
		retRows[i] = row
	}

	return retRows, nil
}

func (db *OPEventDB[T, PT]) Save(contract common.Address, e any) (OPEvent, error) {
	row := PT(new(T))

	err := db.db.Transaction(func(txdb *Database) error {
		if c, err := txdb.OPContract.FindOrCreate(contract); err != nil {
			return err
		} else if err := row.AssignEvent(c, e); err != nil {
			return err
		}
		return txdb.rawdb.Create(&row).Error
	})
	if err != nil {
		return nil, err
	}
	return row, err
}

func (db *OPEventDB[T, PT]) Deletes(contract common.Address, rollupIndex uint64) (int64, error) {
	_contract, err := db.db.OPContract.FindOrCreate(contract)
	if err != nil {
		return -1, err
	}

	var affected int64
	err = db.db.rawdb.Transaction(func(s *gorm.DB) error {
		var ids []uint64
		tx := s.
			Model(db.m).
			Joins("Contract").
			Where(fmt.Sprintf("%s = ?", db.m.contractCol()), _contract.ID).
			Where(fmt.Sprintf("%s >= ?", db.m.rollupIndexCol()), rollupIndex).
			Pluck(db.m.idCol(), &ids)
		if tx.Error != nil {
			return tx.Error
		}

		tx = s.Where("id IN (?)", ids).Delete(db.m)
		if tx.Error != nil {
			return tx.Error
		}

		affected = tx.RowsAffected
		return nil
	})
	if err != nil {
		return -1, err
	}

	return affected, nil
}

func NewMessage(event OPEvent, l1ChainID *big.Int, approved bool) *ethutil.Message {
	return ethutil.NewMessage(
		l1ChainID,
		event.GetContract().Address,
		new(big.Int).SetUint64(event.GetRollupIndex()),
		event.GetRollupHash(),
		approved,
	)
}

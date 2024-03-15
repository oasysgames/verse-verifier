package database

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type OptimismContractDB db

func (db *OptimismContractDB) FindOrCreate(addr common.Address) (row *OptimismContract, err error) {
	err = db.rawdb.Transaction(func(txdb *gorm.DB) error {
		tx := txdb.Where("address = ?", addr).First(&row)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			row.Address = addr
			row.NextIndex = 0
			return txdb.Create(&row).Error
		}
		return tx.Error
	})
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (db *OptimismContractDB) SaveNextIndex(addr common.Address, nextIndex uint64) error {
	row, err := db.FindOrCreate(addr)
	if err != nil {
		return err
	}

	row.NextIndex = nextIndex
	return db.rawdb.Save(row).Error
}

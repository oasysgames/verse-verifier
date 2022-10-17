package database

import (
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type BlockDatabase struct {
	db *gorm.DB
}

// Return the specific block.
func (db *BlockDatabase) Find(number uint64) (*Block, error) {
	var row Block
	tx := db.db.
		Where("number = ?", number).
		First(&row)

	if err := errconv(tx.Error); err != nil {
		return nil, err
	}
	return &row, nil
}

// Return the highest block.
func (db *BlockDatabase) FindHighest() (*Block, error) {
	var row Block
	tx := db.db.
		Order("number DESC").
		First(&row)

	if err := errconv(tx.Error); err != nil {
		return nil, err
	}
	return &row, nil
}

// Returns blocks for uncollected event logs(order by block number).
func (db *BlockDatabase) FindUncollecteds(limit int) ([]*Block, error) {
	var rows []*Block
	tx := db.db.
		Where("log_collected IS FALSE").
		Order("number ASC").
		Limit(limit).
		Find(&rows)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

// Save the new block.
func (db *BlockDatabase) SaveNewBlock(number uint64, hash common.Hash) error {
	tx := db.db.Create(&Block{Number: number, Hash: hash})
	return tx.Error
}

// Save event log collected block.
func (db *BlockDatabase) SaveLogCollected(number uint64) error {
	tx := db.db.
		Model(&Block{}).
		Where("number = ?", number).
		Update("log_collected", true)
	return tx.Error
}

// Delete the specific block.
func (db *BlockDatabase) Delete(number uint64) (rows int64, err error) {
	tx := db.db.
		Where("number = ?", number).
		Delete(&Block{})
	if tx.Error != nil {
		return -1, tx.Error
	}
	return tx.RowsAffected, nil
}

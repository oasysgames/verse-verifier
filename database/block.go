package database

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	var rows []*Block
	tx := db.db.
		Order("number DESC").
		Limit(1).
		Find(&rows) // `First()` is sorted by id, so `Find()` is used.

	if tx.Error != nil {
		return nil, tx.Error
	} else if len(rows) == 0 {
		return nil, ErrNotFound
	}
	return rows[0], nil
}

// Returns blocks for uncollected event logs(order by block number).
func (db *BlockDatabase) FindUncollecteds(limit int) ([]*Block, error) {
	var misc Misc
	tx := db.db.
		Where("id = ?", MISC_COLLECTED_BLOCK).
		First(&misc)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		tx = db.db.Where("log_collected IS FALSE")
	} else if tx.Error == nil {
		tx = db.db.Where("number > ?", new(big.Int).SetBytes(misc.Value).Uint64())
	}

	var rows []*Block
	tx = tx.
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
	misc := &Misc{
		ID:    MISC_COLLECTED_BLOCK,
		Value: new(big.Int).SetUint64(number).Bytes(),
	}
	tx := db.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(misc)
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

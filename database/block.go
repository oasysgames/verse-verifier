package database

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type BlockDB db

// Return the specific block.
func (db *BlockDB) Find(number uint64) (*Block, error) {
	var row Block
	tx := db.rawdb.
		Where("number = ?", number).
		First(&row)

	if err := errconv(tx.Error); err != nil {
		return nil, err
	}
	return &row, nil
}

// Return the highest block.
func (db *BlockDB) FindHighest() (*Block, error) {
	var rows []*Block
	tx := db.rawdb.
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
func (db *BlockDB) FindUncollecteds(limit int) ([]*Block, error) {
	tx := db.rawdb.
		Order("number ASC").
		Limit(limit)

	if number, err := findCollectedBlock(db.rawdb); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		tx = tx.Where("log_collected IS FALSE")
	} else {
		tx = tx.Where("number > ?", number)
	}

	var rows []*Block
	tx = tx.Find(&rows)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

// Save the new block.
func (db *BlockDB) Save(number uint64, hash common.Hash) error {
	tx := db.rawdb.Create(&Block{Number: number, Hash: hash})
	return tx.Error
}

// Save event log collected block.
func (db *BlockDB) SaveCollected(number uint64, hash common.Hash) error {
	return db.rawdb.Transaction(func(tx *gorm.DB) error {
		block, err := newDB(tx).Block.Find(number)
		if err != nil {
			return fmt.Errorf("failed to find the target block: %w", err)
		} else if block.Hash != hash {
			return fmt.Errorf("this block was removed due to reorganization, number=%d old-hash=%s new-hash=%s",
				number, hash, block.Hash)
		}

		return saveCollectedBlock(tx, number)
	})
}

// Delete blocks after the number.
func (db *BlockDB) Deletes(after uint64) error {
	return db.rawdb.Transaction(func(txdb *gorm.DB) error {
		tx := txdb.
			Where("number >= ?", after).
			Delete(&Block{})
		if tx.Error != nil {
			return tx.Error
		}

		collected, err := findCollectedBlock(txdb)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
		} else if collected >= after {
			return saveCollectedBlock(txdb, after-1)
		}

		return nil
	})
}

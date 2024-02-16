package database

import (
	"math/big"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	MISC_COLLECTED_BLOCK = "COLLECTED_BLOCK"
)

func findCollectedBlock(db *gorm.DB) (uint64, error) {
	var misc Misc
	tx := db.
		Where("id = ?", MISC_COLLECTED_BLOCK).
		First(&misc)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return new(big.Int).SetBytes(misc.Value).Uint64(), nil
}

func saveCollectedBlock(db *gorm.DB, number uint64) error {
	misc := &Misc{
		ID:    MISC_COLLECTED_BLOCK,
		Value: new(big.Int).SetUint64(number).Bytes(),
	}
	tx := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(misc)
	return tx.Error
}

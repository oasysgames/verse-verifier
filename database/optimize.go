package database

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Repair signatures with previous_id overtaking id.
func (db *OptimismDatabase) RepairOvertakingSignatures(signer common.Address) {
	var rows []*OptimismSignature
	tx := db.db.
		Joins("Signer").
		Joins("OptimismScc").
		Where("Signer.address = ?", signer).
		Where("optimism_signatures.id < optimism_signatures.previous_id").
		Find(&rows)
	if tx.Error != nil {
		log.Error("Failed to find overtaking signatures", "err", tx.Error)
	} else if tx.RowsAffected == 0 {
		return
	}

	for _, row := range rows {
		logCtx := []interface{}{
			"scc", row.OptimismScc.Address,
			"index", row.BatchIndex,
			"id", row.ID,
			"previous-id", row.PreviousID,
		}

		var prevRow *OptimismSignature
		tx := db.db.
			Joins("Signer").
			Where("Signer.address = ?", signer).
			Where("optimism_signatures.id < ?", row.ID).
			Order("optimism_signatures.id DESC").
			Limit(1).
			First(&prevRow)
		if tx.Error != nil {
			log.Error("Previous signature does not exist", append(logCtx, "err", tx.Error)...)
			continue
		}

		row.PreviousID = prevRow.ID
		if tx := db.db.Save(row); tx.Error != nil {
			log.Error("Failed to save new previous id", append(logCtx, "err", tx.Error)...)
			continue
		}

		log.Warn("Repaired overtaking signature",
			append(logCtx, "new-previous-id", row.PreviousID)...)
	}
}

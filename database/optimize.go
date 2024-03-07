package database

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func (db *OptimismSignatureDB) RepairPreviousID(signer common.Address) {
	db.repairOvertakingSignatures(signer)
	db.repairMissingPrevID(signer)
}

// Repair signatures with previous_id overtaking id.
func (db *OptimismSignatureDB) repairOvertakingSignatures(signer common.Address) {
	var rows []*OptimismSignature
	tx := db.rawdb.
		Joins("Signer").
		Joins("Contract").
		Where("Signer.address = ?", signer).
		Where("optimism_signatures.id < optimism_signatures.previous_id").
		Find(&rows)
	if tx.Error != nil {
		log.Error("Failed to find signatures", "err", tx.Error)
		return
	}

	for _, row := range rows {
		db.repairPrevID(row, "overtaking")
	}
}

func (db *OptimismSignatureDB) repairMissingPrevID(signer common.Address) {
	var rows []*OptimismSignature
	tx := db.rawdb.
		Joins("Signer").
		Joins("Contract").
		Joins("LEFT JOIN optimism_signatures AS t2 ON optimism_signatures.previous_id = t2.id").
		Where("Signer.address = ?", signer).
		Where("optimism_signatures.previous_id != ''").
		Where("t2.id IS NULL").
		Find(&rows)
	if tx.Error != nil {
		log.Error("Failed to find signatures", "err", tx.Error)
		return
	}

	for _, row := range rows {
		db.repairPrevID(row, "missing")
	}
}

func (db *OptimismSignatureDB) repairPrevID(row *OptimismSignature, reason string) {
	logCtx := []interface{}{
		"reason", reason, "signer", row.Signer.Address,
		"contract", row.Contract.Address, "index", row.RollupIndex,
		"id", row.ID, "old-previous-id", row.PreviousID,
	}

	var prevRow []*OptimismSignature
	tx := db.rawdb.
		Where("optimism_signatures.signer_id = ?", row.SignerID).
		Where("optimism_signatures.id < ?", row.ID).
		Order("optimism_signatures.id DESC").
		Limit(1).
		Find(&prevRow)

	if tx.Error == nil {
		if len(prevRow) > 0 {
			row.PreviousID = prevRow[0].ID
		} else {
			row.PreviousID = ""
		}
	} else {
		log.Error("Previous signature does not exist", append(logCtx, "err", tx.Error)...)
		return
	}

	if tx := db.rawdb.Save(row); tx.Error != nil {
		log.Error("Failed to save new previous id", append(logCtx, "err", tx.Error)...)
		return
	}

	log.Warn("Repaired previous id",
		append(logCtx, "new-previous-id", row.PreviousID)...)
}

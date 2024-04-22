package database

import (
	"bytes"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"gorm.io/gorm"
)

type OptimismSignatureDB db

func (db *OptimismSignatureDB) FindByID(id string) (*OptimismSignature, error) {
	var row OptimismSignature
	tx := db.rawdb.
		Joins("Signer").
		Joins("Contract").
		Where("optimism_signatures.id = ?", id).
		First(&row)

	if err := errconv(tx.Error); err != nil {
		return nil, err
	}
	return &row, nil
}

func (db *OptimismSignatureDB) Find(
	idAfter *string,
	signer *common.Address,
	contract *common.Address,
	index *uint64,
	limit, offset int,
) ([]*OptimismSignature, error) {
	tx := db.rawdb.
		Joins("Signer").
		Joins("Contract").
		Order("optimism_signatures.id").
		Limit(limit).
		Offset(offset)

	if idAfter != nil {
		tx = tx.Where("optimism_signatures.id >= ?", *idAfter)
	}
	if signer != nil {
		_signer, err := db.db.Signer.FindOrCreate(*signer)
		if err != nil {
			return nil, err
		}
		tx = tx.Where("optimism_signatures.signer_id = ?", _signer.ID)
	}
	if contract != nil {
		_contract, err := db.db.OPContract.FindOrCreate(*contract)
		if err != nil {
			return nil, err
		}
		tx = tx.Where("optimism_signatures.optimism_scc_id = ?", _contract.ID)
	}
	if index != nil {
		tx = tx.Where("optimism_signatures.batch_index = ?", *index)
	}

	var rows []*OptimismSignature
	tx = tx.Find(&rows)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

func (db *OptimismSignatureDB) FindLatestsPerSigners() ([]*OptimismSignature, error) {
	// search foolishly because group by is slow
	var signers []uint64
	tx := db.rawdb.Model(&Signer{}).
		Select("id").
		Find(&signers)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var rows []*OptimismSignature
	for _, signer := range signers {
		sub := db.rawdb.
			Table("optimism_signatures").
			Select("MAX(id)").
			Where("signer_id = ?", signer)

		var row OptimismSignature
		tx := db.rawdb.
			Joins("Signer").
			Joins("Contract").
			Where("optimism_signatures.id = (?)", sub).
			First(&row)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			continue
		} else if tx.Error != nil {
			return nil, tx.Error
		}
		rows = append(rows, &row)
	}
	return rows, nil
}

func (db *OptimismSignatureDB) FindLatestsBySigner(
	signer common.Address,
	limit, offset int,
) ([]*OptimismSignature, error) {
	_signer, err := db.db.Signer.FindOrCreate(signer)
	if err != nil {
		return nil, err
	}

	var rows []*OptimismSignature
	tx := db.rawdb.
		Joins("Signer").
		Joins("Contract").
		Where("optimism_signatures.signer_id = ?", _signer.ID).
		Order("optimism_signatures.id DESC").
		Limit(limit).
		Offset(offset).
		Find(&rows)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

func (db *OptimismSignatureDB) Save(
	id, previousID *string,
	signer common.Address,
	contract common.Address,
	rollupIndex uint64,
	rollupHash common.Hash,
	approved bool,
	signature Signature,
) (*OptimismSignature, error) {
	_signer, err := db.db.Signer.FindOrCreate(signer)
	if err != nil {
		return nil, err
	}

	_contract, err := db.db.OPContract.FindOrCreate(contract)
	if err != nil {
		return nil, err
	}

	values := map[string]interface{}{
		"signer_id":       _signer.ID,
		"optimism_scc_id": _contract.ID,
		"batch_index":     rollupIndex,
		"batch_root":      rollupHash,
		"approved":        approved,
		"signature":       signature,
	}

	if previousID != nil {
		values["previous_id"] = *previousID
	} else {
		values["previous_id"] = gorm.Expr(`(SELECT IFNULL(
			(SELECT MAX(t.id) FROM optimism_signatures AS t WHERE t.signer_id = ?),
			""
		))`, _signer.ID)
	}

	var created OptimismSignature
	err = db.rawdb.Transaction(func(s *gorm.DB) error {
		// Delete the same batch index signature as it may be recreated for reasons such as chain reorganization.
		if tx := s.Model(&OptimismSignature{}).
			Where("signer_id = ? AND optimism_scc_id = ?", _signer.ID, _contract.ID).
			// WARNING: Do not condition on signature comparison as this will result in a UNIQUE constraint error.
			Where("batch_index = ?", rollupIndex).
			Delete(&OptimismSignature{}); tx.Error != nil {
			return tx.Error
		}

		if id != nil {
			values["id"] = *id
		} else {
			values["id"] = util.ULID(nil).String()
		}

		if tx := s.Model(&OptimismSignature{}).Create(values); tx.Error != nil {
			return tx.Error
		}

		if tx := s.
			Joins("Signer").
			Joins("Contract").
			First(&created, "optimism_signatures.id = ?", values["id"]); tx.Error != nil {
			return tx.Error
		}

		if strings.Compare(created.ID, created.PreviousID) <= 0 {
			return errors.New("previous id is overtaking")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &created, nil
}

// Delete signatures after the specified rollup index.
func (db *OptimismSignatureDB) Deletes(
	signer common.Address,
	contract common.Address,
	rollupIndex uint64,
) (int64, error) {
	var affected int64
	err := db.rawdb.Transaction(func(s *gorm.DB) error {
		var ids []string
		tx := s.
			Model(&OptimismSignature{}).
			Joins("Signer").
			Joins("Contract").
			Where("Signer.address = ? AND Contract.address = ?", signer, contract).
			Where("optimism_signatures.batch_index >= ?", rollupIndex).
			Pluck("optimism_signatures.id", &ids)
		if tx.Error != nil {
			return tx.Error
		}

		tx = s.Where("id IN (?)", ids).Delete(&OptimismSignature{})
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

// for debug
func (db *OptimismSignatureDB) SequentialFinder(startPrevID string) func() ([]*OptimismSignature, error) {
	var prevRows []*OptimismSignature
	return func() ([]*OptimismSignature, error) {
		var prevIDs []string
		if prevRows == nil {
			prevIDs = append(prevIDs, startPrevID)
		} else {
			prevIDs = make([]string, len(prevRows))
			for i, row := range prevRows {
				prevIDs[i] = row.ID
			}
		}
		if len(prevIDs) == 0 {
			return nil, nil // reached the head
		}

		var rows []*OptimismSignature
		tx := db.rawdb.
			Joins("Signer").
			Joins("Contract").
			Where("optimism_signatures.previous_id IN (?)", prevIDs).
			Find(&rows)
		if tx.Error != nil {
			return nil, tx.Error
		}

		prevRows = rows
		return rows, nil
	}
}

// Implementation of sort.Interface
type OptimismSignatures []*OptimismSignature

func (sigs OptimismSignatures) Len() int      { return len(sigs) }
func (sigs OptimismSignatures) Swap(i, j int) { sigs[i], sigs[j] = sigs[j], sigs[i] }
func (sigs OptimismSignatures) Less(i, j int) bool {
	a, b := sigs[i], sigs[j]
	return bytes.Compare(a.Signer.Address[:], b.Signer.Address[:]) == -1
}

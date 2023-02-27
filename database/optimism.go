package database

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/scc"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OptimismDatabase struct {
	db *gorm.DB
}

func (db *OptimismDatabase) FindOrCreateSigner(signer common.Address) (*Signer, error) {
	row := &Signer{Address: signer}
	tx := db.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "address"}},
	}).Create(row)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return row, nil
}

func (db *OptimismDatabase) FindOrCreateSCC(scc common.Address) (*OptimismScc, error) {
	var row OptimismScc
	tx := db.db.Where("address = ?", scc).First(&row)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		row.Address = scc
		row.NextIndex = 0

		tx = db.db.Create(&row)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if tx.Error != nil {
		return nil, tx.Error
	}

	return &row, nil
}

func (db *OptimismDatabase) FindSCCs() ([]*OptimismScc, error) {
	var rows []*OptimismScc
	tx := db.db.Find(&rows)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

func (db *OptimismDatabase) FindState(
	scc common.Address,
	batchIndex uint64,
) (*OptimismState, error) {
	var row OptimismState
	tx := db.db.
		Joins("OptimismScc").
		Where("OptimismScc.address = ? AND batch_index = ?", scc, batchIndex).
		First(&row)

	if err := errconv(tx.Error); err != nil {
		return nil, err
	}
	return &row, nil
}

// Return events waiting verification(order by BatchIndex).
func (db *OptimismDatabase) FindVerificationWaitingStates(
	signer common.Address,
	scc common.Address,
	nextIndex uint64,
	limit int,
) ([]*OptimismState, error) {
	_scc, err := db.FindOrCreateSCC(scc)
	if err != nil {
		return nil, err
	}

	if _scc.NextIndex > nextIndex {
		nextIndex = _scc.NextIndex
	}

	var latestSig OptimismSignature
	tx := db.db.
		Joins("Signer").
		Joins("OptimismScc").
		Where("Signer.address = ? AND OptimismScc.id = ?", signer, _scc.ID).
		Order("optimism_signatures.batch_index DESC").
		First(&latestSig)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		// noop
	} else if tx.Error != nil {
		return nil, tx.Error
	} else if latestSig.BatchIndex >= nextIndex {
		nextIndex = latestSig.BatchIndex + 1
	}

	var rows []*OptimismState
	tx = db.db.
		Joins("OptimismScc").
		Where("optimism_scc_id = ? AND batch_index >= ?", _scc.ID, nextIndex).
		Order("optimism_states.batch_index ASC").
		Limit(limit).
		Find(&rows)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return rows, nil
}

func (db *OptimismDatabase) SaveNextIndex(scc common.Address, nextIndex uint64) error {
	scc_, err := db.FindOrCreateSCC(scc)
	if err != nil {
		return err
	}

	scc_.NextIndex = nextIndex
	return db.db.Save(scc_).Error
}

// Save new state batch appended event to database.
func (db *OptimismDatabase) SaveState(e *scc.SccStateBatchAppended) (*OptimismState, error) {
	row := &OptimismState{
		BatchIndex:        e.BatchIndex.Uint64(),
		BatchRoot:         e.BatchRoot,
		BatchSize:         e.BatchSize.Uint64(),
		PrevTotalElements: e.PrevTotalElements.Uint64(),
		ExtraData:         e.ExtraData,
	}

	err := db.db.Transaction(func(s *gorm.DB) error {
		scc, err := newDB(s).Optimism.FindOrCreateSCC(e.Raw.Address)
		if err != nil {
			return err
		}

		row.OptimismScc = *scc
		return s.Create(row).Error
	})
	if err != nil {
		return nil, err
	}

	return row, nil
}

func (db *OptimismDatabase) SaveSignature(
	id, previousID *string,
	signer common.Address,
	scc common.Address,
	batchIndex uint64,
	batchRoot common.Hash,
	batchSize uint64,
	prevTotalElements uint64,
	extraData []byte,
	approved bool,
	signature Signature,
) (*OptimismSignature, error) {
	_signer, err := db.FindOrCreateSigner(signer)
	if err != nil {
		return nil, err
	}

	_scc, err := db.FindOrCreateSCC(scc)
	if err != nil {
		return nil, err
	}

	values := map[string]interface{}{
		"signer_id":           _signer.ID,
		"optimism_scc_id":     _scc.ID,
		"batch_index":         batchIndex,
		"batch_root":          batchRoot,
		"batch_size":          batchSize,
		"prev_total_elements": prevTotalElements,
		"extra_data":          extraData,
		"approved":            approved,
		"signature":           signature,
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
	err = db.db.Transaction(func(s *gorm.DB) error {
		// Delete the same batch index signature as it may be recreated for reasons such as chain reorganization.
		if tx := s.Model(&OptimismSignature{}).
			Where("signer_id = ? AND optimism_scc_id = ?", _signer.ID, _scc.ID).
			// WARNING: Do not condition on signature comparison as this will result in a UNIQUE constraint error.
			Where("batch_index = ?", batchIndex).
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
			Joins("OptimismScc").
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

func (db *OptimismDatabase) FindLatestSignaturePerSigners() ([]*OptimismSignature, error) {
	sub := db.db.Model(&OptimismSignature{}).
		Select("signer_id, MAX(id) AS id").
		Group("signer_id")

	var rows []*OptimismSignature
	tx := db.db.
		Table("optimism_signatures AS t1").
		Joins("Signer").
		Joins("OptimismScc").
		Joins("INNER JOIN (?) AS t2 ON t1.signer_id = t2.signer_id AND t1.id = t2.id", sub).
		Find(&rows)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

func (db *OptimismDatabase) FindLatestSignatureBySigner(
	signer common.Address,
) (*OptimismSignature, error) {
	var row OptimismSignature
	tx := db.db.
		Joins("Signer").
		Joins("OptimismScc").
		Where("Signer.address = ?", signer).
		Order("optimism_signatures.id DESC").
		First(&row)

	if err := errconv(tx.Error); err != nil {
		return nil, err
	}
	return &row, nil
}

func (db *OptimismDatabase) FindSignatureByID(id string) (*OptimismSignature, error) {
	var row OptimismSignature
	tx := db.db.
		Joins("Signer").
		Joins("OptimismScc").
		Where("optimism_signatures.id = ?", id).
		First(&row)

	if err := errconv(tx.Error); err != nil {
		return nil, err
	}
	return &row, nil
}

func (db *OptimismDatabase) FindSignatures(
	idAfter *string,
	signer *common.Address,
	scc *common.Address,
	index *uint64,
	limit, offset int,
) ([]*OptimismSignature, error) {
	tx := db.db.
		Joins("Signer").
		Joins("OptimismScc").
		Order("optimism_signatures.id").
		Limit(limit).
		Offset(offset)

	if idAfter != nil {
		tx = tx.Where("optimism_signatures.id >= ?", *idAfter)
	}
	if signer != nil {
		tx = tx.Where("Signer.address = ?", *signer)
	}
	if scc != nil {
		tx = tx.Where("OptimismScc.address = ?", *scc)
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

// Delete states after the specified batchIndex.
func (db *OptimismDatabase) DeleteStates(scc common.Address, batchIndex uint64) (int64, error) {
	var affected int64
	err := db.db.Transaction(func(s *gorm.DB) error {
		var ids []uint64
		tx := s.
			Model(&OptimismState{}).
			Joins("OptimismScc").
			Where("OptimismScc.address = ? AND batch_index >= ?", scc, batchIndex).
			Pluck("optimism_states.id", &ids)
		if tx.Error != nil {
			return tx.Error
		}

		tx = s.Where("id IN ?", ids).Delete(&OptimismState{})
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

// Delete signatures after the specified batchIndex.
func (db *OptimismDatabase) DeleteSignatures(
	signer common.Address,
	scc common.Address,
	batchIndex uint64,
) (int64, error) {
	var affected int64
	err := db.db.Transaction(func(s *gorm.DB) error {
		var ids []string
		tx := s.
			Model(&OptimismSignature{}).
			Joins("Signer").
			Joins("OptimismScc").
			Where("Signer.address = ? AND OptimismScc.address = ?", signer, scc).
			Where("optimism_signatures.batch_index >= ?", batchIndex).
			Pluck("optimism_signatures.id", &ids)
		if tx.Error != nil {
			return tx.Error
		}

		tx = s.Where("id IN ?", ids).Delete(&OptimismSignature{})
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

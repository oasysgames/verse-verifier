package database

import (
	"errors"

	"github.com/oasysgames/oasys-optimism-verifier/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

var (
	ErrNotFound = errors.New("not found")

	models = []interface{}{
		&Block{},
		&Signer{},
		&OptimismContract{},
		&OptimismState{},
		&OpstackProposal{},
		&OptimismSignature{},
		&Misc{},
	}
)

type Database struct {
	rawdb *gorm.DB

	Block       *BlockDB
	Signer      *SignerDB
	OPContract  *OptimismContractDB
	OPSignature *OptimismSignatureDB
}

type db struct {
	rawdb *gorm.DB
	db    *Database
}

func NewDatabase(cfg *config.Database) (*Database, error) {
	config := &gorm.Config{Logger: &mylogger{
		LogLevel:            gormlog.Info,
		LongQueryTime:       cfg.LongQueryTime,
		MinExaminedRowLimit: cfg.MinExaminedRowLimit,
	}}
	db, err := gorm.Open(sqlite.Open(cfg.Path), config)
	if err != nil {
		return nil, err
	}

	// workaround for "database is locked" error
	if rawdb, err := db.DB(); err != nil {
		return nil, err
	} else {
		rawdb.SetMaxOpenConns(1)
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return nil, err
		}
	}

	return newDB(db), nil
}

func (db *Database) Transaction(fn func(txdb *Database) error) error {
	return db.rawdb.Transaction(func(rawtxdb *gorm.DB) error {
		return fn(newDB(rawtxdb))
	})
}

func newDB(rawdb *gorm.DB) *Database {
	var db Database
	db = Database{
		rawdb:       rawdb,
		Block:       &BlockDB{rawdb: rawdb, db: &db},
		Signer:      &SignerDB{rawdb: rawdb, db: &db},
		OPContract:  &OptimismContractDB{rawdb: rawdb, db: &db},
		OPSignature: &OptimismSignatureDB{rawdb: rawdb, db: &db},
	}
	return &db
}

func errconv(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	return err
}

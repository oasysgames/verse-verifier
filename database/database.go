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
		&OptimismScc{},
		&OptimismState{},
		&OptimismSignature{},
	}
)

type Database struct {
	db *gorm.DB

	Block    *BlockDatabase
	Optimism *OptimismDatabase
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

func (db *Database) Transaction(fn func(*Database) error) error {
	return db.db.Transaction(func(tx *gorm.DB) error {
		return fn(newDB(tx))
	})
}

func newDB(db *gorm.DB) *Database {
	return &Database{
		db:       db,
		Block:    &BlockDatabase{db: db},
		Optimism: &OptimismDatabase{db: db},
	}
}

func errconv(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	return err
}

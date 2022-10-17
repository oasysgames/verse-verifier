package database

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

type mylogger struct{}

func (m *mylogger) LogMode(gormlog.LogLevel) gormlog.Interface {
	return &mylogger{}
}

func (m *mylogger) Info(ctx context.Context, msg string, args ...interface{}) {
	log.Info(msg, args)
}

func (m *mylogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	log.Warn(msg, args)
}

func (m *mylogger) Error(ctx context.Context, msg string, args ...interface{}) {
	log.Error(msg, args)
}

func (m *mylogger) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (sql string, rowsAffected int64),
	err error,
) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error(
			"gorm",
			"sql",
			sql,
			"affected",
			rows,
			"time",
			float64(elapsed.Nanoseconds())/1e6,
			"err",
			err,
		)
	} else {
		log.Debug("gorm", "sql", sql, "affected", rows, "time", float64(elapsed.Nanoseconds())/1e6)
	}
}

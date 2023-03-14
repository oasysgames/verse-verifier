package database

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

type mylogger struct {
	LogLevel gormlog.LogLevel
}

func (l *mylogger) LogMode(level gormlog.LogLevel) gormlog.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l *mylogger) Info(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel >= gormlog.Info {
		log.Info(msg, args)
	}
}

func (l *mylogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel >= gormlog.Warn {
		log.Warn(msg, args)
	}
}

func (l *mylogger) Error(ctx context.Context, msg string, args ...interface{}) {
	if l.LogLevel >= gormlog.Error {
		log.Error(msg, args)
	}
}

func (l *mylogger) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (sql string, rowsAffected int64),
	err error,
) {
	if l.LogLevel <= gormlog.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error(
			"gorm",
			"sql", sql,
			"affected", rows,
			"elapsed", elapsed,
			"err", err,
		)
	} else {
		log.Debug(
			"gorm",
			"sql", sql,
			"affected", rows,
			"elapsed", elapsed,
		)
	}
}

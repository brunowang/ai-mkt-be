package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
	*log.Helper
	LogLevel logger.LogLevel
}

func NewGormLogger(lg *log.Helper) *GormLogger {
	return &GormLogger{
		Helper:   lg,
		LogLevel: logger.Info,
	}
}

func (lg *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return logger.Default.LogMode(level)
}

func (lg *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if lg.LogLevel < logger.Info {
		return
	}
	lg.WithContext(ctx).Info(append([]interface{}{msg}, data)...)
}

func (lg *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if lg.LogLevel < logger.Warn {
		return
	}
	lg.WithContext(ctx).Warn(append([]interface{}{msg}, data)...)
}

func (lg *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if lg.LogLevel < logger.Error {
		return
	}
	lg.WithContext(ctx).Error(append([]interface{}{msg}, data)...)
}

func (lg *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {

}

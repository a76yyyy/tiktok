package dlog

import (
	"context"
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type ZapGorm2Logger struct {
	ZapLogger                 *kitexzap.Logger
	LogLevel                  gormlogger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
}

func NewZapGorm2(zapLogger *kitexzap.Logger) ZapGorm2Logger {
	return ZapGorm2Logger{
		ZapLogger:                 zapLogger,
		LogLevel:                  gormlogger.Warn,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
	}
}

func (l ZapGorm2Logger) SetAsDefault() {
	gormlogger.Default = l
}

func (l ZapGorm2Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return ZapGorm2Logger{
		ZapLogger:                 l.ZapLogger,
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
}

func (l ZapGorm2Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}
	l.logger().Debugf(str, args...)
}

func (l ZapGorm2Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}
	l.logger().Warnf(str, args...)
}

func (l ZapGorm2Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}
	l.logger().Errorf(str, args...)
}

func (l ZapGorm2Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		l.logger().Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		l.logger().Warn("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.LogLevel >= gormlogger.Info:
		sql, rows := fc()
		l.logger().Debug("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	}
}

var (
	gormPackage    = filepath.Join("gorm.io", "gorm")
	zapgormPackage = filepath.Join("moul.io", "zapgorm2")
)

func (l ZapGorm2Logger) logger() *kitexzap.Logger {
	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			return l.ZapLogger
		}
	}
	return l.ZapLogger
}

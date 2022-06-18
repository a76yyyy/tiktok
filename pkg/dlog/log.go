/*
 * @Author: a76yyyy q981331502@163.com
 * @Date: 2022-06-11 22:33:13
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:48:00
 * @FilePath: /tiktok/pkg/dlog/log.go
 * @Description: 基于 klog 和 zap 封装的 Logger 及其接口
 */

// 基于 klog 和 zap 封装的 Logger 及其接口
package dlog

import (
	"context"
	"io"

	"github.com/a76yyyy/tiktok/pkg/ttviper"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.uber.org/zap"
)

var (
	logger klog.FullLogger
	config = ttviper.ConfigInit("TIKTOK_LOG", "logConfig")
)

// Init Logger config
func InitLog() *zap.Logger {
	return config.InitLogger()
}

// SetOutput sets the output of default logger. By default, it is stderr.
func SetOutput(w io.Writer) {
	logger.SetOutput(w)
}

// SetLevel sets the level of logs below which logs will not be output.
// The default log level is LevelTrace.
// Note that this method is not concurrent-safe.
func SetLevel(lv klog.Level) {
	logger.SetLevel(lv)
}

// Fatal calls the default logger's Fatal method and then os.Exit(1).
func Fatal(v ...any) {
	logger.Fatal(v...)
}

// Error calls the default logger's Error method.
func Error(v ...any) {
	logger.Error(v...)
}

// Warn calls the default logger's Warn method.
func Warn(v ...any) {
	logger.Warn(v...)
}

// Notice calls the default logger's Notice method.
func Notice(v ...any) {
	logger.Notice(v...)
}

// Info calls the default logger's Info method.
func Info(v ...any) {
	logger.Info(v...)
}

// Debug calls the default logger's Debug method.
func Debug(v ...any) {
	logger.Debug(v...)
}

// Trace calls the default logger's Trace method.
func Trace(v ...any) {
	logger.Trace(v...)
}

// Fatalf calls the default logger's Fatalf method and then os.Exit(1).
func Fatalf(format string, v ...any) {
	logger.Fatalf(format, v...)
}

// Errorf calls the default logger's Errorf method.
func Errorf(format string, v ...any) {
	logger.Errorf(format, v...)
}

// Warnf calls the default logger's Warnf method.
func Warnf(format string, v ...any) {
	logger.Warnf(format, v...)
}

// Noticef calls the default logger's Noticef method.
func Noticef(format string, v ...any) {
	logger.Noticef(format, v...)
}

// Infof calls the default logger's Infof method.
func Infof(format string, v ...any) {
	logger.Infof(format, v...)
}

// Debugf calls the default logger's Debugf method.
func Debugf(format string, v ...any) {
	logger.Debugf(format, v...)
}

// Tracef calls the default logger's Tracef method.
func Tracef(format string, v ...any) {
	logger.Tracef(format, v...)
}

// CtxFatalf calls the default logger's CtxFatalf method and then os.Exit(1).
func CtxFatalf(ctx context.Context, format string, v ...any) {
	logger.CtxFatalf(ctx, format, v...)
}

// CtxErrorf calls the default logger's CtxErrorf method.
func CtxErrorf(ctx context.Context, format string, v ...any) {
	logger.CtxErrorf(ctx, format, v...)
}

// CtxWarnf calls the default logger's CtxWarnf method.
func CtxWarnf(ctx context.Context, format string, v ...any) {
	logger.CtxWarnf(ctx, format, v...)
}

// CtxNoticef calls the default logger's CtxNoticef method.
func CtxNoticef(ctx context.Context, format string, v ...any) {
	logger.CtxNoticef(ctx, format, v...)
}

// CtxInfof calls the default logger's CtxInfof method.
func CtxInfof(ctx context.Context, format string, v ...any) {
	logger.CtxInfof(ctx, format, v...)
}

// CtxDebugf calls the default logger's CtxDebugf method.
func CtxDebugf(ctx context.Context, format string, v ...any) {
	logger.CtxDebugf(ctx, format, v...)
}

// CtxTracef calls the default logger's CtxTracef method.
func CtxTracef(ctx context.Context, format string, v ...any) {
	logger.CtxTracef(ctx, format, v...)
}

type ZapLogger struct {
	SugaredLogger
	Level klog.Level
}

/**
	Control
**/
func (ll *ZapLogger) SetOutput(w io.Writer) {
}

func (ll *ZapLogger) SetLevel(lv klog.Level) {
	ll.Level = lv
}

/**
	Logger
**/
// func (ll *ZapLogger) Fatal(v ...any) {
// 	ll.Fatal(v...)
// }

// func (ll *ZapLogger) Error(v ...any) {
// 	ll.Error(v...)
// }

// func (ll *ZapLogger) Warn(v ...any) {
// 	ll.Warn(v...)
// }

// func (ll *ZapLogger) Notice(v ...any) {
// 	ll.DPanic(v...)
// }

func (s *SugaredLogger) Notice(args ...interface{}) {
	s.log(zap.DPanicLevel, "", args, nil)
}

// func (ll *ZapLogger) Info(v ...any) {
// 	ll.Info(v...)
// }

// func (ll *ZapLogger) Debug(v ...any) {
// 	ll.Debug(v...)
// }

// func (ll *ZapLogger) Trace(v ...any) {
// 	ll.Info(v...)
// }

func (s *SugaredLogger) Trace(args ...interface{}) {
	s.log(zap.InfoLevel, "", args, nil)
}

// /**
// 	FormatLogger
// **/
// func (ll *ZapLogger) Fatalf(format string, v ...any) {
// 	ll.Fatalf(format, v...)
// }

// func (ll *ZapLogger) Errorf(format string, v ...any) {
// 	ll.Errorf(format, v...)
// }

// func (ll *ZapLogger) Warnf(format string, v ...any) {
// 	ll.Warnf(format, v...)
// }

// func (ll *ZapLogger) Noticef(format string, v ...any) {
// 	ll.DPanicf(format, v...)
// }

func (s *SugaredLogger) Noticef(template string, args ...interface{}) {
	s.log(zap.DPanicLevel, template, args, nil)
}

// func (ll *ZapLogger) Infof(format string, v ...any) {
// 	ll.Infof(format, v...)
// }

// func (ll *ZapLogger) Debugf(format string, v ...any) {
// 	ll.Debugf(format, v...)
// }

// func (ll *ZapLogger) Tracef(format string, v ...any) {
// 	ll.Infof(format, v...)
// }

func (s *SugaredLogger) Tracef(template string, args ...interface{}) {
	s.log(zap.InfoLevel, template, args, nil)
}

/**
	CtxLogger
**/
func (ll *ZapLogger) CtxFatalf(ctx context.Context, format string, v ...any) {
	ll.With("ctx", ctx).Fatalw(format, v...)
}

func (ll *ZapLogger) CtxErrorf(ctx context.Context, format string, v ...any) {
	ll.With("ctx", ctx).Errorw(format, v...)
}

func (ll *ZapLogger) CtxWarnf(ctx context.Context, format string, v ...any) {
	ll.With("ctx", ctx).Warnw(format, v...)
}

func (ll *ZapLogger) CtxNoticef(ctx context.Context, format string, v ...any) {
	ll.With("ctx", ctx).DPanicw(format, v...)
}

func (ll *ZapLogger) CtxInfof(ctx context.Context, format string, v ...any) {
	ll.With("ctx", ctx).Infow(format, v...)
}

func (ll *ZapLogger) CtxDebugf(ctx context.Context, format string, v ...any) {
	ll.With("ctx", ctx).Debugw(format, v...)
}

func (ll *ZapLogger) CtxTracef(ctx context.Context, format string, v ...any) {
	ll.With("ctx", ctx).Infow(format, v...)
}

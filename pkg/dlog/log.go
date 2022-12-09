// Copyright 2022 a76yyyy && CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	"encoding/json"
	"errors"
	"io"
	"sort"
	"time"

	"github.com/a76yyyy/tiktok/pkg/ttviper"
	"github.com/cloudwego/kitex/pkg/klog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger klog.FullLogger
	config = ttviper.ConfigInit("TIKTOK_LOG", "logConfig")
)

// build Logger Core from config
func loggerBuild(callerSkip int) (zapcore.Encoder, zapcore.WriteSyncer, zap.AtomicLevel, []zap.Option, error) {
	var cfg zap.Config
	if err := json.Unmarshal(config.ZapLogConfig(), &cfg); err != nil {
		return nil, nil, zap.AtomicLevel{}, nil, err
	}

	enc, err := newEncoder(cfg.Encoding, cfg.EncoderConfig)
	if err != nil {
		return nil, nil, zap.AtomicLevel{}, nil, err
	}

	if cfg.Level == (zap.AtomicLevel{}) {
		return nil, nil, zap.AtomicLevel{}, nil, errors.New("missing Level")
	}

	sink, closeOut, err := zap.Open(cfg.OutputPaths...)
	if err != nil {
		return nil, nil, zap.AtomicLevel{}, nil, err
	}

	errSink, _, err := zap.Open(cfg.ErrorOutputPaths...)
	if err != nil {
		closeOut()
		return nil, nil, zap.AtomicLevel{}, nil, err
	}

	opts := []zap.Option{zap.ErrorOutput(errSink)}

	if cfg.Development {
		opts = append(opts, zap.Development())
	}

	if !cfg.DisableCaller {
		opts = append(opts, zap.AddCaller())
	}

	stackLevel := zap.ErrorLevel
	if cfg.Development {
		stackLevel = zap.WarnLevel
	}
	if !cfg.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(stackLevel))
	}

	if scfg := cfg.Sampling; scfg != nil {
		opts = append(opts, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			var samplerOpts []zapcore.SamplerOption
			if scfg.Hook != nil {
				samplerOpts = append(samplerOpts, zapcore.SamplerHook(scfg.Hook))
			}
			return zapcore.NewSamplerWithOptions(
				core,
				time.Second,
				cfg.Sampling.Initial,
				cfg.Sampling.Thereafter,
				samplerOpts...,
			)
		}))
	}

	if len(cfg.InitialFields) > 0 {
		fs := make([]zap.Field, 0, len(cfg.InitialFields))
		keys := make([]string, 0, len(cfg.InitialFields))
		for k := range cfg.InitialFields {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fs = append(fs, zap.Any(k, cfg.InitialFields[k]))
		}
		opts = append(opts, zap.Fields(fs...))
	}

	opts = append(opts, zap.AddCallerSkip(callerSkip))

	return enc, sink, cfg.Level, opts, nil
}

// Init Logger config
func InitLog(callerSkip int) *kitexzap.Logger {
	enc, sink, lvl, opts, err := loggerBuild(callerSkip)
	if err != nil {
		panic(err)
	}

	logger := kitexzap.NewLogger(kitexzap.WithCoreEnc(enc), kitexzap.WithCoreWs(sink), kitexzap.WithCoreLevel(lvl), kitexzap.WithZapOptions(opts...))
	return logger
}

// Init Hertz Logger config
func InitHertzLog(callerSkip int) *hertzzap.Logger {
	enc, sink, lvl, opts, err := loggerBuild(callerSkip)
	if err != nil {
		panic(err)
	}

	logger := hertzzap.NewLogger(hertzzap.WithCoreEnc(enc), hertzzap.WithCoreWs(sink), hertzzap.WithCoreLevel(lvl), hertzzap.WithZapOptions(opts...))
	return logger
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

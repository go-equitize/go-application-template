package logger

import (
	gocontext "context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go-template/internal/pkg/config"
	"go-template/internal/pkg/constant"
)

var logger *zap.Logger

func InitLogger() error {
	commonCfg := config.Instance().Common
	if logger != nil {
		return nil
	}

	logLevel := commonCfg.LogLevel
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.Level(logLevel)),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		DisableCaller: true,
		Encoding:      "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	l, err := cfg.Build()
	if err != nil {
		return err
	}

	logger = l
	return nil
}

func fromCtx(ctx gocontext.Context) *zap.Logger {
	l := ctx.Value(constant.CtxLoggerKey)
	if l == nil {
		return logger
	}
	return l.(*zap.Logger)
}

func Error(ctx gocontext.Context, args ...interface{}) {
	l := fromCtx(ctx)
	l.Sugar().Error(args...)
}

func Errorf(ctx gocontext.Context, format string, args ...interface{}) {
	l := fromCtx(ctx)
	l.Sugar().Errorf(format, args...)
}

func Info(ctx gocontext.Context, args ...interface{}) {
	l := fromCtx(ctx)
	l.Sugar().Info(args...)
}

func Infof(ctx gocontext.Context, format string, args ...interface{}) {
	l := fromCtx(ctx)
	l.Sugar().Infof(format, args...)
}

func Warn(ctx gocontext.Context, args ...interface{}) {
	l := fromCtx(ctx)
	l.Sugar().Warn(args...)
}

func Warnf(ctx gocontext.Context, format string, args ...interface{}) {
	l := fromCtx(ctx)
	l.Sugar().Warnf(format, args...)
}

func L() *zap.Logger {
	return logger
}

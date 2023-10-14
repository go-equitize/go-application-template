package logger

import (
	"context"

	"go.uber.org/zap/zapcore"

	"go-template/internal/pkg/constant"
)

type ctxLoggerKey string

func BindLogger(ctx context.Context, fields map[string]string) context.Context {
	l := logger
	for k, v := range fields {
		f := zapcore.Field{
			Key:    k,
			Type:   zapcore.StringType,
			String: v,
		}
		l = l.With(f)
	}
	ctx = context.WithValue(ctx, ctxLoggerKey(constant.CtxLoggerKey), l)
	return ctx
}

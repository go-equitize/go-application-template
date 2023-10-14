package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap/zapcore"

	"go-template/internal/pkg/constant"
	"go-template/internal/pkg/util/logger"
)

func buildLogFields(c *gin.Context) (zapcore.Field, zapcore.Field) {
	traceIDField := zapcore.Field{
		Key:    constant.CtxTraceIDKey,
		Type:   zapcore.StringType,
		String: uuid.New().String(),
	}

	builder := strings.Builder{}
	builder.WriteString(c.Request.Method)
	builder.WriteString(" ")
	builder.WriteString(c.Request.URL.Path)
	raw := c.Request.URL.RawQuery
	if raw != "" {
		builder.WriteString("?")
		builder.WriteString(raw)
	}
	apiField := zapcore.Field{
		Key:    constant.CtxAPIRequestKey,
		Type:   zapcore.StringType,
		String: builder.String(),
	}
	return traceIDField, apiField
}

// Logger add a logger to gin context with metadata like requestID, etc.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestIDField, apiField := buildLogFields(c)
		l := logger.L().With(requestIDField).With(apiField)

		c.Set(constant.CtxLoggerKey, l)
		c.Set(constant.CtxTraceIDKey, requestIDField.String)

		// Process request
		c.Next()
	}
}

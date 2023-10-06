package middlewares

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	"go-template/internal/pkg/config"
	"go-template/internal/pkg/constant"
	"go-template/internal/pkg/utils/logger"
)

func TestBuildLogFields(t *testing.T) {
	name := "TestLoggerMiddleWare_BuildLogFields"
	t.Log(name)
	ctx := &gin.Context{}
	ctx.Request = &http.Request{}
	ctx.Request.URL = &url.URL{}
	ctx.Request.Method = "GET"
	ctx.Request.URL.Path = "/api/v1/orders"
	ctx.Request.URL.RawQuery = "chainId=1"

	requestIdField, apiField := buildLogFields(ctx)
	assert.NotNil(t, requestIdField)
	assert.EqualValues(t, constant.CtxRequestIDKey, requestIdField.Key)
	assert.EqualValues(t, zapcore.StringType, requestIdField.Type)
	assert.NotEqual(t, "", requestIdField.String)
	assert.EqualValues(t, constant.CtxAPIRequestKey, apiField.Key)
	assert.EqualValues(t, zapcore.StringType, apiField.Type)
	assert.EqualValues(t, "GET /api/v1/orders?chainId=1", apiField.String)
}

func TestLoggerMiddleWare(t *testing.T) {
	name := "TestLoggerMiddleWare"
	t.Log(name)
	ctx := &gin.Context{}
	ctx.Request = &http.Request{}
	ctx.Request.URL = &url.URL{}
	ctx.Request.Method = "GET"
	ctx.Request.URL.Path = "/api/v1/tokens"
	ctx.Request.URL.RawQuery = "chainId=1"

	err := config.Load("../configs/files/test.yaml")
	if err != nil {
		panic(err)
	}
	err = logger.InitLogger()
	if err != nil {
		panic(err)
	}

	f := Logger()
	f(ctx)
	iLogger := ctx.Value(constant.CtxLoggerKey)
	assert.NotNil(t, iLogger)
}

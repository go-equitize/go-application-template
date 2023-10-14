package route

import (
	"github.com/gin-gonic/gin"

	"go-template/internal/app/api/route/common"
	v1 "go-template/internal/app/api/route/v1"
)

func Register(engine *gin.Engine) {
	common.RegisterAPI(engine, newHealthAPI(), "health")
	common.RegisterAPIGroup(engine, v1.New(), "/api/v1")
}

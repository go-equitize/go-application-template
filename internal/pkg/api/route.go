package api

import (
	"github.com/gin-gonic/gin"

	"go-template/internal/pkg/api/v1"
)

type IAPI interface {
	SetupRoute(rg *gin.RouterGroup)
}

func AddHealthRouter(server *gin.Engine) {
	router := server.Group("health")
	AddAPI(newHealthAPI(), "", router)

	// Add API v1
	AddAPI(v1.NewAPI(), "", router)

	// Add other API
}

func AddAPI(api IAPI, path string, rg *gin.RouterGroup) {
	apiRg := rg.Group(path)
	api.SetupRoute(apiRg)
}

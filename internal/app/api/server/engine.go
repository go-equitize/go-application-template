package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-template/internal/app/api/middleware"
	"go-template/internal/pkg/config"
)

func newEngine() *gin.Engine {
	gin.SetMode(config.Instance().Http.Mode)

	healthcheck := []string{
		"/health",
	}

	newEngine := gin.New()
	newEngine.Use(gin.LoggerWithWriter(gin.DefaultWriter, healthcheck...))
	newEngine.Use(gin.Recovery())
	newEngine.Use(middleware.Logger())
	setCORS(newEngine)

	return newEngine
}

func setCORS(engine *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowMethods(http.MethodOptions)
	corsConfig.AllowAllOrigins = true
	engine.Use(cors.New(corsConfig))
}

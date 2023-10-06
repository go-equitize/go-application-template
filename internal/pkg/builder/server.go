package builder

import (
	"net/http"
	"os"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"

	"go-template/internal/pkg/config"
	"go-template/internal/pkg/middlewares"
)

func newServer() (*gin.Engine, error) {
	gin.SetMode(config.Instance().Http.Mode)

	server := gin.Default()
	server.Use(middlewares.Logger())
	server.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// just enable datadog if $DD_ENABLED is not empty
	if os.Getenv("DD_ENABLED") != "" {
		server.Use(gintrace.Middleware(os.Getenv("DD_SERVICE")))
	}

	setCORS(server)

	return server, nil
}

func setCORS(engine *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowMethods(http.MethodOptions)
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowHeaders("x-request-id")
	corsConfig.AddAllowHeaders("X-Request-Id")
	corsConfig.AddAllowHeaders("Accept-Version")
	engine.Use(cors.New(corsConfig))
}

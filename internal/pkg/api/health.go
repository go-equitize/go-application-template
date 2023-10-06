package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthAPI struct{}

func newHealthAPI() *HealthAPI {
	return &HealthAPI{}
}

func (t *HealthAPI) SetupRoute(rg *gin.RouterGroup) {
	rg.GET("/live", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, struct {
			Ok bool `json:"ok"`
		}{true})
	})

	rg.GET("/ready", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, struct {
			Ok bool `json:"ok"`
		}{true})
	})
}

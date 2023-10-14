package v1

import (
	"github.com/gin-gonic/gin"

	"go-template/internal/app/api/dto"
)

type GreetingAPI struct{}

func newGreetingAPI() *GreetingAPI {
	return &GreetingAPI{}
}

func (t *GreetingAPI) SetupRoute(rg *gin.RouterGroup) {
	rg.GET("", func(c *gin.Context) {
		dto.RespondSuccess(c, "Hello World")
	})
}

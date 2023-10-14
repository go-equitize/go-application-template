package v1

import (
	"github.com/gin-gonic/gin"

	"go-template/internal/app/api/route/common"
)

type APIv1 struct{}

func New() *APIv1 {
	return &APIv1{}
}

func (a APIv1) RegisterAPIs(rg *gin.RouterGroup) {
	common.RegisterAPI(rg, newGreetingAPI(), "/greetings")
}

package v1

import "github.com/gin-gonic/gin"

type RouteAPIV1 struct {
}

func NewAPI() *RouteAPIV1 {
	return &RouteAPIV1{}
}

func (t *RouteAPIV1) SetupRoute(rg *gin.RouterGroup) {

}

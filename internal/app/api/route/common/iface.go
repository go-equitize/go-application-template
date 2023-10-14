package common

import "github.com/gin-gonic/gin"

type IAPI interface {
	SetupRoute(*gin.RouterGroup)
}

type IAPIGroup interface {
	RegisterAPIs(*gin.RouterGroup)
}

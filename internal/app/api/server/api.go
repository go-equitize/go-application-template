package server

import (
	"github.com/gin-gonic/gin"

	"go-template/internal/app/api/route"
	"go-template/internal/pkg/config"
)

type APIServer struct {
	engine *gin.Engine
}

func NewAPIServer() *APIServer {
	engine := newEngine()
	route.Register(engine)
	return &APIServer{engine: engine}
}

func (f *APIServer) Run() error {
	return f.engine.Run(config.Instance().Http.BindAddress)
}

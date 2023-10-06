package builder

import (
	"github.com/gin-gonic/gin"

	"go-template/internal/pkg/api"
	"go-template/internal/pkg/config"
)

type apiBuilder struct {
	server *gin.Engine
}

func NewAPIBuilder() (IRunner, error) {
	server, _ := newServer()

	api.AddHealthRouter(server)

	return &apiBuilder{server: server}, nil
}

func (f *apiBuilder) Run() error {
	return f.server.Run(config.Instance().Http.BindAddress)
}

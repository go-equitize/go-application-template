package v1

import (
	"log"

	"github.com/anhvietnguyennva/go-error/pkg/errors"
	"github.com/anhvietnguyennva/go-error/pkg/transformer"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"go-template/internal/app/api/dto"
	"go-template/internal/app/api/dto/request"
	"go-template/internal/app/api/service"
)

type GreetingAPI struct {
	service service.ITemplateService
}

func newGreetingAPI() *GreetingAPI {
	return &GreetingAPI{
		service: service.TemplateServiceInstance(),
	}
}

func (t *GreetingAPI) SetupRoute(rg *gin.RouterGroup) {
	rg.GET("", t.handleGet)
	rg.POST("", t.handlePost)
}

func (t *GreetingAPI) handleGet(c *gin.Context) {
	params := request.TemplateRequest{}
	if err := c.ShouldBindQuery(&params); err != nil {
		log.Printf("failed to bind params, err: %v\n", err)
		apiErr := errors.NewRestAPIErrInvalidFormat(nil)
		dto.RespondError(c, apiErr)

		return
	}

	if err := t.service.Template(c, &params); err != nil {
		dto.RespondError(c, transformer.RestTransformerInstance().DomainErrToRestAPIErr(err))
	}

	dto.RespondSuccess(c, "GET Success")
}

func (t *GreetingAPI) handlePost(c *gin.Context) {
	payload := request.TemplateRequest{}
	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		log.Printf("failed to bind payload, err: %v\n", err)
		apiErr := errors.NewRestAPIErrInvalidFormat(nil)
		dto.RespondError(c, apiErr)

		return
	}

	if err := t.service.Template(c, &payload); err != nil {
		apiErr := transformer.RestTransformerInstance().DomainErrToRestAPIErr(err)
		dto.RespondError(c, apiErr)
	}

	dto.RespondSuccess(c, "POST Success")
}

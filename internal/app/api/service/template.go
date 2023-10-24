package service

import (
	"context"

	"github.com/anhvietnguyennva/go-error/pkg/errors"

	"go-template/internal/app/api/dto/request"
)

type ITemplateService interface {
	Template(context.Context, *request.TemplateRequest) *errors.DomainError
}

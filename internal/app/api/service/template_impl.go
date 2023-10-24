package service

import (
	"context"

	"github.com/anhvietnguyennva/go-error/pkg/errors"
	"github.com/anhvietnguyennva/go-error/pkg/transformer"

	"go-template/internal/app/api/dto/request"
	"go-template/internal/app/api/repository"
)

type TemplateService struct {
	repository repository.ITemplateRepository
}

var templateService *TemplateService

func initTemplateService(repository repository.ITemplateRepository) {
	if templateService == nil {
		templateService = &TemplateService{
			repository: repository,
		}
	}
}

func TemplateServiceInstance() ITemplateService {
	return templateService
}

func (s *TemplateService) Template(ctx context.Context, request *request.TemplateRequest) *errors.DomainError {
	if err := s.repository.Template(request.ShouldReturnSuccess); err != nil {
		return transformer.DomainTransformerInstance().InfraErrToDomainErr(err)
	}

	return nil
}

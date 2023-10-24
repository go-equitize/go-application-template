package repository

import "github.com/anhvietnguyennva/go-error/pkg/errors"

type ITemplateRepository interface {
	Template(uint) *errors.InfraError
}

package repository

import (
	"github.com/anhvietnguyennva/go-error/pkg/errors"
	"gorm.io/gorm"

	e "errors"
)

type TestRepository struct {
	db *gorm.DB
}

var testRepository *TestRepository

func initTemplateRepository(db *gorm.DB) {
	if testRepository == nil {
		testRepository = &TestRepository{
			db,
		}
	}
}

func TemplateRepositoryInstance() ITemplateRepository {
	return testRepository
}

func (r *TestRepository) Template(success uint) *errors.InfraError {
	switch success {
	case 1:
		return nil
	default:
		return errors.NewInfraErrorUnknown(e.New("unknown error"))
	}
}

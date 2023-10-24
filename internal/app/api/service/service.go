package service

import "go-template/internal/app/api/repository"

func InitServices() {
	testRepository := repository.TemplateRepositoryInstance()

	initTemplateService(testRepository)
}

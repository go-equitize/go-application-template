package component

import (
	"go-template/internal/app/api/repository"
	"go-template/internal/app/api/service"
	"go-template/internal/pkg/db"
	"go-template/internal/pkg/redis"
	"go-template/internal/pkg/util/logger"
)

func InitComponents() error {
	var err error

	err = logger.InitLogger()
	if err != nil {
		return err
	}

	err = redis.InitClient()
	if err != nil {
		return err
	}

	err = db.InitDB()
	if err != nil {
		return err
	}

	repository.InitRepos()

	service.InitServices()

	return err
}

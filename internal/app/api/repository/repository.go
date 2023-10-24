package repository

import (
	"go-template/internal/pkg/db"
	"go-template/internal/pkg/redis"
)

func InitRepos() {
	database := db.Instance()
	_ = redis.Instance()

	initTemplateRepository(database)
}

package redis

import (
	gocontext "context"
	"errors"
	"fmt"
	goerror "github.com/anhvietnguyennva/go-error/pkg/errors"
	"github.com/go-redis/redis/v8"
	"go-template/internal/pkg/config"
	"strings"
)

var redisClient redis.UniversalClient

func InitClient() error {
	if redisClient == nil {
		cfg := config.Instance().Redis
		redisAddresses := strings.Split(cfg.RedisAddresses, ",")
		if len(redisAddresses) == 0 {
			return errors.New("redis host is empty")
		}

		redisClient = redis.NewUniversalClient(&redis.UniversalOptions{
			Password:   cfg.Password,
			Addrs:      redisAddresses,
			MasterName: cfg.MasterName,
		})

		if _, err := redisClient.Ping(gocontext.Background()).Result(); err != nil {
			return goerror.NewInfraErrorRedisConnect(err)
		} else {
			fmt.Println("Successfully connect to Redis")
		}
	}
	return nil
}

func Instance() redis.UniversalClient {
	return redisClient
}

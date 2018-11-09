package services

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

type RedisService struct {
	client *redis.Client
}

func (service *RedisService) Initialize() (RedisService, error) {

	connectionString := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	databaseNumber, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if nil != err {
		return RedisService{}, err
	}

	service.client = redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: os.Getenv("REDIS_PASS"),
		DB:       databaseNumber,
	})

	return *service, nil

}

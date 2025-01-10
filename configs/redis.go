package configs

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type redisConfig struct {
	Addr     string
	Password string
	DB       string
	Protocol int
}

func LoadRedisconfig() *redisConfig {
	protocol , err := strconv.Atoi(os.Getenv("redisPortocol"))
	if err != nil{
		panic(err)
	}
	return &redisConfig{
		Addr: os.Getenv("residAddr"),
		Password: os.Getenv("redisPassword"),
		DB: os.Getenv("residDB"),
		Protocol: protocol,
	}
}

func RedisConnectio(config *redis.Options) *redis.Client {
	redisClient := redis.NewClient(config)
	return redisClient
}
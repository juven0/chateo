package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type redisConfig struct {
	Addr     string
	Password string
	DB       string
	Protocol int
}

func LoadRedisconfig() *redisConfig {

	err := godotenv.Load()
    if err != nil {
        panic(err)
    }

	// protocol , err := strconv.Atoi(os.Getenv("redisPortocol"))
	return &redisConfig{
		Addr: os.Getenv("residAddr"),
		Password: os.Getenv("redisPassword"),
		DB: os.Getenv("residDB"),
		Protocol: 2,
	}
}

func RedisConnection() *redis.Client {
	config := LoadRedisconfig()
	opt, err := redis.ParseURL(config.Addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to redis")
	return redis.NewClient(opt)
}
package config

import (
	"github.com/redis/go-redis/v9"
	"os"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS.TCP"),
		DB:   0,
	})

	return client
}

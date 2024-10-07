package config

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

func NewRedisClient() *redis.Client {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS.HOST"), os.Getenv("REDIS.PORT"))
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	return client
}

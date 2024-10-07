package main

import (
	"context"
	"github.com/joho/godotenv"
	"wheater-api/config"
)

func main() {
	app := config.NewApp()
	log := config.NewLogger()

	if err := godotenv.Load("./config.env"); err != nil {
		log.WithError(err).Panic("Error loading .env file")
	}

	rdb := config.NewRedisClient()
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	config.NewBootstrap(&config.Bootstrap{
		App:     app,
		Log:     log,
		RedisDB: rdb,
	})

	panic(app.Listen(":8080"))
}

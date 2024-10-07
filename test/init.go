package test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"wheater-api/config"
)

var app *fiber.App
var log *logrus.Logger
var rdb *redis.Client

func init() {
	app = config.NewApp()
	log = config.NewLogger()

	if err := godotenv.Load("./../config.env"); err != nil {
		log.WithError(err).Panic("Error loading .env file")
	}

	rdb = config.NewRedisClient()

	config.NewBootstrap(&config.Bootstrap{
		App:     app,
		Log:     log,
		RedisDB: rdb,
	})
}

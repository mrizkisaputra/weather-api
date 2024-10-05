package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"os"
	"wheater-api/internal/delivery/http"
	"wheater-api/internal/delivery/http/routes"
	"wheater-api/internal/repositories"
	"wheater-api/internal/usecase"
)

type Bootstrap struct {
	App     *fiber.App
	Log     *logrus.Logger
	RedisDB *redis.Client
}

func NewBootstrap(config *Bootstrap) {
	//setup repository
	weatherRepo := repositories.NewWeatherRepository(os.Getenv("WEATHER.API_KEY"), config.RedisDB)

	//setup usecase
	weatherUseCase := usecase.NewWeatherUseCase(config.Log, *weatherRepo)

	//setup controller
	weatherHandler := http.NewWeatherController(*weatherUseCase)

	//setup middleware

	//setup route
	route := routes.Route{
		App:               config.App,
		WeatherController: weatherHandler,
	}
	route.SetupRoutes()
}

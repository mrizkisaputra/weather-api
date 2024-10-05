package routes

import (
	"github.com/gofiber/fiber/v2"
	"wheater-api/internal/delivery/http"
)

type Route struct {
	App               *fiber.App
	WeatherController http.WeatherController
}

func (r *Route) SetupRoutes() {
	r.guestRoutes()
}

func (r *Route) guestRoutes() {
	r.App.Get("/api/weather/:location", r.WeatherController.GetCurrentAndForecastedWeather)
	r.App.Get("/api/weather/:location/:date1/", r.WeatherController.GetCurrentAndForecastedWeather)
	r.App.Get("/api/weather/:location/:date1/:date2", r.WeatherController.GetCurrentAndForecastedWeather)
}

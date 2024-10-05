package repositories

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var baseURL = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline"

type WeatherRepository struct {
	ApiKey string
}

func NewWeatherRepository(apiKey string) *WeatherRepository {
	return &WeatherRepository{
		ApiKey: apiKey,
	}
}

func (w *WeatherRepository) FetchWeather(location, date1 string, date2 string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s/%s?key=%s", baseURL, location, date1, date2, w.ApiKey)
	client := fiber.AcquireClient()
	defer fiber.ReleaseClient(client)

	status, responseBody, _ := client.Get(url).String()
	switch status {
	case 400:
		return "", fiber.NewError(fiber.StatusBadRequest, responseBody)
	case 429:
		return "", fiber.NewError(fiber.StatusTooManyRequests, responseBody)
	case 500:
		return "", fiber.NewError(fiber.StatusInternalServerError, responseBody)
	}

	return responseBody, nil
}

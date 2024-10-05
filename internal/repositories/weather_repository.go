package repositories

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"time"
)

var baseURL = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline"

type WeatherRepository struct {
	ApiKey  string
	RedisDB *redis.Client
}

func NewWeatherRepository(apiKey string, rdb *redis.Client) *WeatherRepository {
	return &WeatherRepository{
		ApiKey:  apiKey,
		RedisDB: rdb,
	}
}

func (w *WeatherRepository) FetchWeather(location, date1 string, date2 string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s/%s?key=%s", baseURL, location, date1, date2, w.ApiKey)
	key := fmt.Sprintf("weather:location:%s:date1:%s:date2:%s", location, date1, date2)

	weather, err := w.RedisDB.Get(context.Background(), key).Result()
	if err == nil { //check cache
		return weather, nil
	}

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

	if err := w.RedisDB.Set(context.Background(), key, responseBody, time.Hour*24).Err(); err != nil {
		return "", err
	}
	return responseBody, nil
}

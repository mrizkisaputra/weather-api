package usecase

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"wheater-api/internal/models/dto"
	. "wheater-api/internal/repositories"
)

type WeatherUseCase struct {
	WeatherRepo WeatherRepository
	Log         *logrus.Logger
}

func NewWeatherUseCase(log *logrus.Logger, repository WeatherRepository) *WeatherUseCase {
	return &WeatherUseCase{
		Log:         log,
		WeatherRepo: repository,
	}
}

func (w *WeatherUseCase) Get(request *dto.WeatherRequest) (*dto.WeatherResponse, error) {
	if request.Location == "" { //validate:required
		return nil, fiber.NewError(fiber.StatusBadRequest, "Location is required")
	}

	responseBody, err := w.WeatherRepo.FetchWeather(request.Location, request.Date1, request.Date2)
	if err != nil {
		w.Log.WithError(err).Error("error fetching weather data from API Weather")
		return nil, err
	}

	var weatherResponse dto.WeatherResponse
	if err := json.Unmarshal([]byte(responseBody), &weatherResponse); err != nil {
		w.Log.WithError(err).Error("error unmarshalling response body")
		return nil, fiber.ErrInternalServerError
	}

	return &weatherResponse, nil
}

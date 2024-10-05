package http

import (
	"github.com/gofiber/fiber/v2"
	"wheater-api/internal/models/dto"
	. "wheater-api/internal/usecase"
)

type WeatherController struct {
	WeatherUseCase WeatherUseCase
}

func NewWeatherController(usecase WeatherUseCase) WeatherController {
	return WeatherController{
		WeatherUseCase: usecase,
	}
}

func (w WeatherController) GetCurrentAndForecastedWeather(ctx *fiber.Ctx) error {
	requestParam := dto.WeatherRequest{
		Location: ctx.Params("location", ""),
		Date1:    ctx.Params("date1", ""),
		Date2:    ctx.Params("date2", ""),
	}
	response, err := w.WeatherUseCase.Get(&requestParam)
	if err != nil {
		return err
	}

	apiResponse := dto.ApiResponse[dto.WeatherResponse]{
		Status: fiber.StatusOK,
		Data:   response,
	}
	return ctx.Status(fiber.StatusOK).JSON(&apiResponse, "application/json; charset=utf-8")
}

package config

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"wheater-api/internal/models/dto"
)

func NewApp() *fiber.App {
	return fiber.New(fiber.Config{
		ErrorHandler: errorHandle,
	})
}

func errorHandle(ctx *fiber.Ctx, err error) error {
	var e *fiber.Error
	if errors.As(err, &e) {
		return ctx.Status(e.Code).JSON(dto.ApiResponseError{
			Status:  e.Code,
			Message: e.Message,
		})
	}
	return nil
}

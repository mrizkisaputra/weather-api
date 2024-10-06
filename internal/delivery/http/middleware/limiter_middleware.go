package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
	"wheater-api/internal/models/dto"
)

func RateLimiterMiddleware() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        1000,
		Expiration: time.Hour * 24,
		KeyGenerator: func(ctx *fiber.Ctx) string {
			return ctx.IP()
		},
		LimitReached: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusTooManyRequests).JSON(dto.ApiResponseError{
				Status:  fiber.StatusTooManyRequests,
				Message: "Too many requests",
			})
		},
	})
}

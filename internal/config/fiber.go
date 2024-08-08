package config

import (
	"net/http"

	"github.com/Lezonn/fin-tools-api/internal/model"
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.GetString("app.name"),
		ErrorHandler: NewErrorHandler(),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		message := fiber.ErrInternalServerError.Message

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		}

		return ctx.Status(code).JSON(model.ErrorResponse{
			Code:    code,
			Status:  http.StatusText(code),
			Message: message,
		})
	}
}

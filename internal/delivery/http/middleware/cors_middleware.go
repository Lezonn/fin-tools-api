package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func NewCors() fiber.Handler {
	corsConfig := cors.ConfigDefault
	corsConfig.AllowHeaders = "*"
	corsConfig.AllowMethods = "GET, POST, PUT, DELETE"
	corsConfig.AllowOrigins = "http://localhost:5173"
	corsConfig.AllowCredentials = true

	return cors.New(corsConfig)
}

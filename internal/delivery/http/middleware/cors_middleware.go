package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

var corsConfig = cors.Config{
	AllowOrigins: "http://localhost:5173",
	AllowHeaders: "Origin, Content-Type, Accept",
}

func NewCors() fiber.Handler {
	return cors.New(corsConfig)
}

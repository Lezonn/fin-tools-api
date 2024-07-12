package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func NewCors() fiber.Handler {
	return cors.New(cors.ConfigDefault)
}

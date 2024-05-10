package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

var logConfig = logger.Config{
	Next:          nil,
	Done:          nil,
	Format:        "[${time}] ${ip} ${status} - ${latency} ${method} ${path} ${error}\n",
	TimeFormat:    "15:04:05",
	TimeZone:      "Local",
	TimeInterval:  500 * time.Millisecond,
	Output:        os.Stdout,
	DisableColors: false,
}

func NewLogger() fiber.Handler {
	return logger.New(logConfig)
}

package exception

import "github.com/gofiber/fiber/v3"

func InternalServerError(message string) *fiber.Error {
	return &fiber.Error{Code: 500, Message: message}
}

func BadRequest(message string) *fiber.Error {
	return &fiber.Error{Code: 400, Message: message}
}

func NotFound(message string) *fiber.Error {
	return &fiber.Error{Code: 404, Message: message}
}

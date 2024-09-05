package helper

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func GetIdFromParam(ctx fiber.Ctx) (int64, error) {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 16)
	if err != nil {
		return 0, fiber.ErrBadRequest
	}

	return id, nil
}

package middleware

import (
	"github.com/Lezonn/fin-tools-api/internal/model"
	"github.com/Lezonn/fin-tools-api/internal/service"
	"github.com/gofiber/fiber/v3"
)

func NewAuth(userService *service.UserService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return fiber.ErrUnauthorized
		}

		token := authHeader[7:]
		if token == "" {
			return fiber.ErrUnauthorized
		}

		auth, err := userService.Verify(ctx.UserContext(), token)
		if err != nil {
			return err
		}

		ctx.Locals("auth", auth)

		return ctx.Next()
	}
}

func GetUser(ctx fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}

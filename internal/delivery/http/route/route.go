package route

import (
	"github.com/Lezonn/fin-tools-api/internal/delivery/http"
	"github.com/gofiber/fiber/v3"
)

type RouteConfig struct {
	App             *fiber.App
	LoginController *http.LoginController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	// Auth
	c.App.Get("/auth/google/login", c.LoginController.OAuthGoogleLogin)
	c.App.Get("/auth/google/callback", c.LoginController.OAuthGoogleCallback)
}

func (c *RouteConfig) SetupAuthRoute() {

}

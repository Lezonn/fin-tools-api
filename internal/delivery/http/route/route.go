package route

import (
	"github.com/Lezonn/fin-tools-api/internal/delivery/http"
	"github.com/Lezonn/fin-tools-api/internal/delivery/http/middleware"
	"github.com/gofiber/fiber/v3"
)

type RouteConfig struct {
	App                       *fiber.App
	TestController            *http.TestController
	LoginController           *http.UserController
	ExpenseController         *http.ExpenseController
	ExpenseCategoryController *http.ExpenseCategoryController
	AuthMiddleware            fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.App.Use(middleware.NewLogger())
	c.App.Use(middleware.NewCors())

	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Static("/static", "./public")

	c.App.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	c.App.Get("/auth/google/callback", c.LoginController.OAuthGoogleCallback)
	// c.App.Get("/auth/google/logout", c.LoginController.OAuthGoogleLogout)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Get("/test-auth-resource", c.TestController.GetMessage)
	c.App.Post("/api/expenses", c.ExpenseController.Create)
	c.App.Delete("/api/expenses/:id", c.ExpenseController.Delete)
	c.App.Patch("/api/expenses/:id", c.ExpenseController.Update)
	c.App.Get("/api/expenses", c.ExpenseController.List)
	c.App.Get("/api/expense-categories", c.ExpenseCategoryController.List)
}

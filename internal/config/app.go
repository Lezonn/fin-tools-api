package config

import (
	"github.com/Lezonn/fin-tools-api/internal/delivery/http"
	"github.com/Lezonn/fin-tools-api/internal/delivery/http/middleware"
	"github.com/Lezonn/fin-tools-api/internal/delivery/http/route"
	"github.com/Lezonn/fin-tools-api/internal/repository"
	"github.com/Lezonn/fin-tools-api/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB                *gorm.DB
	App               *fiber.App
	Log               *logrus.Logger
	Validate          *validator.Validate
	Config            *viper.Viper
	GoogleLoginConfig *oauth2.Config
}

func Bootstrap(config *BootstrapConfig) {
	// setup repository
	userRepository := repository.NewUserRepository(config.Log)
	expenseRepository := repository.NewExpenseRepository(config.Log)
	expenseCategoryRepository := repository.NewExpenseCategoryRepository(config.Log)

	// setup service
	userService := service.NewUserService(config.Config, config.DB, config.Log, config.Validate, userRepository)
	expenseService := service.NewExpenseService(config.DB, config.Log, config.Validate, expenseRepository)
	expenseCategoryService := service.NewExpenseCategoryService(config.DB, config.Log, config.Validate, expenseCategoryRepository)

	// setup controller
	loginController := http.NewUserController(config.Config, config.GoogleLoginConfig, config.Log, userService)
	expenseController := http.NewExpenseController(config.Log, expenseService)
	expenseCategoryController := http.NewExpenseCategoryController(config.Log, expenseCategoryService)
	testController := http.NewTestController(config.Config, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userService)

	// setup route
	routeConfig := route.RouteConfig{
		App:                       config.App,
		LoginController:           loginController,
		ExpenseController:         expenseController,
		ExpenseCategoryController: expenseCategoryController,
		TestController:            testController,
		AuthMiddleware:            authMiddleware,
	}

	routeConfig.Setup()
}

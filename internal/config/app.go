package config

import (
	"github.com/Lezonn/fin-tools-api/internal/delivery/http"
	"github.com/Lezonn/fin-tools-api/internal/delivery/http/route"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type BootstrapConfig struct {
	App               *fiber.App
	Log               *logrus.Logger
	Validate          *validator.Validate
	Config            *viper.Viper
	GoogleLoginConfig *oauth2.Config
}

func Bootstrap(config *BootstrapConfig) {
	// setup controller
	loginController := http.NewLoginController(config.Config, config.GoogleLoginConfig, config.Log)

	routeConfig := route.RouteConfig{
		App:             config.App,
		LoginController: loginController,
	}

	routeConfig.Setup()
}

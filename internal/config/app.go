package config

import (
	nethttp "net/http"

	"github.com/Lezonn/fin-tools-api/internal/delivery/http"
	"github.com/Lezonn/fin-tools-api/internal/delivery/http/route"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type BootstrapConfig struct {
	Log               *logrus.Logger
	Validate          *validator.Validate
	Config            *viper.Viper
	GoogleLoginConfig *oauth2.Config
	Server            *nethttp.Server
}

func Bootstrap(config *BootstrapConfig) {
	// setup controller
	loginController := http.NewLoginController(config.Config, config.GoogleLoginConfig, config.Log)

	routeConfig := route.RouteConfig{
		Server:          config.Server,
		LoginController: loginController,
	}

	routeConfig.Setup()
}

package http

import (
	"net/http"

	"github.com/Lezonn/fin-tools-api/internal/model/web"
	"github.com/Lezonn/fin-tools-api/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type UserController struct {
	Log               *logrus.Logger
	Config            *viper.Viper
	GoogleLoginConfig *oauth2.Config
	Service           *service.UserService
}

func NewUserController(config *viper.Viper, googleLoginConfig *oauth2.Config, logger *logrus.Logger) *UserController {
	return &UserController{
		GoogleLoginConfig: googleLoginConfig,
		Log:               logger,
		Config:            config,
	}
}

func (c *UserController) OAuthGoogleCallback(ctx fiber.Ctx) error {
	authCode := ctx.FormValue("code")

	jwtToken, err := c.Service.LoginWithGoogle(authCode, c.Config, c.GoogleLoginConfig)
	if err != nil {
		c.Log.Error(err.Error())
		return fiber.ErrBadRequest
	}

	ctx.JSON(web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   jwtToken,
	})
	return nil

}

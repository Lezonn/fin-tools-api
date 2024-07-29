package http

import (
	"net/http"

	"github.com/Lezonn/fin-tools-api/internal/model"
	"github.com/Lezonn/fin-tools-api/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type UserController struct {
	Config            *viper.Viper
	GoogleLoginConfig *oauth2.Config
	Log               *logrus.Logger
	Service           *service.UserService
}

func NewUserController(config *viper.Viper, googleLoginConfig *oauth2.Config, logger *logrus.Logger,
	service *service.UserService) *UserController {
	return &UserController{
		Config:            config,
		GoogleLoginConfig: googleLoginConfig,
		Log:               logger,
		Service:           service,
	}
}

func (c *UserController) OAuthGoogleCallback(ctx fiber.Ctx) error {
	authCode := ctx.FormValue("code")

	jwtToken, err := c.Service.LoginWithGoogle(
		ctx.UserContext(),
		authCode,
		c.Config,
		c.GoogleLoginConfig,
	)

	if err != nil {
		c.Log.Error(err.Error())
		return fiber.ErrBadRequest
	}

	response := model.UserResponse{
		Token: jwtToken,
	}

	ctx.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})

	return nil

}

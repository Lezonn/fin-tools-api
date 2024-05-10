package http

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Lezonn/fin-tools-api/internal/model/web"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type LoginController struct {
	GoogleLoginConfig *oauth2.Config
	Log               *logrus.Logger
	Config            *viper.Viper
}

func NewLoginController(config *viper.Viper, googleLoginConfig *oauth2.Config, logger *logrus.Logger) *LoginController {
	return &LoginController{
		GoogleLoginConfig: googleLoginConfig,
		Log:               logger,
		Config:            config,
	}
}

func (l *LoginController) OAuthGoogleCallback(ctx fiber.Ctx) error {
	// Read oauthState from Cookie
	code := ctx.FormValue("code")

	token, err := l.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		return fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	_, err = getUserDataFromGoogle(l.Config, token.AccessToken)
	if err != nil {
		logrus.Error(err.Error())
		err := ctx.Redirect().Status(fiber.StatusTemporaryRedirect).To("/")
		return err
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....

	ctx.JSON(web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   token,
	})
	return nil

}

func getUserDataFromGoogle(config *viper.Viper, accessToken string) ([]byte, error) {
	// Use code to get token and get user info from Google.
	response, err := http.Get(config.GetString("google.oauth.url_api") + accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

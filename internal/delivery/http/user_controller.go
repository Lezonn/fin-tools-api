package http

import (
	"fmt"
	"net/http"

	"github.com/Lezonn/fin-tools-api/internal/model/web"
	"github.com/Lezonn/fin-tools-api/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
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
	authCookie := ctx.Cookies("authentication")

	fmt.Println(authCookie)
	if authCookie != "" {
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(authCookie, claims, func(token *jwt.Token) (any, error) {
			return []byte(c.Config.GetString("jwt_secret")), nil
		})
		if err != nil {
			c.Log.Error(err.Error())
			return fiber.ErrBadRequest
		}
		fmt.Println(claims)
	} else {
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

		ctx.Cookie(&fiber.Cookie{
			Name:     "authentication",
			Value:    jwtToken,
			HTTPOnly: true,
			MaxAge:   60 * 60 * 24 * 30,
			SameSite: fiber.CookieSameSiteStrictMode,
			Secure:   true,
		})
	}

	ctx.JSON(web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   true,
	})

	return nil

}

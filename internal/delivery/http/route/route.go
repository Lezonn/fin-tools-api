package route

import (
	nethttp "net/http"

	"github.com/Lezonn/fin-tools-api/internal/delivery/http"
	"github.com/Lezonn/fin-tools-api/internal/exception"

	"github.com/julienschmidt/httprouter"
)

type RouteConfig struct {
	LoginController *http.LoginController
	Server          *nethttp.Server
	Router          *httprouter.Router
}

func (c *RouteConfig) Setup() {
	c.Router = httprouter.New()

	c.SetupGuestRoute()
	c.SetupAuthRoute()

	c.Router.PanicHandler = exception.ErrorHandler

	c.Server.Handler = c.Router
}

func (c *RouteConfig) SetupGuestRoute() {
	// OauthGoogle
	c.Router.GET("/auth/google/login", c.LoginController.OAuthGoogleLogin)
	c.Router.GET("/auth/google/callback", c.LoginController.OAuthGoogleCallback)
}

func (c *RouteConfig) SetupAuthRoute() {

}

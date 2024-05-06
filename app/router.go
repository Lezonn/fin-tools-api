package app

import (
	"github.com/Lezonn/fin-tools-api/controller"
	"github.com/Lezonn/fin-tools-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	loginController := initiateLoginController()

	// OauthGoogle
	router.GET("/auth/google/login", loginController.OAuthGoogleLogin)
	router.GET("/auth/google/callback", loginController.OAuthGoogleCallback)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func initiateLoginController() controller.LoginController {
	return controller.NewLoginController()
}

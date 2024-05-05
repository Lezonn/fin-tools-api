package app

import (
	"github.com/Lezonn/fin-tools-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	// OauthGoogle
	router.POST("/auth/google/login", oauthGoogleLogin)
	router.POST("/auth/google/callback", oauthGoogleCallback)

	router.PanicHandler = exception.ErrorHandler

	return router
}

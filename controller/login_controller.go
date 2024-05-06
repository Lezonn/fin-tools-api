package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LoginController interface {
	OAuthGoogleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	OAuthGoogleCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

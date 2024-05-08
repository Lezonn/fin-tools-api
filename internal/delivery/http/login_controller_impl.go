package http

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
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

func (l *LoginController) OAuthGoogleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/
	u := l.GoogleLoginConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func (l *LoginController) OAuthGoogleCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		logrus.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(l.Config, l.GoogleLoginConfig, r.FormValue("code"))
	if err != nil {
		logrus.Error(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	fmt.Fprintf(w, "UserInfo: %s\n", data)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func getUserDataFromGoogle(config *viper.Viper, googleLoginConfig *oauth2.Config, code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := googleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(config.GetString("google.oauth.url_api") + token.AccessToken)
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

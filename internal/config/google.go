package config

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewGoogleLoginConfig(config *viper.Viper) *oauth2.Config {
	// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
	return &oauth2.Config{
		RedirectURL:  config.GetString("google.oauth.redirect_url"),
		ClientID:     config.GetString("google.oauth.client_id"),
		ClientSecret: config.GetString("google.oauth.client_secret"),
		Scopes:       []string{config.GetString("google.oauth.scopes")},
		Endpoint:     google.Endpoint,
	}
}

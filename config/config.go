package config

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleLoginConfig oauth2.Config
}

var AppConfig Config

func NewConfig() error {
	if err := viperConfig(); err != nil {
		return err
	}
	googleConfig()
	return nil
}

func googleConfig() {
	// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
	AppConfig.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  viper.GetString("GOOGLE_OAUTH_REDIRECT_URL"),
		ClientID:     viper.GetString("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: viper.GetString("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{viper.GetString("GOOGLE_OAUTH_SCOPES")},
		Endpoint:     google.Endpoint,
	}
}

func viperConfig() error {
	viper.SetConfigFile("config.env")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

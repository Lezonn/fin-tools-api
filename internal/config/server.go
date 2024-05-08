package config

import (
	"net/http"

	"github.com/spf13/viper"
)

func NewServer(config *viper.Viper) *http.Server {
	domain := config.GetString("server.domain")
	port := config.GetString("server.port")

	return &http.Server{
		Addr: domain + ":" + port,
	}
}

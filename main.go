package main

import (
	"net/http"

	"github.com/Lezonn/fin-tools-api/app"
	"github.com/Lezonn/fin-tools-api/config"
	"github.com/Lezonn/fin-tools-api/helper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	err := config.NewConfig()
	helper.PanicIfError(err)

	server := http.Server{
		Addr:    viper.GetString("SERVER_DOMAIN") + ":" + viper.GetString("SERVER_PORT"),
		Handler: app.NewRouter(),
	}
	logrus.Info("Starting HTTP Server. Listening at " + server.Addr)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}

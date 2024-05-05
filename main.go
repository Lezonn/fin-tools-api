package main

import (
	"log"
	"net/http"

	"github.com/Lezonn/fin-tools-api/app"
	"github.com/Lezonn/fin-tools-api/helper"
	"github.com/spf13/viper"
)

func main() {
	err := app.SetupConfiguration()
	helper.PanicIfError(err)

	server := http.Server{
		Addr:    viper.GetString("SERVER_DOMAIN") + ":" + viper.GetString("SERVER_PORT"),
		Handler: app.NewRouter(),
	}
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}

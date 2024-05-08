package main

import (
	"github.com/Lezonn/fin-tools-api/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	validate := config.NewValidator()
	googleLoginConfig := config.NewGoogleLoginConfig(viperConfig)
	server := config.NewServer(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		Log:               log,
		Validate:          validate,
		Config:            viperConfig,
		GoogleLoginConfig: googleLoginConfig,
		Server:            server,
	})

	logrus.Info("Starting HTTP Server. Listening at " + server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

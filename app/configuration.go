package app

import (
	"github.com/spf13/viper"
)

func SetupConfiguration() error {
	viper.SetConfigFile("config.env")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

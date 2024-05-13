package main

import (
	"fmt"

	"github.com/Lezonn/fin-tools-api/internal/config"
	"github.com/gofiber/fiber/v3"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log, "dev")
	validate := config.NewValidator()
	googleLoginConfig := config.NewGoogleLoginConfig(viperConfig)
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:                db,
		App:               app,
		Log:               log,
		Validate:          validate,
		Config:            viperConfig,
		GoogleLoginConfig: googleLoginConfig,
	})

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort), fiber.ListenConfig{
		EnablePrefork: viperConfig.GetBool("web.prefork"),
	})
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

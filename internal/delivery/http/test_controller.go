package http

import (
	"fmt"
	"net/http"

	"github.com/Lezonn/fin-tools-api/internal/model"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type TestController struct {
	Log    *logrus.Logger
	Config *viper.Viper
}

func NewTestController(config *viper.Viper, logger *logrus.Logger) *TestController {
	return &TestController{
		Log:    logger,
		Config: config,
	}
}

func (c *TestController) GetMessage(ctx fiber.Ctx) error {
	message := model.TestResponse{
		Message: "Hello World!",
	}
	fmt.Println(ctx.Locals("auth"))

	ctx.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   message,
	})

	return nil
}

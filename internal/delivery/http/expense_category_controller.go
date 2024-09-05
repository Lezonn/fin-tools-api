package http

import (
	"net/http"

	"github.com/Lezonn/fin-tools-api/internal/model"
	"github.com/Lezonn/fin-tools-api/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ExpenseCategoryController struct {
	Log     *logrus.Logger
	Config  *viper.Viper
	Service *service.ExpenseCategoryService
}

func NewExpenseCategoryController(logger *logrus.Logger, service *service.ExpenseCategoryService) *ExpenseCategoryController {
	return &ExpenseCategoryController{
		Log:     logger,
		Service: service,
	}
}

func (c *ExpenseCategoryController) List(ctx fiber.Ctx) error {
	response, err := c.Service.List(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Error("failed to list expenses category")
		return err
	}

	return ctx.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}

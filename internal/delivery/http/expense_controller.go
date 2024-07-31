package http

import (
	"net/http"

	"github.com/Lezonn/fin-tools-api/internal/delivery/http/middleware"
	"github.com/Lezonn/fin-tools-api/internal/model"
	"github.com/Lezonn/fin-tools-api/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ExpenseController struct {
	Log     *logrus.Logger
	Config  *viper.Viper
	Service *service.ExpenseService
}

func NewExpenseController(logger *logrus.Logger, service *service.ExpenseService) *ExpenseController {
	return &ExpenseController{
		Log:     logger,
		Service: service,
	}
}

func (c *ExpenseController) Create(ctx fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.CreateExpenseRequest{}
	if err := ctx.Bind().Body(request); err != nil {
		c.Log.WithError(err).Error("failed to create expense")
		return fiber.ErrBadRequest
	}

	request.UserID = auth.ID

	response, err := c.Service.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create expense")
		return err
	}

	return ctx.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}

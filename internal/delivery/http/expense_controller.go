package http

import (
	"net/http"
	"strconv"

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

	if err := c.Service.Create(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to create expense")
		return err
	}

	return ctx.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   true,
	})
}

func (c *ExpenseController) Delete(ctx fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	expenseID, err := strconv.ParseInt(ctx.Params("id"), 10, 16)
	if err != nil {
		c.Log.WithError(err).Error("invalid expense ID")
		return fiber.ErrBadRequest
	}

	request := &model.DeleteExpenseRequest{}
	request.UserID = auth.ID
	request.ExpenseID = expenseID

	if err := c.Service.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete expense")
		return err
	}

	return ctx.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   true,
	})
}

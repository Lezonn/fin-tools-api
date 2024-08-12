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

	request := &model.CreateExpenseRequest{
		UserID: auth.ID,
	}

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.WithError(err).Error("failed to create expense")
		return fiber.ErrBadRequest
	}

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

	expenseID, err := getIdFromParam(ctx)
	if err != nil {
		c.Log.WithError(err).Error("failed to parse expense ID")
		return fiber.ErrBadRequest
	}

	request := &model.DeleteExpenseRequest{
		UserID:    auth.ID,
		ExpenseID: expenseID,
	}

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

func (c *ExpenseController) Update(ctx fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	expenseID, err := getIdFromParam(ctx)
	if err != nil {
		c.Log.WithError(err).Error("failed to parse expense ID")
		return err
	}

	request := &model.UpdateExpenseRequest{
		UserID:    auth.ID,
		ExpenseID: expenseID,
	}

	if err := ctx.Bind().Body(request); err != nil {
		c.Log.WithError(err).Error("failed to create expense")
		return fiber.ErrBadRequest
	}

	if err := c.Service.Update(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to update expense")
		return err
	}

	return ctx.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   true,
	})
}

func (c *ExpenseController) List(ctx fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.ListExpenseRequest{
		UserID: auth.ID,
	}

	response, err := c.Service.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list expenses")
		return err
	}

	return ctx.JSON(model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}

func getIdFromParam(ctx fiber.Ctx) (int64, error) {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 16)
	if err != nil {
		return 0, fiber.ErrBadRequest
	}

	return id, nil
}

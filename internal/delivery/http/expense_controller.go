package http

import (
	"github.com/Lezonn/fin-tools-api/internal/service"
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

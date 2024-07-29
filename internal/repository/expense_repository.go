package repository

import (
	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/sirupsen/logrus"
)

type ExpenseRepository struct {
	Repository[entity.Expense]
	Log *logrus.Logger
}

func NewExpenseRepository(log *logrus.Logger) *ExpenseRepository {
	return &ExpenseRepository{
		Log: log,
	}
}

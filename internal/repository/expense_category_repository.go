package repository

import (
	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ExpenseCategoryRepository struct {
	Repository[entity.Expense]
	Log *logrus.Logger
}

func NewExpenseCategoryRepository(log *logrus.Logger) *ExpenseCategoryRepository {
	return &ExpenseCategoryRepository{
		Log: log,
	}
}

func (r *ExpenseCategoryRepository) GetAll(db *gorm.DB, entities *[]entity.ExpenseCategory) error {
	return db.Find(entities).Error
}

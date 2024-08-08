package repository

import (
	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (r *ExpenseRepository) FindByIdAndUserId(db *gorm.DB, entity *entity.Expense, id int64, userId int64) error {
	return db.Where("id = ? AND user_id = ?", id, userId).Take(entity).Error
}

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

func (r *ExpenseRepository) FindByIdAndUserId(db *gorm.DB, entity *entity.Expense) error {
	return db.Where("user_id = ?", entity.UserID).Take(entity).Error
}

func (r *ExpenseRepository) GetListByUserId(db *gorm.DB, userId int64) ([]entity.Expense, error) {
	var expenses []entity.Expense
	if err := db.Where("user_id", userId).Find(&expenses).Error; err != nil {
		return nil, err
	}

	return expenses, nil
}

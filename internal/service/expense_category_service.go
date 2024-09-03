package service

import (
	"context"

	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/Lezonn/fin-tools-api/internal/model"
	"github.com/Lezonn/fin-tools-api/internal/model/converter"
	"github.com/Lezonn/fin-tools-api/internal/model/exception"
	"github.com/Lezonn/fin-tools-api/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ExpenseCategoryService struct {
	DB                        *gorm.DB
	Log                       *logrus.Logger
	Validate                  *validator.Validate
	ExpenseCategoryRepository *repository.ExpenseCategoryRepository
}

func NewExpenseCategoryService(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	expenseCategoryRepository *repository.ExpenseCategoryRepository) *ExpenseCategoryService {
	return &ExpenseCategoryService{
		DB:                        db,
		Log:                       logger,
		Validate:                  validate,
		ExpenseCategoryRepository: expenseCategoryRepository,
	}
}

func (s *ExpenseCategoryService) List(ctx context.Context) ([]model.ExpenseCategoryResponse, error) {
	expenseCategories := []entity.ExpenseCategory{}
	if err := s.ExpenseCategoryRepository.GetAll(s.DB, &expenseCategories); err != nil {
		s.Log.WithError(err).Error("failed to get expenses category")
		return nil, exception.InternalServerError("failed to get expenses category")
	}

	response := make([]model.ExpenseCategoryResponse, len(expenseCategories))
	for i, expenseCategory := range expenseCategories {
		response[i] = *converter.ExpenseCategoryToResponse(&expenseCategory)
	}

	return response, nil
}

package service

import (
	"context"

	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/Lezonn/fin-tools-api/internal/model"
	"github.com/Lezonn/fin-tools-api/internal/model/converter"
	"github.com/Lezonn/fin-tools-api/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ExpenseService struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	ExpenseRepository *repository.ExpenseRepository
}

func NewExpenseService(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	expenseRepository *repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		DB:                db,
		Log:               logger,
		Validate:          validate,
		ExpenseRepository: expenseRepository,
	}
}

func (s *ExpenseService) Create(ctx context.Context, request *model.CreateExpenseRequest) (*model.ExpenseResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	expense := &entity.Expense{
		UserID:            request.UserID,
		ExpenseCategoryID: request.ExpenseCategoryID,
		Amount:            request.Amount,
		Note:              request.Note,
		ExpenseDate:       request.ExpenseDate,
	}

	if err := s.ExpenseRepository.Create(tx, expense); err != nil {
		s.Log.WithError(err).Error("failed to create expense")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ExpenseToResponse(expense), nil
}

package service

import (
	"context"

	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/Lezonn/fin-tools-api/internal/model"
	"github.com/Lezonn/fin-tools-api/internal/model/exception"
	"github.com/Lezonn/fin-tools-api/internal/repository"
	"github.com/go-playground/validator/v10"
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

func (s *ExpenseService) Create(ctx context.Context, request *model.CreateExpenseRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validate request body")
		return exception.BadRequest("failed to validate request body")
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
		return exception.BadRequest("failed to create expense")
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return exception.InternalServerError("failed to commit transaction")
	}

	return nil
}

func (s *ExpenseService) Delete(ctx context.Context, request *model.DeleteExpenseRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	expense := &entity.Expense{}

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validate request")
		return exception.BadRequest("failed to validate request")
	}

	if err := s.ExpenseRepository.FindByIdAndUserId(tx, expense, request.ExpenseID, request.UserID); err != nil {
		s.Log.WithError(err).Error("expense not found")
		return exception.NotFound("expense not found")
	}

	if err := s.ExpenseRepository.Delete(tx, expense); err != nil {
		s.Log.WithError(err).Error("failed to delete expense")
		return exception.InternalServerError("failed to delete expense")
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return exception.InternalServerError("failed to commit transaction")
	}

	return nil
}

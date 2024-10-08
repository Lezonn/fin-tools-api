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

	expense := &entity.Expense{
		ID:     request.ExpenseID,
		UserID: request.UserID,
	}

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validate request")
		return exception.BadRequest("failed to validate request")
	}

	if err := s.ExpenseRepository.FindByIdAndUserId(tx, expense); err != nil {
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

func (s *ExpenseService) Update(ctx context.Context, request *model.UpdateExpenseRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	expense := &entity.Expense{
		ID:     request.ExpenseID,
		UserID: request.UserID,
	}

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validate request")
		return exception.BadRequest("failed to validate request")
	}

	if err := s.ExpenseRepository.FindByIdAndUserId(tx, expense); err != nil {
		s.Log.WithError(err).Error("expense not found")
		return exception.NotFound("expense not found")
	}

	expense.ExpenseCategoryID = request.ExpenseCategoryID
	expense.Amount = request.Amount
	expense.Note = request.Note
	expense.ExpenseDate = request.ExpenseDate

	if err := s.ExpenseRepository.Update(tx, expense); err != nil {
		s.Log.WithError(err).Error("failed to update expense")
		return exception.InternalServerError("failed to update expense")
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return exception.InternalServerError("failed to commit transaction")
	}

	return nil
}

func (s *ExpenseService) List(ctx context.Context, request *model.ListExpenseRequest) ([]model.ExpenseResponse, error) {
	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validate request")
		return nil, exception.BadRequest("failed to validate request")
	}

	expenses, err := s.ExpenseRepository.GetListByUserId(s.DB, request.UserID)
	if err != nil {
		s.Log.WithError(err).Error("failed to get expenses")
		return nil, exception.InternalServerError("failed to get expenses")
	}

	response := make([]model.ExpenseResponse, len(expenses))
	for i, expense := range expenses {
		response[i] = *converter.ExpenseToResponse(&expense)
	}

	return response, nil
}

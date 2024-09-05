package converter

import (
	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/Lezonn/fin-tools-api/internal/model"
)

func ExpenseToResponse(expense *entity.Expense) *model.ExpenseResponse {
	return &model.ExpenseResponse{
		ID:                  expense.ID,
		ExpenseCategoryName: expense.ExpenseCategory.ExpenseCategoryName,
		Amount:              expense.Amount,
		Note:                expense.Note,
		ExpenseDate:         expense.ExpenseDate,
	}
}

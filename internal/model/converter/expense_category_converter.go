package converter

import (
	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/Lezonn/fin-tools-api/internal/model"
)

func ExpenseCategoryToResponse(expenseCategory *entity.ExpenseCategory) *model.ExpenseCategoryResponse {
	return &model.ExpenseCategoryResponse{
		ID:                  expenseCategory.ID,
		ExpenseCategoryName: expenseCategory.ExpenseCategoryName,
	}
}

package model

type ExpenseResponse struct {
	ID                  int64  `json:"id"`
	ExpenseCategoryName string `json:"expense_category_name"`
	Amount              int64  `json:"amount"`
	Note                string `json:"note"`
	ExpenseDate         int64  `json:"expense_date"`
}

type CreateExpenseRequest struct {
	UserID            int64  `json:"user_id" validate:"required"`
	ExpenseCategoryID int64  `json:"expense_category_id" validate:"required"`
	Amount            int64  `json:"amount" validate:"required,min=0"`
	Note              string `json:"note"`
	ExpenseDate       int64  `json:"expense_date" validate:"required"`
}

type DeleteExpenseRequest struct {
	ExpenseID int64 `json:"expense_id" validate:"required"`
	UserID    int64 `json:"user_id" validate:"required"`
}

type UpdateExpenseRequest struct {
	ExpenseID         int64  `json:"expense_id" validate:"required"`
	UserID            int64  `json:"user_id" validate:"required"`
	ExpenseCategoryID int64  `json:"expense_category_id"`
	Amount            int64  `json:"amount"`
	Note              string `json:"note"`
	ExpenseDate       int64  `json:"expense_date"`
}

type ListExpenseRequest struct {
	UserID int64 `json:"user_id" validate:"required"`
}

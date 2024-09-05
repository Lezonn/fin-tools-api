package entity

type Expense struct {
	ID                int64           `gorm:"primaryKey;column:id"`
	UserID            int64           `gorm:"column:user_id"`
	ExpenseCategoryID int64           `gorm:"column:expense_category_id"`
	Amount            int64           `gorm:"column:amount"`
	Note              string          `gorm:"column:note"`
	ExpenseDate       int64           `gorm:"column:expense_date"`
	CreatedAt         int64           `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt         int64           `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	ExpenseCategory   ExpenseCategory `gorm:"foreignKey:ExpenseCategoryID;references:ID"`
}

func (e *Expense) TableName() string {
	return "expenses"
}

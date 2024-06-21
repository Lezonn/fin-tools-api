package entity

type ExpenseCategory struct {
	ID                  int64  `gorm:"primaryKey;column:id"`
	ExpenseCategoryName string `gorm:"column:expense_category_name"`
	CreatedAt           int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt           int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (e *ExpenseCategory) TableName() string {
	return "expense_categories"
}

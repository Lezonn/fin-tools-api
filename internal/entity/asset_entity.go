package entity

type Asset struct {
	ID              int64 `gorm:"primaryKey;column:id"`
	UserID          int64 `gorm:"column:user_id"`
	AssetCategoryID int64 `gorm:"column:asset_category_id"`
	Amount          int64 `gorm:"column:amount"`
	Month           int   `gorm:"column:month"`
	Year            int   `gorm:"column:year"`
	CreatedAt       int64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt       int64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (a *Asset) TableName() string {
	return "assets"
}

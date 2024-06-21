package entity

type AssetCategory struct {
	ID                int64  `gorm:"primaryKey;column:id"`
	AssetCategoryName string `gorm:"column:asset_category_name"`
	CreatedAt         int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt         int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (a *AssetCategory) TableName() string {
	return "asset_categories"
}

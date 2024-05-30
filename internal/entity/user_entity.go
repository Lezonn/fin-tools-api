package entity

type User struct {
	ID        int64  `gorm:"primaryKey;column:id"`
	Email     string `gorm:"column:email"`
	Name      string `gorm:"column:name"`
	Password  string `gorm:"column:password"`
	GoogleID  string `gorm:"column:google_id"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}

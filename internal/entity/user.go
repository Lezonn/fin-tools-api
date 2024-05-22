package entity

import (
	"time"
)

type User struct {
	ID        int64           `gorm:"primaryKey;column:id"`
	Email     string          `gorm:"column:email"`
	Name      string          `gorm:"column:name"`
	Password  string          `gorm:"column:password"`
	CreatedAt time.Time       `gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time       `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	OauthInfo []UserOauthInfo `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}

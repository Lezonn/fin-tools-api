package entity

import (
	"time"
)

type UserOauthInfo struct {
	ID            int64     `gorm:"primaryKey;column:id"`
	UserId        int64     `gorm:"column:user_id"`
	OauthProvider string    `gorm:"column:oauth_provider"`
	AccessToken   string    `gorm:"column:access_token"`
	RefreshToken  string    `gorm:"column:refresh_token"`
	ExpiryDate    time.Time `gorm:"column:expiry_date"`
	User          User      `gorm:"foreignKey:user_id;references:id"`
}

func (u *UserOauthInfo) TableName() string {
	return "user_oauth_info"
}

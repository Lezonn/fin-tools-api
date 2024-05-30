package repository

import (
	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (u *UserRepository) GetByGoogleID(db *gorm.DB, user *entity.User, googleID string) error {
	return db.Where("google_id = ?", googleID).First(user).Error
}

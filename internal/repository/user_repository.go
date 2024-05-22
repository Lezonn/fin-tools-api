package repository

import (
	"context"
	"database/sql"

	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/sirupsen/logrus"
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

func (u *UserRepository) GetOrCreate(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {

	return user
}

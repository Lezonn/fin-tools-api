package helper

import (
	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/Lezonn/fin-tools-api/internal/model"
)

func ToUserResponse(userEntity *entity.User) model.UserResponse {
	return model.UserResponse{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}
}

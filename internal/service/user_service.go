package service

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Lezonn/fin-tools-api/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type UserService struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserService(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	userRepository *repository.UserRepository) *UserService {
	return &UserService{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

func (s *UserService) LoginWithGoogle(authCode string, config *viper.Viper, googleConfig *oauth2.Config) (any, error) {
	token, err := googleConfig.Exchange(context.Background(), authCode)
	if err != nil {
		s.Log.Warnf("code exchange failed: %s", err.Error())
		return "", fiber.ErrInternalServerError
	}

	userData, err := s.getUserDataFromGoogle(config, token.AccessToken)
	if err != nil {
		s.Log.Warn("get user data from google failed")
		return "", fiber.ErrInternalServerError
	}

	fmt.Println(string(userData))

	return token, nil
}

func (s *UserService) getUserDataFromGoogle(config *viper.Viper, accessToken string) ([]byte, error) {
	// Use code to get token and get user info from Google.
	response, err := http.Get(config.GetString("google.oauth.url_api") + accessToken)
	if err != nil {
		s.Log.Warnf("failed getting user info: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		s.Log.Warnf("failed read response: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	return contents, nil
}

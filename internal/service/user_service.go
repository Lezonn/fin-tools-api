package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Lezonn/fin-tools-api/internal/entity"
	"github.com/Lezonn/fin-tools-api/internal/model"
	"github.com/Lezonn/fin-tools-api/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type UserService struct {
	Config         *viper.Viper
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserService(config *viper.Viper, db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	userRepository *repository.UserRepository) *UserService {
	return &UserService{
		Config:         config,
		DB:             db,
		Log:            logger,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

func (s *UserService) Verify(ctx context.Context, token string) (*model.Auth, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return []byte(s.Config.GetString("jwt_secret")), nil
	})
	if err != nil {
		s.Log.Error(err.Error())
		return nil, fiber.ErrUnauthorized
	}

	userID := int64(claims["sub"].(float64))
	return &model.Auth{ID: userID}, nil
}

func (s *UserService) LoginWithGoogle(ctx context.Context, authCode string, config *viper.Viper, googleConfig *oauth2.Config) (string, error) {
	token, err := googleConfig.Exchange(ctx, authCode)
	if err != nil {
		s.Log.WithError(err).Error("code exchange failed")
		return "", fiber.ErrInternalServerError
	}

	userInfo, err := s.getUserDataFromGoogle(ctx, config, token.AccessToken)
	if err != nil {
		s.Log.WithError(err).Error("get user data from google failed")
		return "", fiber.ErrInternalServerError
	}

	user, err := s.registerUser(&userInfo)
	if err != nil {
		s.Log.WithError(err).Error("register user failed")
		return "", fiber.ErrInternalServerError
	}

	jwtToken, err := s.generateJWT(user, config)
	if err != nil {
		s.Log.WithError(err).Error("failed to generate jwt")
		return "", fiber.ErrInternalServerError
	}

	return jwtToken, nil
}

func (s *UserService) getUserDataFromGoogle(ctx context.Context, config *viper.Viper, accessToken string) (model.GoogleUserInfo, error) {
	googleOauthUrlApi := config.GetString("google.oauth.url_api")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, googleOauthUrlApi+accessToken, nil)
	if err != nil {
		s.Log.WithError(err).Error("failed creating request")
		return model.GoogleUserInfo{}, fiber.ErrInternalServerError
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		s.Log.WithError(err).Error("failed getting user info")
		return model.GoogleUserInfo{}, fiber.ErrInternalServerError
	}
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		s.Log.WithError(err).Error("failed read response")
		return model.GoogleUserInfo{}, fiber.ErrInternalServerError
	}

	var user model.GoogleUserInfo
	err = json.Unmarshal(contents, &user)
	if err != nil {
		s.Log.WithError(err).Error("failed to parse user data")
		return model.GoogleUserInfo{}, err
	}

	return user, nil
}

func (s *UserService) registerUser(userInfo *model.GoogleUserInfo) (*entity.User, error) {
	user := &entity.User{}

	err := s.UserRepository.GetByGoogleID(s.DB, user, userInfo.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Email = userInfo.Email
		user.Name = userInfo.Name
		user.GoogleID = userInfo.ID

		err = s.UserRepository.Repository.Create(s.DB, user)
		if err != nil {
			s.Log.WithError(err).Error("failed to create user")
			return nil, err
		}
	} else if err != nil {
		s.Log.WithError(err).Error("failed to get user by google id")
		return nil, err
	}

	return user, nil
}

func (s *UserService) generateJWT(user *entity.User, config *viper.Viper) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"iss": "fin-tools",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := config.GetString("jwt_secret")
	jwtToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		s.Log.WithError(err).Error("failed to sign token")
		return "", err
	}

	return jwtToken, nil
}

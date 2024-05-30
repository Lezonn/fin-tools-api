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

// LoginWithGoogle is a method that handles user login with Google OAuth2.
// It exchanges the authorization code for an access token, retrieves user data from Google,
// and logs any errors that occur during the process.
//
// Parameters:
//   - ctx: A context.Context object that provides a cancellation signal and deadline for the request.
//   - authCode: A string representing the authorization code received from Google.
//   - config: A viper.Viper object containing configuration settings for the application.
//   - googleConfig: An oauth2.Config object representing the Google OAuth2 configuration.
//
// Returns:
//   - An string representing the JSON web token if the login is successful.
//   - An error if any error occurs during the login process.
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

	if err = s.RegisterUser(&userInfo); err != nil {
		s.Log.WithError(err).Error("register user failed")
		return "", fiber.ErrInternalServerError
	}

	jwtToken, err := s.generateJWT(&userInfo, config)
	if err != nil {
		s.Log.WithError(err).Error("failed to generate jwt")
		return "", fiber.ErrInternalServerError
	}

	return jwtToken, nil
}

// getUserDataFromGoogle retrieves user data from Google using the provided access token.
// It sends a GET request to the Google OAuth2 API endpoint with the access token,
// reads the response, and returns the user data as a byte slice.
//
// Parameters:
//   - ctx context.Context: A context.Context object that provides a cancellation signal and deadline for the request.
//   - config *viper.Viper: A viper.Viper object containing configuration settings for the application.
//     The "google.oauth.url_api" key should be present in the configuration.
//   - accessToken string: A string representing the access token received from Google.
//
// Returns:
//   - []byte: A byte slice containing the user data from Google.
//   - error: An error if any error occurs during the request or response handling.
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

// generateJWT generates a JSON Web Token (JWT) using the provided user data and configuration settings.
// It creates JWT claims, signs the token with the provided secret key, and returns the JWT token as a string.
//
// Parameters:
//   - user map[string]any: A map containing the user data. The map should have keys "id", "name", and "email".
//   - config *viper.Viper: A viper.Viper object containing configuration settings for the application.
//     The "jwt_secret" key should be present in the configuration.
//
// Returns:
// - A string representing the JSON Web Token (JWT).
// - An error if any error occurs during the process, such as signing the token.
func (s *UserService) generateJWT(user *model.GoogleUserInfo, config *viper.Viper) (string, error) {
	claims := jwt.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"iss":   "fin-tools",
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
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

func (s *UserService) RegisterUser(userInfo *model.GoogleUserInfo) error {
	user := &entity.User{}

	err := s.UserRepository.GetByGoogleID(s.DB, user, userInfo.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Email = userInfo.Email
		user.Name = userInfo.Name
		user.GoogleID = userInfo.ID

		err = s.UserRepository.Repository.Create(s.DB, user)
		if err != nil {
			s.Log.WithError(err).Error("failed to create user")
			return err
		}
	} else if err != nil {
		s.Log.WithError(err).Error("failed to get user by google id")
		return err
	}

	return nil
}

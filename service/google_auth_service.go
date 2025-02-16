package services

import (
	"context"
	"city_barber.com/configs"
	"city_barber.com/internal/helpers"
	"city_barber.com/internal/models"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type GoogleAuthService struct {
	config *oauth2.Config
	db     *gorm.DB
}

func NewGoogleAuthService(db *gorm.DB) *GoogleAuthService {
	config := &oauth2.Config{
		ClientID:     configs.LoadConfig().GoogleClientID,
		ClientSecret: configs.LoadConfig().GoogleClientSecret,
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	return &GoogleAuthService{config: config, db: db}
}

func (gas *GoogleAuthService) GoogleLogin(code string) (string, error) {
	token, err := gas.config.Exchange(context.Background(), code)
	if err != nil {
		return "", errors.New("failed to exchange code for token")
	}

	oauth2Service, err := oauth2.NewService(context.Background(), option.WithTokenSource(gas.config.TokenSource(context.Background(), token)))
	if err != nil {
		return "", errors.New("failed to create OAuth2 service")
	}

	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		return "", errors.New("failed to get user info")
	}

	// Check if user exists, otherwise create a new user
	var user models.User
	if err := gas.db.Where("email = ?", userInfo.Email).First(&user).Error; err != nil {
		user = models.User{
			Email:        userInfo.Email,
			FirstName:    userInfo.GivenName,
			LastName:     userInfo.FamilyName,
			PasswordHash: "", // No password for Google users
		}
		if err := gas.db.Create(&user).Error; err != nil {
			return "", errors.New("failed to create user")
		}
	}

	// Generate JWT token
	tokenString, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}
package services

import (
	"errors"
	"github.com/abik1221/city_barber.com/internal/helpers"
	"github.com/abik1221/city_barber.com/internal/models"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (as *AuthService) Login(email, password string) (string, error) {
	var user models.User
	if err := as.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := helpers.ComparePassword (user.PasswordHash, password); err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (as *AuthService) ForgotPassword(email, phone string) error {
	var user models.User
	if err := as.db.Where("email = ? OR phone_number = ?", email, phone).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	tempPassword := helpers.GenerateTempPassword()

	hashedPassword, err := helpers.HashPassword(tempPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}


	user.PasswordHash = hashedPassword
	if err := as.db.Save(&user).Error; err != nil {
		return errors.New("failed to update password")
	}


	if email != "" {
		helpers.SendEmail(email, "Password Reset", "Your temporary password is: "+tempPassword)
	} else if phone != "" {
		helpers.SendSMS(phone, "Your temporary password is: "+tempPassword)
	}

	return nil
}

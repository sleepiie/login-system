package usecase

import (
	"time"

	"github.com/sleepiie/login-system/domain"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Login(username, password string) (string, error)
}

type authService struct {
	userRepo  domain.UserRepository
	secretKey string
}

func NewAuthService(repo domain.UserRepository, secret string) AuthUseCase {
	return &authService{
		userRepo:  repo,
		secretKey: secret,
	}
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", domain.ErrInternalServer
	}

	return tokenString, nil
}

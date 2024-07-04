package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/zmskv/sales-app/internal/model"
	"github.com/zmskv/sales-app/internal/repository"
)

const (
	signingKey = "dajdaji213&$31d1^5g"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}
type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) CreateUser(user model.User) (string, error) {
	user.Password = generatePasswordHash(user.Password)
	user.Id = uuid.New().String()
	return s.repos.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repos.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		username,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hashed_password := sha1.New()
	hashed_password.Write([]byte(password))

	return fmt.Sprintf("%x", hashed_password.Sum(nil))
}

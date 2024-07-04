package service

import (
	"crypto/sha1"
	"errors"
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

func (s *AuthService) ParseToken(accesstoken string) (string, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hashed_password := sha1.New()
	hashed_password.Write([]byte(password))

	return fmt.Sprintf("%x", hashed_password.Sum(nil))
}

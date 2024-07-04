package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/google/uuid"

	"github.com/zmskv/sales-app/internal/model"
	"github.com/zmskv/sales-app/internal/repository"
)

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

func generatePasswordHash(password string) string {
	hashed_password := sha1.New()
	hashed_password.Write([]byte(password))

	return fmt.Sprintf("%x", hashed_password.Sum(nil))
}

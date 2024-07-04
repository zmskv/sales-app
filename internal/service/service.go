package service

import (
	"github.com/zmskv/sales-app/internal/model"
	"github.com/zmskv/sales-app/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
}

type SalesList interface {
}

type Service struct {
	Authorization
	SalesList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

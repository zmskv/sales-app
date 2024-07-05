package service

import (
	"github.com/zmskv/sales-app/internal/model"
	"github.com/zmskv/sales-app/internal/repository"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, string, error)
}

type SalesList interface {
	CreateRecord(record model.Product) (int, error)
}

type Service struct {
	Authorization
	SalesList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		SalesList:     NewSalesService(repos.SalesList),
	}
}

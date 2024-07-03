package service

import "github.com/zmskv/sales-app/internal/repository"

type Authorization interface {
}

type SalesList interface {
}

type Service struct {
	Authorization
	SalesList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

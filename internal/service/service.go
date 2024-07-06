package service

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/zmskv/sales-app/internal/model"
	"github.com/zmskv/sales-app/internal/repository"
)

type User interface {
	CreateUser(user model.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, string, error)
}

type SalesList interface {
	CreateRecord(record model.Product) (int, error)
	GetRecord(id string) (model.Product, error)
	DeleteRecord(id string) (string, error)
	GetAllRecords() ([]model.Product, error)
	ExportToPDF(productsWithIndex []ProductWithIndex) (*gofpdf.Fpdf, error)
}

type Service struct {
	User
	SalesList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:      NewAuthService(repos.User),
		SalesList: NewSalesService(repos.SalesList),
	}
}

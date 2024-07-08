package service

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/zmskv/sales-app/internal/model"
	"github.com/zmskv/sales-app/internal/repository"
)

type User interface {
	CreateUser(user model.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, string, string, error)
	GetUserInfo(id string) (model.User, error)
	UpdateUserInfo(user model.User) (string, error)
	DeleteUser(id string) (string, error)
}

type SalesList interface {
	CreateRecord(record model.Product) (int, error)
	GetRecord(id string) (model.Product, error)
	DeleteRecord(id string) (string, error)
	GetAllRecords() ([]model.Product, error)
	ExportToPDF(productsWithIndex []ProductWithIndex) (*gofpdf.Fpdf, error)
	UpdateRecord(record model.Product) (string, error)
}

type Service struct {
	User
	SalesList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:      NewUserService(repos.User),
		SalesList: NewSalesService(repos.SalesList),
	}
}

package repository

import (
	"github.com/zmskv/sales-app/internal/model"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
	GetUser(username, password string) (model.User, error)
}

type SalesList interface {
	CreateRecord(record model.Product) (int, error)
	GetRecord(id string) (model.Product, error)
	DeleteRecord(id string) (string, error)
	GetAllRecords() ([]model.Product, error)
}

type Repository struct {
	Authorization
	SalesList
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		SalesList:     NewSalesPostgres(db),
	}
}

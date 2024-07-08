package repository

import (
	"github.com/zmskv/sales-app/internal/model"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(user model.User) (string, error)
	GetUser(username, password string) (model.User, error)
	GetUserInfo(id string) (model.User, error)
	UpdateUserInfo(user model.User) (string, error)
	DeleteUser(id string) (string, error)
}

type SalesList interface {
	CreateRecord(record model.Product) (int, error)
	GetRecord(id string) (model.Product, error)
	DeleteRecord(id string) (string, error)
	GetAllRecords() ([]model.Product, error)
	UpdateRecord(record model.Product) (string, error)
}

type Repository struct {
	User
	SalesList
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:      NewUserPostgres(db),
		SalesList: NewSalesPostgres(db),
	}
}

package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type SalesList interface {
}

type Repository struct {
	Authorization
	SalesList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}

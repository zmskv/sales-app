package repository

type Authorization interface {
}

type SalesList interface {
}

type Repository struct {
	Authorization
	SalesList
}

func NewRepository() *Repository {
	return &Repository{}
}

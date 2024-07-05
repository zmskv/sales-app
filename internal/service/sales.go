package service

import (
	"github.com/zmskv/sales-app/internal/model"
	"github.com/zmskv/sales-app/internal/repository"
)

type SalesService struct {
	repos repository.SalesList
}

func NewSalesService(repos repository.SalesList) *SalesService {
	return &SalesService{repos: repos}
}

func (s *SalesService) CreateRecord(record model.Product) (int, error) {
	return s.repos.CreateRecord(record)
}

func (s *SalesService) GetRecord(id string) (model.Product, error) {
	return s.repos.GetRecord(id)
}

func (s *SalesService) DeleteRecord(id string) (string, error) {
	return s.repos.DeleteRecord(id)
}

func (s *SalesService) GetAllRecords() ([]model.Product, error) {
	return s.repos.GetAllRecords()
}

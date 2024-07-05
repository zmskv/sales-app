package repository

import (
	"github.com/zmskv/sales-app/internal/model"
	"gorm.io/gorm"
)

type SalesPostgres struct {
	db *gorm.DB
}

func NewSalesPostgres(db *gorm.DB) *SalesPostgres {
	return &SalesPostgres{db: db}
}

func (r *SalesPostgres) CreateRecord(record model.Product) (int, error) {

	result := r.db.Table("sales_list").Create(&record)
	if result.Error != nil {
		return 0, result.Error
	}

	return record.Id, nil
}

func (r *SalesPostgres) GetRecord(id string) (model.Product, error) {
	var data model.Product
	request := r.db.Table("sales_list").Where("id = ?", id).First(&data)
	if request.Error != nil {
		return model.Product{}, request.Error
	}

	return data, nil
}

package repository

import (
	"github.com/zmskv/sales-app/internal/model"
	"gorm.io/gorm"
)

type SalesRepository struct {
	db *gorm.DB
}

func NewSalesRepository(db *gorm.DB) *SalesRepository {
	return &SalesRepository{db: db}
}

func (r *SalesRepository) CreateRecord(record model.Product) (int, error) {

	result := r.db.Table("sales_list").Create(&record)
	if result.Error != nil {
		return 0, result.Error
	}

	return record.Id, nil
}

func (r *SalesRepository) GetRecord(id string) (model.Product, error) {
	var data model.Product
	request := r.db.Table("sales_list").Where("id = ?", id).First(&data)
	if request.Error != nil {
		return model.Product{}, request.Error
	}

	return data, nil
}

func (r *SalesRepository) DeleteRecord(id string) (string, error) {
	request := r.db.Table("sales_list").Where("id = ?", id).Delete(&model.Product{})
	if request.Error != nil {
		return "", request.Error
	}

	return "Successful deleted", nil
}

func (r *SalesRepository) GetAllRecords() ([]model.Product, error) {
	var products []model.Product
	request := r.db.Table("sales_list").Find(&products)
	if request.Error != nil {
		return nil, request.Error
	}

	return products, nil
}

func (r *SalesRepository) UpdateRecord(record model.Product) (string, error) {
	result := r.db.Table("sales_list").Where("id = ?", record.Id).Updates(record)
	if result.Error != nil {
		return "", result.Error
	}
	return "Succesful Updated", nil
}

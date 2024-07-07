package repository

import (
	"github.com/zmskv/sales-app/internal/model"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user model.User) (string, error) {

	result := r.db.Table("users").Create(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return user.Id, nil
}

func (r *UserPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	result := r.db.Table("users").Select("id, email").Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

package repository

import (
	"github.com/zmskv/sales-app/internal/model"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (string, error) {

	result := r.db.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}

	return user.Id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	result := r.db.Table("users").Select("id").Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

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

func (r *UserPostgres) GetUserInfo(id string) (model.User, error) {
	var user model.User
	result := r.db.Table("users").Select("id, username, email, created_at").Where("id = ?", id).First(&user)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return user, nil
}

func (r *UserPostgres) UpdateUserInfo(user model.User) (string, error) {
	result := r.db.Table("users").Where("id = ?", user.Id).Updates(user)
	if result.Error != nil {
		return "", result.Error
	}
	return "Succesful Updated", nil
}

func (r *UserPostgres) DeleteUser(id string) (string, error) {
	current_user, err := r.GetUserInfo(id)
	if err != nil {
		return "", err
	}
	sales := r.db.Table("sales_list").Where("username = ?", current_user.Username).Delete(&model.Product{})
	if sales.Error != nil {
		return "", sales.Error
	}
	result := r.db.Table("users").Where("id = ?", id).Delete(&model.User{})
	if result.Error != nil {
		return "", result.Error
	}

	return "Succesful Deleted", nil
}

package handler

import (
	"github.com/zmskv/sales-app/internal/repository"
	"github.com/zmskv/sales-app/internal/service"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB) *Handler {
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := NewHandler(services)
	return handlers
}

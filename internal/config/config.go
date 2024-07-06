package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmskv/sales-app/internal/handler"
	"github.com/zmskv/sales-app/internal/repository"
	"github.com/zmskv/sales-app/internal/service"
)

func InitConfig() *handler.Handler {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Username:   os.Getenv("DB_USERNAME"),
		Password:   os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SSLMode:    os.Getenv("SSL_MODE"),
		ServerPort: os.Getenv("SERVER_PORT"),
	})

	if err != nil {
		log.Fatalf("Error initialization db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	return handlers

}

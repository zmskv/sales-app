package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zmskv/sales-app/internal/repository"
	"gorm.io/gorm"
)

func InitConfig() *gorm.DB {
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
	return db

}

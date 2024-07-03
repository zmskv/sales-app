package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/zmskv/sales-app"
	"github.com/zmskv/sales-app/internal/handler"
	"github.com/zmskv/sales-app/internal/repository"
	"github.com/zmskv/sales-app/internal/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error Initializing config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(sales.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

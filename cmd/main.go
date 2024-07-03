package main

import (
	"log"

	"github.com/zmskv/sales-app"
	"github.com/zmskv/sales-app/internal/handler"
)

func main() {
	handlers := new(handler.Handler)
	server := new(sales.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error %s", err.Error())
	}

}

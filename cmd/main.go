package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/zmskv/sales-app/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(c context.Context) error {
	return s.httpServer.Shutdown(c)
}

func main() {
	handlers := config.InitConfig()
	server := new(Server)
	if err := server.Run(os.Getenv("SERVER_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error %s", err.Error())
	}

}

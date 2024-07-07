package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/zmskv/sales-app/internal/config"
	"github.com/zmskv/sales-app/internal/handler"
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

//	@title			Sales API
//	@version		1.0
//	@description	API Server SalesList Application

//@host localhost:8000
//@BasePath /

//@securityDefinitions.apikey ApiKeyAuth
//@in Header
//@name Authorization

func main() {
	db := config.InitConfig()
	app := handler.InitApp(db)
	server := new(Server)
	if err := server.Run(os.Getenv("SERVER_PORT"), app.InitRoutes()); err != nil {
		log.Fatalf("Error %s", err.Error())
	}

}

package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zmskv/sales-app/internal/service"

	_ "github.com/zmskv/sales-app/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/all_sales", h.getAllRecords)
		api.POST("/export_to_pdf", h.exportToPDF)
		list := api.Group("/list")
		{
			list.POST("/add", h.createRecord)
			list.GET("/:id", h.getRecord)
			list.DELETE("/delete", h.deleteRecord)

		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

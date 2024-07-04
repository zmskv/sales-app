package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zmskv/sales-app/internal/service"
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
		auth.POST("/sign-in", h.signIn)
		auth.POST("/login", h.Login)
	}

	api := router.Group("/api")
	{
		list := api.Group("/list", h.createList)
		{
			list.POST("/")
			list.POST("/export_to_pdf", h.exportToPDF)
			list.GET("/:id", h.getProduct)
			list.PATCH("/:id/edit", h.editProduct)
			list.DELETE("/:id", h.deleteProduct)

		}
	}

	return router
}

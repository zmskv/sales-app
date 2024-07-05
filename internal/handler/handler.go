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

	api := router.Group("/api", h.userIdentity)
	{
		list := api.Group("/list")
		{
			list.POST("/add", h.createRecord)
			list.POST("/export_to_pdf", h.exportToPDF)
			list.GET("/:id", h.getRecord)
			list.PATCH("/:id/edit", h.editRecord)
			list.DELETE("/:id", h.deleteRecord)

		}
	}

	return router
}

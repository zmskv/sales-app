package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
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

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

	account := router.Group("/account")
	{
		account.POST("/sign-up", h.signUp)
		account.POST("/sign-in", h.signIn)
		account.POST("/logout", h.Logout)
		account.GET("/info", h.getUserInfo)
		account.PATCH("/update_info", h.updateUserInfo)
		account.DELETE("/delete_user", h.deleteUser)

	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/all_sales", h.getAllRecords)
		api.GET("/export_to_pdf", h.exportToPDF)
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

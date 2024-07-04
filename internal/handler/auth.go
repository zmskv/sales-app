package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmskv/sales-app/internal/model"
)

func (h *Handler) Login(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		NewValidationResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	msg, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewValidationResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": msg,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewValidationResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewValidationResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}

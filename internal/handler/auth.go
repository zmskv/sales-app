package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmskv/sales-app/internal/model"
)

type signUpInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// signUp godoc
// @Summary		Sign Up
// @Security		ApiKeyAuth
// @Tags			auth
// @Description	Sign Up
// @ID				sign-up
// @Accept			json
// @Produce			json
// @Param		input	body signUpInput true "account info"
// @Success		200		{object}	SuccessResponse
// @Failure		400			{object}	ErrorResponse
// @Failure		500		{object}	ErrorResponse
// @Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input signUpInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var user model.User
	user.Username = input.Username
	user.Password = input.Password
	msg, err := h.services.User.CreateUser(user)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, msg)

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// signIn godoc
// @Summary		Sign In
// @Security		ApiKeyAuth
// @Tags			auth
// @Description	Sign In
// @ID				sign-in
// @Accept			json
// @Produce		json
// @Param		input	body signInInput true "account info"
// @Success		200		{object}	SuccessResponse
// @Failure		400		{object}	ErrorResponse
// @Failure		500		{object}	ErrorResponse
// @Router			/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.User.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, token)
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmskv/sales-app/internal/model"
)

type signUpInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// signUp godoc
//
//	@Summary		Sign Up
//	@Tags			Account
//	@Description	Сreates a new user in the system.
//	@ID				sign-up
//	@Accept			json
//	@Produce		json
//	@Param			query	body		signUpInput	true	"Account info"
//	@Success		200		{object}	SuccessResponse
//	@Failure		400		{object}	ErrorResponse
//	@Router			/account/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input signUpInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var user model.User
	user.Username = input.Username
	user.Password = input.Password
	user.Email = input.Email
	msg, err := h.services.User.CreateUser(user)

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Username or Email already used")
		return
	}

	NewSuccessResponse(c, http.StatusOK, msg)

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// signIn godoc
//
//	@Summary		Sign In
//	@Tags			Account
//	@Description	Sign In
//	@ID				sign-in
//	@Accept			json
//	@Produce		json
//	@Param			query	body		signInInput	true	"Account info"
//	@Success		200		{object}	SuccessResponse
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/account/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.User.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Incorrect login or password")
		return
	}
	c.SetCookie("token", token, 3600*12, "/", "", false, true)
	c.Header("Authorization", "Bearer "+token)
	NewSuccessResponse(c, http.StatusOK, token)
}

// Logout godoc
//
//	@Summary		Logout
//	@Security		ApiKeyAuth
//	@Tags			Account
//	@Description	Logout
//	@ID				Logout
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SuccessResponse
//	@Failure		401	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/account/logout [post]
func (h *Handler) Logout(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil || cookie == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "No token found in cookies")
		return
	}

	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Header("Authorization", "")

	NewSuccessResponse(c, http.StatusOK, "Successfully logged out")
}

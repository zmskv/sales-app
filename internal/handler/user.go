package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmskv/sales-app/internal/model"
)

type inputUpdateUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// getUserInfo godoc
//
//	@Summary		Get User Info
//
//	@Security		ApiKeyAuth
//
//	@Tags			Account
//	@Description	Get User Info
//	@ID				get-user-info
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SuccessResponse
//	@Failure		401	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/account/info [get]
func (h *Handler) getUserInfo(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "Token not found in cookies")
		return
	}

	UserId, _, _, err := h.services.User.ParseToken(cookie)

	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	info, err := h.services.User.GetUserInfo(UserId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, info)

}

// getUserInfo godoc
//
//	@Summary		Update User Info
//
//	@Security		ApiKeyAuth
//
//	@Tags			Account
//	@Description	Update User Info
//	@ID				update-info
//	@Accept			json
//	@Produce		json
//	@Param			query	body		inputUpdateUser	true	"Account info"
//	@Success		200		{object}	SuccessResponse
//	@Failure		400		{object}	ErrorResponse
//	@Failure		401		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/account/update_info [patch]
func (h *Handler) updateUserInfo(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil || cookie == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "Token not found in cookies")
		return
	}
	UserId, _, _, err := h.services.User.ParseToken(cookie)

	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	var user inputUpdateUser
	if err := c.BindJSON(&user); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	msg, err := h.services.User.UpdateUserInfo(model.User{Id: UserId, Username: user.Username, Password: user.Password, Email: user.Email})
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	NewSuccessResponse(c, http.StatusOK, msg)

}

// deleteUser godoc
//
//	@Summary		Delete Current User
//
//	@Security		ApiKeyAuth
//
//	@Tags			Account
//	@Description	Delete Current User
//	@ID				delete-current-user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SuccessResponse
//	@Failure		401	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/account/delete_user [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil || cookie == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "No token found in cookies")
		return
	}
	UserId, _, _, err := h.services.User.ParseToken(cookie)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	msg, err := h.services.DeleteUser(UserId)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Header("Authorization", "")
	NewSuccessResponse(c, http.StatusOK, msg)
}

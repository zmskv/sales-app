package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "Token not found in cookies")
		return
	}

	UserId, Username, Email, err := h.services.User.ParseToken(cookie)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("user_id", UserId)
	c.Set("username", Username)
	c.Set("email", Email)
}

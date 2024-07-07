package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "Auth Header is empty")
		return
	}

	headerParts := strings.Split(header, " ")

	UserId, Username, Email, err := h.services.User.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("user_id", UserId)
	c.Set("username", Username)
	c.Set("email", Email)

}

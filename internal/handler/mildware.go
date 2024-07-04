package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		NewValidationResponse(c, http.StatusUnauthorized, "Auth Header is empty")
		return
	}

	headerParts := strings.Split(header, " ")

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewValidationResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)
}

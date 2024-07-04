package handler

import "github.com/gin-gonic/gin"

type Validation struct {
	Message string `json:"message"`
}

func NewValidationResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, Validation{message})
}

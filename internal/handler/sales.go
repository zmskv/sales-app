package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zmskv/sales-app/internal/model"
)

func (h *Handler) createRecord(c *gin.Context) {
	var input model.Product
	username, _ := c.Get("username")
	input.Username = username.(string)
	input.Date = time.Now()

	if err := c.BindJSON(&input); err != nil {
		NewValidationResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	msg, err := h.services.SalesList.CreateRecord(input)
	if err != nil {
		NewValidationResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"created record": msg,
	})
}

func (h *Handler) exportToPDF(c *gin.Context) {

}

func (h *Handler) getRecord(c *gin.Context) {

}

func (h *Handler) editRecord(c *gin.Context) {

}

func (h *Handler) deleteRecord(c *gin.Context) {

}

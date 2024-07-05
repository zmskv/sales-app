package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zmskv/sales-app/internal/model"
	"github.com/zmskv/sales-app/internal/service"
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

func (h *Handler) getRecord(c *gin.Context) {
	id := c.Param("id")
	data, err := h.services.SalesList.GetRecord(id)
	if err != nil {
		NewValidationResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       data.Id,
		"username": data.Username,
		"title":    data.Title,
		"amount":   data.Amount,
		"price":    data.Price,
		"date":     data.Date,
	})

}

type DeleteRequest struct {
	ID string `json:"id" binding:"required"`
}

func (h *Handler) deleteRecord(c *gin.Context) {
	var req DeleteRequest
	if err := c.BindJSON(&req); err != nil {
		NewValidationResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id := req.ID
	username, _ := c.Get("username")
	data, err := h.services.SalesList.GetRecord(id)
	if err != nil {
		NewValidationResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if data.Username != username {
		NewValidationResponse(c, http.StatusForbidden, "User does not have permission to delete this record")
		return
	}

	msg, err := h.services.SalesList.DeleteRecord(id)
	if err != nil {
		NewValidationResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": msg,
	})
}

func (h *Handler) getAllRecords(c *gin.Context) {
	data, err := h.services.SalesList.GetAllRecords()
	if err != nil {
		NewValidationResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var productsWithIndex []service.ProductWithIndex
	for i, product := range data {
		productsWithIndex = append(productsWithIndex, service.ProductWithIndex{
			Index:   i + 1,
			Product: product,
		})
	}

	c.JSON(http.StatusOK, productsWithIndex)

}

func (h *Handler) exportToPDF(c *gin.Context) {
	sales, err := h.services.SalesList.GetAllRecords()
	if err != nil {
		NewValidationResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var data []service.ProductWithIndex
	for i, product := range sales {
		data = append(data, service.ProductWithIndex{
			Index:   i + 1,
			Product: product,
		})
	}
	docs, err := h.services.SalesList.ExportToPDF(data)
	if err != nil {
		NewValidationResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Disposition", "attachment; filename=sales_report.pdf")
	c.Header("Content-Type", "application/pdf")
	docs.Output(c.Writer)
}

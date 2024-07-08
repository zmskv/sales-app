package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zmskv/sales-app/internal/model"
	"github.com/zmskv/sales-app/internal/service"
)

type recordInput struct {
	Id     int     `json:"id" binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Amount int     `json:"amount" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

// createRecord godoc
//
//	@Summary		Create Record
//	@Security		ApiKeyAuth
//	@Tags			Sales
//	@Description	Create Record
//	@ID				create-record
//	@Accept			json
//	@Produce		json
//	@Param			query	body		recordInput	true	"Record info"
//	@Success		200		{object}	SuccessResponse
//	@Failure		400		{object}	ErrorResponse
//	@Failure		401		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/api/list/add [post]
func (h *Handler) createRecord(c *gin.Context) {
	var input recordInput
	username, _ := c.Get("username")

	record := model.Product{
		Id:       input.Id,
		Title:    input.Title,
		Amount:   input.Amount,
		Price:    input.Price,
		Username: username.(string),
		Date:     time.Now(),
	}

	if err := c.BindJSON(&record); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	msg, err := h.services.SalesList.CreateRecord(record)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, fmt.Sprint(msg))
}

// getRecord godoc
//
//	@Summary		Get Record
//	@Security		ApiKeyAuth
//	@Tags			Sales
//	@Description	Get Record
//	@ID				get-record
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	false	"id"
//
//	@Success		200	{object}	model.Product
//
//	@Failure		400	{object}	ErrorResponse
//
//	@Failure		401	{object}	ErrorResponse
//
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/api/list/{id} [get]
func (h *Handler) getRecord(c *gin.Context) {
	id := c.Param("id")
	data, err := h.services.SalesList.GetRecord(id)
	if err != nil {
		if err.Error() == "record not found" {
			NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
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

// deleteRecord godoc
//
//	@Summary		Delete Record
//	@Security		ApiKeyAuth
//	@Tags			Sales
//	@Description	Delete Record
//	@ID				delete-record
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	false	"id"
//
//	@Success		200	{object}	SuccessResponse
//
//	@Failure		400	{object}	ErrorResponse
//
//	@Failure		401	{object}	ErrorResponse
//	@Failure		403	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/api/list/delete [delete]
func (h *Handler) deleteRecord(c *gin.Context) {
	id := c.Query("id")

	username, _ := c.Get("username")
	data, err := h.services.SalesList.GetRecord(id)
	if err != nil {
		if err.Error() == "record not found" {
			NewErrorResponse(c, http.StatusNotFound, "ID is required")
			return
		}
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if data.Username != username {
		NewErrorResponse(c, http.StatusForbidden, "User does not have permission to delete this record")
		return
	}

	msg, err := h.services.SalesList.DeleteRecord(id)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	NewSuccessResponse(c, http.StatusOK, msg)
}

// getAllRecords godoc
//
// @Summary		Get All Records
// @Security		ApiKeyAuth
// @Tags			Sales
// @Description	Get All Records
// @ID				get-all-records
// @Accept			json
// @Produce		json
//
// @Success		200	{object}	service.ProductWithIndex
//
// @Failure		401	{object}	ErrorResponse
// @Failure		404	{object}	ErrorResponse
// @Failure		500	{object}	ErrorResponse
// @Router			/api/all_sales [get]
func (h *Handler) getAllRecords(c *gin.Context) {
	data, err := h.services.SalesList.GetAllRecords()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if len(data) == 0 {
		NewErrorResponse(c, http.StatusNotFound, "Records not found")
		return
	}
	var productsWithIndex []service.ProductWithIndex
	for _, product := range data {
		productsWithIndex = append(productsWithIndex, service.ProductWithIndex{
			Index:   product.Id,
			Product: product,
		})
	}

	c.JSON(http.StatusOK, productsWithIndex)

}

// exportToPDF godoc
//
// @Summary		Export data to PDF file
// @Security		ApiKeyAuth
// @Tags			Sales
// @Description	Get All Records
// @ID				export-to-pdf
// @Accept			json
// @Produce		json
//
// @Success		200	{object}	SuccessResponse
//
// @Failure		401	{object}	ErrorResponse
// @Failure		404	{object}	ErrorResponse
// @Failure		500	{object}	ErrorResponse
// @Router			/api/export_to_pdf [get]
func (h *Handler) exportToPDF(c *gin.Context) {
	sales, err := h.services.SalesList.GetAllRecords()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var data []service.ProductWithIndex
	for _, product := range sales {
		data = append(data, service.ProductWithIndex{
			Index:   product.Id,
			Product: product,
		})
	}
	docs, err := h.services.SalesList.ExportToPDF(data)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Disposition", "attachment; filename=sales_report.pdf")
	c.Header("Content-Type", "application/pdf")
	docs.Output(c.Writer)
}

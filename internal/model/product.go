package model

import (
	"time"
)

type Product struct {
	Id       int       `json:"id" binding:"required"`
	Username string    `json:"username"`
	Title    string    `json:"title" binding:"required"`
	Amount   int       `json:"amount" binding:"required"`
	Price    float64   `json:"price" binding:"required"`
	Date     time.Time `json:"date"`
}

package model

import (
	"time"
)

type Product struct {
	Id     int       `json:"id"`
	Title  string    `json:"title"`
	Amount int       `json:"amount"`
	Price  int       `json:"price"`
	Total  int       `json:"total"`
	Date   time.Time `json:"date"`
}

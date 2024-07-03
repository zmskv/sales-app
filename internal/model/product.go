package model

import (
	"time"
)

type Product struct {
	Id     int       `json:"id"`
	Title  string    `json:"title"`
	Amount int       `json:"amount"`
	Total  int       `json:"summa"`
	Date   time.Time `json:"date"`
}

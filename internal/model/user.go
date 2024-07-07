package model

import "time"

type User struct {
	Id        string
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	CreatedAt time.Time
}

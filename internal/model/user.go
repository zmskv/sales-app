package model

type User struct {
	Id        string `json:"-"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	IsCreated bool   `json:"-"`
}

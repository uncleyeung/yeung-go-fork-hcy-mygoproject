package models

import "myProject/corll"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthUser struct {
	Token string
	User  corll.User
}
package auth

import (
	"time"
)

type Entity interface {
	LoginSerialize()
}

type InputRegister struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

type InputLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

type UserResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type TokenResponse struct {
	AccessToken    string    `json:"AccessToken"`
	AccessExpired  time.Time `json:"AccessExpired"`
	RefreshToken   string    `json:"RefreshToken"`
	RefreshExpired time.Time `json:"RefreshExpired"`
}

package domain

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Hashed   string `json:"password"`
	Salted   string
	Verified bool
}

type Claims struct {
	UserID uint
	Role   string
	jwt.StandardClaims
}

package domain

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Claims struct {
	UserID uint
	Name   string
	Role   string
	jwt.StandardClaims
}

type Admin struct {
	gorm.Model
	Email  string `gorm:"unique"`
	Hashed string
	Salted string
	Super  bool
}

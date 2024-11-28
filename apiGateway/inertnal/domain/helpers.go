package domain

import "github.com/dgrijalva/jwt-go"

type Response struct {
	Message    any `json:"message"`
	StatusCode int `json:"code"`
}

type Claims struct {
	UserID uint
	Name   string
	Role   string
	jwt.StandardClaims
}

type Razor struct {
	Order     string `json:"OrderID"`
	Payment   string `json:"PaymentID"`
	Signature string `json:"Signature"`
}

type Message struct {
	Amount int `json:"amount"`
}

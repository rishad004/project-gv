package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Phone     string
	Gender    string
	Following string
	Hashed    string
	Salted    string
	Verified  bool
}

type Claims struct {
	UserID uint
	Name   string
	Role   string
	jwt.StandardClaims
}

type Subscribed struct {
	gorm.Model
	UserId         int
	SubscriptionId int
	Expiry         time.Time
	PaymentId      string
}

type Wallet struct {
	gorm.Model
	Amount int
	UserId int
}

type Kafka struct {
	Message string `json:"message"`
}

type WalletTransactions struct {
	gorm.Model
	WalletId int
	Amount   int
	Type     string
}

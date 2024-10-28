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

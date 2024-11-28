package domain

import (
	"gorm.io/gorm"
)

type Streamer struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
	StreamKey   string `gorm:"unique"`
	UserId      uint   `gorm:"unique"`
	Followers   string
}

type Subscription struct {
	gorm.Model
	Amount     int
	StreamerId int `gorm:"unique"`
}

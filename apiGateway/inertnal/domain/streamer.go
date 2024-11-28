package domain

import "gorm.io/gorm"

type Streamer struct {
	gorm.Model
	Name        string
	Description string
	StreamKey   string
	UserId      uint
}

type Subscription struct {
	gorm.Model
	Amount     int
	StreamerId int
}
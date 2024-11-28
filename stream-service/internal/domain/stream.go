package domain

import (
	"time"

	"gorm.io/gorm"
)

type StreamData struct {
	gorm.Model
	StreamerId  int
	Title       string
	Description string
}

type Stream struct {
	gorm.Model
	StreamerId  int
	Status      bool
	StreamStart time.Time
	StreamEnd   time.Time
}

package domain

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	RazorId string
	OrderId string
	Amount  int
	Type    string
	Status  bool
}

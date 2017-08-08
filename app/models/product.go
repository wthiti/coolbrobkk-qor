package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ID          int
	Name        string
	Description string
	Image       string
	Price       int

	OrderProducts []OrderProduct
}

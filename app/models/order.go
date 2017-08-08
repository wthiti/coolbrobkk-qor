package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	ID int

	// Order
	Status       string
	OrderDate    time.Time
	DeliveryDate time.Time
	DeliveryFee  int

	// User
	Name    string
	Address string
	Phone   string
	Email   string

	OrderProducts []OrderProduct
}

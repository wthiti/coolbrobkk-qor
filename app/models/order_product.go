package models

import (
	"github.com/jinzhu/gorm"
)

type OrderProduct struct {
	gorm.Model
	ID int

	Order   Order
	OrderID int

	Product   Product
	ProductID int

	Quitity int
	Price   int
}

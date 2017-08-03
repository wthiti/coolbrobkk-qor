package models

import (
  "github.com/jinzhu/gorm"
)

type User struct {
  gorm.Model
  Email string
  Password string
  Name string
  Gender string
  Role string
}

type Product struct {
  gorm.Model
  Name string
  Description string
}

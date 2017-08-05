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

func (u User) DisplayName() string {
  return u.Name
}

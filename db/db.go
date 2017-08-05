package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  "github.com/qor/publish2"

  "bitbucket.org/wthiti/coolbrobkk-qor/app/models"
)

var (
  DB *gorm.DB
)

func init() {
  DB, _ = gorm.Open("sqlite3", "test.db")
  DB.AutoMigrate(&models.User{}, &models.Product{})

  publish2.RegisterCallbacks(DB)
}

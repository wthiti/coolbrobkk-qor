package main

import (
  "fmt"
  "net/http"

  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
  "github.com/qor/qor"
  "github.com/qor/admin"

  "bitbucket.org/wthiti/coolbrobkk-qor/app/models"
)

func main() {
  DB, _ := gorm.Open("sqlite3", "demo.db")
  DB.AutoMigrate(&models.User{}, &models.Product{})

  Admin := admin.New(&qor.Config{DB: DB})

  Admin.SetSiteName("CoolBroBKK")

  Admin.AddResource(&models.User{})
  Admin.AddResource(&models.Product{})

  mux := http.NewServeMux()
  Admin.MountTo("/admin", mux)

  fmt.Println("Listening on: 9000")
  http.ListenAndServe(":9000", mux)
}

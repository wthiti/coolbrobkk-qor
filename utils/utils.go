package utils

import (
  "net/http"

  "github.com/go-chi/chi"
  "github.com/jinzhu/gorm"
  "github.com/qor/qor/utils"

  "bitbucket.org/wthiti/coolbrobkk-qor/db"
)

func GetDB(req *http.Request) *gorm.DB {
  if db := utils.GetDBFromRequest(req); db != nil {
    return db
  }
  return db.DB
}

func URLParam(name string, req *http.Request) string {
  return chi.URLParam(req, name)
}

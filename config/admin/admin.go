package admin

import (
  "github.com/qor/qor"
  "github.com/qor/admin"

  "bitbucket.org/wthiti/coolbrobkk-qor/app/models"
  "bitbucket.org/wthiti/coolbrobkk-qor/db"
  "bitbucket.org/wthiti/coolbrobkk-qor/config/auth"
)

var (
  QORAdmin *admin.Admin
)

func init() {
  QORAdmin = admin.New(&qor.Config{DB: db.DB})

  QORAdmin.SetSiteName("CoolBroBKK")

  QORAdmin.SetAuth(&auth.AdminAuth{})

  QORAdmin.AddResource(&models.User{})
  QORAdmin.AddResource(&models.Product{})
}

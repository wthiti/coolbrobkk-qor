package admin

import (
	"github.com/qor/admin"
	"github.com/qor/qor"

	"bitbucket.org/wthiti/coolbrobkk-qor/app/models"
	"bitbucket.org/wthiti/coolbrobkk-qor/db"
)

var (
	QORAdmin *admin.Admin
)

func init() {
	QORAdmin = admin.New(&qor.Config{DB: db.DB})

	QORAdmin.SetSiteName("CoolBroBKK")

	QORAdmin.AddResource(&models.User{})
	QORAdmin.AddResource(&models.Product{})
}

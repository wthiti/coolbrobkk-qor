package auth

import (
  "github.com/qor/admin"
  "github.com/qor/qor"

  "bitbucket.org/wthiti/coolbrobkk-qor/app/models"
)

type AdminAuth struct {}

func (AdminAuth) LoginURL(c *admin.Context) string {
  return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
  return "/auth/logout"
}

func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
  return &models.User{}
}

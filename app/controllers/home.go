package controllers

import (
  "net/http"

  "bitbucket.org/wthiti/coolbrobkk-qor/config"
)

func HomeIndex(w http.ResponseWriter, req *http.Request) {
  config.View.Execute("home_index", map[string]interface{}{}, req, w)
}

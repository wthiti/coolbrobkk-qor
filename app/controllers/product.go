package controllers

import (
  "net/http"

  "bitbucket.org/wthiti/coolbrobkk-qor/config"
)

func ProductIndex(w http.ResponseWriter, req *http.Request) {
  var (
    // db = utils.GetDB(req)
  )

  config.View.Execute("product/product_index", map[string]interface{}{}, req, w)
}

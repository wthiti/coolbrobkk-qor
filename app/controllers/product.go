package controllers

import (
	"log"
	"net/http"

	"bitbucket.org/wthiti/coolbrobkk-qor/config/view"
	"bitbucket.org/wthiti/coolbrobkk-qor/utils"
)

func ProductIndex(w http.ResponseWriter, req *http.Request) {
	var (
		db = utils.GetDB(req)
		v  = view.ViewValueMap()
	)

	log.Println(db)
	log.Println(req.Context())
	view.View.Execute("product/product_index", v, req, w)
}

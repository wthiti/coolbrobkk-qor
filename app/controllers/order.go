package controllers

import (
	"net/http"

	"bitbucket.org/wthiti/coolbrobkk-qor/config/view"
)

func OrderIndex(w http.ResponseWriter, req *http.Request) {
	var (
	// db = utils.GetDB(req)
	)

	view.View.Execute("order/order_index", req.Context(), req, w)
}

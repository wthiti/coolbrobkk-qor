package controllers

import (
	"net/http"

	"bitbucket.org/wthiti/coolbrobkk-qor/config/view"
)

func HomeIndex(w http.ResponseWriter, req *http.Request) {
	view.View.Execute("home_index", map[string]interface{}{}, req, w)
}

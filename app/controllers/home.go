package controllers

import (
	"net/http"

	"bitbucket.org/wthiti/coolbrobkk-qor/config/view"
)

func HomeIndex(w http.ResponseWriter, req *http.Request) {
	view.View.Execute("home_index", map[string]interface{}{}, req, w)
}

func ProcessIndex(w http.ResponseWriter, req *http.Request) {
	view.View.Execute("process_index", map[string]interface{}{}, req, w)
}

func FindUsIndex(w http.ResponseWriter, req *http.Request) {
	view.View.Execute("findus_index", map[string]interface{}{}, req, w)
}

func ContactUsIndex(w http.ResponseWriter, req *http.Request) {
	view.View.Execute("contactus_index", map[string]interface{}{}, req, w)
}

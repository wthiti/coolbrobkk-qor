package controllers

import (
	"net/http"
	"net/mail"

	"github.com/qor/mailer"

	"bitbucket.org/wthiti/coolbrobkk-qor/config"
	"bitbucket.org/wthiti/coolbrobkk-qor/config/view"
)

type ContactInformation struct {
	Name    string
	Address string
	Email   string
	Phone   string
	Comment string
}

func HomeIndex(w http.ResponseWriter, req *http.Request) {
	view.View.Execute("home_index", map[string]interface{}{}, req, w)
}

func ProcessIndex(w http.ResponseWriter, req *http.Request) {
	view.View.Execute("home_index", map[string]interface{}{}, req, w)
}

func FindUsIndex(w http.ResponseWriter, req *http.Request) {
	view.View.Execute("home_index", map[string]interface{}{}, req, w)
}

func ContactUsIndex(w http.ResponseWriter, req *http.Request) {
	view.View.Execute("contactus_index", map[string]interface{}{}, req, w)
}

func ContactUsProcess(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	contactInformation := ContactInformation{}
	contactInformation.Name = req.FormValue("name")
	contactInformation.Address = req.FormValue("address")
	contactInformation.Email = req.FormValue("email")
	contactInformation.Phone = req.FormValue("phone")
	contactInformation.Comment = req.FormValue("comment")
	createContactMail(&contactInformation, w, req)
	view.View.Execute("contactus_complete", map[string]interface{}{}, req, w)
}

func createContactMail(contactInformation *ContactInformation, w http.ResponseWriter, req *http.Request) {
	v := view.ViewValueMap()
	v["info"] = contactInformation
	template := mailer.Template{
		Name:    "contact",
		Data:    v,
		Request: req,
		Writer:  w,
	}
	infoEmail := "info@coolbrobkk.com"
	email := config.Mailer.Render(template)
	email.TO = []mail.Address{mail.Address{Address: infoEmail}}
	email.From = &mail.Address{Address: infoEmail}
	email.BCC = []mail.Address{mail.Address{Address: infoEmail}}
	email.Subject = "Coolbro Coldbrew Inquiry"
	config.Mailer.Send(email)
}

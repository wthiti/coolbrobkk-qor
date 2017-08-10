package config

import (
	"github.com/qor/mailer"
	"github.com/qor/mailer/gomailer"
	"gopkg.in/gomail.v2"
)

var (
	Mailer *mailer.Mailer
)

func init() {
	dialer := gomail.NewDialer("smtp.zoho.com", 587, "daemon@coolbrobkk.com", "iQBEW7zY#C3z")
	sender, _ := dialer.Dial()

	Mailer = mailer.New(&mailer.Config{
		Sender: gomailer.New(&gomailer.Config{
			Sender: sender,
		}),
	})

}

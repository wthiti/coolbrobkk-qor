package auth

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/justinas/nosurf"
	"github.com/volatiletech/authboss"
	_ "github.com/volatiletech/authboss/auth"

	"bitbucket.org/wthiti/coolbrobkk-qor/app/models"
	"bitbucket.org/wthiti/coolbrobkk-qor/config"
)

var funcs = template.FuncMap{
	"formatDate": func(date time.Time) string {
		return date.Format("2006/01/02 03:04pm")
	},
	"yield": func() string { return "" },
}

var (
	ab = authboss.New()
)

func GetAuthBoss() *authboss.Authboss {
	ab.Storer = NewMemStorer()
	ab.MountPath = "/auth"
	ab.LogWriter = os.Stdout
	ab.ViewsPath = "app/views/ab_views"
	ab.RootURL = `http://localhost:9000`
	ab.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Not Found in AuthBoss")
	})

	ab.LayoutDataMaker = layoutData
	ab.XSRFName = "csrf_token"
	ab.XSRFMaker = func(_ http.ResponseWriter, r *http.Request) string {
		return nosurf.Token(r)
	}

	ab.CookieStoreMaker = config.NewCookieStorer
	ab.SessionStoreMaker = config.NewSessionStorer

	ab.Mailer = authboss.LogMailer(os.Stdout)

	ab.Policies = []authboss.Validator{
		authboss.Rules{
			FieldName:       "email",
			Required:        true,
			AllowWhitespace: false,
		},
		authboss.Rules{
			FieldName:       "password",
			Required:        true,
			MinLength:       4,
			MaxLength:       8,
			AllowWhitespace: false,
		},
	}

	b, err := ioutil.ReadFile(filepath.Join("app/views/layouts", "layout.html.tpl"))
	if err != nil {
		log.Println("in auth")
		panic(err)
	}
	ab.Layout = template.Must(template.New("layout").Funcs(funcs).Parse(string(b)))

	if err := ab.Init(); err != nil {
		log.Fatalln(err)
	}

	return ab
}

func layoutData(w http.ResponseWriter, r *http.Request) authboss.HTMLData {
	currentUserName := ""
	userInter, err := ab.CurrentUser(w, r)
	if userInter != nil && err == nil {
		currentUserName = userInter.(*models.User).DisplayName()
	}

	return authboss.HTMLData{
		"loggedin":               userInter != nil,
		"username":               "",
		authboss.FlashSuccessKey: ab.FlashSuccess(w, r),
		authboss.FlashErrorKey:   ab.FlashError(w, r),
		"current_user_name":      currentUserName,
	}
}

package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/wthiti/coolbrobkk-qor/config/view"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var (
	conf *oauth2.Config
)

type quizUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func init() {
	conf = &oauth2.Config{
		ClientID:     "1760831864216176",
		ClientSecret: "da5d4de292b7c858ced59b1cebf8805d",
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
		RedirectURL:  "http://localhost:9000/quiz/facebook",
	}
}

func generateFacebookLoginURL() string {
	url := conf.AuthCodeURL("coolbroquiz")
	return url
}

func Quiz(w http.ResponseWriter, req *http.Request) {
	var (
		// db = utils.GetDB(req)
		v = view.ViewValueMap()
	)
	v["facebook_login"] = generateFacebookLoginURL()

	view.View.Execute("quiz/quiz_login", v, req, w)
}

func QuizFacebook(w http.ResponseWriter, req *http.Request) {
	var (
		v = view.ViewValueMap()
	)
	ctx := context.Background()
	code := req.FormValue("code")
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(ctx, tok)
	resp, err := client.Get("https://graph.facebook.com/v2.10/me?fields=id,name")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	user := quizUser{}
	jsonDecode := json.NewDecoder(resp.Body)
	err = jsonDecode.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	v["id"] = user.Id
	v["name"] = user.Name

	view.View.Execute("quiz/quiz", v, req, w)
}

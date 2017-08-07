package middleware

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"

	"bitbucket.org/wthiti/coolbrobkk-qor/config/auth"
)

func AuthProtect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if u, err := auth.GetAuthBoss().CurrentUser(w, r); err != nil {
			log.Println("Error fetching current user:", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else if u == nil {
			log.Println("Redirecting unauthorized user from:", r.URL.Path)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func Nosurfing(next http.Handler) http.Handler {
	surfing := nosurf.New(next)
	surfing.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Failed to validate XSRF Token:", nosurf.Reason(r))
		w.WriteHeader(http.StatusBadRequest)
	}))
	return surfing
}

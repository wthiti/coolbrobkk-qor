package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/qor/qor/utils"

	"bitbucket.org/wthiti/coolbrobkk-qor/app/controllers"
	"bitbucket.org/wthiti/coolbrobkk-qor/config/admin"
	"bitbucket.org/wthiti/coolbrobkk-qor/config/auth"
	"bitbucket.org/wthiti/coolbrobkk-qor/config/middleware"
	"bitbucket.org/wthiti/coolbrobkk-qor/db"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Nosurfing)
	router.Use(qorContextMiddleware)
	router.NotFound(notFoundHandler())

	fs := http.FileServer(http.Dir("public"))

	router.Get("/", controllers.HomeIndex)
	router.Get("/process", controllers.ProcessIndex)
	router.Get("/findus", controllers.FindUsIndex)
	router.Get("/contactus", controllers.ContactUsIndex)

	router.Mount("/public", http.StripPrefix("/public", fs))

	router.Mount("/admin", middleware.AuthProtect(admin.QORAdmin.NewServeMux("/admin")))

	router.Mount("/auth", logger(auth.GetAuthBoss().NewRouter()))

	router.Get("/product/cart", controllers.CartView)
	router.Post("/cart", controllers.CartAdd)
	router.Post("/cart/edit", controllers.CartEditQuantity)
	router.Post("/cart/delete", controllers.CartDelete)
	router.Post("/product/cart/confirm", controllers.CartConfirmOrder)

	//app
	router.Get("/product", controllers.ProductIndex)

	return router
}

func qorContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), utils.ContextDBName, db.DB)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n%s %s %s\n", r.Method, r.URL.Path, r.Proto)
		next.ServeHTTP(w, r)
	})
}

func notFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path not Found")
	}
}

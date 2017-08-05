package routes

import (
  "context"
  "net/http"
  "github.com/go-chi/chi"
  "github.com/qor/qor/utils"

  "bitbucket.org/wthiti/coolbrobkk-qor/db"
  "bitbucket.org/wthiti/coolbrobkk-qor/config/admin"
  "bitbucket.org/wthiti/coolbrobkk-qor/app/controllers"
)

func Router() http.Handler {
  router := chi.NewRouter()

  router.Use(qorContextMiddleware)

  fs := http.FileServer(http.Dir("public"))

  router.Get("/", controllers.HomeIndex)
  router.Mount("/public", http.StripPrefix("/public", fs))

  mux := http.NewServeMux()
  admin.QORAdmin.MountTo("/admin", mux)
  router.Mount("/admin", mux)

  return router
}

func qorContextMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    ctx := context.WithValue(r.Context(), utils.ContextDBName, db.DB)
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

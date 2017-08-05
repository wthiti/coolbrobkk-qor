package main

import (
  "fmt"
  "net/http"

  _ "bitbucket.org/wthiti/coolbrobkk-qor/db"
  "bitbucket.org/wthiti/coolbrobkk-qor/config/routes"
)

func main() {
  router := routes.Router()

  fmt.Println("Listening on: 9000")
  http.ListenAndServe(":9000", router)
}

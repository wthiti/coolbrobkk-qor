package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/wthiti/coolbrobkk-qor/config/routes"
	_ "bitbucket.org/wthiti/coolbrobkk-qor/db"
	"github.com/gorilla/context"
)

func main() {
	router := routes.Router()

	fmt.Println("Listening on: 80")
	log.Println(http.ListenAndServe(":80", context.ClearHandler(router)))
}

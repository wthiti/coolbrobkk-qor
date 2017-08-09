package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/wthiti/coolbrobkk-qor/app/models"
	"bitbucket.org/wthiti/coolbrobkk-qor/config"
	"bitbucket.org/wthiti/coolbrobkk-qor/config/view"
	"bitbucket.org/wthiti/coolbrobkk-qor/utils"
)

type Cart struct {
	Items []Item
}

type Item struct {
	Product  models.Product
	Quantity int
}

func CartView(w http.ResponseWriter, req *http.Request) {
	var (
		// db = utils.GetDB(req)
		v = view.ViewValueMap()
	)

	view.View.Execute("cart/cart_view", v, req, w)
}

func CartAdd(w http.ResponseWriter, req *http.Request) {
	productID, _ := strconv.Atoi(utils.URLParam("ProductID", req))
	product := GetTempProductList()[productID]
	item := []Item{Item{Product: product, Quantity: 1}}
	cart := &Cart{Items: item}
	jsonStr, _ := json.Marshal(cart)

	cookieStorer := config.GetCookieStorer(w, req)
	log.Print(cookieStorer.Get("cart"))
	cookieStorer.Put("cart", string(jsonStr))
	log.Println(string(jsonStr))
	http.Redirect(w, req, "/product", http.StatusOK)
}

func CardEdit(w http.ResponseWriter, req *http.Request) {

}

package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"strconv"

	"bitbucket.org/wthiti/coolbrobkk-qor/app/models"
	"bitbucket.org/wthiti/coolbrobkk-qor/config"
	"bitbucket.org/wthiti/coolbrobkk-qor/config/view"
	"github.com/qor/mailer"
)

type Cart struct {
	Items       []Item
	TotalAmount int
	Name        string
	Address     string
	Email       string
	Phone       string
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
	cookieStorer := config.GetCookieStorer(w, req)
	jsonStr, _ := cookieStorer.Get("cart")
	customerName, _ := cookieStorer.Get("customer_name")
	customerAddress, _ := cookieStorer.Get("customer_address")
	customerEmail, _ := cookieStorer.Get("customer_email")
	customerPhone, _ := cookieStorer.Get("customer_phone")

	cart := Cart{}
	json.Unmarshal([]byte(jsonStr), &cart)
	cart.Name = customerName
	cart.Address = customerAddress
	cart.Email = customerEmail
	cart.Phone = customerPhone

	v["cart"] = cart
	view.View.Execute("cart/cart_view", v, req, w)
}

func CartAdd(w http.ResponseWriter, req *http.Request) {
	productID, _ := strconv.Atoi(req.FormValue("ID"))
	product := GetTempProductList()[productID-1]
	item := []Item{Item{Product: product, Quantity: 1}}
	cart := &Cart{Items: item}
	jsonStr, _ := json.Marshal(cart)

	cookieStorer := config.GetCookieStorer(w, req)
	cookieStorer.Put("cart", string(jsonStr))
	http.Redirect(w, req, "product/cart", http.StatusFound)
}

func CartEditQuantity(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	log.Println(req.Form)
	http.Redirect(w, req, "product/cart", http.StatusFound)
}

func CartDelete(w http.ResponseWriter, req *http.Request) {

}

func CartConfirmOrder(w http.ResponseWriter, req *http.Request) {
	v := view.ViewValueMap()
	cookieStorer := config.GetCookieStorer(w, req)
	jsonStr, _ := cookieStorer.Get("cart")
	cart := Cart{}
	json.Unmarshal([]byte(jsonStr), &cart)

	cart.Name = req.FormValue("name")
	cart.Address = req.FormValue("address")
	cart.Email = req.FormValue("email")
	cart.Phone = req.FormValue("phone")

	storeAddressInCookie(&cart, w, req)
	createPreorderMail(&cart, w, req)
	view.View.Execute("cart/cart_complete", v, req, w)
}

func createPreorderMail(cart *Cart, w http.ResponseWriter, req *http.Request) {
	v := view.ViewValueMap()
	v["cart"] = cart
	template := mailer.Template{
		Name:    "preorder",
		Data:    v,
		Request: req,
		Writer:  w,
	}
	email := config.Mailer.Render(template)
	email.TO = []mail.Address{mail.Address{Address: "thiti.wa@coolbrobkk.com"}}
	email.From = &mail.Address{Address: "sales@coolbrobkk.com"}
	email.Subject = "Test Preorder"
	config.Mailer.Send(email)
}

func storeAddressInCookie(cart *Cart, w http.ResponseWriter, req *http.Request) {
	cookieStorer := config.GetCookieStorer(w, req)
	cookieStorer.Put("customer_name", cart.Name)
	cookieStorer.Put("customer_address", cart.Address)
	cookieStorer.Put("customer_email", cart.Email)
	cookieStorer.Put("customer_phone", cart.Phone)
}

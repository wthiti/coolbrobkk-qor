package controllers

import (
	"net/http"

	"bitbucket.org/wthiti/coolbrobkk-qor/app/models"
	"bitbucket.org/wthiti/coolbrobkk-qor/config/view"
)

var tempProduct []models.Product

func init() {
	tempProduct = make([]models.Product, 2)
	tempProduct[0] = models.Product{
		ID:   1,
		Name: "ORIGINAL BLACK",
		Description: `<ul><li>INGREDIENTS : WATER, COFFEE GROUND</li>
									<li>CALORIES ~ ZERO</li>
									<li>CAFFEINE ~ 240 MG/BOTTLE</li></ul>`,
		Image: "/public/assets/img/black_mini.jpg",
		Price: 100,
	}
	tempProduct[1] = models.Product{
		ID:   2,
		Name: "HONEY MILK",
		Description: `<ul><li>INGREDIENTS : WATER, COFFEE GROUND, MILK, HONEY</li>
									<li>CALORIES ~ 47  CAL/100ML</li>
									<li>CAFFEINE ~ 157 MG/BOTTLE</li></ul>`,
		Image: "/public/assets/img/white_mini.jpg",
		Price: 100,
	}
}

func GetTempProductList() []models.Product {
	return tempProduct
}

func ProductIndex(w http.ResponseWriter, req *http.Request) {
	var (
		// db = utils.GetDB(req)
		v = view.ViewValueMap()
	)
	v["product_list"] = tempProduct
	view.View.Execute("product/product_index", v, req, w)
}

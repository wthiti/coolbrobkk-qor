package controllers

import (
	"net/http"

	"bitbucket.org/wthiti/coolbrobkk-qor/app/models"
	"bitbucket.org/wthiti/coolbrobkk-qor/config/view"
)

var tempProduct []models.Product

func init() {
	tempProduct = make([]models.Product, 4)
	tempProduct[0] = models.Product{
		ID:          1,
		Name:        "Test1",
		Description: "Description1",
		Image:       "Image1",
		Price:       100,
	}
	tempProduct[1] = models.Product{
		ID:          2,
		Name:        "Test2",
		Description: "Description2",
		Image:       "Image2",
		Price:       200,
	}
	tempProduct[2] = models.Product{
		ID:          3,
		Name:        "Test3",
		Description: "Description3",
		Image:       "Image3",
		Price:       300,
	}
	tempProduct[3] = models.Product{
		ID:          4,
		Name:        "Test4",
		Description: "Description4",
		Image:       "Image4",
		Price:       400,
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

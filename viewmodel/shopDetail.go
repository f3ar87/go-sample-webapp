package viewmodel

import (
	"github.com/f3ar87/go-sample-webapp/model"
)

//ShopDetail test comment
type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

//NewShopDetail test comment
func NewShopDetail(products []model.Product) ShopDetail {
	result := ShopDetail{
		Title:    "Lemonade Stand Supply",
		Active:   "shop",
		Products: []Product{},
	}
	for _, p := range products {
		result.Products = append(result.Products, productToVM(p))
	}
	return result
}

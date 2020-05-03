package controllers

import (
	"net/http"
	"regexp"
	"strconv"
	"text/template"

	"github.com/f3ar87/go-sample-webapp/model"
	"github.com/f3ar87/go-sample-webapp/viewmodel"
)

type shop struct {
	shopTemplate     *template.Template
	categoryTemplate *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
	http.HandleFunc("/shop/", s.handleShop)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	categoryPattern, _ := regexp.Compile(`/shop/(\d+)`)
	matches := categoryPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		categoryID, _ := strconv.Atoi(matches[1])
		s.handleCategory(w, r, categoryID)
	} else {
		categories := model.GetCategories()
		vm := viewmodel.NewShop(categories)
		s.shopTemplate.Execute(w, vm)
	}
}

func (s shop) handleCategory(w http.ResponseWriter, r *http.Request, categoryID int) {
	products := model.GetProductsForCategory(categoryID)
	vm := viewmodel.NewShopDetail(products)
	s.categoryTemplate.Execute(w, vm)
}

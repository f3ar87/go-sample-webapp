package controllers

import (
	"net/http"
	"text/template"
)

var (
	homeController home
	shopController shop
)

//Startup test comment
func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	shopController.shopTemplate = templates["shop.html"]
	shopController.categoryTemplate = templates["shop_details.html"]
	homeController.registerRoutes()
	shopController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}

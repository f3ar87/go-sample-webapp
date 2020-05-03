package viewmodel

import (
	"fmt"

	"github.com/f3ar87/go-sample-webapp/model"
)

//Shop is a test comment
type Shop struct {
	Title      string
	Active     string
	Categories []Category
}

//Category is a test comment
type Category struct {
	URL           string
	ImageURL      string
	Title         string
	Description   string
	IsOrientRight bool
}

//NewShop is a test comment
func NewShop(categories []model.Category) Shop {
	result := Shop{
		Title:  "Lemonade Stand Supply - Shop",
		Active: "shop",
	}
	result.Categories = make([]Category, len(categories))
	for i := 0; i < len(categories); i++ {
		vm := categorytoVM(categories[i])
		vm.IsOrientRight = i%2 == 1
		result.Categories[i] = vm
	}
	return result
}

func categorytoVM(c model.Category) Category {
	return Category{
		URL:         fmt.Sprintf("/shop/%v", c.ID),
		ImageURL:    c.ImageURL,
		Title:       c.Title,
		Description: c.Description,
	}
}

package viewmodel

import (
	"fmt"

	"github.com/bokjo/go-web-app/webapp/model"
)

// Shop struct
type Shop struct {
	Title      string
	Active     string
	Categories []Category
}

// Category struct
type Category struct {
	URL           string
	ImageURL      string
	Title         string
	Description   string
	IsOrientRight bool
}

// NewShop constructor function
func NewShop(categories []model.Category) Shop {

	result := Shop{
		Active: "shop",
		Title:  "Lemonade Stand Supply - Shop",
	}

	result.Categories = make([]Category, len(categories))

	for i, c := range categories {
		vm := categorytoVM(c)
		vm.IsOrientRight = i%2 == 1
		result.Categories = append(result.Categories, vm)
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

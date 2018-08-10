package viewmodel

import (
	"github.com/bokjo/go-web-app/webapp/model"
)

// ShopDetail struct
type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

// NewShopDetail - () - function
func NewShopDetail(products []model.Product) ShopDetail {

	result := ShopDetail{
		Active:   "shop",
		Title:    "Lemonade Stand Supply",
		Products: []Product{},
	}

	for _, p := range products {
		result.Products = append(result.Products, productToVM(&p))
	}

	return result
}

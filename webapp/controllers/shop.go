package controllers

import (
	"html/template"
	"net/http"

	"github.com/bokjo/go-web-app/webapp/viewmodel"
)

type shop struct {
	shopTemplate *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	shopVM := viewmodel.NewShop()

	s.shopTemplate.Execute(w, shopVM)
}

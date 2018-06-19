package controllers

import (
	"html/template"
	"net/http"

	"github.com/bokjo/go-web-app/webapp/viewmodel"
)

type home struct {
	homeTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	homeVM := viewmodel.NewBase()

	h.homeTemplate.Execute(w, homeVM)
}

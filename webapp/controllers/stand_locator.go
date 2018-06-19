package controllers

import (
	"html/template"
	"net/http"

	"github.com/bokjo/go-web-app/webapp/viewmodel"
)

type standLocator struct {
	standLocatorTemplate *template.Template
}

func (sl standLocator) registerRoutes() {
	http.HandleFunc("/stand_locator", sl.handleStandLocator)
}

func (sl standLocator) handleStandLocator(w http.ResponseWriter, r *http.Request) {
	standLocatorVM := viewmodel.NewStandLocator()

	sl.standLocatorTemplate.Execute(w, standLocatorVM)
}

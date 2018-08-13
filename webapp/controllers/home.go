package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bokjo/go-web-app/webapp/viewmodel"
)

type home struct {
	homeTemplate  *template.Template
	loginTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	homeVM := viewmodel.NewBase()
	w.Header().Add("Content-Type", "text/html")

	// TEST middleware timeout
	// time.Sleep(4 * time.Second)

	h.homeTemplate.Execute(w, homeVM)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	loginVM := viewmodel.NewLogin()

	if r.Method == http.MethodPost {
		err := r.ParseForm()

		if err != nil {
			log.Println(fmt.Errorf("Error login in: %v", err))
		}

		// SECURITY 101
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		if email == "admin@admin.com" && password == "password" {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			loginVM.Email = email
			loginVM.Password = password
		}

	}

	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, loginVM)
}

package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bokjo/go-web-app/webapp/model"
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

	// ###########################################################
	// ### Removed due to pusher http2 race condition and panicing
	// ###########################################################

	// if pusher, ok := w.(http.Pusher); ok {
	// 	pusher.Push("/css/app.css", &http.PushOptions{
	// 		Header: http.Header{"Content-Type": []string{"text/css"}},
	// 	})
	// }

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

		if user, err := model.Login(email, password); err == nil {
			log.Printf("User has logged in: %#v!\n", user)
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		} else {
			log.Printf("Failed to login the user with email: %v\n[Error] %v\n", email, err)
			loginVM.Email = email
			loginVM.Password = password
		}

	}

	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, loginVM)
}

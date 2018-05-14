package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	templates := populateTemplates()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]

		t := templates.Lookup(requestedFile + ".html")

		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.Handle("/img/", http.FileServer(http.Dir("../public")))
	http.Handle("/css/", http.FileServer(http.Dir("../public")))

	http.ListenAndServe(":1234", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	basePath := "../templates"

	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}

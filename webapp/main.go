package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bokjo/go-web-app/webapp/controllers"
	"github.com/bokjo/go-web-app/webapp/middleware"
	"github.com/bokjo/go-web-app/webapp/model"

	_ "github.com/lib/pq"
)

func main() {
	templates := populateTemplates()

	db := connectToDB()
	defer db.Close()

	controllers.Startup(templates)

	http.ListenAndServeTLS(":1234", "cert.pem", "key.pem", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

func connectToDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/go-web-app?sslmode=disable")

	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to the database: %v", err))
	}
	model.SetDatabase(db)
	return db
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "../templates"

	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))

	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Cannot open the content template directory: " + err.Error())
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed the read the contents of the content directory: " + err.Error())
	}

	for _, file := range files {
		f, err := os.Open(basePath + "/content/" + file.Name())
		if err != nil {
			panic("Failed to open template: '" + file.Name() + "'")
		}

		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read the content from the template file: '" + file.Name() + "'")
		}

		f.Close()

		temp := template.Must(layout.Clone())
		_, err = temp.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + file.Name() + "' as template")
		}
		result[file.Name()] = temp
	}

	return result
}

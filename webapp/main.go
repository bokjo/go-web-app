package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {

	f, err := os.Open("../public" + r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	defer f.Close()

	var contentType string

	switch {
	case strings.HasSuffix(r.URL.Path, "css"):
		contentType = "text/css"
	case strings.HasSuffix(r.URL.Path, "html"):
		contentType = "text/html"
	case strings.HasSuffix(r.URL.Path, "png"):
		contentType = "image/png"
	default:
		contentType = "text/plain"
	}
	w.Header().Add("Content-Type", contentType)
	io.Copy(w, f)

}

func main() {
	http.HandleFunc("/", serveStatic)
	http.ListenAndServe(":1234", nil)
}

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from GO!"))
	})

	http.ListenAndServe(":1234", nil)
}

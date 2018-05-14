package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	templateString := `Lemonade Stand Supply Co.`
	//http.ListenAndServe(":1234", http.FileServer(http.Dir("../public")))

	t, err := template.New("template").Parse(templateString)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}

}

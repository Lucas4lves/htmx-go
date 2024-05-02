package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/home" {
		templ := template.Must(template.ParseFiles("./index.html"))
		w.Header().Set("Content-type", "html")
		templ.Execute(w, nil)
		w.Write([]byte("<p>Welcome to HTMX Jungle</p>"))
		fmt.Printf("[%s] : %s \n", r.URL.Path, r.Method)
	}
}

func main() {
	fmt.Println("HTMX GO SERVER")

	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-type", "text/plain")
	// 	w.Write([]byte("Welcome to HTMX Jungle"))
	// }

	http.HandleFunc("/home", RequestHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

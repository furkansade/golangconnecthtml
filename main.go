package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", Home)
	http.ListenAndServe(":9090", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("index.html")
	view.Execute(w, nil)
}

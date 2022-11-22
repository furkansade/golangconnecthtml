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

	data := make(map[string]interface{})
	data["test0"] = "This is a test0!"                                 // string
	data["test1"] = 111                                                // int
	data["test2"] = []string{"Ankara", "London", "New York", "Berlin"} // array

	view.Execute(w, data)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "got foo request")
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"subst": "injectedval"}
    // the template must have {{define "main"}} in it for this to work
	templates.ExecuteTemplate(w, "main", data)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

var templates = template.Must(template.ParseGlob(`/tmp/templates/*`))

func main() {
	fmt.Printf("hello world")
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/static/", staticHandler)
	http.ListenAndServe(":9876", nil)
}

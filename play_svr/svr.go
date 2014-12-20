package main

import (
	"fmt"
	"html/template"
	"net/http"
    "os"
    "path"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "got foo request")
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"subst": "injectedval"}
	htmlTemplate.Execute(w, data)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

var templateFile string
var htmlTemplate *template.Template

func main() {
    curWd, _ := os.Getwd()
    templateFile = path.Join(curWd, "templates", "foo.html")
    htmlTemplate = template.Must(template.ParseFiles(templateFile))

	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/static/", staticHandler)
	http.ListenAndServe(":9876", nil)
}

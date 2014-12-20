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
	// the template must have {{define "main"}} in it for this to work
	htmlTemplates.ExecuteTemplate(w, "main", data)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

var templateDir string
var htmlTemplates *template.Template

func main() {
    curWd, _ := os.Getwd()
    templateDir = path.Join(curWd, "templates")
    fmt.Printf("Using html from <%v\n>", templateDir)
    htmlTemplates = template.Must(template.ParseGlob(templateDir))

	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/static/", staticHandler)
	http.ListenAndServe(":9876", nil)
}

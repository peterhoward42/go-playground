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
	t := template.Must(template.New("myTemplate").Parse(HTML_TEMPLATE))
	data := map[string]string{"subst": "injectedval"}
	t.Execute(w, data)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

const HTML_TEMPLATE = `
    <html>
        <body>
            this is the template param with <b>{{.subst}}</b>
        </body>
    </html>
`

func main() {
	fmt.Printf("hello world")
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/static/", staticHandler)
	http.ListenAndServe(":9876", nil)
}

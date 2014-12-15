package main

import (
	"fmt"
	"net/http"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "got foo request")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "got bar request")
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	fmt.Printf("hello world")
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/static/", staticHandler)
	http.ListenAndServe(":9876", nil)
}

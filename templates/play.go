package main

import (
	"os"
	"text/template"
)

type Thing struct {
	D string
}

func main() {
	canonical()
	ifelse()
}

func canonical() {
	t := template.Must(template.New("foo").Parse(
		"this is a directive {{.D}} postscript\n"))
	d := Thing{"here I am"}
	err := t.Execute(os.Stdout, d)
	if err != nil {
		panic(err)
	}
}

func ifelse() {
	t := template.Must(template.New("foo").Parse(
		"{{if .D}}dTrue{{else}}dFalse{{end}}\n"))
	d := Thing{"here I am"}
	err := t.Execute(os.Stdout, d)
	if err != nil {
		panic(err)
	}
}

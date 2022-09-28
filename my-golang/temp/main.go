package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", Gopher)
	http.ListenAndServe(":8080", nil)
}

var tmpl = template.Must(template.New("msg").
	Parse("<html><body>{{.}}さんの運勢は「大吉」です！</body></html>"))

func Gopher(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, r.FormValue("p"))
}

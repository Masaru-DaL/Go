package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct { // テンプレート展開用のデータ構造
	Title String
	Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World.", 1}                 // テンプレート用のデータ
	tmpl, err := template.ParseFiles("layout.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/txt", viewHandler)
	http.ListenAndServe(":8080", nil)

	http.Handle("/", String("Hello World."))
	http.ListenAndServe("localhost:8000", nil)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

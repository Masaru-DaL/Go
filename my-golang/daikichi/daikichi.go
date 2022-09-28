package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Gopher)
	http.ListenAndServe(":8080", nil)
}

// Gopherハンドラ
func Gopher(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.FormValue("p"), "さんの運勢は「大吉」です！")
	fmt.Fprint(w, r.FormValue("p"), "さんの運勢は「大吉」です！")
}

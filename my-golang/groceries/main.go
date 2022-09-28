package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/allgroceries", AllGroceries)
	r.HandleFunc("/groceries/{name}", SingleGrocery)
	r.HandleFunc("/groceries", GroceriesToBuy).Methods("POST")
	r.HandleFunc("/groceries/{name}", UpdateGrocery).Methods("PUT")
	r.HandleFunc("/groceries/{name}", DeleteGrocery).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", r))
}

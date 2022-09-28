package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/sample1", Sample1)
    http.ListenAndServe(":8080", nil)
}

// Sample1 ハンドラ
func Sample1(w http.ResponseWriter, r *http.Request) {
    fmt.Println("hoge", r.FormValue("hoge"))
    fmt.Println("foo", r.FormValue("foo"))
}

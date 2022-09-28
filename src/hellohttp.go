package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r * http.Request)  {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

/* 実行結果 */
// terminalは起動中
// http://localhost:8080にアクセスすると
// Hello, HTTPサーバと出ている。

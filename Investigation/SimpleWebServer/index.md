# Golang: シンプルなWebサーバを立てる

参考: [Go言語でhttpサーバーを立ち上げてHello Worldをする](https://qiita.com/taizo/items/bf1ec35a65ad5f608d45)

## 1. Hello World

```go: main.go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

/* Hello, World */
```

## 2. ServeHTTP

```go: main.go
package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.Handle("/", String("Hello World."))
	http.ListenAndServe("localhost:8000", nil)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

/* Hello World. */
```

`type String string`で型を定義し、これに対してServeHTTPメソッドを定義している。
   いまいちよくわからなかったので調べてみると、

> Goではあらゆる非インタフェース型の定義に対し、メソッドを追加することができる。

   ということのようです。[参考](https://zenn.dev/skonb/articles/0bad1d59371d09#:~:text=Go%E3%81%A7%E3%81%AF%E3%81%82%E3%82%89%E3%82%86%E3%82%8B%E9%9D%9E%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%83%95%E3%82%A7%E3%83%BC%E3%82%B9%E5%9E%8B%E3%81%AE%E5%AE%9A%E7%BE%A9%E3%81%AB%E5%AF%BE%E3%81%97%E3%83%A1%E3%82%BD%E3%83%83%E3%83%89%E3%82%92%E8%BF%BD%E5%8A%A0%E3%81%99%E3%82%8B%E3%81%93%E3%81%A8%E3%81%8C%E3%81%A7%E3%81%8D%E3%81%BE%E3%81%99%E3%80%82)

また、[こちら](https://teratail.com/questions/224653)から、

> Goのhttpサーバー実装はHTTPリクエストのURLにしたがってhttp.Handlerインターフェースと互換のあるインスタンスのServeHTTPメソッドを呼ぶ仕掛けになっています。

ServeHTTPメソッドを持たせたString型は、Handleの第2引数に登録する事が出来るようになり、リクエストのURLに従って関数が実行される。というようなことだと思われます。

## 3. text/templateを使う

必ずしも`ServeHTTP`を既存の型などに定義しなくても良い。
`http.ResponseWriter`と、`*http.Request`を引数に取る関数を用意すれば**ハンドラとして登録することができる**。

```go:
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
	page := Page{"Hello World.", 1}                                       // テンプレート用のデータ(雛形)

  /* template.new(<テンプレート名>).Parse(<文字列>) */
	tmpl, err := template.New("new").Parse("{{.Title}} {{.Count}} count") // テンプレート文
	if err != nil {
		panic(err)
	}

  /* Execute(io.Writer(出力先), <データ>) */
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

/* Hello World. 1 count */
```

text/templateの使い方がいまいち分からなかったのでそこから。

1. templateの雛形の作成
2. template.Newメソッドでテンプレート名を指定
3. template.Parseメソッドでテンプレート内容を読み込ませる
4. template.Executeメソッドで出力するために、出力先と出力する値を指定する
5. `{{.Title}} {{.Count}} count`このプレースホルダに1のテンプレートのTitleとCountが入ってくる
6. 出力 -> Hello World. 1 count

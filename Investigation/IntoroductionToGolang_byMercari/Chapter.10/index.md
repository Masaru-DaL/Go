- [メルカリ作のプログラミング言語Go完全入門 読破](#メルカリ作のプログラミング言語go完全入門-読破)
- [10. HTTPサーバとクライアント](#10-httpサーバとクライアント)
  - [10-1. HTTPサーバを立てる](#10-1-httpサーバを立てる)
      - [10-1-1. net/httpパッケージ](#10-1-1-nethttpパッケージ)
      - [10-1-2. HTTPサーバの作成の流れ](#10-1-2-httpサーバの作成の流れ)
      - [10-1-3. TRY Hello, HTTPサーバ](#10-1-3-try-hello-httpサーバ)
      - [10-1-4. HTTPハンドラ](#10-1-4-httpハンドラ)
# メルカリ作のプログラミング言語Go完全入門 読破
# 10. HTTPサーバとクライアント
## 10-1. HTTPサーバを立てる
- HTTPリクエスト
  - クライアント -> サーバ
  - GET, POSTなど

- HTTPレスポンス
  - サーバ -> クライアント
  - 200(OK), 404(NotFound)など

#### 10-1-1. net/httpパッケージ
- net/httpパッケージの機能
  - HTTP関連の型や関数を提供する

- HTTPサーバ
  - クライアントのリクエストに応じてレスポンスを返す
  - HTMPを返せばブラウザで表示できる

- HTTPクライアント
  - HTTPサーバにリクエストを送り、レスポンスを受け取る
    - HTTPを採用しているサーバであればOK
      - Goで書いてある必要性はない

#### 10-1-2. HTTPサーバの作成の流れ
1. HTTPハンドラの作成
1つのHTTPリクエストを処理する関数(->ハンドラ)を作成する
1リクエストあたりの処理は並行に動く

2. HTTPハンドラとエントリポイントの結び付け
"/index"などのエントリポイントとハンドラを結びつける
ルーティングとも呼ばれる

3. HTTPサーバの起動
ホスト名(IP)とポート番号、ハンドラを指定してHTTPサーバを起動する

#### 10-1-3. TRY Hello, HTTPサーバ
```go:
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
```

#### 10-1-4. HTTPハンドラ
[http.Handler](https://qiita.com/nirasan/items/2160be0a1d1c7ccb5e65#httphandler-%E3%81%A8%E3%81%AF)
```go:
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```
> ServeHTTP関数を持つだけのインターフェース
> HTTPリクエストを受け、レスポンスを返すことを責務とする

- 引数にレスポンスを書き込む先とリクエストを取る
  - 第1引数 -> レスポンスを書き込む先
    - fmtパッケージの関数で書き込みを行える
  - 第2引数 -> クライアントからのリクエスト
```go:
/*  */
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello, HTTPサーバ")
}
```
参考: [ハンドラによるレスポンス返却の詳細｜Deep Dive into The Go's Web Server](https://zenn.dev/hsaki/books/golang-httpserver-internal/viewer/httphandler#%E3%83%8F%E3%83%B3%E3%83%89%E3%83%A9%E9%96%A2%E6%95%B0%E3%81%AE%E3%81%8A%E3%81%95%E3%82%89%E3%81%84)

- 変数wに`http.ResponseWriter`を指定
ここに書き込まれた内容はネットワークを通じてレスポンスへの書き込みとなる。
-> `Fprint`で書き込み先(w)を指定している。


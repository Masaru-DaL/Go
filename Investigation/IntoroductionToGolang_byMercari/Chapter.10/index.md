- [メルカリ作のプログラミング言語Go完全入門 読破](#メルカリ作のプログラミング言語go完全入門-読破)
- [10. HTTPサーバとクライアント](#10-httpサーバとクライアント)
  - [10-1. HTTPサーバを立てる](#10-1-httpサーバを立てる)
      - [10-1-1. net/httpパッケージ](#10-1-1-nethttpパッケージ)
      - [10-1-2. HTTPサーバの作成の流れ](#10-1-2-httpサーバの作成の流れ)
      - [10-1-3. TRY Hello, HTTPサーバ](#10-1-3-try-hello-httpサーバ)
      - [10-1-4. HTTPハンドラ](#10-1-4-httpハンドラ)
      - [10-1-5. http.Handleでハンドラを登録](#10-1-5-httphandleでハンドラを登録)
      - [10-1-6. http.HandleFuncでハンドラを登録](#10-1-6-httphandlefuncでハンドラを登録)
      - [10-1-7. http.HandlerFuncとは](#10-1-7-httphandlerfuncとは)
      - [10-1-8. Handler と Handle](#10-1-8-handler-と-handle)
      - [10-1-9. http.ServeMux](#10-1-9-httpservemux)
      - [10-1-10. HTTPサーバの起動](#10-1-10-httpサーバの起動)
  - [10-2. レスポンスとリクエスト](#10-2-レスポンスとリクエスト)
      - [10-2-1. http.ResponseWriterについて](#10-2-1-httpresponsewriterについて)
      - [10-2-2. エラーを返す](#10-2-2-エラーを返す)
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
ス
#### 10-1-5. http.Handleでハンドラを登録
[http.Handle](https://qiita.com/nirasan/items/2160be0a1d1c7ccb5e65#httphandle-%E3%81%A8%E3%81%AF) -> 表示するURLと、URLに対応する`http.Handler`を`DefaultServeMux`に登録する関数

**パターン**と**http.Handler**を指定して登録する
- 第1引数 -> パターンを指定する
- 第2引数 -> http.Handlerを指定する
- 結果 -> http.DefaultServeMuxに登録される
`func Handle(pattern string, <handler Handler>)`
<handler Handler>にはServeHTTPメソッドを持つ型の具体的な値がくる。-> **ServeHTTPメソッドを持つ型がハンドラとして扱われる**

- 具体例
```go:
type AnyHandler struct {}
func (a *AnyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
anyHandler := &AnyHandler{}

/* http.Handleの引数の具体例 */
http.Handle("/any/", anyHandler)
http.ListenAndServe(":8080", nil)
```

#### 10-1-6. http.HandleFuncでハンドラを登録
[http.HandleFunc](https://qiita.com/nirasan/items/2160be0a1d1c7ccb5e65#httphandlefunc-%E3%81%A8%E3%81%AF) -> URLと`func(ResponseWriter, *Request)`を渡して`DefaultServeMux`に登録する関数
内部で`func(ResponseWriter, *Request)`から`http.HandleFunc`へのキャストが行われている。

**パターン**と**関数**を指定して登録する
- 第1引数としてパターンを指定する
- 第2引数として関数を指定する
- 結果 -> http.DefaultServeMuxに登録される
`func HandleFunc(pattern string, handler func(ResponseWriter, *Request))`
"handler func(ResponseWriter, *Request)" -> http.HandlerのServeHTTPメソッドと同じ引数の関数

#### 10-1-7. http.HandlerFuncとは
[http.HandlerFunc](https://qiita.com/nirasan/items/2160be0a1d1c7ccb5e65#httphandlerfunc-%E3%81%A8%E3%81%AF)
- `func(ResponseWriter, *Request)`の別名の型
- `ServeHTTP`関数を持つ
関数を定義して`http.HandlerFunc`にキャストするだけで構造体を宣言することなく`http.Handler`を用意することができる。

```go:
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

#### 10-1-8. Handler と Handle
- Handler
  - **登録されるもの**

- Handle
  - **登録するもの**

#### 10-1-9. http.ServeMux
- 複数のハンドラをまとめる(登録されたものをまとめる)
- パス(URL)によって使うハンドラを切り替える
- 自信もhttp.Handlerを実装している
- http.Handleとhttp.HandleFuncはデフォルトのhttp.ServeMuxであるhttp.DefaultServeMuxを使用している

#### 10-1-10. HTTPサーバの起動
http.ListenAndServeを使う
`http.ListenAndServe(":8080", nil)`
- 第1引数 -> ホスト名とポート番号を指定
  - ホスト名を省略した場合localhost
- 第2引数でHTTPハンドラを指定
  - nilで省略した場合はhttp.HandleFuncなどで登録したハンドラが使用される

## 10-2. レスポンスとリクエスト
#### 10-2-1. http.ResponseWriterについて
- io.Writerと同じWriteメソッドを持つ
  - ResponseWriterを満たすとio.Writerを満たす
- io.Writerとしても振る舞える
  - fmt.Fprint*の引数に取れる
  - json.NewEncoderの引数に取れる

**インタフェースなのでモックも作りやすい = テストが簡単**

#### 10-2-2. エラーを返す
http.Error関数を使う
`func Error(w ResponseWriter, error string, code int) {}`

[func Error](https://pkg.go.dev/net/http#Error:~:text=application/octet%2Dstream%22.-,func%20Error%20%C2%B6,-func%20Error(w)
> Errorは指定されたエラーメッセージとHTTPコードでリクエストに応答する。

error string -> エラーメッセージの指定
code int -> ステータスコードの指定
ステータスコードは定義済み定数を使用する。



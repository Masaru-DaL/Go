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
      - [10-2-3. JSONを返す](#10-2-3-jsonを返す)
      - [10-2-4. JSONデコード](#10-2-4-jsonデコード)
      - [10-2-5. レスポンスヘッダーを設定する](#10-2-5-レスポンスヘッダーを設定する)
      - [10-2-6. リクエストパラメタの取得](#10-2-6-リクエストパラメタの取得)
      - [10-2-7. リクエストボディの取得](#10-2-7-リクエストボディの取得)
      - [10-2-8. リクエストヘッダー](#10-2-8-リクエストヘッダー)
      - [10-2-9. TRY リクエストパラメタの使用](#10-2-9-try-リクエストパラメタの使用)
      - [10-2-10. テンプレートエンジンの使用](#10-2-10-テンプレートエンジンの使用)
      - [10-2-11. TRY テンプレートエンジンの使用](#10-2-11-try-テンプレートエンジンの使用)
      - [10-2-12. リダイレクト](#10-2-12-リダイレクト)
      - [10-2-13. ミドルウェアを作る](#10-2-13-ミドルウェアを作る)
  - [10-3. HTTPハンドラのテスト](#10-3-httpハンドラのテスト)
      - [10-3-1. ハンドラのテストの例](#10-3-1-ハンドラのテストの例)
  - [10-4. HTTPクライアント](#10-4-httpクライアント)
      - [10-4-1. HTTPリクエストを送る](#10-4-1-httpリクエストを送る)
      - [10-4-2. レスポンスを読み取る](#10-4-2-レスポンスを読み取る)
      - [10-4-3. リクエストを指定する](#10-4-3-リクエストを指定する)
      - [10-4-4. リクエストとコンテキスト](#10-4-4-リクエストとコンテキスト)
      - [10-4-5. http.Clientとhttp.Transport](#10-4-5-httpclientとhttptransport)
      - [10-4-6. http.DefaultTransport](#10-4-6-httpdefaulttransport)
      - [10-4-7. http.RoundTripperを実装する場合のTIPS](#10-4-7-httproundtripperを実装する場合のtips)
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

#### 10-2-3. JSONを返す
encoding/jsonパッケージを使う
- 機械的に処理しやすいJSONをレスポンスに用いる場合も多い
- JSONエンコーダを使ってGoの値をJSONに変換する
- 構造体やスライスをJSONのオブジェクトや配列にできる

```mermaid
graph LR
  A[サーバ] --> B[Goのデータ構造]
  B -->|エンコード|C[JSON]
  C -->|デコード|D[JavaやSwiftのデータ構造]
  D --> E
  A -->|レスポンス|E[クライアント]
  E -->|リクエスト|A
```

#### 10-2-4. JSONデコード
json.Decoder型を使う

```go:
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := &Person{Name: "tenntenn", Age: 31}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf) // jsonに変換
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	var p2 Person
	dec := json.NewDecoder(&buf) // jsonから変換
	if err := dec.Decode(&p2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p2)
}

/* 実行結果 */
/*
{"name":"tenntenn","age":31}

{tenntenn 31}
*/
```

#### 10-2-5. レスポンスヘッダーを設定する
[引用: レスポンスヘッダー image](https://itsakura.com/wp-content/uploads/2017/11/http-response1.svg)
ResponseWriterのHeaderメソッドを使う
- WriteやWriteHeaderを呼び出した後に設定しても効果がない

```go:
func handler(w http.ResponseWriter, req *http.Request) {
  /* 設定 */
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	v := struct {
		Msg string `json:"msg"`
	}{
		Msg: "hello",
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error:", err)
	}
}
```

#### 10-2-6. リクエストパラメタの取得
```go:
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello", r.FormValue("msg"))
})
```
`r.FromValue("msg")` r = `*http.Request`
(*http.Request).FormValueから名前を指定して取得する。
-> `http://localhost:8080?msg=Gophers`のようにパラメタを指定する。
複数ある場合は`&`でつなぐ。
`http://localhost:8080?a=100&b=200`

#### 10-2-7. リクエストボディの取得
  [引用: リクエストボディ image](https://itsakura.com/wp-content/uploads/2017/11/http-request-post1.svg)
```go:
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
var p Person
dec := json.NewDecoder(r.Body)
if err := dec.Decode(&p); err != nil {
	// ... エラー処理 …
}
fmt.Println(p)
})
```
`r.Body` -> (*http.Request).Body
io.Readerを実装している

#### 10-2-8. リクエストヘッダー
[引用: リクエストヘッダー(Get) image](https://itsakura.com/wp-content/uploads/2017/11/http-request-get1.svg)
```go:
func handler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	fmt.Fprintln(w, contentType)
}
```
`req.Header.Get("Content-Type")`
-> (*http.Request).Header.Get(<ヘッダー名>)
Getメソッドを使うとヘッダー名を指定して取得できる

#### 10-2-9. TRY リクエストパラメタの使用
```go:
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

/* 実行結果 */
// fmt.Println -> 標準出力に出力
// fmt.Fprint -> レスポンスへの書き込み(ブラウザ表示)
```

#### 10-2-10. テンプレートエンジンの使用
html/templateを使う
- Go標準のテンプレートエンジン
- text/templateのHTML特化版
[template](https://pkg.go.dev/html/template)

```go:
/* テンプレートの生成 */
// sign -> テンプレート名
template.Must(template.New("sign").Parse("<html><body>{{.}}</body></html>"))

/* ------------------------------- */
/* テンプレートに埋め込む */
tmpl.Execute(w, r.FormValue("content"))
// w -> io.Writer
// r.FormValue("content") -> リクエストから貰った値を埋め込む
```

```go: sample
var tmpl = template.Must(template.New("msg").
	Parse("<html><body>{{.}}さん こんにちは</body></html>"))

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, r.FormValue("p"))
}
```

#### 10-2-11. TRY テンプレートエンジンの使用
```go:
package main

import (
	"net/http"
	"html/template"
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
```

#### 10-2-12. リダイレクト
http.Redirect関数を使う
```go:
func handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

/* 第3引数 -> 遷移したいパス
    第4引数 -> 3xx系のステータスコード(レスポンス)*/
```

#### 10-2-13. ミドルウェアを作る
ハンドラより前に行う共通処理
ライブラリを使ってもOK
参考: [Goで始めるMiddleware - Qiita](https://qiita.com/giraffate/items/ea962f1cdad21c2f68aa)

## 10-3. HTTPハンドラのテスト
net/http/httptestを使う
- ハンドラのテストのための機能などを提供
- httptest.ResponseRecorder
  - http.ResponseWriterインタフェースを実装
- NewRequestメソッド(Go1.7以上)
  - 簡単にテスト用のリクエストが作れる

#### 10-3-1. ハンドラのテストの例
```go:
/* テスト対象 */
func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello, net/http!")
}

/* テストコード */
func TestSample(t *testing.T) {

  /* ここから */
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler(w, r)
	res := w.Result()
  /* ここまで */
	defer res.Body.Close()

  /* テスト結果によって処理を行う */
	if res.StatusCode != http.StatusOK { t.Fatal("unexpected status code") }
	b, err := ioutil.ReadAll(res.Body)
	if err != nil { t.Fatal("unexpected error") }
	const expected = "Hello, net/http!"
	if s := string(b); s != expected { t.Fatalf("unexpected response: %s", s) }
}
```

## 10-4. HTTPクライアント
参考: [Goでnet/httpを使う時のこまごまとした注意](https://qiita.com/ono_matope/items/60e96c01b43c64ed1d18)
#### 10-4-1. HTTPリクエストを送る
http.DefaultClientを用いる
http.DefaultClientはゼロ値。
http.Getやhttp.Postはhttp.DefaultClientのラッパーになる。
参考: [How to issue HTTP request](https://tutuz-tech.hatenablog.com/entry/2020/03/22/160529#How-to-issue-HTTP-request:~:text=%E5%8F%82%E8%80%83-,How%20to%20issue%20HTTP%20request,-%E3%82%AF%E3%83%A9%E3%82%A4%E3%82%A2%E3%83%B3%E3%83%88%E3%81%AE%E5%AE%9F%E8%A3%85)

#### 10-4-2. レスポンスを読み取る
(*http.Response).Bodyを使う
```go:
func main() {
resp, err := http.Get("http://example.com/")
if err != nil { /* エラー処理 */ }
defer resp.Body.Close() // 必ずcloseメソッドを呼ぶ
var p Person
dec := json.NewDecoder(resp.Body)
if err := dec.Decode(&p); err != nil {
	// ... エラー処理 …
}
fmt.Println(p)
}
```

#### 10-4-3. リクエストを指定する
http.Client.Doを用いる
```go:
req, err := http.NewRequest("GET", "http://example.com", nil)
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req) // req -> *http.Request
```

#### 10-4-4. リクエストとコンテキスト
- *http.Requestから取得する(サーバ)
`ctx := req.Context()`

- Contextを更新する(クライアント)
  - 新しい*http.Requestが生成される
`req = req.WithContext(ctx)`

#### 10-4-5. http.Clientとhttp.Transport
https://pkg.go.dev/net/http
> クライアントとトランスポートは、複数のゴールーチンによる同時使用に対して安全であり、効率のために一度だけ作成して再利用する必要がある。

- http.Transport型
  - [http.RoundTripper](https://pkg.go.dev/net/http#RoundTripper:~:text=%E3%82%BF%E3%82%A4%E3%83%97RoundTripper%20%C2%B6,-type%20RoundTripper%20interface)を実装した型
    - https://pkg.go.dev/net/http#RoundTripper
    - > リクエストに対してレスポンスを返す
    - > レスポンスを解釈しない
    - > レスポンスのHTTPステータスコードに関係なく、`err == nil`を返す必要がある
  - [http.Transport](https://pkg.go.dev/net/http#Transport:~:text=%E3%81%A6%E3%81%84%E3%81%BE%E3%81%99%0A%7D-,Transport,-%E3%81%AF%E3%80%81HTTP%E3%80%81HTTPS)
    - > Transport は、HTTP、HTTPS、および HTTP プロキシ (HTTP または CONNECT を使用した HTTPS のいずれか) をサポートする RoundTripper の実装です。
    - HTTP/HTTPS/HTTPプロキシに対応している
    - コネクションのキャッシュを行う

#### 10-4-6. http.DefaultTransport
http.ClientのTransportフィールドがnilの時に使われる
-> Clientへ明示的にRoundTripperが指定されなければ、デフォルトとして以下の[DefaultTransport](https://pkg.go.dev/net/http#RoundTripper:~:text=%E7%AC%AC%E4%BA%8C%E3%81%AB%E3%80%81%0A%7D-,DefaultTransport,-%E3%81%AF%20Transport%20%E3%81%AE)が使われる。
```go:
var DefaultTransport RoundTripper = &Transport{
	Proxy: ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}
```

#### 10-4-7. http.RoundTripperを実装する場合のTIPS
https://journal.lampetty.net/entry/mocking-http-access-in-golang
- 注意点
  - レスポンスを返す場合、エラーはnilにすること
  - リクエストは変更しない
  - リクエストは他のゴールーチンから参照される可能性があることを考慮する

- TIPS
  - 元になるhttp.RoundTripperをラップしておく
    - フィールドで設定できるようにしておく
    - HTTP通信の部分は親のRoundTripメソッドを呼ぶ
    - フィールドがnilの場合はhttp.DefaultTransportを使う


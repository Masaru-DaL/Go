- [2. Hello World](#2-hello-world)
    - [2-1. Introduction](#2-1-introduction)
    - [2-2. Registering a Request Handler: リクエストハンドラの登録](#2-2-registering-a-request-handler-リクエストハンドラの登録)
    - [2-3. Listen for HTTP Connections: サーバをリッスン状態にする](#2-3-listen-for-http-connections-サーバをリッスン状態にする)
    - [2-4. The Code (for copy/paste):](#2-4-the-code-for-copypaste)
- [3. HTTP Server](#3-http-server)
    - [3-1. Introduction](#3-1-introduction)
    - [3-2. Process dynamic requests: 動的リクエストの処理](#3-2-process-dynamic-requests-動的リクエストの処理)
    - [3-3. Serving static assets: 静的アセットの提供](#3-3-serving-static-assets-静的アセットの提供)
    - [3-4. Accept connections: サーバをリッスン状態にする](#3-4-accept-connections-サーバをリッスン状態にする)
    - [3-5. The Code(for copy/paste)](#3-5-the-codefor-copypaste)
- [4. Routing(using gorilla/mux)](#4-routingusing-gorillamux)
    - [4-1. Introduction](#4-1-introduction)
    - [4-2. Installing the gorilla/mux package](#4-2-installing-the-gorillamux-package)
    - [4-3. Create a new Router](#4-3-create-a-new-router)
    - [4-4. Registering a Request Handler](#4-4-registering-a-request-handler)
    - [4-5. URL Parameters](#4-5-url-parameters)
    - [4-6. Setting the HTTP server's router](#4-6-setting-the-http-servers-router)
    - [4-7. The Code (for copy/paste)](#4-7-the-code-for-copypaste)
- [5. Building a Simple REST API in Go With Gorilla/Mux](#5-building-a-simple-rest-api-in-go-with-gorillamux)
    - [5-1. 概要](#5-1-概要)
    - [5-2. ディレクトリ構造とファイル概要](#5-2-ディレクトリ構造とファイル概要)
    - [5-3. Gorilla/muxのインストール](#5-3-gorillamuxのインストール)
    - [5-4. grocery.go](#5-4-grocerygo)
    - [5-5. main.go](#5-5-maingo)
    - [5-6. handler.go](#5-6-handlergo)

### 1. 参考資料

[Go Web Examples - Learn Web Programming in Go by Examples](https://gowebexamples.com/)

## 2. Hello World

標準ライブラリのnet/httpパッケージを使用して、HTTPサーバを作成する。

#### 2-1. Introduction

net/httpパッケージには、HTTPプロトコルに関する全ての機能が備わっている。
サーバクライアントモデルが含まれる(他にもあるが)
この章で簡単にウェブサーバを作ることができる。

#### 2-2. Registering a Request Handler: リクエストハンドラの登録

まず、ブラウザ、HTTPクライアント(PC)、APIリクエストからのすべてのHTTP接続を受け取るハンドラを作成する。

- ハンドラ関数
`func (w http.ResponseWriter, r *http.Request)`
ハンドラ関数は2つのパラメータを受け取る。
  - `http.ResponseWriter` -> レスポンスを受け取る内容を書き込む。textまたはhtmlで受け取れる。
  - `*http.Request` -> HTTPリクエストに関する全ての情報を受け取る。例えばURLや、ヘッダーフィールドなど。

- `/`(デフォルト)にアクセスした際のハンドラ

```go:
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
})
```

#### 2-3. Listen for HTTP Connections: サーバをリッスン状態にする

リクエストハンドラだけではサーバは外部からのHTTP接続を一切受け付けません。
そのためにサーバに**ポート番号**を指定し、そのポートへの接続を受け付ける状態にします。
サーバに通信したいポート番号を登録して、サーバはそのポート番号に接続要求があると通知を受けて処理を行います。この動作の事を"**ポートをリッスンする**"と表現します。
この節で、サーバにポート80を指定してリッスン状態にします。

[ListenAndServe](https://cs.opensource.google/go/go/+/go1.19.1:src/net/http/server.go;l=3253)
`func ListenAndServe(addr string, handler Handler) error`

第1引数: ポート番号
第2引数: ハンドラ
※2目の引数にnilが渡された場合、デフォルトでDefaultServeMuxというServeMux型のハンドラが使用されます。
基本的にはnilを渡すのが正解のようですので、nilを渡します。

`http.ListenAndServe(":80", nil)`

#### 2-4. The Code (for copy/paste):

```go: hello.go
package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
  })

  http.ListenAndServe(":80", nil)
}
```

`$ go run hello.go`
`http://localhost:80`
Hello, you've requested: /

## 3. HTTP Server

Goでの基本的なHTTPサーバを作成する方法を学ぶ。
2章と違う部分は静的アセットを処理する項がある部分。
静的アセット: 画像、CSS、JSなどのリクエストよって表示する内容が変わらないもの。

#### 3-1. Introduction

HTTPサーバには、いくつかの重要な役割を担っている
- 動的なリクエストの処理(リクエストの内容はリクエストの度に違う)
  - ウェブサイトを閲覧
  - アカウントにログインする
  - 画像を投稿したりするユーザからの受信リクエスト

- 静的アセットの処理
  - JavaScript, CSS, 画像などをクライアントに送り、ユーザにダイナミックな体験を提供する

- クライアントからの接続を受け入れる
  - クライアントからリクエストを受け、レスポンスを返すためには特定のポートでリッスンする必要がある。

#### 3-2. Process dynamic requests: 動的リクエストの処理

リクエストを受け付け、処理するためのハンドラを登録する。登録するには"http.HandleFunc"関数を使用する。
第1引数: path(URL)
第2引数: 第1引数にアクセスした際に実行する関数

```go:
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome to my website!")
})
```

`http.Request`に、リクエストとそのパラメータに関する全ての情報が含まれている。
各パラメータの読み取りを行うには以下のように行う。
GETパラメータ: `r.URL.Query().Get("token")`
POSTパラメータ(HTMLフォームのフィールド): `r.FormValue("email")`

#### 3-3. Serving static assets: 静的アセットの提供

1. net/httpパッケージのhttp.FileServerを使用し、URLパスを指定する。

FileServerは以下のようになっている。
`func FileServer(root FileSystem) Handler`
引数にFileSystemを受け取る必要があり、戻り値はハンドラです。
このハンドラは、ルートにあるファイルシステムの内容をHTTPリクエストに返す。

以下のコードでファイルサーバを設置している。
`fs := http.FileServer(http.Dir("static/"))`

これはファイルシステムをDirメソッドを用いてstaticディレクトと定義している。
Dirメソッドを使用するとOSのファイルシステム実装を使用できる。(OS内のディレクトリを指定できる。)

2. ハンドラを登録するハンドルを登録する

ハンドラは、DefaultServeMuxのパターンとしてハンドラを登録します。
正しくファイルを提供するために、StripPrefixメソッドを使用してURLパスの一部を削除する必要がある。

`http.Handle("/static/, http.StripPrefix("/static/), fs)`

#### 3-4. Accept connections: サーバをリッスン状態にする

`http.ListenAndServe(":80", nil)`

#### 3-5. The Code(for copy/paste)

```go:
package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to my website!")
  })

  fs := http.FileServer(http.Dir("static/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.ListenAndServe(":80", nil)
}
```

## 4. Routing(using gorilla/mux)

gorilla/muxパッケージを使用し、RESTfulなサーバとのやり取りを学ぶ。

#### 4-1. Introduction

net/httpのあまり得意ではない事の1つが、リクエストURLをリクエスト内容によって分割するような複雑なリクエストルーティングである。(RESTの設計は得意ではない)
そのため、gorilla/muxパッケージを使用する。

この章では名前付きパラメータ、GET/POSTハンドラ、ドメイン制限のあるルートを作成する方法を学ぶ。

#### 4-2. Installing the gorilla/mux package

- gorilla/muxパッケージの概要
net/httpパッケージのルーティングに適用できるパッケージ。
**gorilla/muxはルーティング機能を提供する**
Webアプリケーションを書く時の生産性を上げるための機能が多く備わっている。
ミドルウェアなどの他のHTTPライブラリや既存のアプリケーションと混在させることが可能

インストールするには以下のコマンドを使用する。
`$ go get -u github.com/gorilla/mux`

#### 4-3. Create a new Router

新しいリクエストルータを作成する。
このルータは、ウェブアプリケーションのメインルータになる。
全てのHTTP接続を受信し、登録したリクエストハンドラを介してサーバにパラメータとして渡される。

新しいルータの作成するには以下のコマンドを使用する。
`r := mux.NewRouter`

#### 4-4. Registering a Request Handler

新しいルータを作成したら、通常と同じように(gorilla/muxを使用しなかった時と同じように)リクエストハンドラを登録する。
違いは`http.HandleFunc(...)`のようにhttpメソッドを呼ぶ代わりに、作成したルータ上でHandleFuncを呼ぶ所です。

`r.HandleFunc(...)`

#### 4-5. URL Parameters

gorilla/mux Routerの最大の強みは、リクエストURLからセグメント(下で説明)を抽出することが出来る点。

`/books/go-programming-blueprint/page/10`
このURLを元に理解します。
このURLには2つのダイナミック(動的な)セグメントがある。

1. go-programming-blueprint
bookから続き、本のタイトルを表すセグメント。

2. page(**10**)
本の何ページ目かを表すセグメント

リクエストハンドラがこのような動的に変わるURLを処理するためには、ダイナミックセグメントをプレースホルダに設定して、次のようにリクエストハンドラを変更する。

```go:
r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
  // get the book
  // navigate to the page
})
```

セグメントからデータを取得するには、gorilla/muxパッケージに不随するmex.Vars(r)関数を使用する。(rは作成したルータ)
これはhttp.Requestをパラメータとして受け取り、セグメントのマップを返す。

```go:
func(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  vars["title"] // the book title slug
  vars["page"] // the page
}
```

#### 4-6. Setting the HTTP server's router

前述したが、`http.ListenAndServe(":80", nil)`のnilはデフォルトでnilで、nilの場合net/httpパッケージのデフォルトルータを使用する事を意味する。
今回はgorilla/muxパッケージを用いてルータを作成しているので、作成したルータを指定する。

`http.ListenAndServe(":80":, r)`

#### 4-7. The Code (for copy/paste)

```go:
package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    title := vars["title"]
    page := vars["page"]

    fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
  })

  http.ListenAndServe(":80", r)
}
```

## 5. Building a Simple REST API in Go With Gorilla/Mux
[Building a Simple REST API in Go With Gorilla/Mux](https://betterprogramming.pub/building-a-simple-rest-api-in-go-with-gorilla-mux-892ceb128c6f)

ここまででgolangにおけるサーバ構築の基礎を学べたので、上記のサイトを参考にREST APIのサーバを構築したいと思います。

#### 5-1. 概要
Groceries(食料品)APIの概要
1. 特定の食料品とその数量の取得
2. 全ての食料品とその数量の取得
3. 食料品を投稿してそれを更新する

#### 5-2. ディレクトリ構造とファイル概要

```shell:
$ tree

```

- grocery.go: APIのモデルを定義
- handler.go: リクエストを管理する関数
- main.go: URLのpath

#### 5-3. Gorilla/muxのインストール

`go get -u github.com/gorilla/mux`

#### 5-4. grocery.go

```go: grocery.go
package main

type Grocery struct {

  Name string `json: "name"`
  Quantity int `json: "quantity"`
}
```

grocery.goにはAPIのモデルを定義する。
APIのモデルは食料品の名前を表すNameと、その数量を表すQuantityの2つのフィールドだけです。

#### 5-5. main.go

```go: main.go
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
```

- mux.NewRouter()
ルータを指定して各ハンドルを登録する。
第1引数: URL path
第2引数: 第1引数にアクセスされた際に処理する関数

- StrictSlash()
StrictSlashはデフォルトでfalseで、trueを指定すると、"/path/"とパス指定の場合に"/path"にアクセスすると前者にリダイレクトされる。逆の場合も同様。
アプリケーションには常にルートで指定されたパスが表示される。

- Methods
MethodsメソッドでHTTPメソッドを指定できる。
書かない場合はGET？
※今回はデータベースを用意しない。

- log.Fatal
サーバのログを取得し、エラーが発生した場合はメッセージエラーが発生し、プログラムが停止する。

#### 5-6. handler.go

1. func AllGroceries

```go: handler.go
package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"

  "github.com/gorilla/mux"
)

var groceries = []Grocery {
  {Name: "Almod Mild", Quantity: 2},
  {Name: "Apple", Quantity: 6},
}

func AllGroceries(w http.ResponseWriter, r *http.Request) {

  fmt.Println("Endpoint hit: returnAllGroceries")
  json.NewEncoder(w).Encode(groceries)
}
```
今回はデータベースを使用しないので、変数groceriesを定義し、2つの食料品店の情報を含む配列を代入する。
`AllGroceries`が呼び出されると、すべての食料品を含む配列がJSONとして返される。

2. func SingleGrocery

```go: handler.go
func SingleGrocery(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  name := vars["name"]

  for _, grocery := range groceries {

    if grocery.Name == name {
      json.NewEncoder(w).Encode(grocery)
    }
  }
}
```

この関数では、作成したルータから食料品の名前を取得する。
スライスを繰り返し、要求された食料品だけを返す。

3. func GroceriesToBuy

```go: handler.go
func GroceriesToBuy(w http.ResponseWriter, r *http.Request) {
  reqBody, _ := ioutil.ReadAll(r.Body)

  var grocery Grocery
  json.Unmarshal(reqBody, &grocery)
  groceries = append(groceries, grocery)

  json.NewEncoder(w).Encode(groceries)
}

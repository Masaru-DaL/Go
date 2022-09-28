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
    - [5-3. Gorilla/mux のインストール](#5-3-gorillamux-のインストール)
    - [5-4. grocery.go](#5-4-grocerygo)
    - [5-5. main.go](#5-5-maingo)
  - [5-6. handler.go](#5-6-handlergo)
    - [5-6-1. func AllGroceries](#5-6-1-func-allgroceries)
    - [5-6-2. func SingleGrocery](#5-6-2-func-singlegrocery)
    - [5-6-3. func GroceriesToBuy](#5-6-3-func-groceriestobuy)
    - [5-6-4. func DeleteGrocery](#5-6-4-func-deletegrocery)
    - [5-6-5. func UpdateGrocery](#5-6-5-func-updategrocery)

### 1. 参考資料

[Go Web Examples - Learn Web Programming in Go by Examples](https://gowebexamples.com/)

## 2. Hello World

標準ライブラリの net/http パッケージを使用して、HTTP サーバを作成する。

#### 2-1. Introduction

net/http パッケージには、HTTP プロトコルに関する全ての機能が備わっている。
サーバクライアントモデルが含まれる(他にもあるが)
この章で簡単にウェブサーバを作ることができる。

#### 2-2. Registering a Request Handler: リクエストハンドラの登録

まず、ブラウザ、HTTP クライアント(PC)、API リクエストからのすべての HTTP 接続を受け取るハンドラを作成する。

* ハンドラ関数
  `func (w http.ResponseWriter, r *http.Request)`
  ハンドラ関数は 2 つのパラメータを受け取る。

  * `http.ResponseWriter` -> レスポンスを受け取る内容を書き込む。text または html で受け取れる。
  * `*http.Request` -> HTTP リクエストに関する全ての情報を受け取る。例えば URL や、ヘッダーフィールドなど。

* `/`(デフォルト)にアクセスした際のハンドラ

```go:
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
})
```

#### 2-3. Listen for HTTP Connections: サーバをリッスン状態にする

リクエストハンドラだけではサーバは外部からの HTTP 接続を一切受け付けません。
そのためにサーバに**ポート番号**を指定し、そのポートへの接続を受け付ける状態にします。
サーバに通信したいポート番号を登録して、サーバはそのポート番号に接続要求があると通知を受けて処理を行います。この動作の事を"**ポートをリッスンする**"と表現します。
この節で、サーバにポート 80 を指定してリッスン状態にします。

[ListenAndServe](https://cs.opensource.google/go/go/+/go1.19.1:src/net/http/server.go;l=3253)
`func ListenAndServe(addr string, handler Handler) error`

第 1 引数: ポート番号
第 2 引数: ハンドラ
※2 目の引数に nil が渡された場合、デフォルトで DefaultServeMux という ServeMux 型のハンドラが使用されます。
基本的には nil を渡すのが正解のようですので、nil を渡します。

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

Go での基本的な HTTP サーバを作成する方法を学ぶ。
2 章と違う部分は静的アセットを処理する項がある部分。
静的アセット: 画像、CSS、JS などのリクエストよって表示する内容が変わらないもの。

#### 3-1. Introduction

HTTP サーバには、いくつかの重要な役割を担っている

* 動的なリクエストの処理(リクエストの内容はリクエストの度に違う)

  * ウェブサイトを閲覧
  * アカウントにログインする
  * 画像を投稿したりするユーザからの受信リクエスト

* 静的アセットの処理

  * JavaScript, CSS, 画像などをクライアントに送り、ユーザにダイナミックな体験を提供する

* クライアントからの接続を受け入れる
  * クライアントからリクエストを受け、レスポンスを返すためには特定のポートでリッスンする必要がある。

#### 3-2. Process dynamic requests: 動的リクエストの処理

リクエストを受け付け、処理するためのハンドラを登録する。登録するには"http.HandleFunc"関数を使用する。
第 1 引数: path(URL)
第 2 引数: 第 1 引数にアクセスした際に実行する関数

```go:
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome to my website!")
})
```

`http.Request`に、リクエストとそのパラメータに関する全ての情報が含まれている。
各パラメータの読み取りを行うには以下のように行う。
GET パラメータ: `r.URL.Query().Get("token")`
POST パラメータ(HTML フォームのフィールド): `r.FormValue("email")`

#### 3-3. Serving static assets: 静的アセットの提供

1. net/http パッケージの http.FileServer を使用し、URL パスを指定する。

FileServer は以下のようになっている。
`func FileServer(root FileSystem) Handler`
引数に FileSystem を受け取る必要があり、戻り値はハンドラです。
このハンドラは、ルートにあるファイルシステムの内容を HTTP リクエストに返す。

以下のコードでファイルサーバを設置している。
`fs := http.FileServer(http.Dir("static/"))`

これはファイルシステムを Dir メソッドを用いて static ディレクトと定義している。
Dir メソッドを使用すると OS のファイルシステム実装を使用できる。(OS 内のディレクトリを指定できる。)

2. ハンドラを登録するハンドルを登録する

ハンドラは、DefaultServeMux のパターンとしてハンドラを登録します。
正しくファイルを提供するために、StripPrefix メソッドを使用して URL パスの一部を削除する必要がある。

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

1. プロジェクト(ディレクトリ)を作成し、上記コード`main.go`としファイルを置く。

2. 同プロジェクト内に static ディレクトリを用意し、ファイルを用意する。ここでは`index.html`ファイルを置き、以下のコードを書くものとする。

```html:
<!DOCTYPE html>
<head>
  <title>File Server</title>
</head>
<body>
  <h1>Hello World!</h1>
</body>
```

3. `http://localhost:80`へアクセス
   Welcome to my website!と表示される。

4. `http://localhost:80/static`へアクセス
   Hello World!と表示される。

## 4. Routing(using gorilla/mux)

gorilla/mux パッケージを使用し、RESTful なサーバとのやり取りを学ぶ。

#### 4-1. Introduction

net/http のあまり得意ではない事の 1 つが、リクエスト URL をリクエスト内容によって分割するような複雑なリクエストルーティングである。(REST の設計は得意ではない)
そのため、gorilla/mux パッケージを使用する。

この章では名前付きパラメータ、GET/POST ハンドラ、ドメイン制限のあるルートを作成する方法を学ぶ。

#### 4-2. Installing the gorilla/mux package

* gorilla/mux パッケージの概要
  net/http パッケージのルーティングに適用できるパッケージ。
  **gorilla/mux はルーティング機能を提供する**
  Web アプリケーションを書く時の生産性を上げるための機能が多く備わっている。
  ミドルウェアなどの他の HTTP ライブラリや既存のアプリケーションと混在させることが可能

インストールするには以下のコマンドを使用する。
`$ go get -u github.com/gorilla/mux`

#### 4-3. Create a new Router

新しいリクエストルータを作成する。
このルータは、ウェブアプリケーションのメインルータになる。
全ての HTTP 接続を受信し、登録したリクエストハンドラを介してサーバにパラメータとして渡される。

新しいルータの作成するには以下のコマンドを使用する。
`r := mux.NewRouter`

#### 4-4. Registering a Request Handler

新しいルータを作成したら、通常と同じように(gorilla/mux を使用しなかった時と同じように)リクエストハンドラを登録する。
違いは`http.HandleFunc(...)`のように http メソッドを呼ぶ代わりに、作成したルータ上で HandleFunc を呼ぶ所です。

`r.HandleFunc(...)`

#### 4-5. URL Parameters

gorilla/mux Router の最大の強みは、リクエスト URL からセグメント(下で説明)を抽出することが出来る点。

`/books/go-programming-blueprint/page/10`
この URL を元に理解します。
この URL には 2 つのダイナミック(動的な)セグメントがある。

1. go-programming-blueprint
   book から続き、本のタイトルを表すセグメント。

2. page(**10**)
   本の何ページ目かを表すセグメント

リクエストハンドラがこのような動的に変わる URL を処理するためには、ダイナミックセグメントをプレースホルダに設定して、次のようにリクエストハンドラを変更する。

```go:
r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
  // get the book
  // navigate to the page
})
```

セグメントからデータを取得するには、gorilla/mux パッケージに不随する mex.Vars(r)関数を使用する。(r は作成したルータ)
これは http.Request をパラメータとして受け取り、セグメントのマップを返す。

```go:
func(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  vars["title"] // the book title slug
  vars["page"] // the page
}
```

#### 4-6. Setting the HTTP server's router

前述したが、`http.ListenAndServe(":80", nil)`の nil はデフォルトで nil で、nil の場合 net/http パッケージのデフォルトルータを使用する事を意味する。
今回は gorilla/mux パッケージを用いてルータを作成しているので、作成したルータを指定する。

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

1. `http://localhost/books/hoge/page/21`
   URL の"hoge", "21"は動的セグメントなので自由な値で大丈夫です。
   You've requested the book: hoge on page 21 と表示される。

## 5. Building a Simple REST API in Go With Gorilla/Mux

[Building a Simple REST API in Go With Gorilla/Mux](https://betterprogramming.pub/building-a-simple-rest-api-in-go-with-gorilla-mux-892ceb128c6f)

ここまでで golang におけるサーバ構築の基礎を学べたので、上記のサイトを参考に REST API のサーバを構築したいと思います。

#### 5-1. 概要

Groceries(食料品)API の概要

1. 特定の食料品とその数量の取得
2. 全ての食料品とその数量の取得
3. 食料品を投稿してそれを更新する

#### 5-2. ディレクトリ構造とファイル概要

ProjectName -> groceries
`go mod init groceries`

```shell:
$ tree
.
├── go.mod
├── grocery.go
├── handler.go
└── main.go
```

* grocery.go: API のモデルを定義
* handler.go: リクエストを管理する関数
* main.go: URL の path

#### 5-3. Gorilla/mux のインストール

`go get -u github.com/gorilla/mux`

#### 5-4. grocery.go

```go: grocery.go
package main

type Grocery struct {

  Name string `json: "name"`
  Quantity int `json: "quantity"`
}
```

grocery.go には API のモデルを定義する。
API のモデルは食料品の名前を表す Name と、その数量を表す Quantity の 2 つのフィールドだけです。

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

* mux.NewRouter()
  ルータを指定して各ハンドルを登録する。
  第 1 引数: URL path
  第 2 引数: 第 1 引数にアクセスされた際に処理する関数

* StrictSlash()
  StrictSlash はデフォルトで false で、true を指定すると、"/path/"とパス指定の場合に"/path"にアクセスすると前者にリダイレクトされる。逆の場合も同様。
  アプリケーションには常にルートで指定されたパスが表示される。

* Methods
  Methods メソッドで HTTP メソッドを指定できる。
  書かない場合は GET？
  ※今回はデータベースを用意しない。

* log.Fatal
  サーバのログを取得し、エラーが発生した場合はメッセージエラーが発生し、プログラムが停止する。

### 5-6. handler.go

#### 5-6-1. func AllGroceries

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

今回はデータベースを使用しないので、変数 groceries を定義し、2 つの食料品店の情報を含む配列を代入する。
`AllGroceries`が呼び出されると、すべての食料品を含む配列が JSON として返される。

#### 5-6-2. func SingleGrocery

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

#### 5-6-3. func GroceriesToBuy

```go: handler.go
func GroceriesToBuy(w http.ResponseWriter, r *http.Request) {
  reqBody, _ := ioutil.ReadAll(r.Body)

  var grocery Grocery
  json.Unmarshal(reqBody, &grocery)
  groceries = append(groceries, grocery)

  json.NewEncoder(w).Encode(groceries)
}
```

※ioutil パッケージは現在 io, os パッケージに別れて非推奨となっている。
`ioutil.ReadAll()`は、1.16以降は`io.ReadAll()`となっている。
`ioutil.ReadAll()`は一括読み込みする機能。

1. POSTリクエストを受け取ってreqBodyに代入している。(main 関数で methods を POST に指定済み)
2. Groceryタイプとする変数 groceryを定義している。
3. POST リクエストで受け取った jsonデータを変数 grocery に格納する(json.Unmarshalはjsonを構造体に変換する)
4. groceriesに3を追加している。

#### 5-6-4. func DeleteGrocery

```go: handler.go
func DeleteGrocery(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  name := vars("name")

  for index, grocery := range groceries {
    if grocery.Name == name {
      groceries = append(groceries[:index], groceries[index+1:]...)
    }
  }
}
```

groceries(食料品全部)から取り出したgrocery(食料品)の名前が、groceriesの中のいずれかに一致する場合、そのgroceryを削除します。
その後、スライスを更新します。

#### 5-6-5. func UpdateGrocery

```go: handler.go
func UpdateGrocery(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  name := vars["name"]

  for index, grocery := range groceries {
    if grocery.Name == name {
      groceries = append(groceries[:index], groceries[index+1]...)

      var updateGrocery Grocery

      json.NewDecoder(r .Body).Decode(&updateGrocery)
      groceries = append(groceries, updateGrocery)
      fmt.Println("Endpoint hit: UpdateGroceries")
      json.NewEncoder(w).Encode(updateGrocery)
      return
    }
  }
}
```

前項と同じように、名前が一致した場合に処理を行います。
この関数を使用して一致した場合、食料品を更新します。

これはmain関数でPUTメソッドに指定しているので、PUTリクエストを受け取り、デコードし、updateGrocery変数に格納し、groceriesに追加します。

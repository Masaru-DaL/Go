- [2. Hello World](#2-hello-world)
    - [2-1. Introduction](#2-1-introduction)
    - [2-2. Registering a Request Handler: リクエストハンドラの登録](#2-2-registering-a-request-handler-リクエストハンドラの登録)
    - [2-3. Listen for HTTP Connections: サーバをリッスン状態にする](#2-3-listen-for-http-connections-サーバをリッスン状態にする)
    - [2-4. The Code (for copy/paste): Hello Worldの完成コード](#2-4-the-code-for-copypaste-hello-worldの完成コード)
- [3. HTTP Server](#3-http-server)
    - [3-1. Introduction](#3-1-introduction)
    - [3-2. Process dynamic requests: 動的リクエストの処理](#3-2-process-dynamic-requests-動的リクエストの処理)
    - [3-3. Serving static assets: 静的アセットの提供](#3-3-serving-static-assets-静的アセットの提供)

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

#### 2-4. The Code (for copy/paste): Hello Worldの完成コード

```go:
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

net/httpパッケージのhttp.FileServerを使用し、URLパスを指定する。
FileServerは以下のようになっている。
`func FileServer(root FileSystem) Handler`
引数にFileSystemを受け取る必要があり、戻り値はハンドラです。
このハンドラは、ルートにあるファイルシステムの内容をHTTPリクエストに返す。

以下のコードでファイルサーバを設置している。
`fs := http.FileServer(http.Dir("static/"))`

これはファイルシステムをDirメソッドを用いてstaticディレクトと定義している。
Dirメソッドを使用するとOSのファイルシステム実装を使用できる。(OS内のディレクトリを指定できる。)


- [2. Hello World](#2-hello-world)
    - [2-1. Introduction](#2-1-introduction)
    - [2-2. リクエストハンドラの登録](#2-2-リクエストハンドラの登録)
    - [2-3. サーバをリッスン状態にする](#2-3-サーバをリッスン状態にする)
    - [2-4. Hello Worldの完成コード](#2-4-hello-worldの完成コード)
- [3. HTTP Server](#3-http-server)

### 1. 参考資料

[Go Web Examples - Learn Web Programming in Go by Examples](https://gowebexamples.com/)

## 2. Hello World

標準ライブラリのnet/httpパッケージを使用して、HTTPサーバを作成する。

#### 2-1. Introduction

net/httpパッケージには、HTTPプロトコルに関する全ての機能が備わっている。
サーバクライアントモデルが含まれる(他にもあるが)
この章で簡単にウェブサーバを作ることができる。

#### 2-2. リクエストハンドラの登録

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

#### 2-3. サーバをリッスン状態にする
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

#### 2-4. Hello Worldの完成コード

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
2章と違う部分は静的アセットをを処理する部分。
静的アセット: 画像、CSS、JSなどのリクエストよって表示する内容が変わらないもの。


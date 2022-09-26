- [2. Hello World](#2-hello-world)
    - [2-1. Introduction](#2-1-introduction)
    - [2-2. リクエストハンドラの登録](#2-2-リクエストハンドラの登録)

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
ハンドラ関数は2つのパラメータ(引数)をとる。

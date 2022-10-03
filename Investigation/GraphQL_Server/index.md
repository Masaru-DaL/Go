# Building a GraphQL Server with Go Backend Tutorial | Intro

参考: [GraphQL Tutorial](https://www.howtographql.com/graphql-go/0-introduction)

## 1. Introduction

#### 1-1. Motivation

gqlgenは、GoでGraphQLアプリケーションを作成するためのライブラリである。
このチュートリアルでは、HackernewsのGraphQL APIクローンをgolangとgqlgenで実装し、その過程でGraphQLの基礎について学ぶ。

#### 1-2. Schema-Driven Development

GraphQLでは、**APIは全ての型、クエリ、ミューテーションを定義するスキーマから始まる**。つまり、サーバとクライアントの契約のようなものです。(リクエストする形式と返す形式を合わせるような意味合いか)

GraphQL APIに新しい機能を追加する必要がある場合は、スキーマファイルを再定義し、その部分をコードで実装する必要がある。gqlgenは、GraphQLサーバを構築するためのGoライブラリで、スキーマの定義に基づいてコードを生成する。

## 2. Getting Started

### 2-1. Project Setup

1. `$ go mod init github.com/[username]/hackernews`
2. `$ go get github.com/99designs/gqlgen`
3. `$ go run github.com/99designs/gqlgen init`

#### 2-2. 生成されたgqlgenファイルの説明

* `gqlgen.yml`
  + gqlgenの設定ファイル
  + 生成されたコードの制御

* `graph/generated/generated.go`
  + GraphQL実行ランタイム
  + 生成されたコード

* `graph/model/models_gen.go`
  + グラフを構築するために必要な、生成されたモデル
  + 必要に応じて独自のモデルでオーバーライドする

* `graph/schema.graphqls`
  + GraphQLスキーマを追加するファイル

* `graph/schema.resolvers.go`
  + ここにアプリケーションのコードを記述する
  + `generated.go`はここを呼び出し、ユーザがリクエストしたデータを取得する

* `server.go`
  + 最小限のエントリポイント
  + `go run server.go`でサーバーを起動し、ブラウザを開くと、GraphQL playgroundが表示される

### 2-3. Defining Our Schema

* APIに必要なスキーマの定義
  + リンクとユーザをクライアントに表現するための2つのタイプ
    1. **リンクのリストを返すためのリンククエリ**
    2. **リンクを作成するためのミューテーション**
  + 新しいリンクを作成するための入力
  + ログイン, createUser, refreshTokenなどの認証システムのためのミューテーション

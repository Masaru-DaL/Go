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

* 上述を満たすように、スキーマの定義を行う

```go: graph/schema.graphqls
type Link {
  id: ID!
  title: String!
  address: String!
  user: User!
}

type User {
  id: ID!
  name: String!
}

type Query {
  links: [Link!]!
}

input NewLink {
  title: String!
  address: String!
}

input RefreshTokenInput{
  token: String!
}

input NewUser {
  username: String!
  password: String!
}

input Login {
  username: String!
  password: String!
}

type Mutation {
  createLink(input: NewLink!): Link!
  createUser(input: NewUser!): String!
  login(input: Login!): String!
  # we'll talk about this in authentication section
  refreshToken(input: RefreshTokenInput!): String!
}

```

スキーマ定義後、
 `$ go run github.com/99designs/gqlgen generate`

`validation failed: packages.Load` というエラーが出たので、公式の手順通りに `schema.resolvers.go` のCreateTodoとTodosを削除してから再度 `go run github.com/99designs/gqlgen generate` を叩いて無事スキーマで定義した関数が生成される。

## 3. Queries

2章でサーバーのセットアップが完了した。
`schema.graphqls` で定義したQueryを実装する。

### 3-1. What Is A Query

* Queryとは？
GraphQLのクエリとは、**データを要求するもの**。
クエリを使って欲しい情報を指定すると、GraphQLサーバは要求した情報を返す。

### 3-2. Simple Query

スキーマで定義されたものを実装するには、 `schema.resolvers.go` に記述することで実装される。
既に生成されているLinks関数を見てみる。

 `func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {`

この関数は、Contextを受け取り、Linksのスライスとエラー(もしあれば)を返す。
`ctx` 引数には、リクエストを送信した人のデータが含まれている。

1. この関数に対して、ダミーのレスポンスを作成してみる。

```go: schema.resolvers.go
func(r *queryResolver) Links(ctx context. Context) ([]*model. Link, error) {
    var links []*model.Link
    dummyLink := model.Link{
        Title: "our dummy link",
        Address: "https://address.org",
        User: &model.User{Name: "admin"},
    }
    links = append(links, &dummyLink)
    return links, nil
}
```

2. `$ go run server.go`

3. GraphQLサーバにQueryを送る

```graphql:
Query {
  links {

    title,
    address,
    user {
      name
    }

  }
}
```

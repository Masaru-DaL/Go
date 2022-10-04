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

4. GraphQLからのレスポンス

```graphql:
{
  "data": {
    "links": [
      {
        "title": "our dummy link",
        "address": "https://address.org",
        "user": {
          "name": "admin"
        }
      }
    ]
  }
}
```

`resolvers.go` に実装することで、クエリを投げた際にその項目がレスポンスされるその方法が分かった。ここまではあくまでもダミーのレスポンスなので、実際にやりたいことは**他のユーザのリンクを全て照会できるようにしたい**。

次の章で、アプリケーションにデータベースに保存したリンクを取得できるようにする。

## 4. Mutations

### 4-1. What Is A Mutation

**技術的にはQueryもデータの書き込み(POST)に使用出来るが、そのやり方は推奨されていない。**
つまり、ミューテーションはクエリのように使えるということで、名前とパラメータを持っていて**データを返す事ができる**。

### 4-2. A Simple Mutation

まだデータベースをセットアップしていない。(次の章で行う)
そのため、リンクデータを受け取り、リンクオブジェクトを構築し、レスポンスを返すという動作を行う `CreateLink` ミューテーションを実装する。(データベースを介さない)

 `func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {`

この関数は、 `schema.graphqls` で定義したNewLink構造を持つNewLinkを入力として受け取る。(ちょっとよくわからない)

1. `schema.graphqls`で定義したLinkオブジェクトを構築する

```go: schema.resolvers.go
func (r *mutationResolver) CreateLink(ctx context. Context, input model. NewLink) (*model. Link, error) {
  var link model. Link
  var user model. User
  link. Address = input. Address
  link. Title = input. Title
  user. Name = "test"
  link. User = &user
  return &link, nil
}

```

2. `$ go run server.go`

3. ミューテーションを使用して新しいリンクを作成する

```graphql:
mutation {
  createLink(input: {title: "new link", address:"http://address.org"}) {
    title,
    user {
      name
    }
    address
  }
}
```

4. GraphQLからのレスポンス

```graphql:
{
  "data": {

    "createLink": {
      "title": "new link",
      "user": {
        "name": "test"
      },
      "address": "http://address.org"
    }

  }
}

```

`schema.resolvers.go` で実装したように、inputで入力した内容がinput.で記述した箇所と紐づき、デフォルト値(test)を設定した箇所はそれが返っている。
**どういった内容でミューテーションが送られたか、その結果が返ってきている**。

## 5. Database

GraphQLスキーマを実装するために、ユーザとリンクを保存するためのデータベースをセットアップする必要がある。

* MySQLのセットアップ
* MySQLデータベースの作成
* モデルを定義し、マイグレーションを作成する

### 5-1. Setup MySQL

dockerでMySQLのイメージを使用する。

 `docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=hackernews -d mysql:latest`

`$ docker ps` でMySQLのイメージが起動しているのを確認する。

### 5-2. Create MySQL database

* hackernewsデータベースを作成する。

1. MySQLのコンテナに入る
 `$ docker exec -it mysql bash`

2. rootユーザMySQLを使用する
 `$ mysql -u root -p`

3. rootのパスワードを入力
"dbpass" (上で設定したもの)

4. データベースの作成
 `CREATE DATABASE hackernews;`

### 5-3. Models and migrations

アプリを実行する度に必要なテーブルを作成・適切に動作するために、マイグレーションを作成する必要がある。

* マイグレーション
  + テーブル・インデックス更新の自動化が行える。

1. プロジェクトのルートディレクトリに、データベースファイルのためのフォルダ構造を作成する。

```shell:
$ tree

```

2. `go mysql driver`と`golang-migrate`パッケージをインストールし、migrationsを作成する。

```shell:
$ go get -u github.com/go-sql-driver/mysql
$ go build -tags 'mysql' -ldflags="-X main. Version=1.0.0" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/
$ cd internal/pkg/db/migrations/
$ migrate create -ext sql -dir mysql -seq create_users_table
$ migrate create -ext sql -dir mysql -seq create_links_table

```

migrateコマンドは、マイグレーションごとに `.up` と `.down` で終わる2つのファイルを作成する。
up -> マイグレーションを適用する役割
down -> それを反転する役割

3. `000001_create_users_table.up.sq`に、ユーザ用のテーブルを追加する。

```sql:
CREATE TABLE IF NOT EXISTS Users (
  ID INT NOT NULL UNIQUE AUTO_INCREMENT,
  Username VARCHAR (127) NOT NULL UNIQUE,
  Password VARCHAR (127) NOT NULL,
  PRIMARY KEY (ID)
)
```

4. `000002_create_links_table.up.sql`に、リンク用のテーブルを追加する。

```sql:
CREATE TABLE IF NOT EXISTS Links(

    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Title VARCHAR (255) ,
    Address VARCHAR (255) ,
    UserID INT ,
    FOREIGN KEY (UserID) REFERENCES Users(ID) ,
    PRIMARY KEY (ID)

)
```

5. 3, 4で設定した内容を反映させ、それぞれのテーブルを作成する。migrateコマンドで行う。
プロジェクトのルートディレクトリでこのコマンドを実行します。

 `$ migrate -database mysql://root:dbpass@/hackernews -path internal/pkg/db/migrations/mysql up`

6. データベースの接続を行う。
今回はMySQLを使用するので、mysqlフォルダの下にデータベースへの接続を初期化する関数を作成する。
複数のデータベースを持つ場合は、他のフォルダを追加できる。

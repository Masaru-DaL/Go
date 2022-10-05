- [Building a GraphQL Server with Go Backend Tutorial | Intro](#building-a-graphql-server-with-go-backend-tutorial--intro)
  - [1. Introduction](#1-introduction)
      - [1-1. Motivation](#1-1-motivation)
      - [1-2. Schema-Driven Development](#1-2-schema-driven-development)
  - [2. Getting Started](#2-getting-started)
    - [2-1. Project Setup](#2-1-project-setup)
      - [2-2. 生成されたgqlgenファイルの説明](#2-2-生成されたgqlgenファイルの説明)
    - [2-3. Defining Our Schema](#2-3-defining-our-schema)
  - [3. Queries](#3-queries)
    - [3-1. What Is A Query](#3-1-what-is-a-query)
    - [3-2. Simple Query](#3-2-simple-query)
      - [3-1-1. この関数に対して、ダミーのレスポンスを作成してみる。](#3-1-1-この関数に対してダミーのレスポンスを作成してみる)
      - [3-1-2. `$ go run server.go`](#3-1-2--go-run-servergo)
      - [3-1-3. GraphQLサーバにQueryを送る](#3-1-3-graphqlサーバにqueryを送る)
      - [3-1-4. GraphQLからのレスポンス](#3-1-4-graphqlからのレスポンス)
  - [4. Mutations](#4-mutations)
    - [4-1. What Is A Mutation](#4-1-what-is-a-mutation)
    - [4-2. A Simple Mutation](#4-2-a-simple-mutation)
      - [4-2-1. `schema.graphqls` で定義したLinkオブジェクトを構築する](#4-2-1-schemagraphqls-で定義したlinkオブジェクトを構築する)
      - [4-2-2. `$ go run server.go`](#4-2-2--go-run-servergo)
      - [4-2-3. ミューテーションを使用して新しいリンクを作成する](#4-2-3-ミューテーションを使用して新しいリンクを作成する)
      - [4-2-4. GraphQLからのレスポンス](#4-2-4-graphqlからのレスポンス)
  - [5. Database](#5-database)
    - [5-1. Setup MySQL](#5-1-setup-mysql)
    - [5-2. Create MySQL database](#5-2-create-mysql-database)
    - [5-3. Models and migrations](#5-3-models-and-migrations)
      - [5-3-1. プロジェクトのルートディレクトリに、データベースファイルのためのフォルダ構造を作成する。](#5-3-1-プロジェクトのルートディレクトリにデータベースファイルのためのフォルダ構造を作成する)
      - [5-3-2. `go mysql driver` と `golang-migrate` パッケージをインストールし、migrationsを作成する。](#5-3-2-go-mysql-driver-と-golang-migrate-パッケージをインストールしmigrationsを作成する)
      - [5-3-3. `000001_create_users_table.up.sq`に、ユーザ用のテーブルを追加する。](#5-3-3-000001_create_users_tableupsqにユーザ用のテーブルを追加する)
      - [5-3-4. `000002_create_links_table.up.sql` に、リンク用のテーブルを追加する。](#5-3-4-000002_create_links_tableupsql-にリンク用のテーブルを追加する)
      - [5-3-5. 3, 4で設定した内容を反映させ、それぞれのテーブルを作成する。migrateコマンドで行う。](#5-3-5-3-4で設定した内容を反映させそれぞれのテーブルを作成するmigrateコマンドで行う)
      - [5-3-6. データベースの接続を行う。](#5-3-6-データベースの接続を行う)
      - [5-3-7. main関数にInitDBとMigrateを呼び出すように記述し、アプリの開始時にデータベース接続を作成するようにする。](#5-3-7-main関数にinitdbとmigrateを呼び出すように記述しアプリの開始時にデータベース接続を作成するようにする)
  - [6. Create and Retrieve Links](#6-create-and-retrieve-links)
    - [6-1. CreateLinks](#6-1-createlinks)
      - [6-1-1. usersディレクトリ](#6-1-1-usersディレクトリ)
      - [6-1-2. linksディレクトリ](#6-1-2-linksディレクトリ)
      - [6-1-3. Save関数をCreateLinkリゾルバで使用する](#6-1-3-save関数をcreatelinkリゾルバで使用する)
      - [6-1-4. ミューテーションの送信](#6-1-4-ミューテーションの送信)
    - [6-2. links Query](#6-2-links-query)
      - [6-2-1. リンクの取得、サーバに渡す関数](#6-2-1-リンクの取得サーバに渡す関数)
      - [6-2-2. GetAll関数でリンクを取得できるようにする](#6-2-2-getall関数でリンクを取得できるようにする)
      - [6-2-3. Queryの送信](#6-2-3-queryの送信)
  - [7. Authentication](#7-authentication)
    - [7-1. JWT: Json Web Token](#7-1-jwt-json-web-token)
    - [7-2. Setup](#7-2-setup)
    - [7-3. Generating and Parsing JWT Tokens](#7-3-generating-and-parsing-jwt-tokens)
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

```graphqls: graph/schema.graphqls
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

#### 3-1-1. この関数に対して、ダミーのレスポンスを作成してみる。

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

#### 3-1-2. `$ go run server.go`

#### 3-1-3. GraphQLサーバにQueryを送る

```graphql:
query {
  links {

    title,
    address,
    user {
      name
    }

  }
}

```

#### 3-1-4. GraphQLからのレスポンス

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

#### 4-2-1. `schema.graphqls` で定義したLinkオブジェクトを構築する

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

#### 4-2-2. `$ go run server.go`

#### 4-2-3. ミューテーションを使用して新しいリンクを作成する

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

#### 4-2-4. GraphQLからのレスポンス

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

#### 5-3-1. プロジェクトのルートディレクトリに、データベースファイルのためのフォルダ構造を作成する。

```shell:
$ tree

```

#### 5-3-2. `go mysql driver` と `golang-migrate` パッケージをインストールし、migrationsを作成する。

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

#### 5-3-3. `000001_create_users_table.up.sq`に、ユーザ用のテーブルを追加する。

```sql:
CREATE TABLE IF NOT EXISTS Users (
  ID INT NOT NULL UNIQUE AUTO_INCREMENT,
  Username VARCHAR (127) NOT NULL UNIQUE,
  Password VARCHAR (127) NOT NULL,
  PRIMARY KEY (ID)
)
```

#### 5-3-4. `000002_create_links_table.up.sql` に、リンク用のテーブルを追加する。

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

#### 5-3-5. 3, 4で設定した内容を反映させ、それぞれのテーブルを作成する。migrateコマンドで行う。

プロジェクトのルートディレクトリでこのコマンドを実行します。

 `$ migrate -database mysql://root:dbpass@/hackernews -path internal/pkg/db/migrations/mysql up`

#### 5-3-6. データベースの接続を行う。

今回はMySQLを使用するので、mysqlフォルダの下にデータベースへの接続を初期化する関数を作成する。
複数のデータベースを持つ場合は、他のフォルダを追加できる。

```go: internal/pkg/db/mysql/mysql.go
package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
)

var Db *sql.DB

func InitDB() {
	// Use root:dbpass@tcp(172.17.0.2)/hackernews, if you're using Windows.
	db, err := sql.Open("mysql", "root:dbpass@tcp(localhost)/hackernews")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
 		log.Panic(err)
	}
	Db = db
}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := mysql.WithInstance(Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
```

* InitDB関数
  + データベースへの接続を作成する。

* Migrate関数
  + migrationsファイルを実行する
  + コマンドラインで行った事と同じようにmigrationsを適用する
  + この関数を使うことで、アプリケーションは起動前に常に最新のmigrationsを適用する

* CloseDB関数
  + アプリケーションが終了したらデータベース接続を閉じる役割を果たす
  + deferキーワードで呼び出され、main関数で終了したときに実行される
    - (遅延させて、最後に実行される)

#### 5-3-7. main関数にInitDBとMigrateを呼び出すように記述し、アプリの開始時にデータベース接続を作成するようにする。

```go: server.go
func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	database.InitDB()
	defer database.CloseDB()
	database.Migrate()
	server := handler.NewDefaultServer(hackernews.NewExecutableSchema(hackernews.Config{Resolvers: &hackernews.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}

```

## 6. Create and Retrieve Links

### 6-1. CreateLinks

まず、アプリケーションとデータベースを繋げるリンクを作成するための関数が必要。
`internal` ディレクトリに、linksとusersディレクトリを作成する。
作成した2つのディレクトリが、データベースとアプリケーションのやり取りを行う階層となる。

#### 6-1-1. usersディレクトリ

```go: internal/users/users.go
package users

type User struct {
  ID  string `json:"id"`
  Username string `json:"name"`
  Password string `json:"paddword"`
}
```

#### 6-1-2. linksディレクトリ

```go: internal/links/links.go
package links

import (

	database "github.com/glyphack/go-graphql-hackernews/internal/pkg/db/mysql"
	"github.com/glyphack/go-graphql-hackernews/internal/users"
	"log"

)

// #1
type Link struct {

	ID      string
	Title   string
	Address string
	User    *users.User

}

//#2
func (link Link) Save() int64 {

	//#3
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id

}

```

1. リンクを表す構造体の定義
2. リンクをデータベースに挿入し、IDを返す関数を定義
3. "INSERT INTO..."でリンクをLinksテーブルに挿入するSQLクエリ。`prepare`を使うとセキュリティやパフォーマンス向上に役立つ。
4. SQL文の実行
5. 挿入されたリンクのIDを取得する。

#### 6-1-3. Save関数をCreateLinkリゾルバで使用する

```go: schema.resolvers.go
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
  var link links.Link
  link.Title = input.Title
  link.Address = input.Address
  linkID := link.Save()
  return &model.Link{ID: strconv.FormatInt(linkID, 10), Title:link.Title, Address:link.Address}, nil
}
```

このコードは、入力からリンクオブジェクトを作成してデータベースに保存し、新しく作成されたリンクを返している。(strconv. FormatIntでIDを文字列に変換している)

#### 6-1-4. ミューテーションの送信

```graphql:
mutation create{
  createLink(input: {title: "something", address: "somewhere"}){

    title,
    address,
    id,

  }
}

```

- レスポンス

```graphql:
{
  "data": {
    "createLink": {
      "title": "something",
      "address": "somewhere",
      "id": "1" // AutoIncrement
    }
  }
}
```

### 6-2. links Query

CreateLinkミューテーションの次に、Linksクエリを実装する。
データベースからリンクを取得し、リゾルバでそれをGraphQLサーバに渡す関数が必要。

#### 6-2-1. リンクの取得、サーバに渡す関数

関数をGetAllという名前で作成する。

```go: internal/links/links.go
func GetAll() []Link {
  stmt, err := database. Db. Prepare("select id, title, address from Links")
  if err != nil {

    log.Fatal(err)

  }
  defer stmt. Close()

  rows, err := stmt. Query()
  if err != nil {

    log.Fatal(err)

  }
  defer rows. Close()

  var links []link
  for rows. Next() {

    var link Link
    err := rows.Scan(&link.ID, &link.Title, &link.Address)
    if err != nil {
      log.Fatal(err)
    }
    links = append(links, link)

  }
  if err = rows. Err(); err != nil {

    log.Fatal(err)

  }
  return links
}

```

#### 6-2-2. GetAll関数でリンクを取得できるようにする

```go: schema.resolvers.go
func (r *queryResolver) Links(ctx context. Context) ([]*model. Link, error) {
  var resultLinks []*model. Link
  var dbLinks []links. Link
  dbLinks = links. GetAll()
  for _, link := range dbLinks {

    resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title:link.Title, Address:link.Address})
  }
  return resultLinks, nil
}
```

#### 6-2-3. Queryの送信

```graphql:
query {
  links {

    title
    address
    id

  }
}

```

- レスポンス

```graphql:
{
  "data": {
    "links": [
      {
        "title": "something",
        "address": "somewhere",
        "id": "1"
      }
    ]
  }
}
```

## 7. Authentication

ウェブアプリケーションに認証レイヤを追加する。
ユーザを認証する方法として"JWTトークン"を使用する。

### 7-1. JWT: Json Web Token

**ハッシュを含む文字列で、ユーザを認証するためのもの**。
ヘッダー、ペイロード、シグネチャで構成される。

ユーザがアプリケーションにログインするたびに、サーバはトークンを生成する。この時、**サーバはユーザ名などの情報をトークンに含め、後でユーザを認識できるようにする**。
このトークンは秘密鍵で署名されるので、発行者(アプリケーション側)だけがトークンの中身を読むことができる。
今回実装する。

### 7-2. Setup

このアプリでは、**ユーザがサインアップまたはログインするときにトークンを生成できる**ようにする必要がある。また、与えられたトークンを使ってユーザを認証するためのミドルウェアを作成し、誰がサーバに接続しているかを知る必要がある。

JWTトークンの生成とパースには、 `github.com/dgrijalva/jwt-go` ライブラリを使用する。

### 7-3. Generating and Parsing JWT Tokens

アプリケーションのルートにpkgという新しいディレクトリを作成する。対してアプリケーションの内部でのみ使用したいものには `internal` を使用している。
pkgディレクトリは、アプリケーションのどこにでもインポートできるファイルのためのもの。JWT生成スクリプトや検証スクリプトがこれに該当する。

**クレーム**と呼ばれる概念があることを覚えておく。

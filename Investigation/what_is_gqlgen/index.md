# gqlgenについて調べてみる

## 参考

[gqlgen](https://github.com/99designs/gqlgen)
[gqlgenチュートリアルをできるだけわかりやすく解説する](https://zenn.dev/omoterikuto/articles/a43c989ca36073)
[GraphQL](https://graphql.org/)

## 1. what is gqlgen?

**GraphQLサーバを簡単に構築するためのGoライブラリのこと**

* スキーマファースト
* 型の安全性を優先する
* Codegenを可能にする

つまり、"**スキーマファースト**"で、"**型の安全性を保ったまま**"、"**コードを自動生成できる**"という特徴があるようです。

gqlgenの読み方がわからないけど、"graphqlgenerate"的な感じがする。

#### 1-1. what is GraphQL?

そもそもGraphQLをあまり理解していない。
イメージとしてはREST APIと対を成すイメージ。

* APIのためのクエリ言語
  + APIに対してSQLのように命令をする？

* URLは1つ

* クエリを投げて情報を得る

* 使いこなしている人が少ない？

* SQLと混同しやすい？

今のところあまりイメージがしにくいです。

## 2. Quick start

概要は掴めたので、何はともあれ手を動かしてみる。

#### 2-1. 作業ディレクトリの作成

1. `mkdir gqlgen_todos && cd gqlgen_todos`
2. `go mod init gqlgen_todos`
3. `touch tools.go`
4. `printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go`
5. `go mod tidy`
6. `go run github.com/99designs/gqlgen init`
7. `go run server.go`

* この時点でのフォルダ構成

```shell:
-> tree
.
├── go.mod
├── go.sum
├── gqlgen.yml               // コード生成の設定ファイル
├── graph
│   ├── generated            // 自動生成されたパッケージ（基本的にいじらない）
│   │   └── generated.go
│   ├── model                // Goで実装したgraph model用のパッケージ（自動生成されたファイルと自分でもファイルを定義することが可能）
│   │   └── models_gen.go
│   ├── resolver.go          // ルートのresolverの型定義ファイル. 再生成で上書きされない。
│   ├── schema.graphqls      // GraphQLのスキーマ定義ファイル. 実装者が好きに分割してもOK
│   └── schema.resolvers.go  // schema.graphqlから生成されたresolverの実装ファイル
└── server.go                // アプリへのエントリポイント. 再生成で上書きされない。
└── tools.go

3 directories, 10 files

```

#### 2-2. graphフォルダの中身

* graph/generated/generated.go
  GraphQLサーバに対するリクストを解釈し、 `graph/resolver.go` の適切なメソッドを呼ぶ役割を果たす。

* graph/model/models_gen.go
  スキーマで定義したものをgolangの構造体に変換したものが定義される。
  なんとなくイメージ ->  [GraphQLスキーマ設計](https://future-architect.github.io/articles/20200609/#:~:text=%E3%81%A7%E3%81%8A%E3%81%97%E3%81%BE%E3%81%84%E3%81%A7%E3%81%99%E3%80%82-, GraphQL%E3%82%B9%E3%82%AD%E3%83%BC%E3%83%9E%E8%A8%AD%E8%A8%88, -%E3%81%93%E3%81%93%E3%81%8B%E3%82%89GraphQL)

* graph/schema.resolver.go
  リクエストを元に実際の処理を実装する `resolver` ファイル

上記の3つのファイルをもって、スキーマを変更した後、 `go run github.com/99designs/gqlgen generate` を実行することでコードが再生成される。

* graph/resolver.go
  ルートとなるresolver構造体が宣言される。再生成はされない。

* graph/schema.graphqls
  GraphQLスキーマを定義するファイル。このファイルをもとに他のファイルが再生成される。

* gqlgen.yml
  gqlgenの設定ファイル。schemaの分割などの設定もこのファイルで行われる。

## 3. アプリケーションの作成

Todoアプリケーションを動作させるには、`CreateTodo`と、`Todos`のメソッドを実装する必要があるようです。

#### 3-1. graph/resolver.go

```go: resolver.go
type Resolver struct {
  todos []*model.Todo
}
```

#### 3-2. graph/schema.resolvers.go

```go: schema.resolvers.go
func (r *mutationResolver) CreateTodo(ctx context. Context, input model. NewTodo) (*model. Todo, error) {

	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil

}

func (r *queryResolver) Todos(ctx context. Context) ([]*model. Todo, error) {

	return r.todos, nil

}

```

#### 3-4. server.go

```shell:
$ go run server.go
2022/10/01 14:18:45 connect to http://localhost:8080/ for GraphQL playground
```

言われた通りにアクセスします。
すると無事に `GraphQL playground` に繋がりました。

## 4. GraphQL playground で色々やってみる

#### 4-1. todoの作成と取得(確認)

1. todoの作成
以下のコードを左側に記述し、"Execute query"(以降エンター)

```graphql:
mutation {
  createTodo(input: { text: "todo", userId: "1" }) {

    user {
      id
    }
    text
    done

  }
}

```

レスポンス

```graphql:
{
  "data": {

    "createTodo": {
      "user": {
        "id": "1"
      },
      "text": "todo",
      "done": false
    }

  }
}

```

2. 作成したtodoの取得
以下のコードを記述後、エンター

```graphql:
query {
  todos {

    text
    done
    user {
      name
    }

  }
}

```

レスポンス

```graphql:
{
  "data": {

    "todos": [
      {
        "text": "todo",
        "done": false,
        "user": {
          "name": "user 1"
        }
      }
    ]

  }
}

```

ここまでやった事はinitで作成されたgqlgenのひな形の `resolver.go` でresolverを実装しただけです。
それだけでtodoの作成、取得ができるアプリケーションが作成されました。

#### 4-2. GraphQLの恩恵

Todoにユーザのデータ自体 `*User` として読み込ませると、取得する際にコストがかかってしまう。
GraphQLでは特定の情報だけ抽出する事が出来るので、UserIDというただの文字列を返すものをmodelの実装する。

```go:
type Todo struct {

	ID   string `json:"id"`

	Text string `json:"text"`

	Done bool `json:"done"`

	User *User `json:"user"` // Userのポインタを指している

}

```

* 新しく作るTodoのmodel

```go:
type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"` // UserIDの追加
	User   *User  `json:"user"`
}

```

新しく定義するTodoの構造体にはデータではなく、UserIDとしてただの文字列を返すようにしています。
このようにする事によって、不要なデータを返さないようにして、GraphQLの恩恵を最大限受けるのが正しい設計であると言えるようです。

## 5. GraphQLに沿った設計

#### 5-1. 新しいTodo

```go: graph/model/todo.go
package model

type Todo struct {

	ID     string `json:"id"`

	Text   string `json:"text"`

	Done   bool `json:"done"`

	UserID string `json:"userId"`

	User   *User `json:"user"`

}

```

```yml:
# gqlgen will search for any type names in the schema in these go packages
# if they match it will use them, otherwise it will generate them.
autobind:
  - "gqlgen-todos/graph/model"

###################################

# This section declares type mapping between the GraphQL and go type systems
#
# The first line in each type will be used as defaults for resolver arguments and
# modelgen, the others will be allowed when binding to fields. Configure them to
# your liking
models:
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.ID
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
  Int:
    model:
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
  Todo:
    fields:
      user:
        resolver: true

```

新しくmodelディレクトリに `todo.go` を作成する。
次に `yml` ファイルを書き換える。
自動バインドの有効と、ユーザフィールドのリゾルバの生成。
正直あんま分かってない...

そこまで出来たら
`go run github.com/99designs/gqlgen generate` を行う。
再生成される。

#### 5-2. resolver.go

resolver.goのファイルがエラーになっています。
`Todo struct` のUserを `*User` ではなく `UserID` に変更したため、そこに対する修正を行います。

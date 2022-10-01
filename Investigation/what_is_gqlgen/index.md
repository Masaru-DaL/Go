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

1. `mkdir gqlgen_tutorial && cd gqlgen_tutorial`
2. `go mod init gqlgen_tutorial`
3. `go get -u github.com/99designs/gqlgen@v0.17.5`
4. `go run github.com/99designs/gqlgen init`
   このコマンドは初。
   > gqlgenの設定を初期化し、モデルを生成

* この時点でのフォルダ構成

```shell:
-> tree
.
├── go.mod
├── go.sum
├── gqlgen.yml
├── graph
│   ├── generated
│   │   └── generated.go
│   ├── model
│   │   └── models_gen.go
│   ├── resolver.go
│   ├── schema.graphqls
│   └── schema.resolvers.go
└── server.go

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

#### 2-3. アプリケーションの作成

この時点である、 `graph/schema.graphqls` の中身(スキーマ)が、gqlgenがデフォルトで生成してくれているスキーマです。
このスキーマをベースに簡単なtodoアプリケーションの作成を行う。

`init`で作成されたものを正常に動作させるためには`CreateTodo`と`Todos`を`graph/resolver.go` に実装させる必要がある。

1. 構造体の宣言

```go: resolver.go
type Resolver struct {
  todos []*model.Todo
}
```

2. CreateTodo, Todosの関数の実装

```go: resolver.go
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

3. この時点でのresolver.go

formatterにより自動インポートなど

```go: resolver.go
package graph

import (
	"context"
	"math/rand" // crypto/randから変更
	"fmt"

	"github.com/<UserName>/gqlgen-todos/graph/model" // UserName is your User Name
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}
```

4. server.go

公式もこの後に `server.go` を促していきますが、明らかにエラーが出ていて、コマンドを打っても案の定起動できません。

調べていると[こちら](https://stackoverflow.com/questions/60669166/golang-gqlgen-error-trying-to-import-model-into-resolver-go)に当たって、generateコマンドは良く分からなかったので、もう1つの `graph/schema.resolvers.go` の `CreateTodo` と `Todos` をコメントアウトしました。

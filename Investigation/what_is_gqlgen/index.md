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

* URLは1つ

* クエリを投げて情報を得る

* 使いこなしている人が少ない？

* SQLと混同しやすい？

今のところあまりイメージがしにくいです。

## 2. Quick start

概要は掴めたので、何はともあれ手を動かしてみる。

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
├── main.go
└── server.go
```

* graph/generated/generated.go
  GraphQLサーバに対するリクストを解釈し、`graph/resolver.go`の適切なメソッドを呼ぶ役割を果たす。

* graph/model/models_gen.go

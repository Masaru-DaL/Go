# GraphQL Server

参考: [GraphQL](https://graphql.org/)

## 1. Getiing Started

始めようとしたら、まず言語を選ぶところから。

> GraphQLは通信サポートであるため、あらゆる言語でGraphQLをサポートする、作業を始めるためのツールが数多く存在する。

Goを選ぶと更に様々な選択肢が。
今回はGolangのGraphQLを行いたいので [graphql-go](https://github.com/graphql-go/graphql)を選択。

## 2. ライブラリのインストール

1. `$ mkdir graphql-tutorial`
2. `$ cd graphql-tutorial`
3. `$ go mod init graphql-tutorial`
4. `$ go get github.com/graphql-go/graphql`

## 3. Learn Golang + GraphQL + Relay #1

公式にはチュートリアル的なものはなかったので、紹介されていた[Learn Golang + GraphQL + Relay #1](https://wehavefaces.net/learn-golang-graphql-relay-1-e59ea174a902)を見ていきたいと思います。(記事が古いのが気になりますが)

* GraphQLの型システムの簡略表記
上位のクエリの1つにlatestPostフィールドを持ち、それはGraphQLの組み込みのStringの値を返す(Hello World)

```graphql: schema.graphql
//
// In GraphQL's type system shorthand notation.
// (See: http://graphql.org/docs/typesystem)
//

type Query {
  latestPost: String
}
```

これを基にGraphQLスキーマを定義していく。

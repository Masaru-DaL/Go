# GraphQLとREST APIの違い

## 現時点で分かっている事

* URLの扱いの違い
これぐらいしか分からないけど、一番大きな違いだと思っているが...
この認識自体も再確認しつつ、他の部分も比較していく。

REST APIはURLに合わせて行いたい操作を紐付けていく設計思想です。
`books/{title}/page/{page}` のように、URLだけで何を行いたいのか開発者にも分かりやすい点で広く知られています。
ここではREST APIに関しては深掘りせず、GraphQLと比較の対象としてだけフォーカスします。

## 1. GraphQLとは？

* Facebook社が開発・提供
* Web APIのクエリ言語
  + サーバにデータを問い合わせ出来る
* クエリ言語とスキーマ言語で構成
* 様々なOSに対応していて、クロスプラットフォーム開発が可能

## 2. GraphQLの開発の経緯

参考: [GraphQLとは？RESTとの違いや導入事例を紹介](https://udemy.benesse.co.jp/development/system/graphql.html)

REST API -> GraphQLというように、REST APIの方が先にあった。
REST APIに不便性があり、GraphQLが開発された。

**REST APIの不便性**
* 特定のデータだけを参照する事が難しい
RESTは不必要なデータも取得してしまうため、大量のデータを処理する必要が出てくる。それによって、処理の複雑化、パフォーマンス低下が起こる。
この問題を解決するためにGraphQLの開発がスタートした。

何のために開発されたのか、を見ていくとGraphQLの最大の特徴が見えてきました。[公式](https://graphql.org/)でも述べられています。

> GraphQL provides a complete and understandable description of the data in your API, gives clients the power to ask for exactly what they need and nothing more, makes it easier to evolve APIs over time, and enables powerful developer tools.

> GraphQLは、API内のデータについて完全で理解しやすい記述を提供し、クライアントが必要なものだけを要求する力を与え、時間の経過とともにAPIを進化させることを容易にし、強力な開発者ツールを可能にします。

**クライアントが必要なものだけを要求する事に応える。という点**が最大の特徴だということが分かりました。

## 3. RESTとGraphQLの比較

参考: [GraphQL対REST - 考慮すべき点](https://www.infoq.com/jp/news/2017/08/graphql-vs-rest/)

#### 3-1. 設計思想

* REST
  + 1つの操作に対して1つのURLを用いる点で、

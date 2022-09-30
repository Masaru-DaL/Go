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

#### 3-1. APIに対しての操作の違い

* REST
  + 複数URL(1つの操作に対して1つのURLを用いる点)
    - メリット -> URLから予測可能。HTTPメソッドを使用する事により、分かりやすい。
    - デメリット -> URLが増えれば増えるほど管理が困難化

* GraphQL
  + 単一URL(用意するのは1つのURLのみ)
    - メリット -> 一度のリクエストで多くの情報が取得できる
    - デメリット -> 予測可能ではなくなる。(視認性低下)

#### 3-2. パフォーマンス

* REST
  + パフォーマンスは良くない
    - 特定の情報だけの抽出が難しい
    - 処理が重くなる。
    - 不必要なデータがメモリに蓄積し、メモリリークを引き起こす可能性もある

* GraphQL
  + パフォーマンスは良い
    - 特定の情報だけの抽出が可能
    - 更に1度のリクエストで取得できるので負荷が少ない

#### 3-3. 設計思想の違い

RESTは**原則**, GraphQLは**規格**

RESTは自由性がある。
GraphQLは規格に縛られる。

両者のその特徴が良い所でもあり、悪い所でもある。

#### 3-4. APIの規模感

(個人的主観含む)
RESTは小規模~中規模。
GraphQLは中規模~大規模。

#### 3-5. その他

他にもありますが、いまいち理解し切れてないので違いのみ抜粋。

* キャッシュの有無の違い
* クエリの破壊の発生の有無
* 学習コストの違い

## 4. まとめ

両者にはそれぞれ特徴があり、別にRESTからGraphQLに取って代わる、というわけではない。
それぞれに長所と短所があることから、規模感などに合わせてどちらを使うか選ぶということ。

[Amaud Lauret](https://apihandyman.io/about/)氏(APIの世界で著名な方)はこう言っています。

"****"

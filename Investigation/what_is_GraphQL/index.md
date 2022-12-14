- [GraphQL Server](#graphql-server)
  - [1. Basics Tutorial - Introduction](#1-basics-tutorial---introduction)
      - [1-1. GraphQLの特徴](#1-1-graphqlの特徴)
      - [1-2. APIのためのクエリ言語](#1-2-apiのためのクエリ言語)
      - [1-3. REST APIから代替が必要とされた背景](#1-3-rest-apiから代替が必要とされた背景)
      - [1-4. 急成長するコミュニティ](#1-4-急成長するコミュニティ)
  - [2. GraphQL is the better REST](#2-graphql-is-the-better-rest)
    - [2-1. Data Fetching with REST vs GraphQL](#2-1-data-fetching-with-rest-vs-graphql)
      - [2-2. REST APIの場合](#2-2-rest-apiの場合)
      - [2-3. GraphQLの場合](#2-3-graphqlの場合)
    - [2-4. No more Over- and Underfetching](#2-4-no-more-over--and-underfetching)
      - [2-5. Overfetching: Downloading superfluous data](#2-5-overfetching-downloading-superfluous-data)
      - [2-6. Underfetching and the n+1 problem](#2-6-underfetching-and-the-n1-problem)
    - [2-7. Rapid Product Iterations on the Frontend](#2-7-rapid-product-iterations-on-the-frontend)
    - [2-8. Insightful Analytics on the Backend](#2-8-insightful-analytics-on-the-backend)
    - [2-9. Benefits of a Schema & Type System](#2-9-benefits-of-a-schema--type-system)
  - [3. Core Concepts](#3-core-concepts)
    - [3-1. The Schema Definition Language (SDL)](#3-1-the-schema-definition-language-sdl)
    - [3-2. Fetching Data with Queries](#3-2-fetching-data-with-queries)
      - [3-3. Basic Queries](#3-3-basic-queries)
      - [3-4. Queries with Arguments](#3-4-queries-with-arguments)
    - [3-5. Writing Data with Mutations](#3-5-writing-data-with-mutations)
    - [3-6. Realtime Updates with Subscriptions](#3-6-realtime-updates-with-subscriptions)
    - [3-7. Defining a Schema](#3-7-defining-a-schema)
  - [4. Big Picture (Architecture)](#4-big-picture-architecture)
    - [4-1. Use Cases](#4-1-use-cases)
      - [4-1-1. GraphQL server with a connected database](#4-1-1-graphql-server-with-a-connected-database)
      - [4-2. GraphQL layer that integrates existing systems](#4-2-graphql-layer-that-integrates-existing-systems)
      - [4-3. Hybrid approach with connected database and integration of existing system](#4-3-hybrid-approach-with-connected-database-and-integration-of-existing-system)
    - [4-4. Resolver Functions](#4-4-resolver-functions)
    - [4-5. GraphQL Client Libraries](#4-5-graphql-client-libraries)
# GraphQL Server

: [GraphQL](https://graphql.org/)
: [HOW TO GRAPHQL](https://www.howtographql.com/)

## 1. Basics Tutorial - Introduction

#### 1-1. GraphQLの特徴

* RESTよりも効率的で強力かつ柔軟なAPIを提供する**新しいAPI規格**
* **Facebookによって開発、オープンソース化**
* クライアントがAPIから必要とするデータを正確に指定する事ができる
  + 宣言的なデータ取得
* GraphQLサーバは、**単一のエンドポイントのみ**を公開し、クライアントが要求したデータで正確に応答する

#### 1-2. APIのためのクエリ言語

GraphQLはデータベース技術であると混同されるが、誤解である。
**GraphQLはAPIのためのクエリ言語**。

#### 1-3. REST APIから代替が必要とされた背景

RESTのコンセプトが開発された当時は、アプリケーションが比較的シンプルで開発ペースも今日のように速くなかった。そのためRESTの設計は多くのアプリケーションとマッチしていた。
しかし、APIの状況がここ数年で急激に変化している。以下に、変化していった点を3点挙げる。

1. モバイル利用の増加による効率的なデータロードの必要性
FacebookがGraphQLを開発した最初の理由がこれに当たる。

2. 多種多様なフロントエンドフレームワークとプラットフォーム
現在では、アプリケーションを実行するフロンドエンドフレームワークやプラットフォームが異種混在しているため、**RESTではその全ての要件に適合する1つのAPIを構築して維持することが困難**。

3. 迅速な開発 & 迅速な機能開発への期待
現在は多くの企業で継続的なデプロイが標準となり、**迅速なイテレーションと頻繁な製品のアップデートが欠かせなくなっている**。
REST APIでは、クライアント側の特定の要件や設計変更に合わせてサーバがデータを公開する方法を変更する必要があることがよくある。このため、**迅速な開発作業と製品の反復が妨げられる**。

これらの問題を**GraphQLは解決できる**。

#### 1-4. 急成長するコミュニティ

GraphQLはクライアントがAPIと通信するところならどこでも使える技術である。APIとのやりとりをより効率的にするための手段をあらゆる企業が模索していた。(RESTでは解決できなかった。)
GraphQLがオープンソース化した後、名だたる様々な企業で実運用に使われている。

## 2. GraphQL is the better REST

**GraphQLは、より優れたRESTである**

RESTは、過去10年間でWeb APIを設計するための標準となったが、アクセスするクライアントの急速に変化する要件に対応するには**柔軟性に欠ける**ことが分かっている。

GraphQLは、より柔軟で効率的なニーズに対応するために開発された。**REST APIを使用する際に開発者が経験する多くの欠点や非効率性を解決した**。

### 2-1. Data Fetching with REST vs GraphQL

* APIからデータを取得する際のRESTとGraphQLの大きな違いを、簡単なシナリオ例から理解する
  + ブログのアプリケーション
  + 特定のユーザの投稿のタイトルを表示する必要がある
  + 同じ画面に特定のユーザ + 直近の3人のフォロワーの名前も表示される

#### 2-2. REST APIの場合

REST APIでは、複数のエンドポイントにアクセスすることでデータを収集する。(以下のエンドポイントは一例)

1. 特定のユーザのデータを取得する
エンドポイント: `/users/<id>`

2. 特定のユーザのタイトルを取得するために全ての投稿データを取得する
エンドポイント: `/users/<id>/posts`

3. 特定のユーザのフォロワーの名前を表示するためにユーザのフォロワーのリストを取得する
エンドポイント: `/users/<id>/followers`

[引用: REST APIの1~3のデータのやり取り](https://imgur.com/VRyV7Jh.png)
RESTでは必要なデータを取得するために、**異なるエンドポイント**に、**複数回(今回は3回)**のリクエストを送る必要がある。
また、エンドポイントには不必要なデータも含まれているため、**overfetching(過剰取得)**になっている。(引用画像を見ると分かるが、今回のデータ要求に住所や誕生日などは必要がない)

#### 2-3. GraphQLの場合

一方、GraphQLでは、具体的なデータ要件を含む**1つのクエリをGraphQLサーバに送信するだけ**で済む。**GraphQLサーバはJSONオブジェクトで応答する**。

[引用: GraphQLのデータのやり取り](https://imgur.com/z9VKnHs.png)
クエリ内容から分かるように要件を満たすだけの必要なリクエストのみで済んでいる。返ってくる内容も要件を満たし、かつシンプルで見やすい。
※クエリで定義されたネスト構造の通りに正確に返ってくる。

### 2-4. No more Over- and Underfetching

RESTの最も一般的な問題の1つが、オーバーフェッチとアンダーフェッチの問題。これは、クライアントがデータを取得する唯一の手段が**固定的なレスポンスを返すエンドポイントを返す事**によって起こる。
クライアントの要求が都度変わる事に対してRESTの設計で対応することは非常に困難だというのは想像できる。

> "Think in graphs, not endpoints"
> by Lee Byron, GraphQL Co-Inventor

#### 2-5. Overfetching: Downloading superfluous data

Overfetchingは `2-2` の節で説明済みなので割愛する。

#### 2-6. Underfetching and the n+1 problem

* アンダーフェッチ
  + 特定のエンドポイントが必要な情報を十分に提供しないことを指す
この問題が起こる時、要件を満たすまで追加のリクエストを行わなければならない。

n+1問題は例から考えた方が分かりやすかった。
`2-1` の節で用いた例から別の機能を実装しようとするときを考える。

同じアプリでユーザごとの直近3人のフォロワーを表示する必要があった場合。
各ユーザごとを表示させようと思ったら、 `/users` に1回りクエストし、各ユーザの `/users/<id>/followers` にリクエストを送る必要がある。

### 2-7. Rapid Product Iterations on the Frontend

RESTの利点は複数のエンドポイントを用いて、特定のビューに必要な全ての情報を取得できるように出来るため、便利である。
しかし、この設計スタイルはUIを変更する度にエンドポイントの変更をしなくてはならなかったり、バックエンドでデータ取得の調整を行わなければならないなどの必要性が出てくる。

これは現在求められている迅速な開発・迅速な製品のアップデートを行うことを阻害してしまう。

GraphQLを使うとこの問題は解決される。

* クライアントがデータ要件を正確に指定できる
  + サーバ側で余分な作業をしなくてもクライアント側の要求を変えれば良いだけ
* UIが変わってもバックエンドで調整する必要がない

### 2-8. Insightful Analytics on the Backend

GraphQLを使用すると、バックエンド側で要求されたデータについて詳しく知る事が出来る。 -> **分析が出来る**

クライアントからのリクエストは必要なデータを正確に送るため、用意したデータがどのように使われているかを知る事が出来る。あまり要求のされていない特定のフィールドを削除したりなど、**APIを進化させる事に繋がる**。

また、GraphQLを使用すると、サーバで処理されるリクエストの低レベルのパフォーマンス監視を行うことができる。GraphQLではリゾルバ関数の概念を用いてクライアントから要求されたデータを収集するので、このリゾルバのパフォーマンスを計測することで**システムのボトルネックを特定したりすることができる**。

### 2-9. Benefits of a Schema & Type System

GraphQLは強力な型システムを使用している。
ここがGraphQLは**規格**と言う所以だと思う。

APIで公開される全ての型は、"GraphQL Schema Definition Language(SDL)"を使用してスキーマに書き込まれる。**このスキーマは、クライアントがデータにどのようにアクセスできるかを定義する**。

この規格と呼ばれるある種の縛りの恩恵は、これを認識することによりフロントエンドとバックエンドの両チームで**余分なコミュニケーションをとる事なく作業を行う事に繋がる**。

## 3. Core Concepts

この章では、GraphQLの基本的な言語構成について学ぶ。
* 型を定義するための構文
* `queries` & `mutations` を送信するための構文

### 3-1. The Schema Definition Language (SDL)

GraphQLは、APIのスキーマ(Web APIの仕様定義)を定義するために、独自の型システムを持っている。スキーマを記述するための構文は**スキーマ定義言語(SDL)**と呼ばれる。

* SDLを使用してPersonという単純な型を定義する例

```go:
/* Person型は2つのフィールド(name, age)を持つ */
/* それぞれString型, Int型で"!"はこのフィールドが必須であることを意味する */
type Person {
  name: String!
  age: Int!
}

```

- 型と型の間の関係を表現する事も可能

```go:
type Post {
  title: String!
  author: Person! // 型定義したものをフィールドに関連付ける
}
```

* PostにPersonを関連付けたならば、Personにも関連付けなければならない
  + ちょっとよくわからないが、関連付けだけの処理？
  + authorはPersonのデータを持つが、postsはPostと関連付けただけでデータは持たない？

```go:
type Person {
  name: String!
  age: Int!
  posts: [Post!]!
}

```

### 3-2. Fetching Data with Queries

REST APIの場合、データは特定のエンドポイントから読み込まれる。各エンドポイントで、返す情報が決まっている。

GraphQLのアプローチは、**データを返す複数エンドポイントを持っているが、公開するのは単一エンドポイントのみ**。エンドポイントに対して返す情報を固定していないため、これが機能する。その代わり、柔軟性があり、クライアントが実際に必要なデータのみを提供できる。

GraphQLの場合、クライアントは必要なデータのみを取得するために、より多くの情報をサーバに送る必要がある。(=**必要なデータを明示的に指定する**)この情報は、**クエリ**と呼ばれる。

#### 3-3. Basic Queries

* クライアントがサーバに送信するクエリの例

```go:
{
  allPersons {
    name
  }
}
```

allPersonsのフィールドはクエリの**ルートフィールド**と呼ばれる。
ルートフィールドの後に続くものは全て、クエリの**ペイロード**と呼ばれる。
上記のクエリ(allPersonsのname)を指定した場合、レスポンスとして返ってくるのが以下になる。

```go:
{
  "allPersons": [

    { "name": "Johnny" },
    { "name": "Sarah" },
    { "name": "Alice" }

  ]
}

```

各々の名前が返っていて、それ以外の情報は返ってきていない。
**指定された情報がnameだけだったから。**
もし年齢も必要とするなら以下のようにクエリを調整する。

```go:
{
  allPersons {
    name
    age
  }
}
```

GraphQLによる大きな利点の1つが、**ネストした情報を自然にクエリできること。**
例えば、ある人物が書いた全ての投稿を読み込みたい場合、型の構造に従ってこの情報を要求されるだけで良い。つまり、**指定方法がシンプルで分かりやすいという利点がある**。

```go:
{
  allPersons {

    name
    age

    posts {
      title
    }

  }
}

```

#### 3-4. Queries with Arguments

* 引数付きクエリ
allPersonsフィールドに引数を指定した場合、特定の人数までしか返さないようにすることができる。
= ページネーション
[GraphQL Cursor Connections Specification](https://relay.dev/graphql/connections.htm)
参照先: https://relay.dev/graphql/connections.htm

例えば以下の指定方法 `allPersons(last: 2)` は、登録された所から逆順に2人分の名前だけ返すようにクエリを指定している。

```go:
{
  allPersons(last: 2) {
    name
  }
}
```

### 3-5. Writing Data with Mutations

サーバに情報を要求する以外にも、データの変更をする方法が必要。
GraphQLでは、データの変更は**ミューテーション**を使用して行われる。
ミューテーションには以下の3つの種類がある。

1. 新しいデータの作成
2. 既存のデータの更新
3. 既存のデータの削除

ミューテーションはクエリと同じ構文に従うが、**常にミューテーションというキーワードで始まる必要がある**。
以下に新しいPersonを作成する場合の例を挙げる。

```go:
mutation {
  createPerson(name: "Bob", age: 36) {

    name
    age

  }
}

```

前に書いたクエリと同様にこのミューテーションはルートフィールドを持っていて、この場合は `createPerson` が該当する。
この例では、新しくname, ageを定義して、レスポンス要求で更にname, ageを返すように要求しているのであまり役には立たない。(新たに作成したname, ageがそのまま返ってくるだけ)

しかし、**ミューテーションを送信する際に情報の問い合わせが出来る事自体がすごい機能**である。
(いちいちクエリを送る必要がないということ。)

* 上記のミューテーションに対するサーバの応答

```go:
"createPerson": {
  "name": "Bob",
  "age": 36,
}
```

GraphQLのtypeとして定義した型はAutoIncrementの機能がある。
Person型を拡張してidを追加すると、これが一意のIDを持つ。

```go:
type Person {
  id: ID!
  name: String!
  age: Int!
}

```

このようにPersonにidを持たせた事により、ミューテーションで新しくPersonを作成した際にidを求める事ができる。

```go:
mutation {
  createPerson(name: "Alice", age: 36) {
    id // 要求クエリ
  }
}
```

### 3-6. Realtime Updates with Subscriptions

* サブスクリプション
リアルタイムに欲しい情報を得るために**サーバへのリアルタイムな接続**を持ちたい。(例えば重要なイベントなどは即座に確認したい。)
そのためにGraphQLは**サブスクリプションという概念**を提供する。

**クライアントがサブスクリプションをサーバに送信すると、両者の間にコネクションが開かれる。**
サブスクリプションを以下のように送ると、新しいPersonを作成するミューテーションが実行される度に、サーバはこのPersonに関する情報をクライアントに送る。

```go:
subscription {
  newPerson { // Person型にnewというメソッドで新規作成と紐付け？

    name
    age

  }
}

/*  mutation: CreatePerson()が実行された場合のレスポンス */
{
  "newPerson": {

    "name": "Jane",
    "age": 23

  }
}

```

### 3-7. Defining a Schema

これまでのクエリ、ミューテーション、サブスクリプションを用いて具体的にスキーマを定義してみる。
**スキーマは、GraphQL APIを使用する際に最も重要な概念のひとつ**である。

スキーマは、**APIの機能を指定し、クライアントがどのようにデータを要求できるかを定義する**。これはしばしば、**サーバとクライアントとの契約**とみなされる。

スキーマは**GraphQLの型を集めたもの**であるが、APIのスキーマを書く場合には**いくつかの特別なルート型がある**。以下の3つはクライアントから送信されるリクエストのエントリポイントとなる。

* type Query{...} // GET
* type Mutation{...} // POST, PUT, DELETE
* type Subscription{...}

クライアントからQueryでデータ取得を許可する場合は以下のように記述する。(allPersonsに対しての場合)

```go:
type Query {
  allPersons: [Person!]!
}
```

allPersonsは、APIのルートフィールドと呼ばれる。
情報の中から必要な情報を抜き出す際の指定方法を含めて以下のように記述しなければならない。

```go:
type Query {
  allPersons(last: Int): [Person!]!
}

```

同様に、`createPerson-mutation`を使用するには、ミューテーションタイプにもルートフィールドを追加する必要がある。

```go:
type Mutation {
  // 新規作成の際に2つの引数を必要とすることを定義する
  createPerson(name: String!, age:Int!): Person!
}
```

最後に、サブスクリプションを使うために、 `newPerson` をサブスクリプションのルートフィールドに追加する。

```go:
type Subscription {
  newPerson: Person!
}

```

ここまでをまとめた完全なスキーマが以下。

```go:
type Query {
  allPersons(last: Int): [Person!]!
  allPosts(last: Int): [Post!]!
}

type Mutation {
  createPerson(name: String!, age: Int!): Person!
  updatePerson(id: ID!, name: String!, age: String!): Person!
  deletePerson(id: ID!): Person!
}

type Subscription {
  newPerson: Person!
}

type Person {
  id: ID!
  name: String!
  age: Int!
  posts: [Post!]!
}

type Post {
  title: String!
  author: Person!
}
```

## 4. Big Picture (Architecture)

GraphQLは**仕様書のみが公開されている**。
つまり、**サーバーの動作を詳細に記述した文書以上の機能はない**。

### 4-1. Use Cases

このセクションではGraphQLサーバを含む3種類のアーキテクチャを学ぶ。
いずれもGraphQLの主要なユースケースを表していて、GraphQLの柔軟性を示している。

1. データベースと接続されたGraphQLサーバ
2. 既存システムを統合するGraphQLサーバ
3. データベースとの接続、既存システムの統合の2種類の方法でアプローチを行うGraphQLサーバ

#### 4-1-1. GraphQL server with a connected database

GraphQLにおける一般的なアーキテクチャ。
セットアップとして用意するのは、GraphQLの仕様を実装した1台の(Web)サーバ。

1. クライアントからGraphQLサーバにクエリが到着する
2. サーバはクエリのペイロードを読み取る
3. リクエストに応じた必要な情報をデータベースから取得する
1~3の一連の流れを**クエリの解決**と呼ぶ。
4. レスポンスオブジェクト(json形式)を構築し、クライアントに返す。

* GraphQLにおける重要な点
  + 利用可能なあらゆるネットワークプロトコルで使用できる可能性がある点
    - TCPやWebSocketなどをベースにしたサーバの実装の可能性
  + データを保存するために使用されるデータベースや形式を気にしない点
    - SQLデータベース、NoSQLデータベースを使用可能

[引用: 1つのデータベースに接続する、1つのGraphQLサーバを持つアーキテクチャ図](https://imgur.com/cRE6oeb.png)

#### 4-2. GraphQL layer that integrates existing systems

GraphQLのもう1つの主要なユースケースが、既存システムの統合です。
レガシーなインフラストラクチャや、メンテナンスの負担が大きい企業にとっては特に重宝される。

GraphQLは、既存システムを統一してその複雑さを、優れたGraphQL APIの背後に隠すために使用することができる。こうすることで**新しいクライアントアプリケーションを開発し、単にGraphQLサーバと会話して必要なデータを取得することができる**。

[引用: 様々な既存システムの複雑性を単一インターエースで隠蔽する](https://imgur.com/zQggcSX.png)

#### 4-3. Hybrid approach with connected database and integration of existing system

データベース接続 + 既存システム統合のハイブリットアプローチが可能。
**サーバはクエリを受け取るとそれを解決し、接続されたデータベースまたは統合されたAPIの一部から必要なデータを取得する**。

[引用: ハイブリットアプローチ](https://imgur.com/73dByTz.png)

### 4-4. Resolver Functions

GraphQLでこのような柔軟性はどのようにして得られるのか？
GraphQLはなぜこのような異なる種類のユースケースに適合するのか？

GraphQLのクエリ(またはミューテーション)のペイロードは一連のフィールドで構成されている。GraphQLのサーバの実装では、これらの各フィールドは、実際には**リゾルバと呼ばれる1つの関数に対応している**。

リゾルバ関数の唯一の目的は、その**フィールドのデータを取得すること**。

1. サーバはクエリを受け取る。
2. クエリのペイロードで指定されたフィールドに対する全ての関数を呼び出す。
3. クエリを解決し、各フィールドの正しいデータを取得することができる。
4. 全てのリゾルバが返されるとサーバはクエリで記述されたフォーマットでデータをパッケージングし、クライアントに送り返す。

[引用: リゾルバ関数の対応](https://imgur.com/e1gBEP5.png)
サーバにクエリを送り、そのクエリの各フィールドはリゾルバ関数に対応している。つまり、**クエリを送るということは必要なリゾルバを呼び出すということ**。

### 4-5. GraphQL Client Libraries

GraphQLは、オーバーフェッチ、アンダーフェッチを完全に排除するため、フロントエンドの開発者にとっては特に恩恵がある。
クエリは欲しいデータだけを要求すれば良いだけなので、どこにどんなデータがあるからどこのエンドポイントを使用する、などということは**考える必要がない**。

* REST APIの場合のデータ取得の際のステップ

1. HTTPリクエストの作成と送信
2. サーバの応答を受信し、解析する
3. データをローカルに保存する
4. データをUIに表示させる

* 理想的なデータ取得のステップ

1. データ要件の記述
2. データをUIに表示させる

データ保存だけでなく、理想的なデータ取得のステップのように、**低レベルのネットワーキングのタスクは全て抽象化され、データ要求の通りに返されるべき**。
これこそがRelayやApolloのようなGraphQLクライアントライブラリで実現できることである。

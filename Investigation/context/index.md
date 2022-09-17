- [エラー処理・コンテキストについて色々調べてみる](#エラー処理コンテキストについて色々調べてみる)
	- [1. コンテキストとは？](#1-コンテキストとは)
			- [1-1. なんのためにコンテキストはあるのか？](#1-1-なんのためにコンテキストはあるのか)
			- [1-2. **Package contextのドキュメント Overview**を読んで概要を掴む](#1-2-package-contextのドキュメント-overviewを読んで概要を掴む)
			- [1-3. Goroutine](#1-3-goroutine)
			- [1-4. Goroutineの適切なキャンセル](#1-4-goroutineの適切なキャンセル)
			- [1-5. Goroutineのキャンセルパターン](#1-5-goroutineのキャンセルパターン)
			- [1-6. Channels](#1-6-channels)
	- [2. Contextの使い方](#2-contextの使い方)
			- [2-1. Contextの型](#2-1-contextの型)
			- [2-2. Contextの関数](#2-2-contextの関数)
			- [2-3. Contextを使ったサンプルコードで雰囲気を掴む](#2-3-contextを使ったサンプルコードで雰囲気を掴む)
	- [3. リクエスト情報の伝搬](#3-リクエスト情報の伝搬)
			- [3-1. リクエスト情報の伝搬とは？](#3-1-リクエスト情報の伝搬とは)
			- [3-2. どんなデータを伝搬するのか？](#3-2-どんなデータを伝搬するのか)
			- [3-3. なぜ関数にオプションのパラメータを渡すべきではないのか？](#3-3-なぜ関数にオプションのパラメータを渡すべきではないのか)
			- [3-4. context value](#3-4-context-value)
			- [3-5. context value 具体例](#3-5-context-value-具体例)
	- [4. Context Value 使い方](#4-context-value-使い方)
			- [4-1. Context Value 型](#4-1-context-value-型)
			- [4-2. Context Value Function](#4-2-context-value-function)
			- [4-3. 具体例で雰囲気を掴む](#4-3-具体例で雰囲気を掴む)
	- [5. Contextまとめ](#5-contextまとめ)
			- [5-1. キャンセルの役割](#5-1-キャンセルの役割)
			- [5-2. キャンセルの注意点](#5-2-キャンセルの注意点)
			- [5-3. リクエストの情報の伝搬に適した値](#5-3-リクエストの情報の伝搬に適した値)
			- [5-4. ContextValueの注意点](#5-4-contextvalueの注意点)
			- [5-5.適切なContextの使用のために](#5-5適切なcontextの使用のために)
# エラー処理・コンテキストについて色々調べてみる
## 1. コンテキストとは？
Go1.7で標準パッケージになったライブラリのこと
`context`パッケージで定義されている`context`型というものが、何かの処理を行う

#### 1-1. なんのためにコンテキストはあるのか？
引用: [Go Concurrency Patterns: Context](https://free-engineer.life/golang-context/)

> In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request’s deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using.
> At Google, we developed a context package that makes it easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all the goroutines involved in handling a request. The package is publicly available as context. This article describes how to use the package and provides a complete working example.

要約すると、
**キャンセル通知や期限(デッドライン)、その他のリクエストスコープの値を伝達する手段が欲しくて、contextパッケージが作られた**ということです。

#### 1-2. **Package contextのドキュメント Overview**を読んで概要を掴む
引用: [Go context](https://pkg.go.dev/context)

> Package context defines the Context type, **which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.**
- context packageとは以下をするもの
1. 処理の締め切りを伝達
2. キャンセル信号の伝播
3. リクエストスコープ値の伝達

> Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled.
- サーバーへの送受信の際の関数の呼び出しでは、Contextを伝播していくべきである。
- このとき、必要に応じて`WithCancel`, `WithDeadline`, `WithTimeout`, `WithValue`を使用して子Context(派生)を作成する。
- 親Contextがキャンセルされると派生したすべての子Contextもキャンセルされる。

> ...中略
> Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. The Context should be the first parameter, typically named ctx:
- Contextの伝播は、第一引数に`ctx`という名前で渡す

#### 1-3. Goroutine
読み進めると、以下の文が出てきます。
> コンテキストは、複数のゴルーチンによる同時使用に対して安全です。
他の記事も同様に、contextが一番役立つのは、**1つの処理が複数のゴルーチンをまたいで行われる場合**のようです。

- Goroutine
A Tour of Goで学習した内容から。
[Goroutines](https://www.notion.so/var-inc/A-Tour-of-Go-ae4863efa16f49f3b7425944f90425f2#1c3225bcc4354e349a939a8dd98a5c9b)


つまり、
1. 処理の締め切りを伝達
2. キャンセル信号の伝播
この2つは、**Goroutineの適切なキャンセル**という事に言い換えられます。

#### 1-4. Goroutineの適切なキャンセル
- `context`のキャンセルの役割
マルチスレッドの問題点のように、並列実行の場合、パフォーマンスの低下やデッドロックが起こってしまう可能性があります。
`context`には、もしそういった事が起こった場合に**タイムアウトやデッドラインを設定し、後続の処理が停滞するのを防ぎ、不要になったGoroutineはキャンセルしてリソースを解放させる役割**があります

#### 1-5. Goroutineのキャンセルパターン
1. **全ての**Goroutineをキャンセルしたい場合
2. **分岐した先**のGoroutineをキャンセルしたい場合
3. **ブロック(遅い処理)**している処理をキャンセルしたい場合

#### 1-6. Channels
読み進めるとチャネルの概念が出てくるので、A Tour of Goでの学習内容を貼っておきます。

https://www.notion.so/var-inc/A-Tour-of-Go-ae4863efa16f49f3b7425944f90425f2#e7400c3ec42b44bea96e64b8cf00cb26

## 2. Contextの使い方
#### 2-1. Contextの型
```go: context
type Context interface {
  /* Deadlineが設定されている場合はその時刻を返却 */
  /* 設定されていない場合は ok==falseを返す */
  Deadline() (deadline time.Time, ok bool)

  /* このコンテキストがキャンセル、タイムアウトした場合にはcloseされる */
  Done() <-chan struct{}

  /* Doneチャンネルが閉じた後、なぜこのコンテキストがキャンセルされたかを知らせる */
  Err() error

  /* ValueはKeyに紐付いた値を返し、設定した値がない場合はnilを返す */
  Value(key interface{}) interface{}
}
```
上から言える事。
**`Context`インターフェースには、`Deadline`, `Done`, `Err`, `Value`の4つの関数がある。**
- `Deadline`は期限を返す
- `Done`はチャネルを返す。
  - このチャネルはキャンセル・期限切れの場合に閉じられる。
- `Err`は`Done`チャネルが閉じられた理由を返す
- `Value`は`Context`に格納した値を返す

#### 2-2. Contextの関数
1. `func WithCancel`
`func WithCancel(parent Context ) (ctx Context , cancel CancelFunc )`
新しい`Done`チャネルを持った子Context(親Contextのコピー)と、`CancelFunc`を返します。
この子Contextの`Done`チャネルが閉じられるのは、`CancelFunc`が呼び出されたときか、親Contextの`Done`チャネルが閉じられた時です。

2. `func WithDeadline`
`func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)`
`d`以内を起源とした子Context(親Contextのコピー)と`CancelFunc`を返します。
もし親Contextの期限が`d`よりも早ければ親Contextと同じになります。
- 返された子Contextの`Done`チャネルが閉じられるパターン
  - 期限を過ぎた場合
  - `CancelFunc`が呼び出された場合
  - 親Contextの`Done`チャネルが閉じられた場合

3. `func WithTimeout`
`WithTimeout`は`WithDeadline`のエイリアスみたいなものです。
== `WithDeadline(parent, time.Now()Add(timeout))`

4. `func Background`, `func TODO`
どちらも空のContextを生成します。
`Background`が通常使用されます。
一般的にメイン関数や初期化などでトップレベルのContextを生成するために使用します。
- `TODO`
> TODO returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter).
どのContextを使うべきか不明な場合や、Contextがまだ利用できない場合に`TODO`を使用する必要があります。
引用: [Go context](https://pkg.go.dev/context)

5. `type CancelFunc`
`type Cancel = context.CancelFunc`
Contextをキャンセルします。

#### 2-3. Contextを使ったサンプルコードで雰囲気を掴む
```go: Context
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	/* contextを2つ生成 */
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, _ := context.WithTimeout(context.Background(), 3*time.Second)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		childProcess(ctx1, "プロセス1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		childProcess(ctx2, "プロセス2")
	}()

	/* ctx2は3秒でタイムアウトするので、スリープ中にキャンセルされる */
	time.Sleep(5 * time.Second)
	/* 明治的にcancelFuncを呼び出して、ctx1をキャンセルする */
	cancel1()
	wg.Wait()
}

func childProcess(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			/* キャンセル時の処理 */
			fmt.Printf("%s canceld. error: %s\n", name, ctx.Err())
			return
		case <-time.After(1 * time.Second):
			fmt.Printf("%s processing...\n", name)
		}
	}
}

/* 実行結果 */
/*
プロセス1 processing...
プロセス2 processing...
プロセス2 processing...
プロセス1 processing...
プロセス2 processing...
プロセス1 processing...
プロセス2 canceld. error: context deadline exceeded
プロセス1 processing...
プロセス1 processing...
プロセス1 canceld. error: context canceled

Program exited.
*/
```

## 3. リクエスト情報の伝搬
ここまででContextの主な用途
- Goroutineの適切なキャンセル
- リクエスト情報の伝搬
この2つの内、Goroutineの適切なキャンセルについて学びました。
ここからは、**リクエスト情報の伝搬**について学びます。

#### 3-1. リクエスト情報の伝搬とは？
**リクエストが来た際に、userIDなどをcontextに保存し、後続の処理で使う時に使用する**イメージ。

#### 3-2. どんなデータを伝搬するのか？
引用: [Go言語による並列処理](https://www.amazon.co.jp/dp/4873118468)
> Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
> コンテキスト値はプロセスやAPIの境界を通過するリクエストスコープでのデータに絞って使いましょう。
> 関数にオプションのパラメータを渡すために使うべきではありません。

#### 3-3. なぜ関数にオプションのパラメータを渡すべきではないのか？
1. contextは複数のgoroutineから参照される
- **返却されるvalueの値は複数のgoroutineからアクセスがあっても問題がない値にするべきである**
 - イミュータブルな値でないとアクセスした際に予期しない挙動になる可能性がある。
**つまり、contextに入れる値はイミュータブル(不変)な値にする必要がある。**

2. context valueから取得した値を使って関数を呼び出したり、振る舞いが変わってしまう場合、それは関数にオプションのパラメータを渡していると言える。
**そのため、contextに入れる値は、その値によって振る舞いが変わるものを入れてはいけない**

#### 3-4. context value
リクエストで受け取って、その後の処理で共通して使いたくなる値を考えてみます。
どれも後続の処理で共通して使いたくなる値で、一度はcontext valueに入れてみたくなります。

- ユーザID
- リクエストID
- 認証情報
- リクエストトークン
- URL
- Logger
- DBの接続情報
- など

こういった値から、**context valueに格納して良い値**はどれが該当するのでしょうか？
3-3の1, 2の条件に合う値(ユーザID, リクエストID, 認証情報, リクエストトークン)がcontext valueに格納して良い値となります。

#### 3-5. context value 具体例
また、Go Kitの作成者であるPeter Bourgon氏の記事にも具体的な内容が書かれていました。[Peter Bourgon Context](https://peter.bourgon.org/blog/2016/07/11/context.html)

> Good examples of request scoped data include user IDs extracted from headers, authentication tokens tied to cookies or session IDs, distributed tracing IDs, and so on.
> リクエストスコープのデータの良い例としては、ヘッダーから抽出したユーザーID、クッキーやセッションIDに結びついた認証トークン、分散トレースIDなどがあります。

具体例があると分かりやすいです。
- Headerから取得したUserID
- cookieやsessionに紐づく認証トークン
- 分散処理の追跡ID
などが良い例として挙げられています。

> I think a good rule of thumb might be: use context to store values, like strings and data structs; avoid using it to store references, like pointers or handles.
> 文字列やデータ構造などの値を格納するためにコンテキストを使用し、ポインタやハンドルなどの参照を格納するためにコンテキストを使用しない、というのが良い経験則になるかと思います。

context valueには、値を保存するために使用し、**参照のために使用することは避けた方が良い**と書かれています。

## 4. Context Value 使い方
#### 4-1. Context Value 型
```go: context value
type Context interface {
  /* ...今回説明しない箇所は省略 */

  /* Valueはkeyに紐付いた値を返し、設定した値がない場合はnilを返す */
  Value(key interface{}) interface{}
}
```
※返ってくる値がinterface型なので注意が必要
`Context.Value(key)`でkeyに紐付く値を返します。
マッチするkeyが存在しないときはnilを返します。

#### 4-2. Context Value Function
- `func WithValue`
`func WithValue(parent Context, key interface {}, val interface{}) Context`
第一引数に`key-value`を格納するContextを指定し、`key-value`をセットします。
そして、セット済みのcontextが返却されます。

#### 4-3. 具体例で雰囲気を掴む
```go: Context Value
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctxValue1 := context.WithValue(ctx, "hoge", 1)
	ctxValue2 := context.WithValue(ctxValue1, "piyo", 2)
	ctxValue3 := context.WithValue(ctxValue2, "fuga", 3)
	ctxValue4 := context.WithValue(ctxValue3, "fuga", 4) // fugaを上書き
	go sayValue(ctxValue4)

	time.Sleep(2 * time.Second)
}

func sayValue(ctx context.Context) {
	for {
		fmt.Print("hoge", ctx.Value("hoge"), " : ")
		fmt.Print("piyo", ctx.Value("piyo"), " : ")
		fmt.Println("fuga", ctx.Value("fuga"))
		time.Sleep(1 * time.Second)
	}
}

/* 実行結果 */
// hoge1 : piyo2 : fuga 4
// hoge1 : piyo2 : fuga 4
```
`WithValue()`でkey-valueをセットした新たなcontextを生成している。
`ctxValue4`では、key`"fuga"`を指定し、value`4`に上書きしています。
`ctx`で呼び出していますが、子で上書いた値が優先されている事が分かります。

## 5. Contextまとめ
- Contextの主な用途は2つ
1. Goroutineの適切なキャンセル
2. リクエスト情報の伝搬

#### 5-1. キャンセルの役割
**タイムアウトやデッドラインを設定し後続の処理が停滞するのを防ぎ、不要になったGoroutineはキャンセルしてリソースの解放をする役割**を持つ

#### 5-2. キャンセルの注意点
- Contextを構造体にの中に定義することはしてはいけない
contextは複数のgoroutineから参照する事ができます。
contextを含んだ構造体を乱立させると複雑な親子関係や意図しないキャンセルが発生する可能性が高くなります。
なので**第一引数で渡していく方法**が推奨されています。

- contextを第一引数に
**Contextの伝播は、第一引数に`ctx`という名前で渡す**のが決まりです。

#### 5-3. リクエストの情報の伝搬に適した値
ContextのValueはどんな値なら入れて良いのか？
- Headerから取得したUserID
- cookieやsessionに紐づく認証トークン
- 分散処理の追跡ID
これらの例が適切なvalueとして挙げられています。

#### 5-4. ContextValueの注意点
Context.Valueは`set`で値を設定し、`get`で取り出せます。
- `set`
  - `WithValue(parent Context, key, val interface{}) Context`
- `get`
  - `Value(key interface{}) interface{}`

`Value`で取り出す際に注視すべきは型がinterfaceである点です。
**型を区別できないので、コンパイル時にバグを発見することができない**ということです。

そのため、goroutine同士がcontext valueにセットした情報を知る必要があります。

#### 5-5.適切なContextの使用のために
1. contextは第一引数に渡し、構造体の中に定義してはいけない。
2. `context.Value`を使う上で注意することは、**値へのアクセスを制限する**, **ちゃんと型を持たせる**。

# CodeReviewCommentsを読む
## 1. Gofmt
自分が書いたコードに、`gohtm`を実行すると、機械的に直すことのできるスタイルの大部分を自動的に修正してくれる。
機械的に直すことのできない部分に関しては、代わりに`goimports`を使う手段がある。
`gofmt`に加えて、必要に応じてimports内に改行をつけたり消したりする機能がある。

## 2. Comment Sentences
宣言を文書化するcomment文は、多少冗長に見えても完全な文にする必要がある。
これに沿って書く事で適切にフォーマットされる。
commentは、**説明されているものの名前で始まり、ピリオドで終わる必要がある**。

```go: Comments Sentences
// Request represents a request to run a command.
type Request struct { ...

// Encode writes the JSON encoding of req to w.
func Encode(w io.Writer, req *Request) { ...
```

## 3. Contexts
Contextを調べた内容と同じ。

## 4. Copying
別のパッケージから構造体をコピーする時には注意が必要。
(予期しない参照を防ぐため)
メソッドがポインタの値に関連付けられているなら`T`ではなく`*T`を使用するようにする。

## 5. Crypto Rand
鍵の生成に`math/rand`を使用してはいけない。
**`crypto/rand`の`rand.Reader`を使用する**
文字列として扱いたい場合は、**16進数かbase64エンコード**を行う。

## 6. Declaring Empty Slices
- 空のリストを宣言したい時
  - 推奨
    - `var t []string`
  - 非推奨
    - `t := []string{}`

推奨される方はもしスライスに何も追加されなくても、余計なメモリを消費しないため。
空のリストを宣言した場合、推奨 -> nil, 非推奨 -> [] となる。
インターフェースを設計する時は両者は区別しない方が良い(分かりづらいミスを誘発する可能性があるため)

## 7. Doc Comments
"Doc comment"は、**トップレベル**のpackage, const, func, type, varの宣言の直前に**改行を入れずに書かれたもの**である。
**全ての公開された(大文字の)名前には"Doc comment"があるべき**である。

- Goは、`/*  */`と`//`の行コメントを提供する。
- **トップレベルの宣言の前に表示され、間に改行がないコメントは、宣言自体を文書化すると見なされる**。

## 8. Don't Panic
**通常のエラー処理に`panic`を使用してはいけない**。
なるべく`error`型を含んだ複数の値を返すようにするべき。

## 9. Error Strings
エラー文字列は、**頭文字を大文字にしたり、句読点で終わってはいけない**。
なぜなら、この文字列は他のコンテキストで使われるケースがあるからである。

どういう事かと言うと、
`log.Print("Reading %s: %v", filename, err)`が途中で大文字を出力する事がないように、`fmt.Errorf("Something bad")`ではなく、`fmt.Errorf("something bad")`とするべきである。
※ロギングのような暗黙的な行指向や、他のメッセージと組み合わせないものには適用されない。
## 10. Examples
新しいパッケージを追加した時、使い方のサンプルを含めると良い。
実行可能な例や、一連の流れを追えるテストなどが該当します。

## 11. Goroutine Lifetimes
`goroutine`を生成する時は、いつ終了されるかを明確にすべき。
そうしないと、バグが起こるか、他の問題を引き起こす事がある。

並行処理は`goroutine`の生存期間が明確になるように、十分にシンプルに書くべき。
それができない場合は、いつどんな理由で`goroutine`が終了するかドキュメントに書きましょう。

## 12. Handle Errors
アンダースコア変数`_`(棄却処理)で、エラー無視をしてはいけない！きちんとエラー処理を行い、チェックすべき！

## 13. Imports
packageをリネームするのはやめましょう。
importする際は、空行でいくつかのグループに分けましょう。
標準パッケージは最初のグループにしましょう。
**`goimports`がこれをやってくれる！**
```go: Imports
package main

import (
    "fmt"
    "hash/adler32"
    "os"

    "appengine/foo"
    "appengine/user"

    "github.com/foo/bar"
    "rsc.io/goversion/version"
)
```
## 14. Import Blank
`import _ "pkg"`の書き方で、パッケージをインポートした際の副作用だけを利用する方法がある。
この方法は、`メインパッケージ`、または`テスト`でのみ利用しましょう。

## 15. Import Dot
```go: Import Dot
import (
  . "fmt"
  "string"
)
```
このように書くと通常`fmt.Println("hello world")`と書く所を`Println("hello world")`のように書く事ができる。

このようなやり方は
1. 可読性が極端に落ちる
2. period importしたパッケージに含まれる関数と同名の関数を宣言できなくなる
このような理由から、循環参照によってパッケージの一部がテストできないときに利用するのが望ましい。

## 16. In-Band Errors
C言語などでは、一般的なテクニックとして`-1`や`null`を返すことでエラーや結果を発見できなかったことを表現する場合があるが、Goでは推奨されません。
なぜならGoは複数の値を返す事ができるからです。

より良い方法は、`Error`型か`bool`を返して正しく取得できたか示す値を追加して返す方法です。
※戻り値の中で最後に位置するべき。
```go: In-Band Errors
// Lookup returns the value for key or ok=false if there is no mapping for key.
func Lookup(key string) (value string, ok bool)
```
## 17. Indent Error Flow
エラーハンドリングの分岐を最小にしよう！
- 良くない例
  - if -> エラーハンドリング
  - else -> 通常のコード

- 良い例
  - if -> エラーハンドリング
    - return
  - エラー処理と分断して通常のコード
```go: Good code
if err != nil {
    // error handling
    return // or continue, etc.
}
// normal code
```

## 18. Initialisms
`URL`は、`URL`か`url`にする。`Url`にしてはいけない。
同様に、`appId`ではなく、`appID`にするべき。

ProtocolBufferによって生成されたコードは上記のようなルールから免除されます。
人が書いたコードには、機械が書いたコードよりも厳しい基準が要求される。

- Protocol Buffer
  - 簡単に言うと、構造化データをバイト列に変換(シリアライズ)するための技術

## 19. Interfaces
**モック**のためにAPIの実装側にインターフェースを定義してはいけない。
代わりに、実際に実装された公開APIからテストできるようなAPIを設計する。
使われる前にインターフェースを定義してはいけない。

- モック
  - あるクラスと同じインターフェースを持つダミーのインスタンスの事
  - 主にテストコードの中で活用され、一部の処理をテスト用に置き換えるために使用される

## 20. Line Length
Goでは1行の長さを決めていないが、長すぎないようにしましょう。
ただし、**1行を短く保ちたいがために無理に改行を入れる必要はない**。
**行の意味によって改行すべき**。

## 21. Mixed Caps
公開しない定数は`maxLength`としましょう。

## 22. Named Result Parameters
関数内で戻り値のための変数の宣言を省略したいがために、名前付き戻り値を使うのはやめましょう。

```go: Bad
func (f *foo) Location() (float64, float64, error)
```

```go: Good
// Location returns f's latitude and longitude.
// Negative values mean south and west, respectively.
func (f *foo) Location() (lat, long float64, err error)
```
**名前付き戻り値は、戻り値が何を意味をするか分かりづらい時に使う！**

## 23. Naked Returns
変数無しの`return`文は、関数定義で宣言した変数の値を返します。

```go: Naked Returns
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}
```

## 24. Package Comments
Godocのパッケージのコメントは、パッケージ節の前に**改行無し**で書かなくてはいけない！
```go: Package Comments
// Package math provides basic constants and mathematical functions.
package math
```
- Godocのコメントは公開されるものなので
  - **正しい英語で書け！**
  - **先頭は大文字にしろ！**

## 25. Package Names
- パッケージから公開されている全ての識別子への参照は、パッケージ名を通して行われる
  - 識別子にパッケージ名を使っている場合は外すべき
例:
`chubby`というパッケージの中で`ChubbyFile`という名前にすると、呼び出すときに`chubby.ChubbyFile`となるので良くない。
`chubby.File`などが良い例になる。

## 26. Pass Values
たかだが数バイトを節約するために引数にポインタを指定するのはやめましょう。
安易に引数`x`を`*x`として、引数をポインタにすべきではない。

ただし、大きな構造体や今後大きくなりそうなものはこれに限らない。

## 27. Receiver Names
- レシーバとは
  - メソッドを呼び出される対象のことを「レシーバ」と呼ぶ。
  - なぜそう呼ばれるかと言うと、昔のOOP言語の文脈ではメソッド呼び出しの事を「メッセージの送信」と言い、メソッドを呼び出される側は「メッセージの受信側(レシーバー)」だから、という背景があります。
```go: Receiver
p := Person{Name: "Taro"}
p.Greet("Hi")
/* ↑のpがレシーバです。 */
```
**レシーバの名前はそれがなんであるか適切に表したものでなければいけない**
- レシーバ名は大抵1~2文字で足りる
- `Client`であれば、`me`とか`this`とかではなく、`c`とか`cl`にしよう
- レシーバ名は統一しよう
  - `Client`を`c`としたなら、違う所で`cl`などと名付けてはいけない

## 28. Receiver Type
**レシーバのガイドライン**
メソッドのレシーバをポインタにするか、値にするかの基準
- レシーバが`map`, `func`, `channel`ならポインタは使わない
- メソッドが値を変更するならポインタ一択
- `sync.Mutex`などの`sync`系はポインタ一択
- レシーバが大きな構造体や配列ならポインタが良い
- 値型の`time.Time`や`int`, `string`など、基本的な型の場合のレシーバは値の方が良い場合がある。
  - 値型だからという理由だけでプロファイルで確認する前にレシーバにするのはやめましょう。
- 迷った時はポインタにしよう
- レシーバの型は統一しよう

## 29. Synchronous Functions
非同期処理も同期処理(結果を直接返すか、値を返す前にコールバックやchannelの操作が終わっている関数)を好む。
呼び出し側で別のgoroutineから呼び出してあげるのが良い。

## 30. Useful Test Failures
将来の自分のコードをデバッグする人に有効なメッセージを届ける責任が自分にあることを意識すべし！

テストが失敗した時には以下を伝えよう。
- 何が悪かったのか
- どんな入力があったのか
- 実際にどんな値が来たのか(got or actual)
- どんな値が来ることを期待していたのか(want or expected)

## 31. Variable Names
Goでの変数はより短くあるべき。
`lineCount`よりも`c`, `LineIndex`よりも`i`。

**基本的なルール**
- 宣言された箇所よりも離れた場所でその変数が使われるなら、その変数が分かりやすいようにより説明的な名前にするべき
- メソッドレシーバの変数名は1~2文字で十分
- ループの添え字や読み込みのリーダーのような一般的な変数は(`i`, `r`のように)1文字で良い
- より珍しい物や、パッケージ外で使われるものの名前はより説明的にする

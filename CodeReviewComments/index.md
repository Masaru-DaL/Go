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
## 18. Initialisms
## 19. Interfaces
## 20. Line Length
## 21. Mixed Caps
## 22. Named Result Parameters
## 23. Naked Returns
## 24. Package Comments
## 25. Package Names
## 26. Pass Values
## 27. Receiver Names
## 28. Receiver Type
## 29. Synchronous Functions
## 30. Useful Test Failures
## 31. Variable Names

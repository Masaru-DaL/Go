# エラー処理・コンテキストについて色々調べてみる
## 1. エラー処理について
#### 1-1. Goにおけるエラーとは？
- **Go言語には例外が存在しない**
例外 -> pythonで言う所の`try, except`

つまり、
**例外以外でのエラーハンドリング(エラー処理)が必要！**ということ。

#### 1-2. どうやってエラーハンドリングするの？
**Goでは、`error`と呼ばれる型を定義することで、エラーを表す事ができる。**
そして、Goでいうエラーとは、以下のように定義された`error`インターフェースです。
```go: error
type error interface {
  Error() string
}
```
要するに、`string`を返す`Error`メソッドを実装しなさい。という意味のようです。

エラーハンドリングで使用する際は、**返されるerror変数とnilを比較する**ことで、操作が成功したか判断する。(エラーの可否)
ここがスタートであり、Goでのエラーハンドリングの根本の所。

#### 1-3. どういったものがあるか見てみる
1. `os.Open`関数は、ファイルのオープンに失敗した時に、nilではないerror変数を返します。

```go: os.Open
func Open(name string) (file *file, err error)
```

2. `os.Open`を使って、ファイルを1つオープンする。もしエラーが発生した場合、`log.Fatal`(パッケージ`log`で定義されているログを出力する関数の1つ)を呼び出す事で、エラー情報を出力する事ができる。

```go: os.Open(log.Fatal)
f, err := os.Open("filename.ext")
if err != nil {
  log.Fatal(err)
}
```

#### 1-4. Goにおけるエラー処理は簡単に行える
> 標準パッケージの全てのエラーを発生させうるAPIはどれもerror変数を返すので、簡単にエラー処理を行うことができる。
参考: https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/11.1.html

#### 1-5. いったんまとめる
1. Goでは、エラーが発生した時、戻り値として`error`型の値を返す
2. `error`の戻り値を`if`文でチェックする(チェック項目はerrorがnilかどうか！)

この手法がGoのエラーハンドリングの定番である。

## 2. エラーハンドリングの手段
1. 元々用意されている関数を使用する
2. 自分で独自のエラー処理を作成する

この2パターンが主な手段です。

#### 2-1. パターン1を調べる
関数でエラーを返すには、いくつかの方法があるようです。
1. error.New
これはerrorsパッケージで用意された関数です。
`error.New("エラーです！")`
このように、引数に文字列を指定すると、エラーが起こった場合にその文字列が返ります。

```go: error.New
package main

import (
    "errors"
    "fmt"
)

func main() {
    if err := f1(); err != nil {
        fmt.Println(err)
    }
}

func f1() error {
    return errors.New("エラーです")
}

/* 実行結果 */
// エラーです
```

2. fmt.Errorf
こちらはfmtパッケージで用意された関数です。
`fmt.Errorf("err: %s", "エラーです")`
このように、引数の内容で整形した文字列のエラーを返します。(`%s`は文字列にフォーマットするということ)

```go: fmt.Errorf
package main

import "fmt"

func main() {
  if err := f1(); err != nil {
    fmt.Println(err)
  }
}

func f1() error {
  return fmt.Errorf("err: %s", "エラーです")
}

/* 実行結果 */
// err: エラーです
```

#### 2-2. どっち使えばいいの？？
**どっちも同じ事ができる**
`fmt.Errorf`でできることは、出力される結果をフォーマットできるということ。
`error.New`の引数に`fmt.Sprintf`を使用することで同様の事ができる。

- 結論
個人的には出力させる時点で`fmt`パッケージを使うので、フォーマットさせたい場合は素直に`fmt.Errorf`でいいような気がしてます。
ただの文字列で処理したい場合は`error.New`を使用すれば良いと思われます。

#### 2-3. パターン2を調べる
独自で行う場合、所謂カスタム型、カスタム定義と呼ばれる方法です。
実装する場合は、基本をもう一度押さえておく必要があります。

```go: error
type error interface {
  Error() string
}
```
このerror型を元に、エラーの情報を文字列で返す**`Error`メソッドの実装が必要**です。

#### 2-4. カスタムエラー定義を実装してみる
- ファイル検索に引っ掛からなかった場合に出すエラー処理
```go: error(custom)
package main

import (
	"fmt"
)

func main() {
	if _, err := readFile("/invalid/path"); err != nil {
		fmt.Println(err)
	}
}

func readFile(path string) (string, error) {
	// （中略）

	// ファイルが見つからない場合
	if true {
		// error型に準拠した構造体を返却
		return "", &FileNotFoundError{Path: path}
	}

	// （中略）

	return path, nil
}

/* custom error */
type FileNotFoundError struct {
	Path string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("%s: ファイルが見つかりません", e.Path)
}
/* custom error end */

/* 実行結果 */
// /invalid/path: ファイルが見つかりません
```
1. `FileNotFoundError`という型で、Pathを文字列で定義します。
2. すぐ下で`FileNotFoundError`を型に指定してエラー出力するメソッドを実装しています。

## 3. エラーのラッピング
エラーハンドリングでは、ビルドイン関数の使用/自作のエラー定義を作成する、という2つ以外に、**エラーのラッピング**というものがあるようです。

#### 3-1. 概要
- エラーを他のエラーでラップすることが可能
  - 元のエラーの型やフィールド値を保持しつつ新しいエラーを生成できる
  - ラップされたエラーは、あとから取り出す事が可能

- 呼び出し先の深い場所で起きたエラーの種類を判別しやすくなる

#### 3-2. ラッピングの種類
1. エラーをラップする
2. エラーをアンラップする
3. エラーを探索する

大きく分けるとこの3種類に分けられます。

#### 3-3-1. エラーをラップする事の利点
- エラーをラップする事の利点
運用しやすくなるというのが大きな利点だと思います。

クレジットカードの有効期限が切れているエラーのデフォルトエラーが以下だったとします。
`{ "error": "not found book: rpc error: code = NotFound desc = "projects.../databases/(default)/documents/books/not-found-book-id" not found", "message": "not found book: rpc error: code = NotFound desc = "projects/.../databases/(default)/documents/books/not-found-book-id" not found"}"`

こんな冗長的で長いエラーを処理しようとすると大変そうだなぁというのは分かると思います。
このエラーを独自に以下のエラーに出来たらどうでしょうか。
`{ "errorCode":"ErrExpiredCard", "message":"カードの有効期限が切れています"}`

一目見ただけで分かりやすく、またこのメッセージをフロント側で受け取って更に利用しようとした時もデフォルトエラーよりも楽になります。
これがエラーをラップすることの利点です。

#### 3-3-2. エラーをラップしてみる
ビルドイン関数の`fmt.Errorf()`を用いる事でエラーをラップする事が出来ます。

```go: error wrap
innerErr := fmt.Errorf("Inner Error")
newErr := fmt.Errorf("Wrapped: %w", innerErr)
```
 通常、`fmt.Errorf()`は`error.errorString`側のエラーを返すが、以下の条件を満たす場合、`fmt.WrapError`型のエラーを返す。
- フォーマット指定子に`%w`を使用する。(文字列内での位置は問わない)
- その`%w`を置き換えるのが`error`型である(上の`innerErr`)

つまり、最初に定義した`innerErr`を、`newErr`でラップしている。
条件である`%w`を指定し、`innerErr`から`newErr`を定義している事で条件を満たしている。ということ。

#### 3-3-3. `fmt.WrapError`とは？
ラップ機能が追加された`error`型のこと。

通常、エラー処理をラップする処理を行おうと思い、以下のようにしたとする。
```go: error
err := fmt.Errorf("Inner Error")
newErr := fmt.Errorf("Wrapped: %s", err)
fmt.Println(newErr)   // 'Wrapped: Inner Error'
```
これだとエラーが起きた場合に`Inner Error`と返ってくるだけで、どんなエラーが起きたのかの判定が難しい。細かいエラー毎にエラー定義しなくてはいけないの？というのが思い浮かぶ。

しかし、`fmt.wrapError`の構造体を見てみると、以下の仕組みを持っているため、元のエラーがなんであるかを確認する事ができる。
```go: fmt.wrapError(struct)
type wrapError struct {
	msg string
	err error /* ここに元のエラーが入る */
}
```

#### 3-4-1. エラーをアンラップするとは
アンラップするとは、ラップされた中身のエラーを取り出す事を言います。
```go: fmt.wrapError(struct)
type wrapError struct {
	msg string
	err error /* ここに元のエラーが入る */
}
```
上記構造体の`err`を取り出す、ということです。
`fmt.wrapError`は、アンラップ時のエラーの取り出しが簡単になっている、という特徴があります。

#### 3-4-2. アンラップ
アンラップするには、`errors.Unwrap()`を使います。
```go: errors.Unwrap()
errors.Unwrap(newErr error) error
```

#### 3-4-3. アンラップのやり方
```go: Unwrap
iErr := fmt.Errorf("Inner Error")
newErr := fmt.Errorf("Wrapped: %w", iErr)   // "Wrapped: Inner Error"
innerErr := errors.Unwrap(newErr)   // "Inner Error"
```
- `errors.Unwrap()`の引数にラップした`newError`を渡しています。
  - 渡されたエラーが`Unwrap()`を実装していれば、その戻り値を返します。
    - 渡されたエラーは`fmt.wrapError`型となっているので、条件を満たします。
  - `newError.err`、つまりラップ元の`iErr`を返します。
  - もし`Unwrap()`が未実装なら`nil`を返します。

#### 3-5-1. エラーの探索
ラップされた一連のエラーの中から目的のエラーを探し出す事ができます。
そのための2つのメソッドが提供されています。

1. `error.Is(err, target error)`
2. `error.As(err error, target interface{})`

なぜ、この2つのメソッドを使用するのか？と言うと、今までのエラーハンドリングでは、ラップしたエラーに対しては上手くエラーハンドリング出来ないからです。

#### 3-5-2. `errors.Is`
- 特定のエラーを探す
**エラーチェーン(エラーが連続して発生している状態)の中に、特定のエラーが存在するかどうかを調べる。**

出力は`true` or `false`

```go: errors.Is
func main() {
    errA := fmt.Errorf("First Error")   // これをラップしているかどうかを判定したい

    errB := fmt.Errorf("Wrapped: %w", errA)   // B <- A
    fmt.Printf("errB wraps errA = %v\n", errors.Is(errB, errA))

    errC := fmt.Errorf("Wrapped: %w", errB)   // C <- B <- A
    fmt.Printf("errC wraps errA = %v\n", errors.Is(errC, errA))

    errX := fmt.Errorf("Another error")       // A をラップしてない
    fmt.Printf("errX wraps errA = %v\n", errors.Is(errX, errA))

		/* 実行結果 */
		// errB wraps errA = true
		// errC wraps errA = true
		// errX wraps errA = false
}
```
- `errors.Is()`の引数に2つのエラーを指定する
  - 値が同じなら`true`を返す
  - 値が異なるなら`false`を返す

#### 3-5-3. `error.As()`
`error.Is()`は、特定のエラーがエラーチェーン内にあるかどうかを判定するだけです。
その特定のエラーを取得し、さらになんらかの形で利用したい場合に、`error.As`を利用します。

- `error.As`
**特定の型を持つエラーが含まれるか判定し、それを取得する**
`errors.As(err error, target interface{})`
1. エラーチェーンを辿っていき、`target`に代入可能なエラーが見つかったらそれを`target`で受け取る。
2. 戻り値は`bool`で、目的のエラーが見つかった場合に`true`, 見つからなければ`false`が返る。

- `target`の指定方法
  - interface型へのポインタ
  - `error`型を実装した型へのポインタ

```go: errors.As
type myError struct {
    t time.Time
}

func (e myError) Error() string {
    return fmt.Sprintf("[%s] myError found!", e.t.Format("2006/1/2 15:04:05"))
}

func main() {
    var innerErr myError

    err := wrap1()

    if errors.As(err, &innerErr) {
        fmt.Printf("%s\n", innerErr)
    } else {
        fmt.Printf("Unknown error : %s\n", err)
    }
}

func wrap1() error {
    err := wrap2()
    return fmt.Errorf("wrap1: %w", err)
}

func wrap2() error {
    return myError{t: time.Now()}
}

/* 実行結果 */
// [2009/11/10 23:00:00] myError found!
```

## 4. エラーハンドリングのまとめ
- エラーハンドリング方法
1. ビルドイン関数を使う
2. 独自の関数を作成し、使用する
3. エラーのラッピングを使用する

大まかに分けるとこのようになります。
デバッグ程度のエラー処理ならビルドインでいいでしょうし、慣れてきたら自作のエラー処理を行う。
また、大規模になればなるほど、エラー自体を利用した運用というのが必要になってくるので、基本的にはエラーのラッピングを行い、運用するというのが推奨されるやり方です。

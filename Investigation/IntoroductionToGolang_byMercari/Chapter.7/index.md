# メルカリ作のプログラミング言語Go完全入門 読破
# 7. エラー処理
## 7-1. エラー処理
#### 7-1-1. 正常系と異常系
- 正常系
  - エラーが起こらず動作した場合のこと
  - ユーザが意図通り(期待通り)に使った場合の挙動(

- 異常系
  - 意図しない挙動
  - ユーザが意図通り(期待通り)に使わなかった場合の挙動
  - 外部要因による意図しないエラー
    - ネットワーク、ファイル、ライブラリなどのバグ
  - バグが起因のエラー

#### 7-1-2. エラー処理の必要性
- エラーは必ず起きる
  - 外部要因で起きる可能性がある
    - 意図通りに動作したとしても
  - 正常系より異常系の方が難しいし、パターンが多いことがある

- エラー処理
  - 致命的なエラーじゃなければ処理を続けられる場合もある
  - リトライをかけたり、別の方法を取ることもできる
  - エラーに対して適切な処理を行い、出来る限り処理が続けられるようにする

**エラーハンドリングの重要性**

#### 7-1-3. エラー
errorインタフェース
- エラーを表す型
- 最後の戻り値として返すことが多い
```go:
type error interface {
  Error() string
}
```

#### 7-1-4. エラー処理
**nilと比較してエラーが発生したかをチェックする**
```go:
if err := f(); err != nil {
  // エラー処理
}
```

#### 7-1-5. エラー処理のよくあるミス
err変数を使い回すことによるハンドル(エラー処理)ミス
```go:
/* 1. */
f, err := os.Open("file.txt")
if err != nil {
	// エラー処理
}

/* 2. */
// 本来は err = doSomething(f) としたつもり
doSomething(f)
if err != nil {
	// エラー処理
}
```
スコープの違いから起こるエラーだと思われる。
エラーが発生してもハンドルされずに次に進んでしまう。
[errcheck](https://github.com/kisielk/errcheck)などの静的解析ツールで回避できる。
errcheck-> goのプログラムでチェックされていないエラーをチェックするためのプログラム

#### 7-1-6. エラー処理で大事なこと
- 必要十分な正しい情報を伝えること
誰に？ユーザかな...
意図しない使われ方をした場合にエラーを吐いて「こう使ってしまったからエラーが出たんですよ」と、ユーザに伝わるようなエラー情報を吐かなくてはいけない、ということかなと。
1つのエラー情報だけで伝わらなければ情報を増やす。
基本的には無駄に情報を増やさない(ユーザが分かりにくくなるだけ)

- 受け取り手によって伝え方を変える
  - 同じパッケージの別の関数なのか？別のパッケージ？
    - エラーハンドリングによる分岐処理が必要
  - クライアント？
  - エンドユーザ？

#### 7-1-7. 文字列ベースで簡単なエラーの作り方
- errors.Newを使う
  - エラーが起こった場合に引数に指定した文字列が返る
`err := errors.New("Error")`

- fmt.ErrorFを使う
  - 書式を指定する。文字列が返るのは同じ
`err := fmt.Errorf("%s is not found", name)`

#### 7-1-8. エラー型の定義
Errorメソッドを実装している型を定義する
```go:
type PathError struct {
       Op   string
       Path string
       Err  error
}

func (e *PathError) Error() string {
       return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```
対応するエラーがどんなエラーか？どういった物を吐くのか？
対応させるエラー型の構造をエラーに必要なメソッドを実装する。

#### 7-1-9. TRY エラー処理をしてみよう
```go:
package main

import (
	"fmt"
	"os" // 標準エラー出力用
)

/* 文字列を返すStringerインタフェース */
type Stringer interface {
	String() string
}

/* (v interface{})->引数が空のインタフェース
戻り値が文字列とエラーを返す */
func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, MyError("CastError")
}

type MyError string

func (e MyError) Error() string {
	return string(e)
}

type S string

func (s S) String() string {
	return string(s)
}

func main() {
  /* 今回は文字列の場合は正常としている */
	v := S("hoge")
  /* ToStringerの引数が空のインタフェースなので、vに何の型を入れてもエラー処理が行える */
	if s, err := ToStringer(v); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		fmt.Println(s.String())
	}

}
```

#### 7-1-10. エラー処理をまとめる
bufio.Scannerの実装が参考になる
```go:
s := bufio.NewScanner(r)
for s.Scan() {
	fmt.Println(s.Text())
}
if err := s.Err(); err != nil {
	// エラー処理
}
```
参考: [Golang Error Handling lesson by Rob Pike](https://jxck.hatenablog.com/entry/golang-error-handling-lesson-by-rob-pike)
> Goでは多値を返して、その最後がエラーになるという形式が一般的であり、かつ型として定義されている。

ここから来る問題点の解消ができる。
- 途中でエラーが発生したらそれ以降の処理を飛ばす
- 全ての処理が終わったらまとめてエラーを処理
エラー処理が一箇所になるというのが利点で、エラーが発生した後の処理を実行する必要のない場合に利用する。

#### 7-1-11. TRY エラー処理をまとめる
https://docs.google.com/presentation/d/1HW3wG8J_Q2536Iu__7HGr_mhurHajC7IOGjCnn3kZmg/edit#slide=id.g7e7c89dcc9_0_355

回答コードを見ても何を処理しているか分からない。

#### 7-1-12. エラーをまとめる
[multierr](https://github.com/uber-go/multierr)を使う
- 特徴
  - 成功したものは成功させたい -> ok
  - 失敗したものだけエラーとして報告したい -> ok
  - N番目のエラーはどういうエラなのか知れる

```go:
var rerr error

/* 条件分岐の場合の例 */
if err := step1(); err != nil {
	rerr = multierr.Append(rerr, err)
}
if err := step2(); err != nil {
	rerr = multierr.Append(rerr, err)
}
return rerr
/* ------------------------------- */

/* 繰り返しの場合の例 */
for _, err := range multierr.Errors(rerr) {
       fmt.Println(err)
}
```

#### 7-1-13. エラーに文脈を持たせる
fmt.Errorf関数の％wを使う
```go:
/* 変数errのエラー出力"foo"を%wを使用してラップしている */
err := fmt.Errorf("bar: %w", errors.New("foo"))
/* "foo" -ラップ後-> "bar: foo" */
fmt.Println(err)                // bar: foo

/* erros.Unwrap関数で元のエラーが取得できる */
fmt.Println(errors.Unwrap(err)) // foo
```

#### 7-1-14. 振る舞いでエラーを処理する
エラー処理は具象型に依存させない

参考: [Golangのエラー処理とpkg/errors](https://deeeet.com/writing/2016/04/25/go-pkg-errors/)
> fmt.Errorf()はerrorを別のstringに結合して別のerrorを作り出す。
> 原因となったエラーが特定の型を持っていた場合にそれを隠蔽してしまう。

ここの所でなんとなく言いたい事がわかった。
```go:
// 一時的なエラーかどうかを判定する関数
/* err error -> error型であることが重要 */
func IsTemporary(err error) bool {
  te, ok := err.(interface {
      Temporary() bool
  })
  return ok && te.Temporary()
}
```
- エラーの種類で処理を分けたい場合がある
- インタフェースを使い、振る舞いで処理する

#### 7-1-15. 値によって分岐する
error.Is関数を使う
`error.IS(<変数に代入した値が出すエラー>, <第1引数のエラーを何と比較するか>)`
↑ややこしいかも...
第1引数のエラーが、第2引数の値かどうかを判定する
```go:
if errors.Is(err, os.ErrExist) {
	// os.ErrExistだった場合の処理
}
```
判定不能の場合は`errors.Unwrap`関数を読んでアンラップ後に判定

#### 7-1-16. エラーから情報を取り出す
errors.As関数を用いる
第1引数で指定したエラーを第2引数で指定した**ポインタが指す変数**に代入する
```go:
var pathError *os.PathError
if errors.As(err, &pathError) {
	fmt.Println("Failed at path:", pathError.Path)
} else {
	fmt.Println(err)
}
```
キャスト不可の場合は`errors.Unwrap`関数でアンラップ後に試みる

#### 7-1-17. エラーをerrorで扱う
- まずはエラーがあったかなかったで判断
  - `nil`かどうかで判定する
- 種類で分岐しても具象型に依存させない
  - error型を使う

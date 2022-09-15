# メルカリ作のプログラミング言語Go完全入門 読破
# 6. 抽象化
## 6-1. インタフェース
#### 6-1-1. インタフェースと抽象化
- 抽象化
  - 具体的な実装を隠し、振る舞いによって共通化させること
  - 複数の実装を同室のものとして扱う
[引用: インタフェースと抽象化 image](https://cdn-ak.f.st-hatena.com/images/fotolife/y/y-zumi/20190728/20190728023343.jpg)
**Goではインタフェースでしか抽象化をすることができない**

#### 6-1-2. インタフェース
```go:
package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	type Stringer1 interface {
		String() string
	}

	type Stringer2 interface {
		String() string
	}

	var s Stringer1 = Hex(100)
	var m Stringer2 = Hex(1000)
	fmt.Println(s.String())
	fmt.Println(m.String())

}

/* 実行結果 */
// 64
// 3e8
```
型Tが構造体でフィールド3つ4つ持っていたとする。
そこからインターフェースIを実装する際に、型Tのフィールドから実装するフィールドを選んでインターフェースIに実装できる。

引用: [Goのinterfaceがわからない人へ](https://qiita.com/rtok/items/46eadbf7b0b7a1b0eb08)
```go:
// 食べるためのインターフェース
type Eater interface{
    PutIn() // 口に入れる
    Chew() // 噛む
    Swallow() // 飲み込む
}

// 人間用のインターフェースの実装
func (h Human) PutIn(){
    fmt.Println("道具を使って丁寧に口に運ぶ")
}
func (h Human) Chew(){
    fmt.Println("歯でしっかり噛む")
}
func (h Human) Swallow(){
    fmt.Println("よく噛んだら飲み込む")
}
```
実装イメージとしては、ここらへんが凄い分かりやすい。

#### 6-1-3. interface{}
empty(空の) interface
```go:
package main

import "fmt"

func main() {
	/* 変数vを空のinterfaceで宣言する */
	var v interface{}

	/* int型を代入 */
	v = 100
	fmt.Println(v)
	/* string型を代入 */
	v = "hoge"
	fmt.Println(v)
}

/* 実行結果 */
// 100
// hoge
```
空のインタフェースにはどの型の値も実装できる！

#### 6-1-4. 関数にインタフェースを実装させる
参考:
[無名関数](https://blog.y-yuki.net/entry/2017/05/04/000000)
[Golang入門（関数）](https://qiita.com/gorilla0513/items/7e734c4e0680b5ea341d)
[【備忘】Go 言語のデータ型と型変換](https://qiita.com/t-yama-3/items/724f5f76356b814b0b2d)

```go:
package main

import (
	"fmt"
)

/* func([引数]) [戻り値の型] { [関数の本体] } */
// 関数型
type Func func() string

/* レシーバが関数型, String(関数名), 戻り値->関数を実行してstring型で返す */
func (f Func) String() string { return f() }

func main() {
	/* 変数sにfmtパッケージの文字列を返すメソッドを代入(helloを戻り値で返す) */
	var s fmt.Stringer = Func(func() string { return "hello" })
	fmt.Println(s) // 戻り値のhelloを出力する
}

/* 実行結果 */
// hello
```

```go:
/* fmt.Stringer -> インタフェース */
type Stringer interface {
    String() string
}
```
`fmt.Stringer = Func`
ここでメソッドをFunc型に型変換している。

#### 6-1-5. スライスとインタフェース
```go:
package main

func main() {
	ns := []int{1, 2, 3, 4}
	// できない
	var vs []interface{} = ns
}

/* 実行結果 */
/* cannot use ns (variable of type []int) as type []interface{} in variable declaration */
/* 変数宣言で ns (型[]int の変数) を型 []interface{} として使用できない */
```
スライスとインタフェースに互換性がない。
コピーするにはforで要素を取り出して格納するしかない。

#### 6-1-6. インタフェースの実装チェック
参考: [値がインターフェイスを実装しているかどうかの確認の説明](https://www.web-dev-qa-db-ja.com/ja/interface/%E5%80%A4%E3%81%8C%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%83%95%E3%82%A7%E3%82%A4%E3%82%B9%E3%82%92%E5%AE%9F%E8%A3%85%E3%81%97%E3%81%A6%E3%81%84%E3%82%8B%E3%81%8B%E3%81%A9%E3%81%86%E3%81%8B%E3%81%AE%E7%A2%BA%E8%AA%8D%E3%81%AE%E8%AA%AC%E6%98%8E/1049952806/)

- コンパイル時に実装しているかチェックする
```go:
package main

import (
	"fmt"
)

type Func func() string

func (f Func) String() string { return f() }

/* インタフェース型の変数に代入する */
var _ fmt.Stringer = Func(nil)

func main() {}

/* インタフェースが実装されていなかったらエラーが出力されるということだと思う。 */
```

#### 6-1-7. 型アサーション
インタフェースの値の基になる具体的な値を利用する手段を提供する
参考:
[型アサーション](https://blog.y-yuki.net/entry/2017/05/08/000000#:~:text=%22world%22%7D)%0A%7D-,%E5%9E%8B%E3%82%A2%E3%82%B5%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3,-%E3%81%A9%E3%82%93%E3%81%AA%E5%9E%8B%E3%81%AE)
[Go言語入門(Goの型アサーション)](https://qiita.com/knt45/items/2ae6f0dcbf84d0a24c42)

`<インタフェース>.(<型>)`
`value, ok := <変数>.(<型>)`
1番目の変数(value)には**型アサーション成功時に実際の値が格納される**。
2番目の変数(ok)には**型アサーションの成功の有無（true/false）が格納される**。
```go:
package main

import "fmt"

func main() {
	var v interface{}
	v = 100

	n, ok := v.(int)
	fmt.Println(n, ok) // n->vがint型だったのでチェックが成功したので100が格納されている。
	s, ok := v.(string)
	fmt.Println(s, ok) // s->vがstring型ではないのでチェックが成功しなかったため、値に格納されなかった。
}

/* 実行結果 */
// 100 true
// false
```

#### 6-1-8. 型スイッチ
型によって処理をスイッチする
代入文は省略可能
```go:
package main

import "fmt"

func main() {
	var i interface{}
	i = 100
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2) // i=100でint型なので、この処理だけ実行される。
	case string:
		fmt.Println(v + "hoge")
	default:
		fmt.Println("default")
	}
}

/* 実行結果 */
// 200
```

#### 6-1-9. インタフェースの設計
参考:
[Best Practices for Interfaces in Go](https://blog.boot.dev/golang/golang-interfaces/)
[[翻訳] Golang におけるinterface 実装のベストプラクティス](https://qiita.com/takuya_sss/items/ed162724bbce7b4b4bd0)

**interface設計におけるベストプラクティス**
- interfaceは小さく作る
最も重要なトピック
最小の必要な振る舞いの定義
HTTP Packageがinterface設計の良い例

- interfaceをclassのように扱わない
...クラスのように扱うのかと思っていました。
何が問題かというと、**Goでは型階層が作れない点**のようです。
5つinterfaceを作成したらその5つそれぞれが別の振る舞いをすることが求められる -> エラーハンドリングもそれぞれに必要、ということかと思われる。

#### 6-1-10. io.Readerとio.Writer
入出力の抽象化
- `io`のインタフェースを使わなかった場合
  - 入力がファイルからの場合、またはネットワークからの場合という風にパターンに沿って実装をいくつも用意しなければならない
- `io`のインタフェースを使った場合
  - さまざまなパターンをまとめて実装することができる。

- それぞれ1つのメソッドしか持たないので実装が楽
`io.Reader`
> **何かを読み込む機能を持つものをまとめて扱うために抽象化されたもの**

```go: io.Reader
type Reader interface {
  Read(p []byte) (n int, err error)
}
```

`io.Writer`
> **何かに書き込む機能を持つものをまとめて扱うために抽象化されたもの**

```go: io.Writer
type Writer interface {
       Write(p []byte) (n int, err error)
}
```

#### 6-1-11. TRY インタフェースを作ろう
レシーバが分からなくなったので再度整理。
関数の外から他の関数のメソッドを呼び出す際にレシーバを通してアクセスする。
レシーバを通してメソッドの操作(引数の値の指定)なども行える。

```go:
package main

import "fmt"

/* interface型をStringerと名前を付け、
メソッド名Stringがstring型であることを定義する */
type Stringer interface {
	String() string
}

/* interfaceをINT型かどうかをチェックするための関数 */
type I int
// レシーバにint型
func (i I) String() string {
  return "type int"
}

/* bool型 */
type B bool
// レシーバにbool型
func (b B) String() string {
	return "type bool"
}

/* string型 */
type S string
// レシーバにstring型
func (s S) String() string {
	return "type string"
}

/* 引数で渡した型とinterface(Stringer)を比較する=>s.(type) */
func F(s Stringer) {
	switch v := s.(type) {
	case I:
		fmt.Println(int(v), "I")
	case B:
		fmt.Println(bool(v), "B")
	case S:
		fmt.Println(string(v), "S")
	}
}

/* <インタフェース>.(<型>) */
func main() {
	var i I = I(100)
	F(i) // 変数iにはI型(int型)が代入されているので、caseIが出力される
  /* B(true->bool), S("hoge"->string)は型を指定している */
	F(B(true))
	F(S("hoge"))
}

/* 実行結果 */
/*
100 I
true B
hoge S
*/
```

## 6-2. 埋め込みとインタフェース
#### 6-2-1. 構造体の埋め込み
構造体の中に構造体を埋め込む
```go:
type Hoge struct {
	N int
}
// Fuga型にHoge型を埋め込む
type Fuga struct {
	Hoge // 名前のないフィールドになる
}
```
Hoge -> 名前のないフィールド = 匿名フィールド

#### 6-2-2. 埋め込みとフィールド
埋め込んだ値に委譲(継承ではない)
```go:
package main

import "fmt"

type Hoge struct {
	N int
}
type Fuga struct {
	Hoge
}

func main() {
	f := Fuga{Hoge{N: 100}}
	// Hoge型のフィールドにアクセスできる
	fmt.Println(f.N)
	// 型名を指定してアクセスできる
	fmt.Println(f.Hoge.N)
}

/* 実行結果 */
// 100
// 100
```
1. `f := Fuga{Hoge{N: 100}}`を紐解く。
変数`f`に代入されるもの -> `<アクセスしたい構造体の埋め込み先名>:{<アクセスしたい構造体名>}<アクセスしたいフィールド名>: <値の指定>`

2. `f.N`
`f`にはHoge型の構造体(埋め込まれた構造体)が格納している。
`f.N`で、HogeのフィールドNへアクセスできる。

3. `f.Hoge.N`
やっている事は2と変わらない。
明示的にHoge型を指定してNフィールドへアクセスしている。

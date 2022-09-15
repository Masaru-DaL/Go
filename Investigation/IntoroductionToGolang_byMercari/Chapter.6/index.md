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


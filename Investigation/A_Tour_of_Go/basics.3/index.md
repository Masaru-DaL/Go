# A Tour of Go Basics.3
## 3-1. Pointers
Goではポインタを扱える。
ポインタ？
参考: [Goで学ぶポインタとアドレス](https://qiita.com/Sekky0905/items/447efa04a95e3fec217f)
記事が古いのが気になりますが、いいねが多いのでこちらで理解を深めます。

- ポインタを理解するには前提知識が必要
1. プログラムのコンパイルから実行までの流れ
Goはコンパイラでコンパイルが必要である。

2. 変数とメモリとアドレス
- アドレスとは
  - 変数に値を代入し、コンパイルした時には、「メモリ上のある場所」に変数の値が格納される。
  - この「メモリ上のある場所」が「アドレス」というもの
  - アドレス番地は16進数で表される。(例: `0x1040a0d0`)

```go:
package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	// ポインタ型の変数を宣言する
	// pがポインタ変数
	// *Personポインタ型
	var p *Person

	p = &Person{
		Name: "太郎",
		Age: 20,
	}
	fmt.Printf("変数pに格納されているアドレス :%p", p)
}
```
変数pに格納されているアドレス :0xc000118018

- `&`オペレータ
  - 変数のアドレス番地を引き出す
  - `i := 42` -> `p = $i` (&iでiのアドレス番地を引き出す)

3. ポインタとポインタ変数
上記のコードに以下を追加して出力する
`fmt.Printf("p :%+v\n", p)`
-> p :&{Name:太郎 Age:20}

- `%+v`
`%+v`
構造体を出力する際に、%vの内容にフィールド名も加わる
`%v`
デフォルトフォーマットで対象データの情報を埋め込む
`\n`
言わずもがな、改行

ようするに、`p`を出力することで`&{Name:太郎 Age:20}`が出力されることが重要。
**変数`p`がポインタ変数となり、実際に`p`にはアドレスが格納されている。**

まだ続いているが、こんがらがりそうなので一旦終了。

## 3-2. Structs
struct(構造体)
なんとなくやってる事は分かるけど、理解しきれない。
クラスのような、関数のような、何が違うのか。

`type Person`で定義して、人は名前や性別、年齢、身長など様々な要素(フィールド)がある。
structはフィールドの集まり。という説明はなんとなく理解できる。

```go: struct
package main

import "fmt"

type Person struct {
	Name string
	Sex string
	Age int
	Height int
	Weight int
}

func main() {
	fmt.Println(Person{"Musashi", "男", 27, 182, 74})
}
```
{Musashi 男 27 182 74}
引数の順番は、フィールドを定義した順番と同じ

## 3-3. Struct Fields
structのフィールドには、ドット`.`を用いてアクセスする。

```go: struct(.)
package main

import "fmt"

type Person struct {
	Name string
	Sex string
	Age int
	Height int
	Weight int
}

func main() {
	Musashi := Person{"Musashi", "男", 27, 182, 74}
	fmt.Println(Musashi.Name)
}
```
Musashi

## 3-4. Pointers to structs
structのフィールドは、structのポインタを通してアクセスすることもできる。

```go: Pointers to structs
package main

import "fmt"

type Person struct {
	Name string
	Sex string
	Age int
	Height int
	Weight int
}

func main() {
	Musashi := Person{"Musashi", "男", 27, 182, 74}
	p := &Musashi
	p.Age = 1e9
	fmt.Println(Musashi)
}
```
{Musashi 男 1000000000 182 74}

- 1e9
  - よく小数点を省略する時に使う記法と同じで、1の後に0を9個付けるという意味

文字列へのアクセスはまた別途理解が必要。

## 3-5. Struct Literals
1. 構造体を定義する
2. 構造体を使って様々に派生できる
ということが言えるのだと思われる。

調べた結果:
Goはオブジェクト指向言語の様なクラスを持たない。
なのでクラスの継承も存在しない。
構造体でそのような事ができる。ということ。

## 3-6. Arrays
`var a [10]int`
これは変数`a`がint型で10個の配列を宣言しているということ。

```go: Arrays
package main

import "fmt"

func main() {
	// 1
	var a [2]string

	// 2
	a[0] = "Hello"
	a[1] = "World"

	// 3
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 4
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 5
	fmt.Println(primes)
}
```
1. `var a [2]string`で文字列で2個の配列を宣言
2. `a`のindex0, 1に文字列を入れている。
3. `a[0], a[1]`だとそれぞれの要素を出力し、`a`だと配列として出力している
4. `primes`に、配列の宣言と配列への要素の格納を同時に行なっている。
5. `primes`で出力しているので配列で出力されている。

- 出力結果
Hello World
[Hello World]
[2 3 5 7 11 13]

## 3-7. Slices
スライス(の概念)はpythonと同じ感じ

## 3-8. Slices are like references to arrays
```go: slice
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]

スライスで取得した部分の要素の変更が可能。
`b := names[1:3]` -> `b[0] = "XXX"`

## 3-9. Slice literals
配列の流れとしては、
1. 配列の宣言
   1. 何個入る配列なのか
   2. 型を宣言
2. 配列に要素を格納

これを同時に行える時点で、何個入る配列なのかという部分は省略可能。

3-8のコードを、
`[4]string` -> `[]string`にしても良いということ。

実務においては何個入る配列かを明示的にしたほうがいいシーンもあるかもしれない。

## 3-10. Slice defaults
pythonと同様

## 3-11. Slice length and capacity
length -> 実際の要素の個数
capacity -> 基底配列の長さ, 最大収納量
```go: Slice length and capacity
package main

import "fmt"

func main() {
	// 1
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 2
	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// 3
	// Extend its length.
	s = s[:4]
	printSlice(s)

	// 4
	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```
- 前提知識
配列の容量が不足した場合、新しい配列を大きなサイズでメモリ確保して元データ参照範囲をコピーする。

1. 配列の宣言と要素の格納(最初に6個格納しているのでこれがcapとなる)

2. 1の基底配列から取得しているので、実際の個数は0、capは6

3. ここが一番のキモ！
2の配列から取りたい、となった時に要素が何も入っていないため、1の配列から要素を取ってくるという流れになります。
なので、1の配列のindex0, 1, 2, 3を取ってくるため、len4, cap6となる。
参考: [Go言語「A Tour of Go」のsliceでわからないことがあります](https://teratail.com/questions/155045)

4. 3から取得します。
3の配列は[2 3 5 7]
len4、cap4です。
`[2:]`を指定すると5, 7が取得され、len2, cap4となります。

## 3-12. Nil slices
スライスのゼロ値は`nil`である。

`var s []int` (空のリスト)
`if s == nil`が成立する。

## 3-13. Creating a slice with make
`make`(組み込み関数)でスライスを作成することができる。
`make(型, len, cap)`
第2引数に指定するとlenになり、第3引数は省略可能

```go: Creating a slice with make
package main

import "fmt"

func main() {
	// 1.
	a := make([]int, 5)
	printSlice("a", a)

	// 2.
	b := make([]int, 0, 5)
	printSlice("b", b)

	// 3.
	c := b[:2]
	printSlice("c", c)

	// 4.
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```
1. 空のリストの宣言
`make([], 5)` -> len=5
出力結果: len=5 cap=5 [0 0 0 0 0]

2. 空のリストの宣言
`make([]int, 0, 5)` -> len=0, cap=5
出力結果: len=0 cap=5 []

3. lenが0の配列からスライスする
その場合、cap=5に合わせてゼロ値がある配列からスライスをすることになる。
`[0, 0, 0, 0, 0]` これに対して`[:2]`
-> index0, 1をスライス
`[0, 0]`
len=2, cap=5(2の基底配列を引き継ぐ)

4. 3の`len=2 cap=2 [0, 0]`から`[2:5]`のスライスをする
なぜcap=3になるのかが不明。

## 3-14. Slices of slices
多次元スライスの中のスライスの型をそれぞれ違う型に出来るか要確認。
- 可能！

```go: slice
board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
```

## 3-15. Appending to a slice
スライスへ新しい要素を追加するには`append`を使用する。
`append(追加先スライス, 追加する要素)`

```go: append
package main

import "fmt"

func main() {
	// 1.
	var s []int
	printSlice(s)

	// 2.
	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// 3.
	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// 4.
	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```
1. 空のスライスを作成 -> `s`
`len=0 cap=0 []`

2. `s`に0値を追加
`len=1 cap=1 [0]`

3. 2の`s`に1値を追加
`len=2 cap=2 [0, 1]`

4. 3の`s`に複数の値を追加
`len=5 cap=6 [0, 1, 2, 3, 4]`
なぜcap=6となるのか？

メモリを確保する時に「このぐらいがいいだろう」ということで少し大きめに容量を確保している。

## 3-16. Range
- 前提知識
`fmt.Printf("2**%d = %d\n", i, v`
`fmt.Printf`は、フォーマットを指定して引数を出力できる。
`%d`は数値
`%d`が2つあり、引数に`i`, `v`があるので、`2**%d`に`i`が入り、`= %d_n`に`v`が入る。

- range
rangeはforループに利用する。
スライスやマップをひとつずつ反復処理するために使用する。
スライスをrangeで繰り返す場合、rangeは反復毎に2つの変数を返す(1. index番号, 2. indexの場所の要素のコピー)

`for index番号, indexの場所の要素のコピー := range スライス {}`

## 3-17. Range continued
rangeを使用した場合2つの変数が返されるが、アンダーバー`_`へ代入することで捨てることができる。

1. index番号を捨てる場合
`for _, value := range slice`

2. 要素を捨てる場合
`for i, _ := range slice`

3. indexだけ必要な場合、2つの値を省略できる
`for i := range slice`

## 3-18. Exercise: Slices
- 条件の洗い出し
Goal -> pic関数の実装, 画像の表示
  - 生成する画像は好きにして良い
	- 長さ dy のsliceに、各要素が8bitのunsigned int型で長さ dx のsliceを割り当てたものを返すように実装する必要がある

1. 最初のコードの意味を理解する
```go:
package main

// パッケージのインポート(pic.Show関数が定義)
import "golang.org/x/tour/pic"

// Pic関数の定義
func Pic(dx, dy int) [][]uint8 {
}

// pic.Show関数の呼び出し
func main() {
	pic.Show(Pic)
}
```

2. pic.Show関数の振る舞いの理解
`golang.org/x/tour/pic`をググってみる。
[Go pic](https://pkg.go.dev/golang.org/x/tour/pic)

`func Sh​​ow(f func(dx, dy int ) [][] uint8 )`
- Showは、実行されると**関数f**(Show(Pic)のpicのことを指す)によって定義された画像を表示する。
  - `func Pic(dx, dy int) [][]uint8 {}`で定義された画像を表示するという意味

- Pic関数の引数`dx`, `dy`は何を受け取るものか？
Show関数の中に答えがあるようです。
[pic.go](https://cs.opensource.google/go/x/tour/+/refs/tags/v0.1.0:pic/pic.go)

`const`で`dx = 256`, `dy = 256`と指定してます。
0値 -> 完全な青, 255 -> 完全な白

...その他は分かったような分からないような。
時間取られそうなので一旦保留。

## 3-19. Maps
`make`関数を使用して、mapを生成することができる。
`make(map[キーの型]値の型)`

```go: map
package main

import "fmt"

func main() {
    m := make(map[string]int)

    fmt.Println(m)

    m["key1"] = 10
    m["key2"] = 20

    fmt.Println(m)
}
```
map[]
map[key1:10 key2:20]

- 上記を踏まえた上で
```go: map(Tour)
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
  // 1.
	m = make(map[string]Vertex)
	fmt.Println(m)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

  // 2.
	fmt.Println(m)

  // 3.
	fmt.Println(m["Bell Labs"])
}
```
map[]
map[Bell Labs:{40.68433 -74.39967}]
{40.68433 -74.39967}

1. make関数で作成すると初期化して使用可能な状態になる。

2. key: `Bell Labs`, value: `{40.68433 -74.39967}`だと分かる。

3. keyを指定すると、対応したvalueが出力される。

## 3-20. Map literals
mapはkey, valueをセットなので、valueのみは不可

## 3-21. Map literals continued
mapに渡すトップレベルの型が単純な型名の場合は、リテラルの要素から推定できるので、その型名を省略することができる。
※何が省略できて、できないか。頭の片隅に置いておく

## 3-22. Mutating Maps
挿入, 更新 -> `m[key] = elem`
取得 -> `elem = m[key]`
削除 -> `delete(m, key)`
keyに対する要素が存在するかの確認 -> `elem, ok = m[key]`
(もし`m`にkeyがあれば変数`ok`は`true`, 存在しなければ`false`となる)

```go: Mutating Maps
package main

import "fmt"

func main() {
	// 1.
	m := make(map[string]int)
	fmt.Println(m)

	// 2.
	m["Answer"] = 42
	fmt.Println(m)
	fmt.Println("The value:", m["Answer"])

	// 3.
	m["Answer"] = 48
	fmt.Println(m)
	fmt.Println("The value:", m["Answer"])

	// 4.
	delete(m, "Answer")
	fmt.Println(m)
	fmt.Println("The value:", m["Answer"])

	// 5.
	v, ok := m["Answer"]
	fmt.Println(m)
	fmt.Println("The value:", v, "Present?", ok)
}
```
1. mapの作成
`map[]`

2. key -> Answer, value -> 42 で挿入
`map[Answer:42]`
`The value: 42`

3. Key(Answer)のvalueを、48に更新
`map[Answer:48]`
`The value: 48`

4. key(Answer)を削除
`map[]`
`The value: 0`

5. key(Answer)に対する要素があるか確認する
`map[]`
`The value: 0 Present? false` (key(Answer)に対する要素が存在しないのでfalseと出力される)

## 3-23. Exercise Maps

## 3-24. Function values
**関数も変数である**
**他の変数と同じように、関数を渡すことができる**
- 前提知識
`math.pow` -> べき乗を計算できる
`math.sqrt` -> 平方根を計算できる

```go: Function values
import (
	"fmt"
	"math"
)

// 1.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {

	// 2.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// 3.
	fmt.Println(hypot(5, 12))

	// 4.
	fmt.Println(compute(hypot))

	// 5.
	fmt.Println(compute(math.Pow))
}
```
1. compute関数を定義
引数: `(float64, float64)`を引数に、`float64`を返す関数
戻り値: `float64`
らしいのですが、良く意味がわかりません。
そもそもの書き方が分かってないかもしれん...
`func compute(fn func(float64, float64) float64) float64{return fn(3, 4)}`

`fn func(x float64, y float64)`
このように引き数の名前が省略されているという説明でしっくりくる。

1. `hypot`関数を定義
引数2つを`(x*x + y*y)`で計算を行った結果の平方根を求める。

3. (5*5 + 12*12)の平方根を求める。

4. `fmt.Println(math.Sqrt(25))`と同じ事をしている。
25というのは、(3*3 + 4*4)のこと。

5. `fmt.Println(math.Pow(3, 4))` と同様の事をしている。

## 3-25. Function closures
- クロージャとは
1. 変数にはスコープ(参照できる範囲)がある。
2. クロージャを使うと、その変数のスコープをその範囲の中に制限することができる。

```go: Function closures
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```
`sum := 0` は最初だけ？
関数を引数にするのと同じような感じで理解できてない気がする。

## 3-26. Exercise: Fibonacci closure

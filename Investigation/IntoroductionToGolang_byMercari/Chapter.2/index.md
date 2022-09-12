# メルカリ作のプログラミング言語Go完全入門 読破
# 2. 基本構文
## 2-1. 変数
#### 2-1-1. 変数
- 値を保持する領域
- 一時的に保存な保存処理

#### 2-1-2. 変数と型
- 型とは
  - どういう種類の値かを表すもの
    - 整数、文字列など
    - 自分で作ることも可能(ユーザ定義型)

- 動的型付け言語
  - 変数に型がない -> なんでも代入できる
  - **プログラム実行時**に型を検証する

- 静的型付け言語
  - **Goは静的型付け言語**
  - 変数に型があり、型が違うと代入できない
  - **コンパイル時**に型を検証する

#### 2-1-3. 静的型付けの利点
1. 実行前に型の不一致を検出できる
2. 曖昧なものはエラーになる
   - 暗黙の型変換がない
     - 例: `1 + "2" -> "12"`
3. 型推論がある
   - 明示的に型を書く必要がない場合が多い

#### 2-1-4. 変数宣言
予約語`var`(variableの略)を用いる
予約語: 識別子(名前)に使用できない語(if, forなども同じ)

```go:
/* 変数宣言と代入を同時に行う */
var n int = 100

/* 変数宣言と代入を別で行う */
var n int
n = 100

/* 型を省略(int型になる->型推論) */
var n = 100

/* varを使わず変数宣言を行う
   関数内でのみしか出来ない！*/
n := 100

/* まとめて変数宣言を行う */
var (
  n = 100
  m = 200
)
```

#### 2-1-5. 使っていない変数
変数を宣言した場合、その変数が使われていないとコンパイル時にエラーが出る。
```go:
package main

func main() {
    x := 100
    println(x)

    y := 200   // 使ってないのでコンパイルエラー
    println(x) // println(y)の間違い
}
```

#### 2-1-6. 組み込み型
https://docs.google.com/presentation/d/1CIMDenDLZ7NPNgzmfbCNH_W3dYjaTEBdUYfUuXXuMHk/edit#slide=id.g4e29972148_0_32

#### 2-1-7. 変数のゼロ値
- Goの変数は**明示的な初期化をしなくても使える**
  - **ゼロ値という値が設定され、型によって違う**

| 型                | ゼロ値   |
| ---------------- | ----- |
| intやfloat64などの数値 | 0     |
| string           | ""    |
| bool             | false |
| errorなどのインタフェース  | nil      |

#### 2-1-8. 変数の利用
```go:
package main

import "fmt"

func main() {
	var first_word string = "Hello, World"
	fmt.Println(first_word)
}

/* 実行結果 */
// Hello, World
```

## 2-2. 定数
#### 2-2-1. 定数
`const`の説明とは別？
数値や文字列のリテラルそのものを言ってるのかな。

#### 2-2-2. 2進数, 8進数, 16進数表記
ちょっと分からなかったので別で調べる
```go:
package main

import "fmt"

func main() {
    s := ""
    s = fmt.Sprintf("%b", 255)
    fmt.Println(s) // => "11111111"
    s = fmt.Sprintf("%o", 255)
    fmt.Println(s) // => "11111111"
    s = fmt.Sprintf("%x", 255)
    fmt.Println(s) // => "ff"
}
```
こっちの方が分かりやすい感じがあった。

#### 2-2-3. 数字の区切り
- 数値リテラルに`_`を入れても無視される
- 数値リテラルの可読性を上げるために導入された
  - `,`みたいなものか

`5_000_000_000_000_000`
`0b_1010_0110`
`3.1415_9265`
出力される時は`_`が取り除かれた状態で出力。
例: 5000000000000000

#### 2-2-4. 定数式
定数のみからなる演算式
コンパイル時に計算される

- シフト演算
  - `x << n`
  - xのn bit左シフト
1 << 2 -> 100(2進数) -> 4(10進数)

- 関係演算 / 論理演算
  - `!` not演算子
  - !(10 == 20) -> 10と20は等しくない事を否定 -> true

#### 2-2-5. 名前付き定数
- 定数に名前を付けて宣言
  - `const`
  - 変数宣言と同じように定数にも名前が付けられる

```go:
/* 型のある定数 */
const n int = 100

/* 型のない定数 */
const m = 100

/* 定数式の利用 */
const s = "Hello, " + "World"

/* 複数の定数宣言をまとめる */
const (
  x = 100
  y = 200
)

/* 型のない定数の型を確認 */
fmt.Printf("%T", m)
// int
```
型のない定数は型推論されると思われる。
(事項で説明されてた。デフォルトの型として推論される。)
大体変数と同じ。

#### 2-2-6. 定数の型
変数と同じ

#### 2-2-7. 定数のデフォルトの型
```go:
package main

func main() {
	// 定数は型を持たないので無限の精度になる
	const n = 10000000000000000000 / 10000000000000000000
	var m = n  // mはint型
	println(m) // 1
}
```
`const n`が1でint型なので、それを代入した`m`もint型になる。
※型変換できない場合はコンパイルエラーになる。

#### 2-2-8. なぜ型なしの定数が必要か？
- 定数をただの数字のように扱いたい
ちょっと分かりづらかったので別で調べます。
https://qiita.com/hkurokawa/items/a4d402d3182dff387674
```go:
const hello = "Hello, ワールド"
var s string = hello    // OK
type MyString string
var ms MyString = hello // OK
ms = s                  // NG
```
独自定義した`MyString`に定数が代入出来ていると書かれています。もし定数`hello`を"string"として定義していたら代入できていないということです。

- 型無しの定数がないと型変換ばかりのコードになる

- Go言語が提供する単純さの1つが定数
https://gihyo.jp/news/report/01/GoCon2014Autumn/0001#:~:text=%E3%81%84%E3%82%8B%E3%81%9D%E3%81%86%E3%81%A7%E3%81%99%E3%80%82-,%E5%AE%9A%E6%95%B0,-%E5%90%8C%E6%B0%8F%E3%81%AF%E3%80%81

#### 2-2-9. 定数の利用
```go:
package main

func main() {
	const first_word = "Hello, World"
	println(first_word)
}
```

#### 2-2-10. 右辺の省略
```go:
package main

import "fmt"

func main() {
	const (
		a = 1 + 2
    /* 右辺を省略 */
		b
		c
	)
	fmt.Println(a, b, c)
}

/* 実行結果 */
// 3 3 3
/* 全部3になる */
```
これは凄いと思った。
- グループ化で宣言する際に有効。
- 省略された定数定義の右辺は、**最後の省略されていない定数と同じ**になる

#### 2-2-11. iota
```go:
package main

import "fmt"

const (
	zero        = iota
	one         = iota
	two         = iota
	three, four = iota, iota
)

const five = iota

func main() {
	fmt.Printf("zero:%v\n", zero)
	fmt.Printf("one:%v\n", one)
	fmt.Printf("two:%v\n", two)
	fmt.Printf("three:%v\n", three)
	fmt.Printf("four:%v\n", four)
	fmt.Printf("five:%v\n", five)
}

/* 実行結果 */
/*
zero:0
one:1
two:2
three:3
four:3
five:0
 */
```
試したみた感じ、まとめて宣言した場合のindex番号を表している
別で定義するとindex番号はまた0から。

#### 2-2-12. iotaを利用した連番
- 値に意味がない、それぞれが違う値であれば良いなどの時に利用する
  - 逆に値に意味があったり、DBなどに保存する場合はiotaは使わない
- 2つ目以降の右辺を省略すると1つ目の右辺と同じになる

```go:
const (
    StatusOK = iota // 0
    StatusNG        // 1（StatusNG = iotaと同じ意味）
)
```

## 2-3. 演算子
同じオペレータが出てくるから混同しないようにする。

#### 2-3-1. 代入演算
`:=`: **演算の初期化**と代入
`++`, `--`は、**式ではなく文**

#### 2-3-2. ビット演算
| -> 論理和
& -> 論理積
^ -> 否定
^ -> 排他的論理和
&^ -> 論理積の否定
<<, >> -> 左, 右に算出シフト

#### 2-3-3. 論理演算
|| -> または: `a || b`
&& -> かつ: `a && b`
! -> 否定: `!a`

#### 2-3-4. アドレス演算
& -> ポインタを取得: `&a`
* -> ポインタが指す値を取得: `*ptr`

#### 2-3-5. チャネル演算
`<-` -> チャネルへの送受信: `ch<-100, <-ch`

#### 2-3-6. TRY 演算子の利用
```go:
package main

import "fmt"

func main() {
	n := 100 + 200
	m := n + 200
	msg := "hoge" + "huga"

	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(msg)
}

/* 実行結果 */
// 300
// 500
// hogehuga
```

## 2-4. 制御構文
#### 2-4-1. 条件分岐: if
```go:
if x == 1 {
  println("xは1")
} else if x == 2 {
  println("xは2")
} else {
  println("xは1でも2でもない")
}
```

```go:
/* 代入文を書く */
if a := f(); a > 0 {
	fmt.Println(a)
} else {
	fmt.Println(2*a)
}

/* f関数を代入したaはifとelseのブロックで使えるようになる */
/* コンパイル時に後ろに改行があると、自動でセミロンを付与する */
```

#### 2-4-2. 条件分岐: switch
`break`がいらない(書けるけど基本書かない)
if文で大量に書くよりもswitch文の方が見通しがよくなる。
```go:
/* 通常の書き方 */
switch a {
case 1, 2: // 1または2の時
	fmt.Println("a is 1 or 2")
default:
	fmt.Println("default")
}

/* caseに式が使える */
switch {
  case a == 1:
    fmt.Println("a is 1")
}
```

- fallthroughを使った例
https://tech.yyh-gl.dev/blog/go-switch-fallthrough/
```go:
package main

import "fmt"

func main() {
	num := 1
	switch num {
	case 1:
		fmt.Print("I ")
		fallthrough
	case 2:
		fmt.Print("am ")
		fallthrough
	case 3:
		fmt.Println("yyh-gl.")
		// fallthrough // 次の節がなければコンパイルエラー
	}
}

/* 実行結果 */
// I am yyh-gl.
```
`fallthrough` -> 複数のcaseをまたぐ(というよりも次のケースに渡すイメージ)

#### 2-4-3. 繰り返し: for
**whileは無い！**, forのみ！
```go:
/* 初期値; いつまで繰り返すのか; 1回終わった後の処理 */
for i := 0; i <= 100; i++ {
}

/* 継続条件のみでも可能 */
for i <= 100 {
}

/* 無限ループ */
for {
}

/* rangeを使った繰り返し */
/* index番号と配列からの要素がi, vのそれぞれに入る */
for i, v := range []int{1, 2, 3} {
}
```

#### 2-4-4. breakによるループの抜け出し
for文中にbreakを用いるとループから抜け出せる。
```go:
/* ラベル指定のbreak */
func main() {
	var i int
LOOP: // for文にラベルをつける
	for {
		switch {
		case i%2 == 1:
			break LOOP // breakするループを明示的に指定
		}
		i++
	}
}
```

#### 2-4-5. TRY 奇数と偶数
```go:
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			fmt.Println(i, "-奇数")
		} else if i%2 == 0 {
			fmt.Println(i, "-偶数")
		}
	}
}

/* 実行結果 */
/*
1 -奇数
2 -偶数
3 -奇数
4 -偶数
...
*/
```

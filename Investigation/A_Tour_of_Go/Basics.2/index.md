# A Tour of Go Basics
## 2-1. For
- 構文
  - セミコロン`;`で3つに分かれる
  - 初期値; 条件式; 後処理

例: `for i := 0; i < 10; i++ { sum += i }`

## 2-2. For continued
**初期化と後処理ステートメントの記述は任意(書かなくても良い)！**
ただし、直前などに式に使用する変数の宣言は必要だと思われる。

## 2-3. For is Go's "while"
...!!
**セミコロン`;`を省略することもできる!!**
Goはforだけを使用する！！

実際の書き方は後ほど調べる必要有りか

## 2-4. Forever
ループ条件を省略すると、無限ループになる。
`for { }`
これだけで無限ループを表現できる。

## 2-5. If
forと同様に、`if 条件式 { }`と書く
セミコロン`;`も省略可能。

## 2-6. If with a short statement
forと同様に、条件式の前に簡単なステートメントを書く事ができる。
また、ここで宣言された変数は、`if`のスコープ内だけで有効。

## 2-7. If and else
`if`で宣言された変数は、`else`ブロック内でも使用が可能。

## 2-8. Exercise: Loops and Functions
書き方諸々理解不能だったのでサイトを参考にしながら理解する。
[A Tour of Go の Exercise: Loops and Functions 解答例](https://webgroove.work/a-tour-of-go-exercise-loops-and-functions-sample-answer/)

- 条件の洗い出し
1. 何が入力されても`z`の初期値を浮動小数点の`1`とする。
2. 計算を10回繰り返す。また毎回`z`を出力する。
3. `z`が実際の平方根と同じになった時点でループを終わらせる。終わったタイミングが10回より多いか少ないかを確認する。
4. `x`の初期値に他の値を入れて確認する。
5. `z`の結果を`math.Sprt`と比較する。

- 1, 2をクリアする
```go:
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	// 1. 初期値の設定
	z := 1.0

	// 2. 10回繰り返し, zの出力
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(i, "回目の計算値は", z, "です。")
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
```
1 回目の計算値は 1.5 です。
2 回目の計算値は 1.4166666666666667 です。
3 回目の計算値は 1.4142156862745099 です。
4 回目の計算値は 1.4142135623746899 です。
5 回目の計算値は 1.4142135623730951 です。
6 回目の計算値は 1.414213562373095 です。
7 回目の計算値は 1.4142135623730951 です。
8 回目の計算値は 1.414213562373095 です。
9 回目の計算値は 1.4142135623730951 です。
10 回目の計算値は 1.414213562373095 です。
1.414213562373095

***

- 3をクリアする。
```go:
package main

import (
	"fmt"

  // mathのインポート
	"math"
)

func Sqrt(x float64) float64 {
	// 1. 初期値の設定
	z := 1.0

	// 2. 10回繰り返し, zの出力
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(i, "回目の計算値は", z, "です。")

    // zの値が標準ライブラリの平方根と同じになったら中断する
		if z == math.Sqrt(x) {
			break;
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
```

***

- 4, 5をクリアする
4は`z := 1.0`の値を変更して確認する。
-> `x`, `x/2`に変更しても問題無し。

5はmainの出力で`z`, `math.Sqrt`と同時に出力して比較する。

```go:
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// 1. 初期値の設定
	z := 1.0

	// 2. 10回繰り返し, zの出力
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(i, "回目の計算値は", z, "です。")
		if z == math.Sqrt(x) {
			break;
		}
	}
	return z
}

func main() {
	x := 2.

	fmt.Println("zの結果は: ", Sqrt(x))
	fmt.Println("math関数の結果は: ", math.Sqrt(x))
}
```

## 2-9. Switch
`switch`ステートメントは、`if-else`ステートメントのシーケンスを短く書ける！
`case`は定数である必要はなく、整数だけでなく文字列でも指定することができる。

## 2-10. Switch evaluation order
Switchにおける`case`は、上から下へ`case`を評価する。
`case`の条件が一致すれば、そこで停止(自動的にbreak)する。

- 書き方例
```go: switch
n := 1

switch n {
  case 3:
    fmt.Println("n is", n)
  case 2:
    fmt.Println("n is", n)
  case 1:
    fmt.Println("n is", n)
}
```
この場合、`case 1`が該当し、`n is 1`が出力される。

## 2-11. Switch with no condition
条件のないswitch -> `switch true`と書くことと同じ
いまいち良く分かりませんが、`default`と書く事を指していると思う。要質問。

`default` == `else`

## 2-12. Defer
`defer`ステートメントは、`defer`へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させる。

```go: defer
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```
hello
world

## 2-13. Stacking defers
`defer`へ渡した関数が複数ある場合、その呼び出しはスタックされる。

```go: Stacking defers
package main

import "fmt"

func main() {
  // ①
  fmt.Println("counting")

  // ③
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }

  // ②
  fmt.Println("done")
}
```
1. counting
2. done
3. 9 8 7 ... 0

`i`は0~9とカウントが上がっていくが、`defer`で渡した後は9~0と出力されている。
スタックと同じように`LIFO(Last In First Out)`で実行されることがわかる。

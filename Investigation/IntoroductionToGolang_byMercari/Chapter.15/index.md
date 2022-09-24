# メルカリ作のプログラミング言語Go完全入門 読破
# 15. ジェネリクス
## 15-1. ジェネリクスの必要性
- ジェネリクスとは？
1つの関数で複数の型に対応する。

参考: [Go 1.18で実装されたジェネリクスを使ってみる – もふもふのブログ](https://mome-n.com/posts/golang-generics/)

- Go1.17までは1つの関数に付き、1つの型にしか対応できなかった。

```go:
//int64型しか対応できない…
func SumInt(a, b int64) int64 {
	return a + b
}
//float64型しか対応できない…
func SumFloat(a, b float64) float64 {
	return a + b
}
```

- Go1.18でジェネリクスを使用すると複数の型を指定できる。

```go:
//int64型またはfloat64型が指定できる！
func SumIntOrFloat[T int64 | float64](a, b T) T {
	return a + b
}
```

1. `()`が、`[]`に代わっている
2. 変数に型定義する形から、`T`にどんな型を定義するか決められるという感じに。パイプで複数指定可能か。
3. Tに型定義を行なってから変数、戻り値をTに指定している。(=**Tという型を指定しているということか？**)

- ジェネリクスの必要性
  - 型に**依存しない**汎用的な関数などを作るのが難しい
    - 上でやったように、各々の型にあった関数定義が必要になってしまう点

- Goにおけるジェネリクス
  - Goにはジェネリクスがなかったが、要望の声が多かった
  - Goが出た直後からジェネリクスの議論が出ている

## 15-2. 型パラメタを持つ関数の定義
```go:
package main

import (
	"fmt"
)

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func main() {
	Print[string]([]string{"Hello, ", "playground\n"})
	Print[int]([]int{10, 20})
}
```
関数`Print`は"T"を`any`で定義している。
`any`には、**任意の型**を指定できる。
呼び出し時には`Print[<呼び出す型>]`で任意の型を指定している。
"T"は、**型パラメタ**と呼ばれ、実際は型引数(<呼び出す型>)で指定した型として扱われる。

- 型パラメタを持つ関数のインスタンス化
  - 型引数を指定すると、インスタンス化することができる
    - 変数への代入
      - `var printStr func([]string) = Print[string]`
    - インスタンス化しないと呼び出せない(=型パラメタのままでは呼び出せない)
      - > "T"は、**型パラメタ**と呼ばれ、実際は型引数(<呼び出す型>)で指定した型として扱われる。

- 型推論
  - 型推論により、型引数を省略する事が可能
```go:
package main

import (
	"fmt"
)

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func main() {
	Print([]string{"Hello, ", "playground\n"})
	Print([]int{1, 2, 3})
}
```

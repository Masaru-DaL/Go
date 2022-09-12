# メルカリ作のプログラミング言語Go完全入門 読破
# 3. 関数と型
## 3-1. 型
#### 3-1-1. 組み込み型
- 最初から使える型
  - byte -> unit8, rune -> int32, any -> interface{}の型エイリアス
  - comparableは型制約のみに使えるインターフェース
    - ジェネリクス -> Go1.18から導入された

- 型エイリアス
参考:
https://zenn.dev/numa999/articles/dea7c50fd4329c
https://qiita.com/tenntenn/items/c3afc87a20d9f50998bb

#### 3-1-2. 型変換(型のキャスト)
ある型から別の型に変換する事(型変換を**キャスト**と呼ぶ)
`T(v)` -> 変数vをT型に変換することを表す

```go:
package main

func main() {
	var f float64 = 10
	var n int = int(f)
	println(n)
}

/* 実行結果 */
// 10
```

#### 3-1-3. TRY 組み込み型(数値)
- 修正対象プログラム
```go:
package main
func main() {
	var sum int // 2案
	sum = 5 + 6 + 3
	avg := sum / 3
	if avg > 4.5 { // 1案
		println("good")
	}
}
```
`avg > 4.5`の時点で`int > float`で比較になっているためエラーになる。
1. avgに対する比較をint(整数値)で比較する(例: avg > 4など)
2. `var sum int`を`var sum float64`としてint型ではなく、浮動小数点型にする

#### 3-1-4. TRY 組み込み型(真偽値)
- 検証プログラム
```go:
package main
func main() {
	var a, b, c bool
	if a && b || !c {
		println("true")
	} else {
		println("false")
	}
}
```
a && b -> aとbがtrueの時
b || !c -> bがtrue, cがfalse, どちらかtrueならtrue

- 埋める真理値表
https://docs.google.com/presentation/d/1DtWB-8FcnNb9asxSpIaOLYbAEc9OjBAwMLNxKnPA8pc/edit#slide=id.g4cbe4d134e_0_125


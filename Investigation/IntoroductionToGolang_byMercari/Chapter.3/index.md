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
a, b, cがそれぞれtrue, or falseの時
a && b -> aとbが同じだったらその値になる
b || !c -> a && b の結果がbに入る。bとcの否定を比較する。

- 埋める真理値表
https://docs.google.com/presentation/d/1DtWB-8FcnNb9asxSpIaOLYbAEc9OjBAwMLNxKnPA8pc/edit#slide=id.g4cbe4d134e_0_125

| a   | b   | c   | a && b | !c  | a && b ll !c |
| --- | --- | --- | ------ | --- | ------------ |
| F   | F   | F   | F      | T   | T            |
| F   | F   | T   | F      | F   | F            |
| F   | T   | F   | F      | T   | T            |
| F   | T   | T   | F      | F   | F            |
| T   | F   | F   | F      | T   | T            |
| T   | F   | T   | F      | F   | F            |
| T   | T   | F   | T      | T   | T            |
| T   | T   | T   | T      | F   | T             |

合っていた。

#### 3-1-5. コンポジット型
- 複数のデータ型が集まって1つのデータ型になっている
| 型の種類 | 説明                 |
| ---- | ------------------ |
| 構造体  | 型の異なるデータ型を集めたデータ型  |
| 配列   | 同じ型のデータを集めて並べたデータ型 |
| スライス | 配列の一部を切り出したデータ型    |
| マップ  | キーと値をマッピングさせたデータ型                   |

スライスは処理と思ってたけど何か違いそう。
Pythonと混同しないようにしなくては。
#### 3-1-6. コンポジット型のゼロ値
- 0値に対する表現方法が違う
構造体, 配列 -> **要素(フィールド)がすべてゼロ値**
スライス, マップ -> **makeなどで初期化が必要なためnil**

#### 3-1-7. 型リテラル
- 型の具体的な定義を書き下ろした型の表現方法
- コンポジット型などを表現するために使う
- 変数定義やユーザ定義型などで使用する

```go:
// int型のスライスの型リテラルを使った変数定義
var ns []int
// mapの型リテラルを使った変数定義
var m map[string]int
```
型リテラルは、`[]int`, `[string]int`のように、**型自身(まだ名前が付与される前のもの)こと**を指す。
`[]int`だったら、`[]`がスライスを表し、`int`は中の要素がint型である事を指している。

#### 3-1-8. 構造体
- 型の異なるデータ型の変数を集めたデータ構造
1. 各変数はフィールドと呼ばれる
2. フィールドの型は異なっても良い
3. フィールドの型には組み込み型以外も使える

```go:
var p struct {
  /* 1. name, ageをフィールドと呼ぶ */
  /* 2. name, ageのそれぞれの型は違う */
  name string
  age int
}
```

#### 3-1-9. 構造体リテラル
構造体の初期化
[こっち](https://nishinatoshiharu.com/go-structure-initialize/)の方が分かりやすかった。

#### 3-1-10. 文法で理解しよう


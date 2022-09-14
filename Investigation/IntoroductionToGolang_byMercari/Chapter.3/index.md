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
**プログラミングの言語の文法は決まっている**
- 一見難しい記述方法でも文法上ではそんなに変わらない

```go:
/* 変数定義の文法 */
var 変数名 型

// int型の変数
var n int

// 構造体の変数
var p struct {
	name string
	age int
}
/* 構造はどちらも同じ！ */
```

#### 3-1-11. フィールドの参照
"."(ドット)でアクセスする。

```go:
package main

func main() {
	p := struct {
		name string
		age  int
	}{name: "Gopher", age: 10}
	// フィールドにアクセスする例
	p.age++
	println(p.name, p.age)
}
```
構造体`p`のフィールドname, ageにアクセス -> `p.name`, `p.age`

#### 3-1-12. 配列
**同じ型のデータを集めて並べたデータ構造**
- **要素の型はすべて同じ**
- 要素数が違えば別の型
- 要素数は変更できない

```go:
/* 型と要素数がセット */
// int型, 要素数5個
var ns [5]int
```

#### 3-1-13. 配列の初期化
- 配列の初期化のいろいろ

```go:
/* 配列の宣言時 */
/* 配列自体の宣言: ゼロ値で初期化 */
var ns1 [5]int

/* 新しく要素を追加: 配列リテラルで初期化 */
var ns2 [5]int{10, 20, 30, 40, 50}

/* 要素数を値から推論 */
ns3 := [...]int{10, 20, 30, 40, 50}

/* 5番目が50, 10番目が100, 他の要素の値は0, 全体の要素数は11の配列 */
ns4 := [...]int{5: 50, 10: 100}
```

```go:
package main

import "fmt"

func main() {
	var ns2 = [5]int{10, 20, 30, 40, 50}
	ns3 := [...]int{10, 20, 30, 40, 50}

	fmt.Println(ns2)
	fmt.Println(ns3)
}

/* 実行結果 */
// [10 20 30 40 50]
// [10 20 30 40 50]
```
`[...]` これは特徴的

#### 3-1-14.　配列の操作
配列の操作
```go:
ns := [...]int{10, 20, 30, 40, 50}
println(ns[3])

// この操作はできない
println(ns)
println(ns[:])

// これはできる
println(len(ns)) // 長さ
fmt.Println(ns)
fmt.Println(ns[:])
```
軽く調べたがいい記事にヒットしない。

#### 3-1-15. 配列とスライス
配列とスライスの違いを整理しておく。
- 配列
  - 複数の要素を1つにまとめる
  - 型はすべて同じ
  - 要素数を指定する必要がある(**固定長**)
    - 要素の追加を行った場合、指定した要素数とずれてしまうためエラーが起きる

- スライス
  - 複数の要素を1つにまとめる
  - 型はすべて同じ
  - 要素数を指定する必要がない(**可変長**)
    - 自由に配列の要素を追加する事ができる

#### 3-1-16. スライス
**配列の一部を切り出したデータ構造**
引用: [スライス image](https://tenntenn.dev/images/qiita-5229bce80ddb688a708a-1.png)
スライスした時点でのポインタを持つ(持つという表現が正しいかはあやしい)

#### 3-1-17. スライスの初期化
```go:
package main

import "fmt"

func main() {
	// ゼロ値はnil
	var ns1 []int
	fmt.Println(ns1)

	// 長さと容量を指定して初期化
	// 各要素はゼロ値で初期化される
	ns1 = make([]int, 3, 10)
	fmt.Println(ns1)

	// スライスリテラルで初期化
	// 要素数は指定しなくてよい
	// 自動で配列は作られる
	var ns2 = []int{10, 20, 30, 40, 50}
	fmt.Println(ns2)

	ns3 := []int{5: 50, 10: 100}
	fmt.Println(ns3)
}

/* 実行結果 */
/*
[]
[0 0 0]
[10 20 30 40 50]
[0 0 0 0 0 50 0 0 0 0 100]
*/
```

#### 3-1-18. スライスと配列の関係
**スライスはベースとなる配列が存在する**
- 前提知識
  - make()
`make([]<型名>, <要素数(len)>, <容量(何個入るか: cap)>)`
```go:
/* 下の2つは大体同じ処理 */
ns := make([]int, 3, 10) // 要素数3個のスライス(最大容量10個)

var array [10]int // 要素数10の配列
ns := array[0:3] // or array[:3], 要素数3個のスライス

/* ------------------------------------------- */

/* 下の2つは大体同じ処理 */
ms := []int{10, 20, 30, 40, 50} // 10~50までの要素数が5個

var array2 = [...]int{10, 20, 30, 40, 50}
ms := array2[0:5] // or array[:], index0~4を取り出し、10~50までの要素数が5個
```

#### 3-1-19. スライスの操作
```go:
ns := []int{10, 20, 30, 40, 50}
// 要素にアクセス
println(ns[3])
// 長さ
println(len(ns))
// 容量
println(cap(ns))
// 要素の追加
// 容量が足りない場合は背後の配列が再確保される
ns = append(ns, 60, 70)
println(len(ns), cap(ns)) // 長さと容量

/* 実行結果 */
/*
40
5
5
7 10
*/
```

#### 3-1-20. appendの挙動
- 容量が足りる場合
  - 新しい要素をコピーする
  - lenを更新する

- 容量が足りない場合
  - 元の**およそ2倍**の容量(cap)の配列を確保しなおす
    - 1024を超えた場合はおよそ1/2ずつ増える
  1. 配列へのポインタを貼り直す
  2. 元の配列から要素をコピーする
  3. 新しい要素をコピーする
  4. lenとcapを更新する


```go:
a := []int{10, 20}
// [10 20] 2
fmt.Println(a, cap(a))

/* ここで容量を超える処理を行なったため、capが倍になる */
b := append(a, 30) // (1)
a[0] = 100 // (2)
// [10 20 30] 4
fmt.Println(b, cap(b))

c := append(b, 40) // (3)
b[1] = 200 // (4)
// [10 200 30 40] 4
fmt.Println(c, cap(c))
```

#### 3-1-21. 配列・スライスへのスライス演算
```go:
ns := []int{10, 20, 30, 40, 50}
n, m := 2, 4

// n番目以降のスライスを取得する
fmt.Println(ns[n:]) // [30 40 50]

// 先頭からm-1番目までのスライスを取得する
fmt.Println(ns[:m]) // [10 20 30 40]

// capを指定する
ms := ns[:m:m]
fmt.Println(cap(ms)) // 4
```
`[:m:m]`は何を意味してるのか。
https://qiita.com/Kashiwara/items/e621a4ad8ec00974f025#%E5%AE%8C%E5%85%A8%E3%82%B9%E3%83%A9%E3%82%A4%E3%82%B9%E5%BC%8F:~:text=%E9%81%95%E3%81%86%E3%81%AE%E3%81%A7%E6%B3%A8%E6%84%8F%E3%81%97%E3%81%A6%E3%81%8F%E3%81%A0%E3%81%95%E3%81%84%E3%80%82-,%E5%AE%8C%E5%85%A8,-%E3%82%B9%E3%83%A9%E3%82%A4%E3%82%B9%E5%BC%8F
完全スライス式というそう。
容量を指定できる利点がありそう。

#### 3-1-22. スライスの要素をfor文で取得する
`for range`文を使用する

#### 3-1-23. Slice Tricks
https://ueokande.github.io/go-slice-tricks/
なんとなくやってることは分かる。

#### 3-1-24. x/exp/slicesパッケージ
スライスに関する便利なパッケージ
- ジェネリクス使用
- 標準ライブラリ入りするかも
- `"golang.org/x/exp/slices"`このパッケージをインポートする

```go:
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	ns := []int{10, 20, 30, 40, 50}

	// 削除: [10 40 50]
	ns = slices.Delete(ns, 1, 3)
	fmt.Println(ns)

	// 挿入: [10 60 70 40 50]
	ns = slices.Insert(ns, 1, 60, 70)
	fmt.Println(ns)

	// 要素があるか: true
	ok := slices.Contains(ns, 70)
	fmt.Println(ok)

	// ソート: [10 40 50 60 70]
	slices.Sort(ns)
	fmt.Println(ns)
}
```
試してみると直感的で分かりやすい。

#### 3-1-25. TRY スライスの利用
```go:
package main

func main() {
	// 3つの変数しか使わないように修正してください
	// プログラムの動作はそのままにすること

	l := [4]int{19, 86, 1, 12}
	var sum int

	for i := 0; i < len(l); i++ {
		sum += l[i]
	}

	println(sum)
}
```
要練習。

#### 3-1-26. マップ
マップ -> ハッシュ(連想配列)のこと
`var m map[string]int`
この場合、key->string, value->int
```go:
package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m)
}

/* 実行結果 */
// map[]
```

#### 3-1-27. マップの初期化
- 初期化方法いろいろ

```go:
// ゼロ値はnil
var m map[string]int
// makeで初期化
m = make(map[string]int)
// 容量を指定できる
m = make(map[string]int, 10)
// リテラルで初期化
m := map[string]int{"x": 10, "y": 20}
// 空の場合
m := map[string]int{}
```
試して見た感じ、`m = make(map[string]int)`これはスライスを宣言するのではなく、初期化処理のイメージ。
`var m map[string]int`ここでマップを作成してあるものを初期化する。`var m`とせずに最初から`make`を使用するようなことはできなかった。

makeは変数宣言に使えるが、スライスの場合は**配列を参照する**という説明がしっくりきました。
https://wireless-network.net/go-make-memory/

#### 3-1-28. マップの操作
```go:
package main

func main() {
	m := map[string]int{"x": 10, "y": 20}

	// キーを指定してアクセス
	println(m["x"])

	// キーを指定して入力
	m["z"] = 30

	// 存在を確認する
	n, ok := m["z"]
	println(n, ok)

	// キーを指定して削除する
	delete(m, "z")

	// 削除されていることを確認
	n, ok = m["z"] // ゼロ値とfalseを返す
	println(n, ok)
}
```
何をやってるかは理解できた。

#### 3-1-29. マップの要素をfor文で取得する
`for range`文を使用する

#### 3-1-30. x/exp/maps　パッケージ
マップに関する便利なパッケージ
- ジェネリクス使用
- 標準ライブラリ入りするかも

Playgroundで動かす場所が空白だった。

#### 3-1-31. コンポジット型を要素にする
- コンポジット型とは
  - 複数のデータを1つの集合として表す型

- スライスの要素がスライス(2次元スライス)
  - 例: [][]int
- マップの値がスライスの場合
  - 例: map[string][]int
- 構造体のフィールドの型が構造体

何となく分かる。

####　3-1-32. ユーザ定義型
typeで名前を付けて新しい型を定義する

MyInt, MyWriter, Personがユーザ定義した型(型名)
```go:
// 組み込み型を基にする
type MyInt int

// 他のパッケージの型を基にする
type MyWriter io.Writer

// 型リテラルを基にする
type Person struct {
     Name string
}
```

#### 3-1-33. Underlying type
空欄だったので、意味は調べてみた。
参考記事: https://qiita.com/behiron/items/89bf7292aec48b097fe4

```go:
// 組み込み型を基にする
type MyInt int
```
これを元にすると、intという基底となる型のことを"underlying type"と呼ぶ。

#### 3-1-34. ユーザ定義型の特徴
- 同じUnderlying typeを持つ型同志は型変換できる
```go:
type MyInt int
var n int = 100
m := MyInt(n)
n = int(m)
```

- 型無し定数から明示的な型変換は不要
	-	デフォルトの方からユーザ定義型へ変換できる場合
```go:
// 10秒を表す（time.Duration型）
d := 10 * time.Second
/* type Duration int64 */
```
(型推論から来てるのかなぁと)

#### 3-1-35. TRY ユーザ定義型の利用
構造体でやるのか〜
`type Score int`かなぁという思考まで。
```go:
/*
次の仕様のデータ構造を考えてみてください
とあるゲームの得点を集計をするプログラム
ゲームの結果は0点から100点まで1点刻みで点数が付けられる
集計は複数回のゲームの結果をもとにユーザごとに行う
どういうデータ構造で1回のゲーム結果を表現すべきか
適切だと思うユーザ定義型を定義してください
*/

package main

type Score struct {
	UserName string
	GameID int
	Point  int
}

func main() {
}
```

#### 3-1-36. 型エイリアス(Go1.9以上)
localがgo1.19.1なので関係ある。
- 型エイリアスを定義できる
  - 完全に同じ型
  - キャスト(変数を別の変換すること)不要
`type Applicant = http.Client`

- **型名を出力する%T**が同じ元の型名を出力する
```go:
package main

import (
	"fmt"
	"net/http"
)

type Applicant = http.Client

func main() {
	fmt.Printf("%T", Applicant{})
}

/* 実行結果 */
// http.Client
```

## 3-2. 関数
**一連の処理をまとめたもの**
- 引数で受け取った値を基に処理を行い、戻り値として結果を返す機能
  - 必ずしも引数や戻り値が無くても良い
  - 戻り値(返り値): 関数の出力となるもの

- 関数の種類
  - 組み込み関数
    - 言語の機能として組み込まれている関数(ビルドイン)
  - ユーザ定義関数
    - ユーザが定義した関数

#### 3-2-1. 関数呼び出し
- 引数を指定して呼び出す
  - 引数は、**変数や式を指定しても良い**
  - 引数が複数ある場合はカンマで区切って指定する
  - 戻り値がある場合は、変数に代入したり式中で使う
例: `x := f(10, 1+1, y)`

#### 3-2-2. 組み込み関数
https://docs.google.com/presentation/d/1DtWB-8FcnNb9asxSpIaOLYbAEc9OjBAwMLNxKnPA8pc/edit#slide=id.g4cbe4d134e_0_280

#### 3-2-3. 関数の定義
- 関数の定義方法
```go:
package main

import (
	"fmt"
)

/* add->関数名, x int,y int->引数とその引数の型, int->戻り値の型 */
func add(x int, y int) int {
	// 戻り値(引数を加算して返す)
	return x + y
}

func main() {
	fmt.Println(add(10, 20))
}

/* 実行結果 */
// 30
```

- 複数の戻り値を返す
```go:
/* (x, y int)->xとyがint型であることをまとめて記述 */
/* (int, int)-> 戻り値yとxがそれぞれint型であることを記述 */
func swap(x, y int) (int, int) {
	return y, x
}
```

#### 3-2-4. 多値の受け取り方
```go:
/* x->第1戻り値, y->第2戻り値 */
x, y := swap(10, 20)
```

- 省略したい場合は`_`(ブランク変数)を用いる
```go:
/* 第2戻り値を棄却 */
x, _ := swap(10, 20)
/* 第1戻り値を棄却 */
_, y := swap(10, 20)
```

#### 3-2-5. 関数の定義
- 名前付き戻り値
例題だと理解しにくい。
戻り値に名前を付けることで戻り値が何か分かるといったメリットがある。
視認性の向上や、変数のスコープを狭くすることが狙えるようです。
参考: https://zenn.dev/yuyu_hf/articles/c7ab8e435509d2

#### 3-2-6. 値の入れ替え
- 一時変数無しで値を入れ替えることができる
  - 一時避難の必要がないということ
`x, y = y, x`
なんか矢印が違う気がする。
左辺の`x`に右辺の`y`, 左辺の`y`に右辺の`x`が代入される。

#### 3-2-7. 無名関数
- 名前の無い関数のこと
  - **クロージャ**とも呼ばれる
```go:
package main

func main() {
	msg := "Hello, 世界"
	/* 以下、無名関数 */
	func() {
		/* 関数外の変数msgを参照できる */
		println(msg)
	}() // 無名関数を定義後、すぐに自身を呼び出している
}

/* 実行結果 */
// Hello, 世界
```
無名関数のメリットは**記述が短くて済むこと**。

#### 3-2-8. 関数型
```go:
package main

import (
	"fmt"
)

func main() {
	/* stringを返す関数(要素)が2つ格納されたスライスを定義 */
	fs := make([]func() string, 2)
	fmt.Println(fs) // [<nil> <nil>]

	fs[0] = func() string { return "hoge" }
	fs[1] = func() string { return "fuga" }
	for _, f := range fs {
		fmt.Println(f()) // hoge fuga
	}
}
```
- ファーストクラスオブジェクトとは
https://ja.wikipedia.org/wiki/%E7%AC%AC%E4%B8%80%E7%B4%9A%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88
- 数値や文字列のように**関数の引数として渡す**ことができる。
- **複合データ型(配列や連想配列、リスト)に格納できる**性質を持つ

上のコードでは、**関数自体をmake関数の引数として渡している**！
**ファーストクラスオブジェクトとして扱うことができる関数のことをファーストクラス関数**と呼ぶ！！

#### 3-2-9. クロージャと良くあるバグ
- 定義と実行のタイミングを気を付ける
  - 関数外の変数(自由変数)を参照している場合
  - 実行のタイミングでは値が変わっている可能性がある
例題コードだけではわからなかったので、以下を参考にしました。
https://qiita.com/xu1718191411/items/70550780d928dec96cab
言いたい事はわかった。

#### 3-2-10. 値のコピー
- 代入ではコピーが発生する
  - 代入元と同じ値が**コピー**される

```go:
package main

func main() {
	p := struct {
		age  int
		name string
	}{age: 10, name: "Gopher"}

	p2 := p // コピー
	p2.age = 20

	println(p.age, p.name)
	println(p2.age, p2.name)
}

/* 実行結果 */
// 10 Gopher
// 20 Gopher
```
p2のフィールドの値を変えても、コピー元であるpに影響はない。

#### 3-2-11. ポインタ
変数の格納先を表す値
- 値で渡される型の値に対して、**破壊的な操作**を加える際に利用する
  - 破壊的な操作 -> **関数を出てもその提供が残る**
```go:
/* *int -> intのポインタ型 */
func f(xp *int) {
*xp = 100 // 変数xpに"*"を付けることで、ポインタの指す先に値を入れる
}
/* ------------------------------------------- */
func main() {
	var x int
	f(&x) // &x -> &でポインタを取得する
	println(x)
}
```

#### 3-2-12. 内部でポインタを使っているデータ型
```go:
package main

func main() {
	ns := []int{10, 20, 30}
	ns2 := ns
	ns[1] = 200
	println(ns[0], ns[1], ns[2])
	println(ns2[0], ns2[1], ns2[2])
}

/* 実行結果 */
// 10 200 30
// 10 200 30
```
ns[1]の更新に合わせてns2[1]の値も更新されている。
3-2-10.の構造体とは違って内部でポインタが用いられているため。
- 内部でポインタが用いられているデータ型
  - コンポジット型の一部
    - スライス
    - マップ
    - チャネル

#### 3-2-13. TRY 奇数偶数判定関数
- 奇数偶数判定関数の作成
  - 以下のプログラムの条件式の部分で使用
```go:
package main
func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		if i%2 == 0 {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}
```

```go:
package main

import "fmt"

func check_num(x int) string {
	if x%2 == 1 {
		return "-奇数"
	} else {
		return "-偶数"
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		s := check_num(i)
		fmt.Printf("%d - %s\n", i, s)
	}
}
```
if, elseifでやってたらエラーが出てた。
if, elseにしないといけなかった。
変数の命名規則的には短い変数名が推奨されている点も気にする。

#### 3-2-14. TRY 複数戻り値の利用
- 値を入れ替えるswap関数を実装
  - 次のコードが正しく動作するようにする
```go:
package main
func main() {
	n, m := swap(10, 20)
	println(n, m)
}
```

```go:
package main

func swap(y, x int) (int, int) {
	return x, y
}

func main() {
	n, m := swap(10, 20)
	println(n, m)
}
```
syntaxエラーが出た。
基本的な書き方を身につけていく。

#### 3-2-15. TRY ポインタ
- 値を入れ替えるswap2関数を
  - 次のコードが正しく動作するようにする
```go:
package main
func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}
```

```go:
package main

func swap2(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp

}

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}
```
普通の入れ替えだと上手くいかなかった。
`swap(&n, &m)`がポインタの参照というところを理解しなくてはいけない。
関数の内部で入れ替えて、それを参照するということ。

## 3-3. メソッド
#### 3-3-1. メソッド
- レシーバとは？
```go:
p := Person{Name: "Taro"}
p.Greet("Hi")
/* ↑の"p"がレシーバと呼ぶもの */
```
**メソッドを呼び出される対象のこと**

- メソッドとは
  - レシーバと紐づけられた関数
    - データとそれに対する操作を紐づけるために用いる
    - `.`(ドット)でメソッドにアクセスする
```go:
package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	// 100をHex型として代入
	var hex Hex = 100
	// Stringメソッドを呼び出す
	fmt.Println(hex.String())
}

/* 実行結果 */
// 64(100(10進数)を16進数に変換して出力している)
```
- %x
https://qiita.com/rock619/items/14eb2b32f189514b5c3c#x-1:~:text=%25-,X,-16%E9%80%B2%E6%95%B0%E3%81%A7%E3%81%AE%E8%A1%A8%E7%8F%BE(A%2DF%E3%81%AF%E5%A4%A7%E6%96%87%E5%AD%97

- メソッドの書き方がいまいち分かってない。
**関数の場合**
```go: func
// 関数
func Add(q Calc) int {
    return q.atai1 + q.atai2
}
/* ↓ ↓ ↓ */
func 関数名（引数）戻り値の型 {
    // 処理
    return
}
```

**メソッドの場合**
```go: method
// メソッド
func (p Calc) Add() int {
    return p.atai1 + p.atai2
}
/* ↓ ↓ ↓ */
func (レシーバ 型) 関数名(引数) 戻り値の型 {
    // 処理
    return
}
```
大きな違いはレシーバの存在。

```go:
package main

import "fmt"

type Hex int

/* 2進数 */
func (h Hex) Two() string {
	return fmt.Sprintf("%b", int(h))
}

/* 8進数 */
func (h Hex) Eight() string {
	return fmt.Sprintf("%o", int(h))
}

/* 16進数 */
func (h Hex) Sixteen() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	p := Hex(100)

/* それぞれの組みのは同じ事をしている。 */

	fmt.Println(p.Two())
	fmt.Println(Hex.Two(100))

	fmt.Println(p.Eight())
	fmt.Println(Hex.Eight(100))

	fmt.Println(p.Sixteen())
	fmt.Println(Hex.Sixteen(100))
}

/* 実行結果 */
/*
1100100
1100100
144
144
64
64
*/
```
このようにすると理解できた。
参考: https://qiita.com/pei0804/items/2a3010df39623fadc3c6

#### 3-3-2. レシーバ
- メソッドに関連付けられた変数
なんであるかは分かるが、いまいち実態を掴めない。
```go:
type T int
func (t *T) M() {
    fmt.Println("method")
}

func main() {
    var t T
    t.M() // (&t).M()と同じ
}
```
これは分かる。

#### 3-3-3. ポインタ型のメソッドリスト
https://qiita.com/tenntenn/items/49bf8b5cc69c2fcfb627
ん〜〜〜

#### 3-3-4. レシーバに出来る型
- typeで定義した型
  - ユーザ定義型のこと
- ポインタ型
  - `*T`など
- 内部にポインタを持つ型
  - マップやスライスなどもレシーバにできる

#### 3-3-5. TRY レシーバに変更を与える
```go:
/*
- 次のプログラムを正しく動作するようにしてください
Incメソッドは自身を1ずつ加算する
今の実装だと正しく動かない
動かない理由を考え、意図通り動くように修正してください
*/

package main

type MyInt int

func (n *MyInt) Inc() {
	*n += 1
}
func main() {
	var n MyInt
	println(n)
	n.Inc()
	println(n)
	n.Inc()
	println(n)
	n.Inc()
	println(n)
}

// 0 1 2 3
```
いまいち分かった感じがしない。

#### 3-3-6. メソッド値
**メソッドも値として扱える**
レシーバは束縛された状態??
```go:
package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	// 100をHex型として代入
	var hex Hex = 100

	/* 変数fにメソッドを代入している */
	f := hex.String
	fmt.Println(f())
}
```
呼ぶ時に`f`じゃなくて`f()`なのが注意点。
`var hex Hex = 100`のように、固定値として関数を呼び出すメソッドのみ。動的な関数の代入はできない。

#### 3-3-7. メソッド式


- [メルカリ作のプログラミング言語Go完全入門 読破](#メルカリ作のプログラミング言語go完全入門-読破)
- [13. リフレクション](#13-リフレクション)
  - [13-1. リフレクションとは](#13-1-リフレクションとは)
      - [13-1-1. Goにおけるリフレクションとは](#13-1-1-goにおけるリフレクションとは)
      - [13-1-2. ジェネリクスとリフレクション](#13-1-2-ジェネリクスとリフレクション)
      - [13-1-3. 構造体タグの取得](#13-1-3-構造体タグの取得)
      - [13-1-4. encoding/jsonパッケージでの利用](#13-1-4-encodingjsonパッケージでの利用)
      - [13-1-5. templateパッケージでの利用](#13-1-5-templateパッケージでの利用)
  - [13-2. Value型とType型](#13-2-value型とtype型)
      - [13-2-1. Value型](#13-2-1-value型)
      - [13-2-2. 値の種類によって分岐する](#13-2-2-値の種類によって分岐する)
      - [13-2-3. Kind型](#13-2-3-kind型)
      - [13-2-4. 任意の値のリフレクション](#13-2-4-任意の値のリフレクション)
      - [13-2-5. 変数への値のセット](#13-2-5-変数への値のセット)
      - [13-2-6. reflect.Indirect関数](#13-2-6-reflectindirect関数)
      - [13-2-7. Value型の定義](#13-2-7-value型の定義)
      - [13-2-8. Type型](#13-2-8-type型)
      - [13-2-9. Type型からValue型を生成する](#13-2-9-type型からvalue型を生成する)
      - [13-2-10. Type型の定義](#13-2-10-type型の定義)
      - [13-2-11. Typeインタフェースの実装](#13-2-11-typeインタフェースの実装)
      - [13-2-12. パニックが頻繁に発生する場合](#13-2-12-パニックが頻繁に発生する場合)
  - [13-3. 構造体のリフレクション](#13-3-構造体のリフレクション)
      - [13-3-1. フィールド情報を値から取得](#13-3-1-フィールド情報を値から取得)
      - [13-3-2. フィールドに値を設定](#13-3-2-フィールドに値を設定)
      - [13-3-3. 非公開フィールドに値を設定](#13-3-3-非公開フィールドに値を設定)
      - [13-3-4. フィールド情報を型から取得](#13-3-4-フィールド情報を型から取得)
  - [13-4. チャネルのリフレクション](#13-4-チャネルのリフレクション)
      - [13-4-1. selectを実行する](#13-4-1-selectを実行する)
      - [13-4-2. SelectCase型の定義](#13-4-2-selectcase型の定義)
# メルカリ作のプログラミング言語Go完全入門 読破
# 13. リフレクション
## 13-1. リフレクションとは
#### 13-1-1. Goにおけるリフレクションとは
[参考](https://qiita.com/s9i/items/b835634d84bba5574d0a#%E3%83%AA%E3%83%95%E3%83%AC%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3%E3%81%A8%E3%81%AF):
> プログラム実行時に、動的にプログラムの構造を読み取ったり置き換えたりする手法のこと

- reflectパッケージを用いる
実行時に値や型の情報を取得できる
取得した情報を元に関数を実行したり、値を生成できる

- 利点: 実行時にならないと分からない情報を取得できる
  - インタフェース型の変数に入っている実際の値の型
    - ダイナミックタイプ
  - 型によらない汎用的な関数やメソッド作れる

- 使用例
参考: [reflectを使って型についての情報を得る](https://qiita.com/atsaki/items/3554f5a0609c59a3e10d#reflect%E3%82%92%E4%BD%BF%E3%81%A3%E3%81%A6%E5%9E%8B%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6%E3%81%AE%E6%83%85%E5%A0%B1%E3%82%92%E5%BE%97%E3%82%8B)
```go:
package main

import (
    "fmt"
    "reflect"
)

func main() {
    type MyInt int
    var x MyInt
    v := reflect.ValueOf(x) // ValueOfでreflect.Value型のオブジェクトを取得
    fmt.Println(v.Type())   // Typeで変数の型を取得
    fmt.Println(v.Kind())   // Kindで変数に格納されている値の種類を取得
}

// output
// main.MyInt
// int
```

#### 13-1-2. ジェネリクスとリフレクション
ジェネリクスはGo1.18で登場。Go1.17まではなかった。
ジェネリクスはリフレクションで行っていたようなことを(一部)コンパイル時にできるようにしてくれる。

調べてみると、リフレクションはあまり良くない？
可読性の低下、パフォーマンスが悪いなど。
リフレクションでしか実現できないこともあるが、使用を減らしたい。
ジェネリクスを使うことでリフレクションの使用を減らすことができる。

#### 13-1-3. 構造体タグの取得
実行時には**リフレクション**でしか取得できない
(goパッケージを使った静的解析でも取得はできる)
構造体タグの目的はフィールドへの情報付加(=視認的向上?)

```go:
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

#### 13-1-4. encoding/jsonパッケージでの利用
- 任意の型の値をJSONに相互変換するために使われる
  - jsonパッケージはどんな型の値が渡されるか実行時まで分からない
  - 実行時に値や型の情報を取得するために使われる

- フィールドを対応させるために使う
```go:
/* Goの構造体 */
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/* ------------------------------- */

/* JSONのオブジェクト */
{
	"name": "Gopher",
	"age":  11
}
```
  - 構造体のフィールドとJSONのフィールドを対応させる
  - 構造体のタグをリフレクションで取得する

#### 13-1-5. templateパッケージでの利用
任意の型の値をテンプレートに埋め込む
html/templateパッケージとtext/templateパッケージを使用

1. templateに埋め込む
`Hello, {{.Name}}`
2. タグに名前を入れる
`Person{ Name: "Gopher" }`
3. 生成される(テキスト)
`Hello, Gopher`

分かりにくそうではある...

## 13-2. Value型とType型
#### 13-2-1. Value型
任意の値を表す型

```go:
package main

import (
	"fmt"
	"reflect" // 値を全てValue型で表す
)

func main() {
	v := reflect.ValueOf(100) //
	// int 100
	fmt.Println(v.Kind(), v.Int())

}
```

#### 13-2-2. 値の種類によって分岐する
```go:
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 100を"hello"やtrueに変えてみよう
	v := reflect.ValueOf(true)
	switch v.Kind() {
	case reflect.Int:
		fmt.Println(v.Int())
	case reflect.String:
		fmt.Println(v.String())
	case reflect.Bool:
		fmt.Println(v.Bool())
	}
}
```
動的に変わる値を取得して処理をするなどに利用できそう。

#### 13-2-3. Kind型
値の種類を表す型
参照: https://docs.google.com/presentation/d/1QLDxUq5Ne_GnlZzVmINUMjGgwn9QmxUKqaBj3mi944I/edit#slide=id.g8503300ad1_0_100

#### 13-2-4. 任意の値のリフレクション
interface{}型の値をリフレクションする
interface{}型 -> **任意の型が実装していることになるインタフェース**
- リフレクションするとは？
  - Value型を取得するという意味合いかなと

```go:
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 100を"hello"やtrueに変えてみよう
	// 空のインタフェースに値を入れる(100を入れるとint型として実装していることになる)

	var data interface{} = 100
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Int:
		fmt.Println(v.Int())
	case reflect.String:
		fmt.Println(v.String())
	case reflect.Bool:
		fmt.Println(v.Bool())
	}
}
```

#### 13-2-5. 変数への値のセット
**ポインタ経由でないとセットできない**
参考: [[Go言語] reflectパッケージで変数の値を変える - Qiita](https://qiita.com/tenntenn/items/3add893529707c837b4f)

セット-> 変数から変数に値を入れる事、という意味で良さそう
ある変数からリフレクションして新たな変数に代入しようとすることはできない。
ポインタを経由というのはポインタを参照すると言う方がしっくりくる。
[そのイメージ](https://docs.google.com/presentation/d/1QLDxUq5Ne_GnlZzVmINUMjGgwn9QmxUKqaBj3mi944I/edit#slide=id.g873b6bd538_0_6)
ポインタ経由でセットする際には**Elem()メソッド**を呼び出す。

```go:
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n int = 100

	// ValueOfにnの値がコピーされて渡される = 変数の情報は含まれない
	v1 := reflect.ValueOf(n)
	if !v1.CanSet() {
		fmt.Println("v1にはSetできない")
	}

	// Elemメソッドでポインタの指す先を取得
	np := &n
	v2 := reflect.ValueOf(np).Elem()
	if v2.CanSet() {
		v2.Set(reflect.ValueOf(200))
	}
	fmt.Println(n)
}
```

#### 13-2-6. reflect.Indirect関数
デリファンスを行う
デリファンス？ヒットしない...
参考: [Indirect](https://zenn.dev/lken/articles/36073ec89ff0ca#(value)-indirect(reflect.value)-reflect.value:~:text=(Value)%20Indirect(reflect.Value)%20reflect.Value)

```go:
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n int = 100

	v1 := reflect.ValueOf(n)
	fmt.Println(reflect.Indirect(v1)) // 100

	np := &n
	v2 := reflect.ValueOf(np)
	fmt.Println(reflect.Indirect(v2)) // 100
}
```
値をセットせずに、実態を参照している。
ポインタで経由せずに出力が可能のようだ。
ポインタを参照する場合はElemメソッドの結果、それ以外の場合は引数をそのまま返す。

Elemメソッドは要素型(配列、チャネルなど)を取得するメソッドだが、Value型に対しても使用することが出来る。Value型の場合の返却値は値。要素型でもValue型でもない場合はpanicを引き起こす。

- 値の比較
  - reflect.DeepEqual関数を用いる

#### 13-2-7. Value型の定義
Value型 -> 構造体(非公開フィールド: 自分で定義するフィールドではなく、外部パッケージなどで定義されたフィールド)
中身 ->
1. typ: 型情報
2. ptr: ポイント値または値へのポインタ
3. flag: メタデータ(埋め込み)

#### 13-2-8. Type型
任意の型を表す型
- reflectパッケージでは型を全てType型で表す
reflect.TypeOf関数、またはValue.Typeメソッドで取得できる

```go:
t1 := reflect.TypeOf(100)
fmt.Println(t1) // int
t2 := reflect.ValueOf("hello").Type()
fmt.Println(t2) // string
```

#### 13-2-9. Type型からValue型を生成する
- reflect.New関数
```go:
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// new(int)と同じようにポインタになる
	v := reflect.New(reflect.TypeOf(0))
	// ptr int 0
	fmt.Println(v.Kind(), v.Elem().Kind(), v.Elem())
}

/* "0"(int型)がポインタ型になっている。 */
```

- reflect.NewAt関数
```go:
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var n int

	/* reflect.NewAt( */
	v := reflect.NewAt(reflect.TypeOf(0), unsafe.Pointer(&n))
	v.Elem().SetInt(100)
	fmt.Println(n) // 100
}
```
いまいちよくわかってないが、指定した型+ポインタを指定して操作できる感じか。

#### 13-2-10. Type型の定義
型情報を提供するインタフェース
- 構造体のフィールドやメソッドに関する情報を提供する
- Value型と同じようなメソッドもある
- 非公開なメソッドを持つ

#### 13-2-11. Typeインタフェースの実装
rtype型が共通部分を提供する
難しいので後回し

#### 13-2-12. パニックが頻繁に発生する場合
- 対処方法
1. Kindを調べる
Kindを調べて、呼べないメソッドを呼ばないようにする

2. Can系のメソッドを用いて事前にチェックする
CanSetなど

3. しっかりとテストを書く
予想しない型の値が渡される可能性がある
interface{}型で渡される場合は、全てのKindの値でテストする

## 13-3. 構造体のリフレクション
#### 13-3-1. フィールド情報を値から取得
Value型のField*系のメソッドを用いる
※構造体の定義の順番やフィールド名が変わることがある。
アクセス仕方(指定するメソッド)によって引数の指定の仕方も変わる。

```go:
package main

import (
	"fmt"
	"reflect"
)

func main() {
	type A struct{ n int }
	type B struct {
		s string
		a A
	}
	v := reflect.ValueOf(B{s: "hoge", a: A{n: 100}})

	// .s
	fmt.Println(v.Field(0)) // hoge

	// .a.n
	fmt.Println(v.FieldByIndex([]int{1, 0})) // 100

	// .s
	fmt.Println(v.FieldByName("s")) // hoge
}
```

#### 13-3-2. フィールドに値を設定
**ポインタ経由でないとダメ**
非公開なフィールドにも設定できない

```go:
package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := struct{ N int }{N: 100}
	v := reflect.ValueOf(&a)
	v.Elem().Field(0).SetInt(200)
	fmt.Println(a)　// {200}
}
```
値を抜き出すというよりは、構造体のフィールドにアクセスして参照しているイメージ。
ポインタ経由だと参照がカギかもしれぬ。

#### 13-3-3. 非公開フィールドに値を設定
unsafe.Pointerとreflect.NewAtを使う
テスト以外での使用はおすすめしない

```go:
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	/* 非公開フィールドに値を設定できない */
	a := struct{ n int }{n: 100}
	v := reflect.ValueOf(&a)
	fv1 := v.Elem().Field(0)
	fmt.Println(fv1.CanSet()) // false

	/* このやり方だと設定できる */
	ptr := unsafe.Pointer(fv1.UnsafeAddr())
	fv2 := reflect.NewAt(fv1.Type(), ptr)
	if fv2.Elem().CanSet() {
		fv2.Elem().SetInt(200)
	}
	fmt.Println(a) // {200}
}
```

#### 13-3-4. フィールド情報を型から取得
Type型のField*系のメソッドを用いる
構造体タグの情報も取得できる

```go:
package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	p := Person{Name: "Gopher", Age: 11}
	t := reflect.TypeOf(p)
	f, _ := t.FieldByName("Name")　// フィールド名を指定
	fmt.Println(f.Tag.Get("json")) // name -> 返す値がタグ情報

}
```

## 13-4. チャネルのリフレクション
#### 13-4-1. selectを実行する
reflect.Select関数を用いる
`func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)`

- 引数
  - cases -> caseを表すSelectCase型のスライス

- 戻り値
  - chosen -> 選ばれたcaseのインデックス
  - recv -> 受け取った値
  - recvOK -> 受け取れたか

#### 13-4-2. SelectCase型の定義
```go:
type SelectCase struct {
	Dir  SelectDir // 受信/送信/デフォルト
	Chan Value     // Valueにしたチャネル
	Send Value     // 送信する値
}

type SelectDir int
const (
       SelectSend    // case Chan <- Send
       SelectRecv    // case <-Chan:
       SelectDefault // default
)
```

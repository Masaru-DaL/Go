# メルカリ作のプログラミング言語Go完全入門 読破
# 12. テキスト処理
## 12-1. 簡単なテキスト処理
#### 12-1-1. テキスト処理とGo
- Goではコマンドラインツールが作りやすい
  - シングルバイナリ・クロスコンパイル
  - 標準パッケージが充実している
  - バッチ処理(処理をまとめておいて、順次処理すること)を書くことも多い

- テキスト処理
  - CSVやXMLなど、テキストを入力とするバッチ処理も多い
  - 文字コードの変換や半角・全角などの変換
  - Goは**テキスト周りのライブラリが充実している**
    - テキスト処理に向いているということ

#### 12-1-2. 1行ずつ読み込む
[bufio](https://pkg.go.dev/bufio).[Scanner](https://pkg.go.dev/bufio#Scanner:~:text=WriteTo%20without%20buffering.-,type%20Scanner%20%C2%B6,-added%20in%20go1.1)を使用する
> Scannerは、改行で区切られたテキスト行のファイルなどのデータを読み取るための便利なインタフェースを提供します。

```go:
// 標準入力から読み込む
/* NewScanner -> 引数から読み込む新しいScannerを返す */
scanner := bufio.NewScanner(os.Stdin)

/* Scanで呼び出し、文字列で返す */
// 1行ずつ読み込んで繰り返す
for scanner.Scan() {
	//1行分を出力する
	fmt.Println(scanner.Text())
}
// まとめてエラー処理をする
if err := scanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
}
```

#### 12-1-3. SplitFunc型
分割するアルゴリズムを表す型
*bufio.Scanner型の[Split](https://pkg.go.dev/bufio#Scanner.Split:~:text=func%20(*Scanner)%20Split%20%C2%B6)メソッドで設定する
```go:
type SplitFunc func(data []byte, atEOF bool) (
 /* 戻り値 */ advance int, token []byte, err error)

/* ------------------------------- */
scanner := bufio.NewScanner(os.Stdin)
/* 標準入力から読み込んだものを分割処理 */
 scanner.Split(bufio.ScanBytes) // 1バイトごと
 scanner.Split(bufio.ScanLines) // 1行ごと（デフォルト）
 scanner.Split(bufio.ScanRunes) // コードポイントごと
 scanner.Split(bufio.ScanWords) // 1単語ごと
```

#### 12-1-4. strconvパッケージ
[strconv](https://pkg.go.dev/strconv)
文字列と他の型の変換を行うパッケージ
```go:
package main

import (
	"fmt"
	"strconv" // strconvをインポートする
)

func main() {
	// 文字列をint型に変換: 100 <nil>
	fmt.Println(strconv.Atoi("100"))

	// int型を文字列に変換: 100円
	fmt.Println(strconv.Itoa(100) + "円")

	// 100を16進数で文字列にする: 64
	fmt.Println(strconv.FormatInt(100, 16))

	// 文字列をbool型にする: true <nil>
	fmt.Println(strconv.ParseBool("true"))
}

/* 実行結果 */
/*
100 <nil>
100円
64
true <nil>
*/
```

#### 12-1-5. 数値へ変換時の注意点
- strconv.Atoi関数で変換した値をキャストする際の注意点
オーバーフローを起こすサイズにキャストにしてもpanicにならない
変換後のint型から別のint16型などにキャストしてはいけない
する場合は最初からstrconv.ParseInt型を用いる
[gosec](https://securego.io/) -> セキュリティ上の欠陥について、golangの静的コード分析を実行する。

```go:
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// int16より大きな値:"32768"
	s := strconv.FormatInt(math.MaxInt16+1, 10)
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	if int16(n) < 0 { // オーバーフロー
		fmt.Println(n)
	}
}

/* int16 -> 正の最大値は32767まで。
	最大値を超えるとラップアラウンド(0に戻る) */
// 32768
```

#### 12-1-6. stringsパッケージ
[strings](https://pkg.go.dev/strings)
-> 文字列関連の処理を行うパッケージ

```go:
package main

import (
	"fmt"
	"strings"
)

func main() {
	// スペースで分割してスライスにする: [a b c]
	fmt.Println(strings.Split("a b c", " "))

	// スライスを","で結合する: a,b,c
	fmt.Println(strings.Join([]string{"a", "b", "c"}, ","))

	// 繰り返す: hogehoge
	fmt.Println(strings.Repeat("hoge", 2))

	// プリフィックスを持つかどうか: true
	fmt.Println(strings.HasPrefix("hoge_fuga", "hoge"))
}
```

#### 12-1-7. 文字列の置換
strings.Replace関数を使う
`strings.Replace(<置換対象の文字列>, <置換したい文字列>, <置換する文字列>, <置換回数>)`
置換回数を`-1`にすると全て置換となる。

```go:
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := strings.Replace("郷に入っては郷に従え", "郷", "Go", 1) // 1回置換する
	// Goに入っては郷に従え
	fmt.Println(s1)

	s2 := strings.Replace("郷に入っては郷に従え", "郷", "Go", -1)
	// Goに入ってはGoに従え
	fmt.Println(s2)

	s3 := strings.ReplaceAll("郷に入っては郷に従え", "郷", "Go")
	// Goに入ってはGoに従え
	fmt.Println(s3)
}

/* Replace(-1)とReplaceAllは同じ -> 全て置換する */
```

#### 12-1-8. 複数文字列の置換
strings.Replacer型を使う

```go:
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 郷 → Go、入れば → 入っては
	/* 第1, 第2引数、第3, 第4引数といったように文字列のペアで指定する
		第1引数を第2引数に置換する。といった内容を変数rに代入している */
	r := strings.NewReplacer("郷", "Go", "入れば", "入っては")

	// Goに入ってはGoに従え
	/* 置換する内容にヒットした内容部分を置換する */
	s := r.Replace("郷に入れば郷に従え") // 実際に置換するのはReplaceメソッド
	fmt.Println(s)

	// Goに入ってはGoに従え
	_, err := r.WriteString(os.Stdout, "郷に入れば郷に従え")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

// Replacerだけで行おうとする場合
// fmt.Println(strings.NewReplacer("郷", "Go", "入れば", "入っては").Replace("郷に入れば郷に従え"))
```

#### 12-1-9. コードポイント(rune)ごとの置換
参考: [Goのruneを理解するためのUnicode知識](https://qiita.com/seihmd/items/4a878e7fa340d7963fee)
strings.Map関数を使う
- **第1引数にrune型単位で置換する関数**
  - 引数には関数を指定することに注意
- 第2引数に置換したい文字列

```go:
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 小文字を大文字に変換する関数
	// + rune形に変換している
	toUpper := func(r rune) rune {
		if 'a' <= r && r <= 'z' {
			return r - ('a' - 'A')
		}
		return r
	}

	// HELLO, WORLD
	s := strings.Map(toUpper, "Hello, World")
	fmt.Println(s)
}
```

#### 12-1-10. 大文字・小文字の変換
[unicode](https://pkg.go.dev/unicode)
unicode.ToUpper関数 / unicode.ToLower関数
-> **rune**単位で大文字/小文字に変換する関数

[strings](https://pkg.go.dev/strings)
strings.ToUpper / strings.ToLower関数
-> **文字列**単位で大文字/小文字に変換する関数

#### 12-1-11. bytesパッケージ
[bytes](https://pkg.go.dev/bytes)
byte型 -> unit8, 8bit=1バイト分の表現が可能
10進数では0~255を表現できる。
byte型からstring型へのキャストが省ける

```go:
package main

import (
	"bytes"
	"fmt"
)

func main() {
	// olink -> moo
	src := []byte("olink olink olink")
	b := bytes.ReplaceAll(src, []byte("olink"), []byte("moo"))
	fmt.Printf("%s\n", b)

	// fmt.Printf("%s\n", bytes.ReplaceAll([]byte("oink oink oink"), []byte("oink"), []byte("moo")))
}

// moo moo moo
```

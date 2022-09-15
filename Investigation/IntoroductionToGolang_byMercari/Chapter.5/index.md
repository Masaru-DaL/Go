# メルカリ作のプログラミング言語Go完全入門 読破
# コマンドラインツール
## 5-1. プログラム引数
#### 5-1-1. プログラム引数
コマンドライン引数 = プログラム引数
```shell:
$ echo hello // hello -> プログラム引数
hello
```

#### 5-1-2. プログラム引数を取得する
参考:
[golang でコマンドライン引数を使う](https://qiita.com/nakaryooo/items/2d0befa2c1cf347800c3)
[Go言語 - os.Argsでコマンドラインパラメータを受け取る](https://blog.y-yuki.net/entry/2017/04/30/000000)

`os.Args`を使用する or `flag`を使用する
`flag`のが簡単そうなイメージ

## 5-2. flagパッケージ
```go:
package main

import (
	"flag"
	"fmt"
	"strings"
)

// 設定される変数のポインタを取得
var msg = flag.String("msg", "デフォルト値", "説明")
var n int

func init() {
	// ポインタを指定して設定を予約
	flag.IntVar(&n, "n", 1, "回数")
}

func main() {
	// ここで実際に設定される
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, n))
}
```

```shell:
$ go run . -msg=こんにちは -n=2
こんにちはこんにちは
```

## 5-3. 入出力
#### 5-3-1. 標準入力と標準出力
osパッケージで提供されている*os.File型の変数
- 標準入力
  - os.Stdin
- 標準出力
  - os.Stdout
- 標準エラー出力
  - os.Stderr
1. さまざまな関数やメソッドの引数として渡せる
2. エラーを出力する場合は標準エラー出力に出力する

#### 5-3-2. fmt.Fprintln関数
**出力先を指定して出力する**
- 末尾に改行をつけて表示する
- 標準出力(os.Stdout), 標準エラー出力(os.Stderr)に出力できる
- ファイルにも出力できる
```go:
package main

import "fmt"

func main() {
	fmt.Fprintln(os.Stderr, "エラー")  // 標準エラー出力に出力
	fmt.Fprintln(os.Stderr, "Hello”)  // 標準出力に出力
}
```
`Fprintln`の第1引数にどこに出力するかを指定している。

#### 5-3-3. プログラムの終了
os.Exit(code int)
```go:
func main() {
	fmt.Fprintln(os.Stderr, "エラー")
  /* プログラムの呼び出し元に終了状態を伝えられる */
  /* 0: 成功(デフォルト) */
	os.Exit(1) // 終了コード(1)を指定してプログラムを終了
}
```

#### 5-3-4. プログラムの終了(エラー)
log.Fatal
```go:
func main() {
	if err := f(); err != nil {
		log.Fatal(err)
	}
}

/* 実行結果 */
// 標準エラー出力にエラーメッセージを表示
// os.Exit(1)で異常終了させる
```
**終了コードがコントロールできないため(os.Exit(1))あまり多用しない**

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

#### 5-3-5. ファイルを扱う
- osパッケージを用いる
参考: [ファイルの読み書き](https://zenn.dev/hsaki/books/golang-io-package/viewer/file)

#### 5-3-6. defer
関数の遅延実行
```go:
package main

import (
	"fmt"
)

func main() {
	msg := "!!!"
	defer fmt.Println(msg) // msg := "!!!"を保持
	msg = "world"
	defer fmt.Println(msg) // msg = "world"を保持
	fmt.Println("hello")
}

/* 実行結果 */
// hello
// world
// !!!
```
- スタック形式で実行される(LIFO: Last In First Out)
- 引数の評価は関数が呼び出された時

#### 5-3-7. forの中でdeferは避ける
参考: [GolangでForループの中でdeferしない](https://kamatama41.hatenablog.com/entry/2019/05/31/173346)
分かりやすい。

#### 5-3-8. 入出力関連の便利パッケージ
- 一覧
https://docs.google.com/presentation/d/1KiU14z2owLUoiTYz5pKo-gP8RnP2BINmucVYJ6DfxTs/edit#slide=id.g6fee908e9d_0_170

#### 5-3-9. 1行ずつ読み込む
- [bufio.Scanner](https://zenn.dev/hsaki/books/golang-io-package/viewer/bufio#bufio.scanner:~:text=%E3%82%8F%E3%81%8B%E3%82%8B%E7%B5%90%E6%9E%9C%E3%81%A7%E3%81%99%E3%80%82-,bufio.Scanner,-bufio%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E3%81%AB)を使用する
トークン(行、単語ごとなど)ごとに読み込みできるといった利点がある。

#### 5-3-10. ファイルパスを扱う
path/filepathパッケージを使う
参考: [Go言語: path/filepathとの良いお付き合い](https://zenn.dev/foxtail88/books/a5e3c432340c28)
OSごとに扱いの違うセパレータ(`/`や`\`)に恐る事なく立ち向かえる！
```go:
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// パスを結合する
	path := filepath.Join("dir", "main.go")
	fmt.Println(path)

	// 拡張子を取る
	fmt.Println(filepath.Ext(path))

	// ファイル名を取得
	fmt.Println(filepath.Base(path))

	// ディレクトリ名を取得
	fmt.Println(filepath.Dir(path))
}

/* 実行結果 */
/*
dir/main.go
.go
main.go
dir
*/
```

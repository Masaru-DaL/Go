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



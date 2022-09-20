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

## 12-2. ioパッケージ
[io](https://pkg.go.dev/io)

#### 12-2-1. 入出力の抽象化
io.Readerとio.Writer
参考: [ioパッケージによる抽象化](https://zenn.dev/hsaki/books/golang-io-package/viewer/io#io.reader%E3%81%AE%E5%AE%9A%E7%BE%A9)

#### 12-2-2. コピー
ioReaderから呼んだデータを、io.Writerに書き込む
戻り値 -> 書き込めたバイト数, エラー
読み込む最大バイト数を指定したい場合はio.CopyN関数を使う

```go:
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("Hello, 世界")
	if _, err := io.Copy(os.Stdout, r1); err != nil {
		panic(err)
	}
	fmt.Println()
	r2 := strings.NewReader("Hello, 世界")
	// 5バイトだけ標準出力する
	if _, err := io.CopyN(os.Stdout, r2, 5); err != nil {
		panic(err)
	}
}

// Hello, 世界
// Hello
```

#### 12-2-3. io.Seekerインタフェース
io.Readerやio.Writerのオフセットを設定する
参考:
[Seek を試す](https://tokizuoh.dev/posts/lgddm8djtvqm9hlc/)]
[Goでファイルの特定位置から読む](https://reiki4040.hatenablog.com/entry/2018/08/13/080000)
offset -> 配列やデータ構造オブジェクト内の、先頭から所定の要素までの距離を示す整数
公式を見ても分からなかったが、参考にさせて頂いた内容を元に整理すると理解出来た。
`whence`にはオフセットを指定する。
- `whence`の基準
  - <位置> -> <指定値> = <実際の値>
  - 先頭 -> Seekstart = (0)
  - 現在のoffset -> SeekCurrent = (1)
  - 終端 -> SeekEnd = (2)
SeekCurrentの意味が分からなかったが、(0)または(2)で設定したものを指す意味で使われると思って良いと思う。

```go:
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, 世界")
	r.Seek(2, io.SeekStart) // 先頭から2の位置に設定
	io.CopyN(os.Stdout, r, 2) // "llo, 世界"から2文字出力
	fmt.Println()							// ll

	/* コピー後にoffsetが移動する？ */
	r.Seek(-4, io.SeekCurrent) // "Hell<current>o, 世界"から-4 -> "<current>Hello, 世界"
	io.CopyN(os.Stdout, r, 7)	//
	fmt.Println()							// Hello,  -> Hello,+空白

	r.Seek(-6, io.SeekEnd)
	io.Copy(os.Stdout, r)
}
```

分かるようで分からない...

#### 12-2-4. io.Pipe関数
パイプのように接続されたReaderとWriterを作る
参考: [Go言語のio.Pipeでファイルを効率よくアップロードする方法. io.Pipeと非同期処理を活かし、ファイルアップロードのメモリ使用量を減らす | by James Kirk | Eureka Engineering | Medium](https://medium.com/eureka-engineering/file-uploads-in-go-with-io-pipe-75519dfa647b)

```go:
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r, w := io.Pipe()
	go func() {
		fmt.Fprint(w, "Hello, 世界\n")
		w.Close()
	}()
	io.Copy(os.Stdout, r)
}
```

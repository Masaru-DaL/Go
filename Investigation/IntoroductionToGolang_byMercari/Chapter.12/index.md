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
		fmt.Fprint(w, "Hello, 世界\n")	// 2. wにデータを書き込む -> rと同期する
		w.Close()												// 3. クローズ
	}()																// 1. 関数の実行
	io.Copy(os.Stdout, r)							// 4. rをコピーして出力
}
```

#### 12-2-5. 読み込みバイト数を制限する
io.LimitedReader型を用いる
Rフィールド -> 元のio.Readerを設定する
Nフィールド -> 読み込むバイト数を設定する

```go:
package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	r := &io.LimitedReader{
		R: strings.NewReader("Hello, 世界"),
		N: 5,
	}
	// Hello
	io.Copy(os.Stdout, r)
}
```

#### 12-2-6. 複数のio.Writerに書き込む
io.MultiWriter関数を用いる
同じ内容が複数のio.Writerに書き込まれる

```go:
package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2) // wに&buf1, &buf2が設定される
	fmt.Fprint(w, "Hello, 世界")
	// buf1: Hello, 世界
	fmt.Println("buf1:", buf1.String())
	// buf2: Hello, 世界
	fmt.Println("buf2:", buf2.String())
}
```

#### 12-2-7. 複数のio.Readerから読み込む
io.MultiReader関数を用いる
- 複数のio.Readerを直列につなげたようなio.Readerを生成
- 分割された複数のファイルから読み込む場合などに一度にメモリに載せなくて済む
- すでに読み込んだ部分を先頭に詰めるなどに応用できる


```go:
package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("Hello, ")
	r2 := strings.NewReader("世界\n")

	/* MultiReaderの引数に、読み込み対象を複数指定する */
	r := io.MultiReader(r1, r2)
	// Hello, 世界 -> 読み込んだ部分を先頭詰めしている。
	io.Copy(os.Stdout, r)
}
```

#### 12-2-8. io.TeeReader関数
読み込みと同時に書き込むio.Readerを作る
参考: [図解 io.TeeReader(Golang)](https://qiita.com/MasatoraAtarashi/items/42ed48729992eab292c3)
[io.TeeReader](https://christina04.hatenablog.com/entry/golang-io-package-diagrams)
引数のio.Readerのベースに読み込まれると同時に、引数のio.Writerに書き込む

```go:
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var buf bytes.Buffer
	r := strings.NewReader("Hello, 世界\n")
	tee := io.TeeReader(r, &buf) // TeeReaderに読み込ませ、同時に指定先に書き込む -> 出力
	// Hello, 世界
	io.Copy(os.Stdout, tee) // bufにも書き込まれる
	// Hello, 世界
	fmt.Print(buf.String())
}
```

## 12-3. 正規表現
[regexp](https://pkg.go.dev/regexp)

#### 12-3-1. 正規表現のコンパイル
regexp.Compile関数を用いる
Compileは正規表現を解析し、成功すればテキストと整合性を確認できるRegexpオブジェクト(*regexp.Regexp型)を返す。
パッケージ変数で1度しか行わない場合はMustCompile関数を使う。
使えるシンタックス: https://github.com/google/re2/wiki/Syntax

```go:
package main

import (
	"fmt"
	"regexp"
)

// パッケージの初期化時に行う
var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

func main() {
	fmt.Println(validID.MatchString("adam[23]"))

	// 関数内で行う場合はエラー処理をする
	validID2, err := regexp.Compile(`^[a-z]+\[[0-9]+\]$`)
	if err != nil { /* エラー処理 */
	}
	fmt.Println(validID2.MatchString("adam[23]"))
}

/* 実行結果はtrueかfalseで返る */
// true
// true
```

#### 12-3-2. 正規表現のマッチ
指定した文字列などが正規表現にマッチするかどうか
Matchメソッドや、MatchStringメソッドを使う

```go:
package main

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

func main() {
	re, err := regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		panic(err)
	}
	// バイト列（[]byte型）がマッチするか
	fmt.Println(re.Match([]byte("1986年01月12日")))

	// 文字列がマッチするか
	fmt.Println(re.MatchString("1986年01月12日"))

	// io.RuneReaderがマッチするか
	var r io.RuneReader = strings.NewReader("1986年01月12日")
	fmt.Println(re.MatchReader(r))
}

/* 実行結果はtrueかfalseで返る */
// true
// true
// true
```

#### 12-3-3. マッチした部分を返す
正規表現にマッチする文字列などを探す
- FindメソッドやFindStringメソッドを用いる
- FindAllメソッドやFindStringAllメソッドは個数を指定できる
  - "-1"はマッチする全てを取得する


```go:
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile(`\d+`) // \d+(-> 数字列)とマッチするかどうか
	if err != nil {
		panic(err)
	}
	// 最初にマッチするバイト列を取得
	fmt.Printf("%q\n", re.Find([]byte("1986年01月12日")))

	// すべてのマッチするバイト列を取得
	fmt.Printf("%q\n", re.FindAll([]byte("1986年01月12日"), -1))

	// 最初にマッチする文字列を取得
	fmt.Printf("%q\n", re.FindString("1986年01月12日"))

	// すべてのマッチする文字列を取得
	fmt.Printf("%q\n", re.FindAllString("1986年01月12日", -1))
}

/* 実行結果 */
/*
"1986"
["1986" "01" "12"]
"1986"
["1986" "01" "12"]
*/
```

#### 12-3-4. マッチした部分のインデックスを返す
正規表現にマッチする部分のインデックスを返す
- Find*Indexメソッドを用いる
返す内容: [0 4](スライスを返すのでindex4の1個前まで) -> [index0, index3]

```go:
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile(`\d+`)
	if err != nil { /* エラー処理 */
	}
	// [0 4]
	fmt.Println(re.FindIndex([]byte("1986年01月12日")))
	// [[0 4] [7 9] [12 14]]
	fmt.Println(re.FindAllIndex([]byte("1986年01月12日"), -1))
	// [0 4]
	fmt.Println(re.FindStringIndex("1986年01月12日"))
	// [[0 4] [7 9] [12 14]]
	fmt.Println(re.FindAllStringIndex("1986年01月12日", -1))
}
```

#### 12-3-5. キャプチャされた部分を取得
Find* Submatch*メソッドを使う
`func (re *Regexp) FindSubmatch(b []byte) [][]byte`

```go:
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}
	// ["1986"]
	fmt.Printf("%q\n", re.FindSubmatch([]byte("1986年01月12日")))
	fmt.Printf("%q\n", re.FindStringSubmatch("1986年01月12日"))
	// [["1986"] ["01"] ["12"]]
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte("1986年01月12日"), -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("1986年01月12日", -1))
	// [0 4]
	fmt.Println(re.FindSubmatchIndex([]byte("1986年01月12日")))
	fmt.Println(re.FindStringSubmatchIndex("1986年01月12日"))
	// [[0 4] [7 9] [12 14]]
	fmt.Println(re.FindAllSubmatchIndex([]byte("1986年01月12日"), -1))
	fmt.Println(re.FindAllStringSubmatchIndex("1986年01月12日", -1))

}

/* fmt.Println(re.FindSubmatch([]byte("1986年01月12日")))とすると、
[[49 57 56 54]]が返る。これはバイト文字列。49->1, 57->9...
"%q"で自然文字列に変換され、1986と返る。 */
```

#### 12-3-6. キャプチャした部分の展開
キャプチャした部分をテンプレートに展開する
ExpandメソッドやExpandStringメソッドを使う
例題コードはExpandStingなので以下に。
`func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte`
引数を4つ指定する。
テンプレートをdstに追加し、その結果を返す。
テンプレートでは、変数は`$name`または`${name}`という形式の部分文字列で示される。
`$1`, `${1}`のような純粋に数字だけの名前は、対応するインデックスを持つサブマッチを参照する。
src -> 置換対象
match -> 置換するindexを指定
FindAllStringSubmatchIndexメソッドなどでインデックスを取得する。

```go:
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// (?P<var_name>regexp)で名前をつけてキャプチャする
	re, err := regexp.Compile(`(?P<Y>\d+)年(?P<M>\d+)月(?P<D>\d+)日`)
	if err != nil {
		panic(err)
	}
	content := "1986年01月12日\n2020年03月24日"
	template := "$Y/$M/$D\n" // "${1}/${2}/${3}"でも可
	var result []byte
	for _, submatches := range re.FindAllStringSubmatchIndex(content, -1) {
		result = re.ExpandString(result, template, content, submatches)
	}

	// "1986/01/12\n2020/03/24\n"
	fmt.Printf("%q", result)
}
```

#### 12-3-7. マッチした部分を置換する
キャプチャした部分を**展開して**置換する
ReplaceAllStringメソッドを用いる
[]byte型にはReplaceAllメソッドを用いる

```go:
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		panic(err)
	}

	// $1, $2はマッチした順番を表す
	s := re.ReplaceAllString("1986年01月12日", "${2}/${3} ${1}")
	// 01/12 1986
	fmt.Println(s)
}
```

#### 12-3-8. マッチした部分を置換する
キャプチャした部分を**展開せずに**置換する
ReplaceAllLiteralStringメソッドを用いる
[]byte型にはReplaceAllLiteralメソッドを用いる

```go:
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		panic(err)
	}
	// テンプレートを展開しない
	s := re.ReplaceAllLiteralString("1986年01月12日", "${2}/${3} ${1}")
	// ${2}/${3} ${1}
	fmt.Println(s)
}
```

#### 12-3-9. 関数を指定して置換
マッチした部分を置換する関数を指定する
- ReplaceAllStringFuncメソッドを用いる
- []byte型にはReplaceAllFuncメソッドを用いる

```go:
re, err := regexp.Compile(`(^|_)[a-zA-Z]`)
if err != nil { /* エラー処理 */ }
s := re.ReplaceAllStringFunc("hello_world", func(s string) string {
	return strings.ToUpper(strings.TrimLeft(s, "_"))
})
// HelloWorld
fmt.Println(s)

/* TrimLeft -> sの先頭からcutsetに含まれるUnicodeコードポイントを除いた文字列を返す */
/* Trimした結果、hello worldに別れてそれぞれの1文字が大文字に置換しているのかな */
```

## 12-4. UnicodeとURF8
#### 12-4-1. Unicodeとrune型
[Unicode](https://ja.wikipedia.org/wiki/Unicode)
- 全世界共通で使えるようにする定めた文字コードの業界規格。
- 世界中の文字を収録し、**通し番号**を割り当て、同じコード体系で使えるようにしたもの。
  - **文字に数値を割り当てたものの集まり**
- "符号化文字集合"と呼ばれるものの1つ。
  - コンピュータ上で「どのような文字や記号を扱うのか」ということを定義したもの。
  - "変換対象となる文字の一覧"が書いてある表

- 16ビットで文字を表す
  - [Unicode一覧 0000-0FFF](https://ja.wikipedia.org/wiki/Unicode%E4%B8%80%E8%A6%A7_0000-0FFF)
  - 緑色の表の名前が<U+(表左上)+縦軸の上3桁+横軸の1桁>となる。
  - 例えばAだとしたら、U+0041となる。
    - 縦軸の4桁の内の下1桁は横軸の値になる。
このU+0041という値は**16進数で0041という値と対応付けられていることを意味する値**である。
この"0041"という値は16ビットであり、**コードポイント(符合点)**と呼ぶ。

参考: [Unicode -> UTF* への文字変換 image](https://cdn-ak.f.st-hatena.com/images/fotolife/s/shiba_yu36/20150913/20150913141538.jpg)

**rune型**
- Goの組み込み型
- Unicodeのコードポイントを表す
何ができるのか？
rune型を用いることで、文字列を1文字ずつ扱うことができる。
以下、stringの扱い方をやってみる。
参考: [Goのruneを理解するためのUnicode知識](https://qiita.com/seihmd/items/4a878e7fa340d7963fee)

```go:
package main

import "fmt"

func main() {
	s := "あ"

	for i := 0; i < len(s); i++ {
		b := s[i] // byte
		fmt.Println(b)
	}

}

/* 実行結果 */
// 227
// 129
// 130
```

ここでなぜ3回出力されているのかという疑問が湧きました。
試しに"あ"の長さを測ってみました-> `len(s)`
そうすると"3"と出力されました。
調べてみると日本語は1文字3バイトあるようです。
227-> `あ[0]`, 129-> `あ[1]`, 130-> `あ[2]`,
ということが分かりました。
記事を見てみると、このstringにindexでアクセスした時に得られるbyte値は、**文字コードをUTF-8で1byteごとに区切った値**ということが書いてあります。
"あ"のUFT-8での表現は、`E3 81 82`です。
この数値は**16進数表記**かつ**2桁で1byte**です。
これを踏まえて上のコードを16進数で出力してみます。

```go:
package main

import "fmt"

func main() {
	s := "あ"
	for i := 0; i < len(s); i++ {
    /* %xで16進数で出力する */
		fmt.Printf("% x", s[i]) // e3 81 82
	}

}

// "あ"のUTF-8の表現と一致する。
```

分かりづらいですが、UTF-8で表現された"あ"という文字を出力できたということです。
ただ、やりたいことは"あ"という文字列を"あ"と出力したいことです。
そこで利用できるのが"rune型"です。

```go:
package main

import "fmt"

func main() {
	s := "あ"
	/* ループの回数(=i)を棄却し、要素の中身(byte)だけ取り出す */
	for _, r := range s {
		// rune
		fmt.Println(r)
	}
}

// 12354
```
この時出力された"12354"という値は、**Unicodeの番号を10進数に変換したもの**です。
[変換ツール](https://www.marbacka.net/msearch/tool.php#chr2enc:~:text=%E3%81%95%E3%82%8C%E3%81%BE%E3%81%99%E3%80%82-,%EF%BC%91%E6%96%87%E5%AD%97%E5%85%A5%E5%8A%9B%E3%81%97%E3%81%A6%E8%AA%BF%E3%81%B9%E3%82%8B,-%E8%AA%BF%E3%81%B9%E3%81%9F%E3%81%84%E6%96%87%E5%AD%97)で確認します。
1. "あ"と入力して、調べるボタンを押す
2. Unicode文字番号 -> U+3042
3. (HTML数値)文字参照(10進数表記) -> (&w**12354**)

つまり、"あ"という文字列を得るにはもう一段階必要です。
手段としては簡単で、`string()`を使うだけです。

```go:
package main

import "fmt"

func main() {
	s := "あ"
	/* ループの回数(=i)を棄却し、要素の中身(byte)だけ取り出す */
	for _, r := range s {
		// rune
		fmt.Println(string(r)) // string()を使用するだけ
	}
}

// あ
```
無事"あ"と出力することができました！

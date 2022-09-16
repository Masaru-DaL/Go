# メルカリ作のプログラミング言語Go完全入門 読破
# 4. パッケージ
## 4-1. パッケージ
#### 4-1-1. Goのプログラムの構成要素
- パッケージ
  - 関数, 変数, 定数, 型をある単位でまとめたもの
  - **Goのプログラムはパッケージを組み合わせる事で実現される**

- 型
  - データの種類を表したもの

- 関数
  - 処理を意味のある単位でまとめたもの
  - パッケージに属さない組み込み関数も存在する
    - (printlnなど) -> importする必要がない

- 変数
  - 処理中に使用するデータを格納したもの

- 定数(名前付き定数)
  - コンパイル時から変わらない値に名前を付けたもの
  - Pi = 3.14 のようなもの

#### 4-1-2. Goのプログラムの構成とパッケージ
矢印の向きがいまいち理解しにくい感じがする。

[こっち](https://osamu-tech-blog.com/wp-content/uploads/2022/02/go_import.png)の方が分かりやすい感じがする。(多分言ってることは同じ)

#### 4-1-2. 複数のファイル
[複数ファイルの使用 image](https://yttm-work.jp/img/go/go_0001.png)
- 同じディレクトリ内でパッケージが混在するのはNG
- インポート文はファイルごとに記述する
- どういう分け方をしても良い(型の定義、メソッドの定義が別れても良い)

#### 4-1-3. パッケージのインポート
**import自体を"IDE"や"goimports"などのツールに任せる**のが推奨されている。(いちいち自分で打たなくて良いということか)

```go:
import (
        /* サードパーティ製と分けるために改行を挟むことが多い */
       "context"
       "fmt"

       "github.com/tenntenn/greeting"
)
```
`"github.com/tenntenn/greeting"`に変数、関数などが用意されていればこれを利用することもできる。

#### 4-1-4. パッケージ名のエイリアス
- パッケージ名に別名を付ける
  - 同じパッケージ名のパッケージを使いたい場合に使う
  - インポートパスとパッケージ名が一致していない場合に用いる
```go:
import (
  "sync"

  mysync "github.com/tenntenn/sync" // 上のsyncパッケージと名前が衝突している
  greeting "github.com/tenntenn/greeting/v2" // バージョンを指定していてインポートパスとパッケージ名が一致していない
)
func main() {
	fmt.Println(greeting.Do())
	// ...
}
```

#### 4-1-5. パッケージ外へのエクスポート
[参考](https://kazuhira-r.hatenablog.com/entry/2021/01/17/012840)
呼び出されるパッケージとして定義する際には**先頭を大文字にした識別子**を定義しておかないと、他のパッケージから利用できないという意味かと思います。

- ライブラリ
  - "fmt"とかのこと
  - main関数のないGoのプログラム
  - エクスポートされているため、使用できる

#### 4-1-6. GOPATH
- GOPATHとは？
  - **GOPATHという場所**と考えると自然な気がする
  - Goのソースコードや、ビルドされたファイルが入るパスが設定される
  - インポートされるパッケージもここから検索される

- bin, pkg, srcという3つのディレクトリがあり、各ディレクトリに入る中身が違う
  - bin -> ビルドされた実行可能ファイルが入る
  - pkg -> ビルドされたパッケージが入る
    - `pkg/GOARCH/pkgname.a`
  - src -> 実行可能なGoのコード, ライブラリのGoのコード
    - `src/cmdname/*.go`, `src/pkgname/*.go`
(合ってるかあやしい)

#### 4-1-7. GOPATHの設定方法
- 環境変数として設定される
- デフォルトが決まっている
- 複数設定できる
- `go env GOPATH`コマンドで取得可能
**試してみた**
```shell:
$ go env GOPATH
/Users/<myname>/go
```

## 4-2. パッケージ変数とスコープ
#### 4-2-1. スコープ
- 識別子(変数名、関数名など)を参照できる範囲
- 参照元によって所属するスコープが違う
- 親子関係があり親のスコープの識別子は参照できる

- Goのスコープ
  - [レキシカルスコープ](https://ja.wikipedia.org/wiki/%E9%9D%99%E7%9A%84%E3%82%B9%E3%82%B3%E3%83%BC%E3%83%97)
    - なんとなく分かる。
  - [スコープの種類](https://qiita.com/tenntenn/items/ac5940dfbca703183fdf#:~:text=%E5%AE%9F%E3%81%AF%E3%80%81%E3%82%B9%E3%82%B3%E3%83%BC%E3%83%97%E3%81%AB%E3%81%AF%E4%BB%A5%E4%B8%8B%E3%81%AE4%E7%A8%AE%E9%A1%9E%E3%81%8C%E3%81%82%E3%82%8A%E3%81%BE%E3%81%99%E3%80%82)

#### 4-2-2. ブロックスコープ
- ブロックごとのスコープ
  - 関数, if, forなど
  - ブロック内で宣言した識別子は、ブロック内でしか参照できない
```go:
package main

func f() {
	n := 100
	println(n)
}

func g() {
	// fの変数とは別もの
	n := 200
	println(n)
}

func main() {
	f()
	g()
}

/* 実行結果 */
// 100
// 200
```
ブロックが違うため、同じ識別子を用いても良いということ。
スコープが関係している！

#### 4-2-3. ファイルスコープ
```go:
// fmtがファイルスコープになる
import "fmt"
```
- `fmt.Println...`は、インポートしたパッケージをスコープとしている
- パッケージ以外を対象としていない
  - 同じ階層のファイルやモジュールを対象としていないことを言っていると思われる。

#### 4-2-4. パッケージスコープ
- パッケージ単位
- 大文字から始まる識別子は他のパッケージからも参照できる
  - パッケージ外へのエクスポートの話

#### 4-2-5. パッケージ変数
```go:
package main

var msg string = "hello"

func f() {
	println(msg)
}

/* f関数とmain関数で"msg"変数を共有したい */
func main() {

	/* main関数内(同じ関数の中)なら"msg"関数を共有できる */
	f()
	msg = "hi, gophers"
	f()
}

/* 実行結果 */
// hello
// hi, gophers
```

#### 4-2-6. ユニバーススコープ
- 組み込み型や、組み込み関数を保持するスコープのこと
  - プログラム実行時からずっとあるスコープ
  - 他のスコープの一番ルートとなるスコープ
[参考](https://www.twihike.dev/docs/golang-primer/scope#%E3%83%A6%E3%83%8B%E3%83%90%E3%83%BC%E3%82%B9:~:text=%E3%83%AD%E3%83%BC%E3%82%AB%E3%83%AB-,%E3%83%A6%E3%83%8B%E3%83%90%E3%83%BC%E3%82%B9,-%E3%81%99%E3%81%B9%E3%81%A6%E3%81%AE%E3%82%BD%E3%83%BC%E3%82%B9)
int, stringなどの型は既にint型, string型であることが決まっているもの -> ユニバーススコープ

## 4-3. パッケージの初期化
#### 4-3-1. パッケージの初期化
1. importしているパッケージリストを出す
2. 依存関係を解決する
   1. 依存されていないパッケージ -> 依存しているパッケージの順で初期化

- 各パッケージの初期化
  - パッケージ変数の初期化をする
  - init関数の実行を行う

#### 4-3-2. init関数
```go:
package main

import (
	"fmt"
)

var msg = message()

func message() string {
	return "Hello"
}

func init() {
	fmt.Print(msg)
}

func main() {
	fmt.Println(", playground")
}

/* 実行結果 */
// Hello, playground
```
実行結果から分かるように、main関数より先にinit関数が呼ばれている。
[参考](https://qiita.com/tenntenn/items/7c70e3451ac783999b4f)
- 記事も合わせて見て、init関数の特徴
  - **パッケージの初期化を行う関数**
  - init関数は明示的に呼び出せない
    - 実行順がシビアなものはinit関数には書かない
  - 複数用意しても良い

- init関数の用途と注意
  - 複雑な初期化を行う場合に用いる！
    - パッケージ変数への代入分だけでは表現できない場合
  - エラーハンドリングが必要な処理は書かない

## 4-4. ライブラリのバージョン管理
#### 4-4-1. go getでライブラリを取得する
- `go get`コマンド
  - Goのライブラリなどを取得する
  - 依存するライブラリも一緒に取得する
  - 指定した場所からダウンロード&インストールしてくれる
  - 一度取得したものは2度取得しない
  - `-u`オプションでダウンロードを強制する
```shell:
$ go get github.com/tenntenn/greeting
$ ls $GOPATH/src/github.com/tenntenn/greeting
# README.md          greeting.go
```
GOPATHに入る

#### 4-4-2. TRY ライブラリを取得してみよう
```shell:
-> go get github.com/tenntenn/greeting
# go: downloading github.com/tenntenn/greeting v1.0.0
# go: added github.com/tenntenn/greeting v1.0.0
```

```go:
package main

import "github.com/tenntenn/greeting"

func main() {
	println(greeting.Do())
}

/* 実行結果 */
// こんにちは
```

#### 4-4-3. GoDocを生成する
- GoDocとは？
  - ソースコード上に書かれたコメントをドキュメントとして扱う
  - エクスポートされたものに書くことが多い

- pkg.go.devを使ってパッケージドキュメントを読む
  - pkg.go.dev -> 自動でGoDocを生成するサービス
  - 後ろにインポートパスを付けてアクセスするとそのパッケージドキュメントが読める

https://pkg.go.dev/fmt#Println
Printlnに関するドキュメントが読めた。

#### 4-4-4. Go Modules (vgo)
[vgo: Versioned Go Command](https://github.com/golang/vgo)
Go Modules -> 外部パッケージの管理システム(vgoと呼ばれることもある)
**Go1.13~正式導入**
- modulesの特徴
  - ビルド時に依存関係を解決する
  - ベンダリングが不要になる(4-4-8節で解説)
  - GOPATHでの管理から、モジュールという概念単位でバージョン管理する
  - 互換性がなくなる場合は、インポートパスを変える
    - 依存関係絡みだと思われる
  - 可能な限り古いバージョンが優先される

#### 4-4-5. goinstallの登場
2010年2月に登場 -> 2012年2月のGo1.0リリースで`go get`に
- インポートパスのルールの確立
  - ソースコードの中にすべての依存関係が記述される
  - `go vet`などの静的解析ツールが簡単に作れるように

#### 4-4-6. go getの登場で解決したこと
- 簡単にビルドできるようになった
- 作ったものを簡単に公開・再利用できるようになった
  - go getすることで簡単に他の人が作ったパッケージを利用できる
- ビルドシステムを意識しなくて良い
  - 依存関係の解決方法などは勝手にやってくれる

#### 4-4-7. go getで解決できなかったこと
- バージョン付けとAPIの安定性
- ベンダリングとビルドの再現可能性

※v1.16から`go install`と`go get`に機能が別れ、`go install`はバイナリをインストール、`go get`は`go.mod`を編集するだけの機能になったようなので、そこは頭に入れておく。
参考: [【Go】go getは不要？go installとは？](https://zenn.dev/tmk616/articles/383fc3fbb0ec4b)

#### 4-4-8. ベンダリング
プロジェクトで利用している3rdパーティ製のパッケージのコピーを作り、プロジェクトのリポジトリのソースコードして保存すること。
- ライブラリのベンダリング
  - vendor(ディレクトリ)以下に置くとimportで優先される
  - バージョン指定はできない
```shell:
$ tree .
main.go
vendor
└── github.com
   └── mitchellh
       └── cli
```

#### 4-4-9. modules(vgo)までの流れ
[【Go言語】modulesについて理解するために過去から調べてみた](https://qiita.com/yoshinori_hisakawa/items/268ba201611401ca7935)
こちらが分かりやすかった。

#### 4-4-10. Import Compatibility Rule
**インポート互換性ルール**
- importパスが同じ場合は後方互換性を維持する
  - 後方互換性が取れない場合はimportパスを変える
```go:
import "github.com/tenntenn/hoge"

/* 後方互換性が担保できない場合(バージョン情報を付随する) */
import hoge "github.com/tenntenn/hoge/v2"
```

#### 4-4-10. Minimal Version Selection
**最小バージョン選択**
- 選択できるバージョンのうち、最も古いバージョンを選択
- どんどんバージョンアップされても常に同じ(=古い)
- 特定のバージョンを指定すれば新しいものを使う事はできる
- 差分のみの対応で良くなる

#### 4-4-11. モジュール
```go:
// My hello, world.

/* "rsc.io/hello" -> モジュール名 */
module "rsc.io/hello"

require (
    "golang.org/x/text" v0.0.0-20180208041248-4e4a3210bb54 // 依存しているモジュール
    "rsc.io/quote" v1.5.2 // v1.5.2 -> 依存しているモジュールのパージョン
)
```

#### 4-4-12. セマンティックバージョニング
[Semantic Versioning 2.0.0](https://semver.org/#introduction)
モジュール間の依存関係の解決策の1つ

- バージョンの付け方とルール
  - semverと略される
  - `v0.0.2`や、`v1.1.1`などと表記する
  - `X.Y.Z`
    - X: メジャーバージョン
    - Y: マイナーバージョン
    - Z: パッチバージョン

- バージョンの上げ方
  - 互換性が崩れる場合 -> メジャーバージョン
    - 例: **v1**.2.3 → **v2**.0.0
  - 機能追加の場合 -> マイナーバージョン
    - 例: v1.**2**.3 → v1.**3**.0
  - バグ修正の場合 -> パッチバージョン
    - 例: v1.2.**3** → v1.2.**4**

#### 4-4-13. modulesが普及するために必要なこと
- modules(vgo)が登場した背景の理解
[Go & Versioning](https://research.swtch.com/vgo)
経緯と歴史が書いてあって興味深い

- semverによるバージョン管理
  - 自分のライブラリをsemverで管理しよう
  - 自分の使っているライブラリにissueを上げる
issue: 問題、課題の解決のための情報を一元化する

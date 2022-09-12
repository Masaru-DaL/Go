# メルカリ作のプログラミング言語Go完全入門 読破
[プログラミング言語Go完全入門](https://tenn.in/go)を読破します。

[A Tour of Go](https://go-tour-jp.appspot.com/welcome/1)を1周しているので、なんとなくは理解している前提で進めます。

# 1. Goに触れる
## 1-1. Goとは？
- **Googleが開発したプログラミング言語**
  - 2022年5月の時点での最新バージョンは`Go1.18`

- 特徴
  - 強力でシンプルな言語設計と文法
  - 並行プログラミング
    - 並列処理が得意という事を目にする事は多い
  - 豊富なライブラリ群
  - 周辺ツールの充実
  - シングルバイナリ
    - 「一つのバイナリ」
  - クロスコンパイル

####　1-1-1. Goが開発された理由
- **Google内の課題を解決する為**に開発された
  - 開発のスケール面
  - 製品のスケール面

#### 1-1-2. Goの開発サイクル
- Goは毎年8月と2月にリリースされる
  - マイナーバージョンのアップ
  - リリース3ヶ月前からDevelopment Freezeになる
    - バグ修正とドキュメント修正以外は基本的に行われない
  - リリース2ヶ月前からベータ版がリリースされる
  - リリース1ヶ月前からRC(Release Candidate: リリース候補)版がリリースされる

#### 1-1-3. Go1の後方互換
- Go1の間は**言語の後方互換が保たれる**
  - セキュリティ、重大バグなどの一部例外はある
  - Go1.0で書いたコードはGo1.18でもビルドできる
    - 場合によっては`go fix`や`go fmt`を掛ける必要がある
  - Go1を対象に執筆された書籍やブログは言語使用に関しては語感が保たれているためGo1.xの間は有効である
    - ランタイムやgo toolに関しては変更される可能性あり

#### 1-1-4. Goの開発プロセスとGo2
- Goの開発プロセス
  - 実際にGoを使った中で問題を見つける
  - 解決策をコミュニティに提案し議論する
  - 実装して評価する
  - 問題なければリリースされる

- Go2とは？
  - Go1.xとして開発していき、Go1の後方互換が保てなくなるような変更を加える必要になったらGo2としてリリースされる
  - 例えば、必ずしもジェネリクスがGo2でリリースされるわけではない
    - ジェネリクスとは
      - Go1.18でリリースされた
      - ジェネリクス(Generics)を利用すると、"型が異なるだけで同じ処理をもつ複数の関数"を"1つの関数"として定義することができる。

#### 1-1-5. 開発版を試す
- `gotip`コマンドを用いる
  - Goは標準で特定のバージョンをインストールできる
  - 使用方法はgoコマンドと同じ(例: `gotip build main.go`)

- The Go Playgroundを用いる
  - 3つのバージョンが選べる(現在, 1つ前, 開発版)

#### 1-1-6. Goの特徴.1
-**強力でシンプルな言語設計と文法**-
- 読みやすい & 書きやすい
  - 複雑な記述をしにくい(愚直に書く)
  - 冗長な記述が不要

- 曖昧さの排除
  - 逆に曖昧な記述はできない

- 考えられたシンプルさ

#### 1-1-7. Goの特徴.2
-**並行プログラミング**-
- ゴールーチン
  - 並列処理の際に用いる
  - **軽量なスレッド**に近いもの
  - 呼び出す関数に`go`を付ける
```go: goroutine
/* 関数fを別のゴールーチンで呼び出す */
go f()
```

- チャネル
  - ゴールーチン間のデータのやり取り
  - **安全にデータをやり取り**できる
[引用: Channel image](https://learnetutorials.com/assets/images/go/channel/image1.png)

#### 1-1-8. Goの特徴.3
-**豊富なライブラリ**-
[標準パッケージ一覧](https://pkg.go.dev/std)

| パッケージ名            | 機能                 |
| ----------------- | ---------------- |
| fmt               | 書式に関する処理など       |
| net/http          | HTTPサーバなど        |
| archive, compress | zipやgzipなど       |
| encoding          | JSON, XML, CSVなど |
| html/template     | HTMLテンプレート       |
| os, path/filepath | ファイル操作など         |

#### 1-1-9. Goの特徴.4
-**周辺ツールの充実**-
- `go <tool名>`で実行可能, 標準/順標準で提供

| ツール名             | 機能                           |
| ---------------- | ---------------------------- |
| go build         | ビルドを行う                       |
|  go test                | `xxxx_test.go`に書かれたテストコードの実行 |
| go doc, godoc    | ドキュメント生成                     |
| gofmt, goimports | コードフォーマッター                   |
| go vet           | コードチェッカー                     |
|  gopls                | Language Server Protocol(LSP)の実装                             |

#### 1-1-10. Goの特徴.5
-**シングルバイナリ・クロスコンパイル**-
- シングルバイナリになる
  - コンパイルするとデフォルトで単一の実行可能ファイルになる
  - 動作環境を特別に用意しなくてもよい
    - デプロイ(別サーバへの配置・実行)が容易になるメリット

- クロスコンパイルできる
  - 開発環境とは違うOSやアーキテクチャ向けのバイナリが作れる
    - 1つのGoプログラムから複数のOS向け(linux, mac, Windows)のバイナリを作成できる
  - 環境変数の`GOOS`と`GOARCH`を指定する
    - GOOS -> OS指定, GOARCH -> CPUの指定
```shell:
# Windows(32ビット)向けにコンパイル
$ GOOS=windows GOARCH=386 go build
# Linux(64ビット)向けにコンパイル
$ GOOS=linux GOARCH=amd64 go build
```

#### 1-1-11. Goに入ってはGoに従え
- **言語だけではなく、文化も学ぶ**
  - ≠Goの慣習に従いなさい
  - 他の言語にも必要な考え方
  - なぜGoが開発されたのかを知ること


## 1-2. Goが利用できる領域
#### 1-2-1. Goの利用状況
- Go開発者向けのアンケート
  - 回答者のうち75%が職場でGoを使用
  - 用途は、API, CLI, Web, DevOpsなど
    - 前年比として傾向は変わっていない

- サーバサイドでの利用
  - HTTPやらそういった用途

#### 1-2-2. gRPCとGo
- gRPC
  - Googleが開発したプロトコルの1つ
  - RPCを実現するために開発された
  - 高速な通信を実現できる

Goではこれが使われているということ。

#### 1-2-3. Google App Engine for Go
- Googleが提供するPaaS(簡単に言うとGoogleの機能を使用する)
  - 高いスケーラビリティ
  - 簡単に利用でき、メンテナンスコストが低い
  - Goのインスタンスの起動は恐ろしく早い！

- Go1.11~ (gVisorベース)
  - Go Modulesが使える
  - ファイルアクセス可能
  - Go1.12から第2世代と呼ばれる
    - App Engine APIが非推奨(サポートされていない)

gVisor? App Engine API? 実体がよく分からない。

#### 1-2-4. Goで開発されている著名なOSS
- クラウン度関連のOSSが多い
  - Docker
  - Kubernetes
    - コンテナを管理する機能を提供する
  - gVisor
    - 調べた感じ、ざっくりとした理解で言うと"プロキシサーバのコンテナ版みたいな感じ"

#### 1-2-5. コマンドラインツールでの利用
- CLIツールの作成に向いている
  - 環境に依存しない点
  - 標準パッケージが豊富な点

#### 1-2-6. 組み込みやIoTでの利用
- 組み込み向け開発にも向いている
  - クロスコンパイル出来る点
  - バイナリサイズが大きくなる問題もある

#### 1-2-7. GopherJSでフロントエンド開発
- Goでフロントエンド開発が出来る、というだけで主流なわけではない

#### 1-2-8. Go Mobileでモバイルアプリ開発
- Goでモバイルアプリ開発が出来る、というだけで主流なわけではない

## 1-3. Goを学ぶには
#### 1-3-1. 怪しい情報に注意しよう
- 無署名な記事は参考にしない

- 1次情報を確認する
  - 事実確認は公式ドキュメントを確認する
  - 1次情報以外は私見が入っている事を意識して読む
  - **日本語で記載された公式のドキュメントはない**

- 記事が多い = 主流な開発手法ではない

#### 1-3-2. 学習教材
- 公式チュートリアル
https://go.dev/doc/tutorial/

- A Tour of Go
https://go.dev/tour

- Shizuoka.go
A Tour of Goを動画で解説
https://shizuoka-go.appspot.com/entry/7b8e7ff0-28e3-4f58-8561-f757f1305913

- Play with Go
Web上でハンズオン可能
https://play-with-go.dev/

- Microsoftが提供するGoのチュートリアル
https://docs.microsoft.com/ja-jp/learn/paths/go-first-steps/

- プログラミング言語Go完全入門
http://tenn.in/go

- Gopher道場 自習室
Gopher道場の動画資料を誰でも閲覧できる
https://gopherdojo.org/studyroom

- Goのハンズオン
https://docs.google.com/presentation/d/1Z5b5fIA5vqVII7YoIc4IesKuPWNtcU00cWgW08gfdjg/edit#slide=id.gd6675bb4dc_1_25

- 公式ドキュメントで学ぶ
https://docs.google.com/presentation/d/1Z5b5fIA5vqVII7YoIc4IesKuPWNtcU00cWgW08gfdjg/edit#slide=id.g11401cd09fe_1_151

- 書籍で学ぶ
https://docs.google.com/presentation/d/1Z5b5fIA5vqVII7YoIc4IesKuPWNtcU00cWgW08gfdjg/edit#slide=id.g1298f90f44e_41_0


## 1-4. Hello, World
#### 1-4-1. サードパーティパッケージも使える
- 標準パッケージ以外の利用
- `import "github.com/k0kubun/pp"`

#### 1-4-2. 複数ファイルも使える
- txtar(テキストアーカイブ)形式で記述する
  - `-- <ファイルパス> --`
  - or `-- <ディレクトリ名/ファイル名> --`
```shell:
$ tree .
.
├── fuga
│   └── fuga.go
├── go.mod
└── main.go
```
```go:
-- go.mod --
module hogera

go 1.12
-- fuga/fuga.go --
package fuga
func Fuga() string { return "Fuga" }
-- main.go --
package main
import "fmt"
import "hogera/fuga"
func main() {
   fmt.Println(fuga.Fuga())
}

/* 実行結果 */
// Fuga
```

- 上手く行かない例
  - コメントの中に`--`で始まる行を含む場合
    - `--`でファイルが区切られてしまう
```go:
package main

import "fmt"

/*
-- hogera --
*/

func main() {
   fmt.Println("Hello, playground")
}

/* 実行結果 */
// prog.go:5:1: comment not terminated

// Go build failed.
```

#### 1-4-3. 同時編集できるPlayground
- ルームごとに同時に編集が行える
- https://gpgsync.herokuapp.com/

#### 1-4-4. Goのソースコードファイルの構成
```go:
/* パッケージの定義 */
package main

/* main関数の定義 */
fun main() {
  /* println -> 画面に表示する組み込み関数 */
  println("Hello, 世界")
}
```

#### 1-4-5. コメント
`//` -> 行コメント: `//`から行末まで
`/*  */` -> ブロックコメント: `/*`から`*/`まで(複数行可)

- What(何)ではなく、Why(なぜ)を書く

#### 1-4-6. プログラムが実行されるまでの流れ
```Mermaid
graph LR
  A[コーディング作業]
  B[Goのコード]
  C[機械語]
  D[表示]

  A-->|コーディング| B
  B-->|コンパイル| C
  C-->|実行| D
```

#### 1-4-7. プログラムに問題がある場合
- コンパイルエラーになる
  - Goのコンパイラがエラーを出力する

- 解決方法
  - エラーを読んで問題を修正する
  - エラーは英語で表示されるがちゃんと読む事が大切
  - コンパイルでは検出できない実行時エラーもある

#### 1-4-8. プログラムはどこから実行されるのか
- mainパッケージのmain関数
  - 自動で実行される関数
    - main関数に処理がまとまっている
  - 初期化が終わったら必ず実行される
**main関数からプログラムを読むと処理を追いやすい**

#### 1-4-9. プログラムとライブラリ
- ライブラリの力を借りる
  - 信頼されたライブラリを用いる
  - ライブラリはそもそも外部利用されることが目的！

- ライブラリはモジュール単位で管理される
  - モジュールは複数パッケージから成る
  - プログラムから利用するにはパッケージごとに用いる(import)

#### 1-4-10. パッケージの種類
1. mainパッケージ
   - main関数が存在する
   - プログラムの起点
   - 実行可能なGoのプログラムに場合には必ず存在する

2. 標準パッケージ
   - ビルドイン
   - 100以上のパッケージが存在する

3. サードパーティパッケージ
   - 第3者が開発したパッケージ
   - ネットで公開される事が多い
   - インストールすると使える
   - **ライブラリとも呼ばれる**

#### 1-4-11. パッケージドキュメントを読もう
公式で用意されている標準ライブラリのドキュメント
https://pkg.go.dev/

#### 1-4-12. インポートパスとパッケージ名
```go:
package main

/* "path/filepath" -> インポートパス */
import "path/filepath"

func main() {
  // 中略
  /* filepath -> パッケージ名 */
  filepath.Ext("hoge.txt")
  // 後略
}
```

#### 1-4-13. fmt.Print関数 / fmt.Println関数
- 標準出力に出力を行う関数
- 複数の値を渡せる
  - fmt.Print
    - 出力後、改行を行わない
  - fmt.Println
    - 出力後、改行を行う

```go:
package main

import "fmt"

func main() {
	fmt.Print("Hello, ")             // 改行されない
	fmt.Println("世界")               // 改行される
	fmt.Println("A", 100, true, 1.5) // スペース区切りで表示される
}

/* 実行結果 */
// Hello, 世界
// A 100 true 1.5
```

#### 1-4-14. fmt.Pringf関数
- 書式を指定して出力する
- 改行はされない(`\n`が必要)
- `%d`, `%s`などで値の書式を指定する

```go:
package main

import "fmt"

func main() {
	fmt.Printf("Hello, 世界\n")      // \nで改行する
	fmt.Printf("%d-%s", 100, "偶数") // %dは整数、%sは文字列
}

/* 実行結果 */
// Hello, 世界
// 100-偶数
```

#### 1-4-15. fmt.Scanln関数
- 引数に渡した変数の場所に、入力したデータを入れる関数
※理解しきれてないので後回し

## 1-5. 開発環境の構築
#### 1-5-1. ローカルに環境構築
- 動画を元にインストール
- https://youtu.be/lu3_OqhLCmw

pathが通ってなかったので.zshrcにパスを追加
```shell:
export PATH=$PATH:/usr/local/go/bin
```

```shell:
$ go version
# go version go1.19.1 darwin/arm64
```

#### 1-5-2. GOPATHとGOROOTの設定は？
- 特に理由がなければ不要
  - GO Modulesを使う前提なのでGOPATHは気にしなくて良い
  - GOROOTはgoコマンドが知っている

#### 1-5-3. VSCodeでGoを書いてみる
1. helloディレクトリを作成し、その中で`.mod`ファイルを作成する
```shell
$ go mod init hello
```
`go.mod`が作成される。
**Goの開発に必要なファイルなので、毎回作成する。**

2. Goの拡張機能の追加
```m:
Name: Go
Id: golang.go
Description: Rich Go language support for Visual Studio Code
Version: 0.35.2
Publisher: Go Team at Google
VS Marketplace Link: https://marketplace.visualstudio.com/items?itemName=golang.Go
```

3. Goに必要な機能の追加
自動的に聞いてくるので`install ALL`で追加

4. main.go
helloディレクトリ内に作成
```go: main.go
package main

func main() {
  println("Hello, World")
}
```

5. 実行
```shell:
go run .

/* 実行結果 */
// Hello, World
```

## 1-6. 開発支援ツール
#### 1-6-1. コードの書式を揃える
- gofmt
  - 読み方: ごーふむと
  - 絶対に使う！
  - `-s`オプションで冗長な書き方をシンプルにできる

- goimports
  - import宣言と追加/削除してくれる
  - 未使用パッケージのimportはコンパイルエラーなので必須
  - フォーマットもかけてくれる
  - `-s`オプション無し

#### 1-6-2. コードの品質を保つ
- 静的解析ツール各種
  - go vet
    - コンパイラでは発見できないバグを見つける
    - go testを走らせれば自動で実行される(Go1.10~)

#### 1-6-3. PRレビューで静的解析ツールを使用
- reviewdog
  - https://github.com/reviewdog/reviewdog
  - レビュー時に自動で静的解析ツールを実行する
  - 設定ファイルを書いておくと勝手にPRにコメントをくれる

#### 1-6-4. リファクタリング
https://docs.google.com/presentation/d/1Z5b5fIA5vqVII7YoIc4IesKuPWNtcU00cWgW08gfdjg/edit#slide=id.g4e29971f9a_0_569

#### 1-6-5. デバッグ
https://docs.google.com/presentation/d/1Z5b5fIA5vqVII7YoIc4IesKuPWNtcU00cWgW08gfdjg/edit#slide=id.g4e29971f9a_0_563



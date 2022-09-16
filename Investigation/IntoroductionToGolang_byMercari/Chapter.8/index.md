# メルカリ作のプログラミング言語Go完全入門 読破
# 8. テストとテスタビリティ
## 8-1. 単体テスト
#### 8-1-1. 単体テスト
- 関数やメソッドレベルでの動作検証を行う
  - 与えた引数に対する戻り値が期待するものか？
  - 期待するエラーが発生するか？

- テストコード
  - 単体テストを行うコード
  - 検証したい機能を実際に触ってみて期待する挙動になるか検証
  - 引数を網羅的に検証したりできる
  - 機械的に実行することができるので、一度書けば何度でも簡単に検証可能
- [テストコードのルール](https://qiita.com/saya713y/items/f7ee07e8f12ab85ed9bf#-%E3%83%86%E3%82%B9%E3%83%88%E3%82%B3%E3%83%BC%E3%83%89%E3%81%AE%E3%83%AB%E3%83%BC%E3%83%AB)
1. パッケージ: `testingパッケージ`をインポートする
2. ファイル名: `xxx_test.go`
3. テスト関数名: `func TestXxx (t *testing.T)`

- テストの必要性(何のために？)
  - ソフトウェアの品質の担保
  - 変更に対する品質の保証
  - ドキュメントの役割(パッケージの最初のユーザ)

#### 8-1-2. go test
`go test` -> 単体テストを行うためのコマンド
末尾に`_test.go`が付いたファイルを対象にテストを実行する
前節で参考にさせて頂いた記事を元に実際に書いてみます。
```go: hello.go
package hello

func GetHello(s string) string {
  return "こんにちは、" + s + "！"
}
```

- 上のhello.goをインポートして使用するプログラム
```go: main.go
package main

import (
  "fmt"
  "./hello"
)

func main() {
  s := hello.GetHello("山澤さん")
  fmt.Println(s)
}

/* 実行結果 */
// こんにちは、山澤さん！
```

- 上のサンプルコードに対するテストコード
```go: hello_test.go
package hello

import (
  "testing"
)

func TestHello(t *testing.T) {
  result := GetHello("山澤さん")
  expect := "こんにちは、山澤さん！！"
  if result != expect {
    t.Error("\n実際： ", result, "\n理想： ", expect)
  }

  t.Log("TestHello終了")
}
```
1. `testing`をインポートしている
2. テストコードのファイル名を`hello_test.go`にする
3. `TestHello(t *testing.T)`という関数名
ルールに則っている。

#### 8-1-3. testingパッケージ
テストを行うための機能を提供するパッケージ
- testingパッケージで出来ること
  - 失敗理由を出力してテストを失敗させること
    - テスト関数を継続: t.Error, t.Errorf
    - テスト関数を終了. t.Fatal, t.Fatalf
  - テストの並列実行
    - t.Parallel(テスト関数の先頭で呼び出す)
    - go testの-parallelオプションで**並列数を指定**
  - ベンチマーク
    - *testing.B型を使う
  - ブラックボックステスト
    - **あまり積極的には使わない**
    - testing/quickパッケージ

#### 8-1-4. テストの後処理
- t.Cleanup
Go1.14でtestingパッケージに追加
テスト終了時に行う関数を登録
deferの置き換え異常に便利らしい。
参考: [Go の t.Cleanup がとてもべんり](https://syfm.hatenablog.com/entry/2020/05/17/161842)

- t.TempDir
(ちょっとあやしいですが)Go1.15でtestingパッケージに追加
参考: [一時的なファイル出力を伴うテストにtesting.TB.TempDir()が便利そう](https://pod.hatenablog.com/entry/2021/01/31/201137)
一時的にファイルを作成 -> テスト終了時にはきれいに消してくれる
ファイルと書いてありますが、ディレクトリかな？

#### 8-1-5. テスティングフレームを使わない理由
= testingパッケージを使わない理由として、進める
参考: [Go 言語のテスト・フレームワーク](https://text.baldanders.info/golang/testing/)

1. アサーションがない(`t, ok := i.(T)`このような真偽値チェック)
自動でエラーメッセージを作るのではなく、ひとつずつErrorfを使って自前で作る必要がある。

2. テストはGoで書く
> 一般的なテストフレームワークにおいて条件・制御・出力機構を持つ専用のミニ言語が用意される傾向がありますが、Go言語にはすでにこれらが備わっています。これらを再び作成するより、我々はGo言語のテストを進めたかったのです。このようにしたことで余計な言語を覚える必要がなくなり、テストを直接的かつ理解しやすくしています。

3. 比較演算子だけでは辛いのでは？
[go-cmp](https://github.com/google/go-cmp)
より厳格にチェックできる。

4. テスティングフレームは使わない方が良い？
一概にそういうわけではない。
- プロジェクトごとに決める
- ライブラリの導入に伴う複雑度の増加も考慮する
- [matryer/is](https://github.com/matryer/is)
  - 軽量なテスト用ミニフレームワーク

#### 8-1-6. Exampleテスト
参考: [Goのtestingを理解する in 2018 - Examples編 #go - My External Storage](https://budougumi0617.github.io/2018/08/30/go-testing2018-examples/#testable-examples)
Exampleテストがどういったものかは分かった。

#### 8-1-7. テストの並列実行
参考: [Go言語でのテストの並列化 〜t.Parallel()メソッドを理解する〜 | メルカリエンジニアリング](https://engineering.mercari.com/blog/entry/how_to_use_t_parallel/)
並列実行することでテスト時間を短縮できる。

#### 8-1-8. サブテスト
参考: [Go1.7~ subTestsとTableDrivenTest - Qiita](https://qiita.com/marnie_ms4/items/d5233045a084cebeea14)
テストの階層化

#### 8-1-9. テーブル駆動テスト
参考: [テーブル駆動テスト]https://zenn.dev/kimuson13/articles/go_table_driven_test#:~:text=%E3%81%A6%E3%81%84%E3%81%8D%E3%81%BE%E3%81%99%E3%80%82-,%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB%E9%A7%86%E5%8B%95%E3%83%86%E3%82%B9%E3%83%88,-%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB%E9%A7%86%E5%8B%95%E3%83%86%E3%82%B9%E3%83%88

#### 8-1-10. サブテストとテーブル駆動テスト
参考: [Go言語でのテストの並列化 〜t.Parallel()メソッドを理解する〜 | メルカリエンジニアリング](https://engineering.mercari.com/blog/entry/how_to_use_t_parallel/)

- 利点
  - 落ちた箇所がわかりやすい
    - テストケースの名前が表示される
  - 特定のサブテストだけ実行できる
    - テストケースが大量な場合はわかりやすい

#### 8-1-11. テストヘルパー
テスト用のヘルパー関数
参考: [【Go】testingパッケージのHelper関数とは？｜Fresopiya](https://fresopiya.com/2022/07/29/go-helper/)
この関数を記述するとErrorを返す場合にファイルと行情報の表示スキップする。
**共通のアサーションを定義した場合、どのファイルのどの行でエラーが発生したかをわかりやすくするためにHelper関数を使用する**

#### 8-1-12. テストケースを簡潔に書く
- 簡潔に書く事で、テストケースごとの差分が分かる
  - そのテストケースで何を試したいのか簡潔に分かる
  - テストケースが縦に伸びると差分が分かりづらい
    - 横に伸びす(かつ簡潔に)
  - 名前の長い定数や変数があれば再定義する
  - 型エイリアスをうまく使う
  - デフォルト値を用意する
  - 構造体リテラルを減らし、ヘルパー関数を活用する
    - ヘルパー関数はエラーを返さない
  - 可変長引数を上手く使った関数を作る
  - []interface をうまく使う

#### 8-1-13. coverprofile
テストのカバレッジを分析
カバレッジ -> 所定の網羅条件がテストによってどれだけ実行されたかを割合で表したもの
参考: [Goでテストカバレッジを測定する - Qiita](https://qiita.com/takehanKosuke/items/4342ca544d205fb36eb0)

## 8-2. テスタビリティ
#### 8-2-1. テスタブルなコードと抽象化
テスタブル = テストしやすいコード
- 個々の機能が疎結合で単体でテストしやすい
  - 分離が重要
- interfaceの利点を使う
  - 抽象化

#### 8-2-2. インタフェースを使う
外部につながる部分はモックに差し替え可能にする
[Mock object](https://en.wikipedia.org/wiki/Mock_object)
```go:
type DB interface {
   Get(key string) string
   Set(key, value string) error
}
// DBはインタフェースなので実装を入れ替えれる
type Server struct { DB DB }
```
なんとなく分かる。

#### 8-2-3. テストする部分だけ実装する
埋め込みを使って一部分だけモックを用意する
**呼び出さないメソッドは実装しなくてもコンパイルエラーにならない**
```go:
type getTestDB struct {
   // DBを埋め込むことで実装したことになる
   DB
}
// Getだけテスト用に実装する
func (db getTestDB) Get(key string) string {...}
```

#### 8-2-4. 環境変数を使う
環境変数(os.Getenvで取得できる)を使って切り替える
[github.com/jinzhu/configor](https://github.com/jinzhu/configor)

#### 8-2-5. テストデータを用意する
**どの環境でも使用できるテストデータを用意する**
- テストデータは`testdata`というディレクトリに入れる。
`testdata`はパッケージとみなされない。

- テストの再現性を担保する(テストデータに求められる事)
  - ネットワークアクセスを発生させない
  - テストデータ以外のファイルにアクセスしない

#### 8-2-6. 非公開な機能を使ったテスト


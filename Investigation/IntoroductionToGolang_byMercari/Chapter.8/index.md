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


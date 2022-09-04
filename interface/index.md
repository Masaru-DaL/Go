# Go interfaceの存在意義について考える
こちらのサイトが参考になった
[今さら訊けない？ Go言語のInterfaceについてその本質をズバリ解説。](https://thiscalifornianlife.com/2021/01/10/golang-interface/)

## 1. interfaceを使用しないと何が起こる？
#### 1-1. 2つの構造体を定義する
monkey, humanという2つの構造体を定義する。
持っている要素は`Name`だけの簡単なものにする。

```go: not interface
package main

import "fmt"

// struct(monkey)
type monkey struct {Name string}

func (m monkey) imitate(){
  fmt.Println("Imitating")
}
func (m monkey) makeADream(){
  fmt.Println("Make a Dream!")
}

// struct(human)
type human struct {Name string}

func (h human) talk(){
  fmt.Println("Talking")
}
func (h human) makeADream(){
  fmt.Println("Make a Dream!")
}
```

- monkey
  - imitate, makeADream の2つのメソッド

- human
  - talk, makeADream の2つのメソッド


```go: not interface
func main() {
  taro := monkey{Name: "TARO"}
  jack := human{Name: "JACK"}

  taro.imitate()
  jack.talk()
}

// 出力結果
Imitating
Talking
```
ここまでは大丈夫。

#### 1-2. 新しい関数の追加があった場合
月へ行くという関数が追加されたとする。

A君:
月へ行くのは猿では無理だよなぁ。
humanが行けるようにしよ。

```go:
func main() {
  taro := monkey{Name: "TARO"}
  jack := human{Name: "JACK"}

  taro.imitate()
  jack.talk()

  // humanを引数にして関数を実行
  goToTheMoon(jack)
}

// 月へ行くという関数の定義
func goToTheMoon(h human){
  fmt.Println("Go to the moon")
}


// 出力結果
Imitating
Talking
Go to the moon
```

A君: 出来ました！
上司: いやいや、猿も行けるようにしてよ
A君: !!??

#### 1-3. 2つの構造体に関数を対応させることに
```go:
func main() {
  taro := monkey{Name: "TARO"}
  jack := human{Name: "JACK"}

  taro.imitate()
  jack.talk()

  // humanを引数にして関数を実行
  goToTheMoon1(jack)
  goToTheMoon2(taro)
}

// 月へ行くという関数の定義
// human用
func goToTheMoon1(h human){
  fmt.Println("Go to the moon")
}

// monkey用
func goToTheMoon2(m monkey){
  fmt.Println("Go to the moon")
}

// 出力結果
Imitating
Talking
Go to the moon
Go to the moon
```
(名前の衝突を避けるために関数名も若干変えています)
...こうなってくると似たようなコードが増えて、よろしくない感じがヒシヒシとします。
構造体が増える度に対応させるコードを増やし続けなければいけないということになります。

## 2. ここで使えるのがInterfaceのポリモーフィズム
ポリモーフィズムは親クラスから複数のインスタンスを生成する際に、振る舞いだけ変えたい時に使用する、オブジェクト指向概念の1つです。

#### 2-1. 宇宙へ行く人(astronaut)として新しく定義する
`interface`の使い方は、`型名`、`メソッド`を指定し、
関数の引数に`型名`を指定すると、

- interfaceを定義してみる
```go: interface
// 1.
type astronaut interface {
  makeADream()
}

// 2.
func goToTheMoon(a astronaut){
  fmt.Println("Go to the moon")
}

func main(){
  taro := monkey{Name: "TARO"}
  jack := human{Name: "JACK"}

  // 3.
  goToTheMoon(jack)
  goToTheMoon(taro)
}
```
1. `astronaut`という型名に`makeADream()`というメソッドを定義しています。
ここで重要なのは、一番最初にhuman, monkeyの構造体を定義する際に、**両方に`makeADream`を定義している**ことです。

2. 関数`goToTheMoon`の引数に`astronaut`を指定しています。

3. 変数`jack`, `taro`のどちらも`goToTheMoon`の引数になることが出来ている。
`jack`, `taro`は、human, monkeyタイプだけでなく、astronautタイプも持っている為です。


# interfaceのまとめ
1. interfaceを使用するとコードがすっきり、簡潔に実装できる
interfaceを使用しなかった場合を見れば一目瞭然だが、
構造体からインスタンスを作成する毎に関数を対応させなくてはいけなくなる。

2. 共通の処理をまとめることができる
1と同じような意味ではあるが、`makeADream`のように、共通メソッドを持つ場合の処理を共通化することが出来る。

3. 自動的にインターフェースが実装される点
**インターフェースの中にあるメソッドと同一のメソッドが全て実装されている構造体には、自動的にインターフェースが実装される**
つまり、interfaceとして機能させようと思った場合、interfaceのメソッドを定義後、そのメソッドを構造体に定義する必要があるということが言える。

存在意義に関しては他にもありそうではあるが、共通処理を行う際に便利であることが一番に言えると思われる。

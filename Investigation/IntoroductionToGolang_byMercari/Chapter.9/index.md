# メルカリ作のプログラミング言語Go完全入門 読破
# 9. ゴールーチンとチャネル
## 9-1. 並行処理
#### 9-1-1. Concurrency is not Parallelism
**並行と並列は別ものである** by RobPike
[引用: 並列と並行の違い image](https://cdn-ak.f.st-hatena.com/images/fotolife/y/yoshitachi/20181112/20181112114605.png)

- 並行 -> Concurrency(コンカレンシー)
  - 同時にいくつかの**質の異なる**ことを扱う

- 並列 -> Parallelism(パラレリズム)
  - 同時にいくつかの**質の同じこと**を扱う

#### 9-1-2. ゴールーチンとConcurrency
ゴールーチンは、**Concurrency(並行処理)を実現できる**
- 1つ1つ(複数)のゴールーチンは、**同時に複数のタスク**をこなす
- 各々のゴールーチンに**役割を与えて分業**する

軽量なスレッド(一連のプログラムの流れ)のようなもの
- プログラムの1から10の間に、複数のゴールーチンが動く

`go`キーワードを付けて関数を呼び出すとゴールーチンが作成される。
`go f()`

#### 9-1-3. TRY ゴールーチンを使ってみよう
`defer fmt.Println("main done")`
`time.Sleep(5 * time.Second)`
これは何秒遅延させるかということだと思われる。
`time.Sleep`の時間を変更すると"main done"がその時間後に出力される。なお、"main done"が出力された段階で他のゴールーチンは出力されない。

1. `time.Sleep(5 * time.Second)`
```go: time.Sleep(5* time.Second)
package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main done")
	go func() {
		defer fmt.Println("goroutine1 done")
		time.Sleep(3 * time.Second)
	}()

	go func() {
		defer fmt.Println("goroutine2 done")
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(5 * time.Second)
}

/* 実行結果 */
goroutine2 done // 1秒補
goroutine1 done // 上の出力から2秒後(全体で3秒後)
main done // 上の出力から2秒後(全体で5秒後)
```

2. `time.Sleep(3 * time.Second)`
```go: time.Sleep(3* time.Second)
package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main done")
	go func() {
		defer fmt.Println("goroutine1 done")
		time.Sleep(3 * time.Second)
	}()

	go func() {
		defer fmt.Println("goroutine2 done")
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(3 * time.Second)
}

goroutine2 done // 1秒補
goroutine1 done // 上の出力から2秒後(全体で3秒後)
main done // 上の出力と同時(全体で3秒後)
```

3. `time.Sleep(2 * time.Second)`
```go: time.Sleep(2* time.Second)
package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main done")
	go func() {
		defer fmt.Println("goroutine1 done")
		time.Sleep(3 * time.Second)
	}()

	go func() {
		defer fmt.Println("goroutine2 done")
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(2 * time.Second)
}

goroutine2 done // 1秒補
main done // 上の出力の1秒後(全体で2秒後)
/* goroutine1 doneは出力されなかった */
```

## 9-2. チャネルでデータ競合を避ける
#### 9-2-1. ゴールーチン間で共有の変数を使う
```go:
package main

import (
	"fmt"
	"time"
)

func main() {
	var v int
	go func() { // ゴールーチン-1
		time.Sleep(1 * time.Second)
		v = 100
	}()

  /* ゴールーチン-1のv=100が出力される */
	go func() { // ゴールーチン-2
		time.Sleep(1 * time.Second)
		fmt.Println(v)
	}()
	time.Sleep(2 * time.Second)
}

/* 実行結果 */
// 100
```
※もしゴールーチン-2が1秒後、ゴールーチン-1が2秒後に終わるという設定の場合、v=0となって出力された。**共有の変数を使用する場合は、変数値を持っているゴールーチンが先に処理を終わってる必要もありそうだ**

#### 9-2-2. ゴールーチン間のデータ競合とその解決
```go:
package main

import (
	"fmt"
	"time"
)

func main() {
	n := 1
	go func() {
		for i := 2; i <= 5; i++ {
			fmt.Println(n, "*", i)
			n *= i
			time.Sleep(100)

		}
	}()
	for i := 1; i <= 10; i++ {
		fmt.Println(n, "+", i)
		n += 1
		time.Sleep(100)
	}
}

/* 実行結果 */
/*
1 + 1
1 * 2
4 + 2
4 * 3
15 * 4
15 + 3
61 + 4
61 * 5
310 + 5
311 + 6
312 + 7
313 + 8
314 + 9
315 + 10
*/
```
本来やりたいのは並行処理だから、交互にnの値を変更したいが、どちらが先にアクセスするか分からないという状態や、値の変更や参照が競合を生んでいる。

- 解決方法
  - 1つの変数には1つのゴールーチンからアクセスする
  - **チャネル**を使ってゴールーチン間で通信をする
  - または**ロックを取る**(syncパッケージ)

> Do not communicate by sharing memory;
> メモリの共有による通信を行わない。
> instead, share memory by communicating
> メモリを共有するのではなく、通信することでメモリを共有する

#### 9-2-3. チャネル
[引用: Channels image](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F154845%2F61dd65b2-6c97-9dae-23a5-d3eefdd2a0f9.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&w=1400&fit=max&s=8f2237df50fc0c94ba20d83c040ddcf4)

- 送受信できる型
  - チャネルを定義する際に**型を指定**する
    - `make(chan <型>)`

- バッファ
  - チャネルにバッファ(1時保管)を持たせることができる
  - 初期化時に指定できる
    - 例: `make(chan string, 2)`
    - 2つまでバッファリングできるということ
  - 指定しないと**容量0**となる
    - 可変ではないということ。
    - 事前に容量指定が必要

- 送受信時の処理のブロック
  - **送信時**にチャネルのバッファがいっぱいだと**ブロック**
  - **受信時**にチャネル内が空だと**ブロック**
> Goはデフォルトで、送る側と受ける側が準備できるまで送受信がブロックされる。

#### 9-2-4. チャネルの基本 -1
- 初期化
`ch1 := make(chan int)` -> int型, 容量0
`ch2 := make(chan int, 10)` -> 容量を10に指定

- 送信
`ch1 <- 10` -> 受け取られるまでブロック
`ch2 <- 10 + 20` -> `20`はch2の容量がいっぱいだったらブロックされる

- 受信
`n1 := <-ch1` -> 送信されるまでブロックされる
`n2 := <-ch2 + 100` -> 空であればブロック

#### 9-2-5. チャネルの基本 -2
```go:
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // 容量0

	// goroutine-1
	go func() {
		ch <- 100
	}()

	// goroutine-2
	go func() {
		v := <-ch
		fmt.Println(v)
	}()

	time.Sleep(2 * time.Second)
}

/* 実行結果 */
// 100
```
勘違いしていました。容量0というのをチャネルに値を入れられないと思っていました。
普通にチャネルを通して100という値が`v`に入り出力されています。
容量とブロックの関係を整理する必要があります。

ゴールーチンを使わずに検証してみた。
何となくイメージは掴めた。
```go:
package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 2) // 容量を2に指定

	messages <- "Hello"
	messages <- "World"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/* 実行結果 */
// Hello
// World
```

```go:
package main

import (
	"fmt"
)

func main() {
	messages := make(chan string) // 容量が0

	messages <- "Hello"
	messages <- "World"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
/* 実行結果 */
/*
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/tmp/sandbox1713984673/prog.go:10 +0x37
*/
```


#### 9-2-6. 複数のチャネルから同時に受信
select-caseを用いる
```go:
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "hi"
	}()

	select {
	case v2 := <-ch1:
		fmt.Println(v2)
	case v1 := <-ch2:
		fmt.Println(v1)
	}
}

/*
select-caseの検証(同時終了、time.Sleepの遅延)
1. 先に処理が終わった方が出力される
2. 同時に処理が終わった場合、caseで先に合致した時点で終了
*/
```

- nilチャネル
```go:
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	// var ch2 chan string
	ch2 := make(chan string)

	go func() { ch1 <- 100 }()
	go func() { ch2 <- "hi" }()

	select {
	case v2 := <-ch2:
		fmt.Println(v2)
	case v1 := <-ch1:
		fmt.Println(v1)
	}
}

/*
makeでチャネルを作成したものと、変数での初期値が違う。
変数はゼロ値がnilのため、`ch2 <- "hi"`で値を入れて先に処理をしようとしても**無視された**。
*/
```

#### 9-2-7. ファーストクラスオブジェクト
[第一級オブジェクト](https://ja.wikipedia.org/wiki/%E7%AC%AC%E4%B8%80%E7%B4%9A%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88)
- チャネルは...
  - 変数へ代入可能
  - 引数に渡す
  - 戻り値で返す
  - チャネルのチャネル -> `chan chan int`など

- timeパッケージ
[func After](http://golang.org/pkg/time/#After)
> Afterは持続時間の経過を待ち、返されたチャネルで現在の時刻を送信する。

```go:
/* 5分経ったら現在時刻が送られてくる"チャネルを返す" */
<-time.After(5 * time.Minute) // 5分待つ
```

#### 9-2-8. チャネルを引数や戻り値にする
```go:
package main

import "fmt"

/* 戻り値でチャネルを返す(int型, 容量0) */
func makeCh() chan int {
	return make(chan int)
}

/* 引数の型 = チャネル, 引数にチャネルを入れたらそのチャネルが戻り値になる */
func recvCh(recv chan int) int {
	return <-recv
}

func main() {
	ch := makeCh() 			// ch = チャネル
	go func() { ch <- 100 }()	// chに100を格納
	/* 引数にチャネルを指定 */
	fmt.Println(recvCh(ch))
}

/* 実行結果 */
// 100
```

#### 9-2-9. 双方向チャネル
```go:
package main

import "fmt"

func makeCh() chan int {
	return make(chan int)
}

func recvCh(recv chan int) int {
	go func() { recv <- 200 }()
	return <-recv
}

func main() {
	ch := makeCh()
	go func() { ch <- 100 }()
	fmt.Println(recvCh(ch))

/* 実行結果 */
// 200

/* main関数の中でチャネルに値を入れているが、recvCh関数内で戻り値に200の固定値を入れてしまっている。この値が出力されている。
間違った使い方ができる、ということは意図しない使い方ができてしまっている。
双方向チャネルという点からすると、
正: 100 -> ch, ch -> recv, return recv
誤: 100 -> ch, ch -> recv, 200 -> recv, return recv
*/
```


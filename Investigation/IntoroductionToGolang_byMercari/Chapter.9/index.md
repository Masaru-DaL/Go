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

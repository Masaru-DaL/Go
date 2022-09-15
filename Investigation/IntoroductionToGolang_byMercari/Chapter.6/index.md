# メルカリ作のプログラミング言語Go完全入門 読破
# 6. 抽象化
## 6-1. インタフェース
#### 6-1-1. インタフェースと抽象化
- 抽象化
  - 具体的な実装を隠し、振る舞いによって共通化させること
  - 複数の実装を同室のものとして扱う
[引用: インタフェースと抽象化 image](https://cdn-ak.f.st-hatena.com/images/fotolife/y/y-zumi/20190728/20190728023343.jpg)
**Goではインタフェースでしか抽象化をすることができない**

#### 6-1-2. インタフェース
```go:
package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	type Stringer1 interface {
		String() string
	}

	type Stringer2 interface {
		String() string
	}

	var s Stringer1 = Hex(100)
	var m Stringer2 = Hex(1000)
	fmt.Println(s.String())
	fmt.Println(m.String())

}

/* 実行結果 */
// 64
// 3e8
```

#### 6-1-3. interface{}


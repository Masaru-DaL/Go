- [REST API FrameWork(echo)](https://www.notion.so/REST-API-FW-88ac694878ac4382a43e2f415660a304)
- [1. 参考資料](https://www.notion.so/REST-API-FW-88ac694878ac4382a43e2f415660a304)
    - [2. プロジェクトの作成からサーバ起動まで](https://www.notion.so/REST-API-FW-88ac694878ac4382a43e2f415660a304)
    - [3. ルーティングの登録](https://www.notion.so/REST-API-FW-88ac694878ac4382a43e2f415660a304)
    - [4. クエリパラメーターを取得](https://www.notion.so/REST-API-FW-88ac694878ac4382a43e2f415660a304)
    - [5. フォームの送信](https://www.notion.so/REST-API-FW-88ac694878ac4382a43e2f415660a304)
    - [6. 複数の異なるデータの送信](https://www.notion.so/REST-API-FW-88ac694878ac4382a43e2f415660a304)
    - [7. 構造体へのタグ付け](https://www.notion.so/REST-API-FW-88ac694878ac4382a43e2f415660a304)
    - [8. 静的コンテンツ](https://www.notion.so/REST-API-FW-88ac694878ac4382a43e2f415660a304)
    - [9. Template Rendering](https://www.notion.so/REST-API-FW-88ac694878ac4382a43e2f415660a304)

# REST API FrameWork(echo)

### 1. 参考資料

[echo](https://echo.labstack.com/guide/)[Go/echo入門](https://qiita.com/pylor1n/items/36912a47c893ea5782cc)

## 2. プロジェクトの作成からサーバ起動まで

1. `go mod init <プロジェクト名>`
2. バージョン確認 -> `go version` go.1.19.1
3. 1.14以上 -> `go get github.com/labstack/echo/v4`
4. ファイルの作成 -> `server.go`[Context](https://echo.labstack.com/guide/context/#:~:text=guide%20context-,Context,-Extending%20Context)
Contextを介してリクエストを送るということでいいかと思う。

```go:
package main
import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() // Echoのインスタンスの作成

	/* GETリクエスト,  */
	e.GET("/", func(c echo.Context) error {
		/* c.String -> ステータスコードとともに文字列をレスポンスに書き込む */
		return c.String(http.StatusOK, "Hello, World!")
	})
	/* e.Start -> HTTPサーバを開始する(指定したポートへのリクエストの待ち受け開始を行う) */
	// ログ出力を同時に行う
	e.Logger.Fatal(e.Start(":1323"))
}

/*
   ____    __
  / __/___/ /  ___
 / _// __/ _ \\/ _ \\
/___/\\__/_//_/\\___/ v4.9.0
High performance, minimalist Go web framework
<https://echo.labstack.com>
____________________________________O/_______
                                    O\\
⇨ http server started on [::]:1323
*/

```

1. `http://localhost:1323`へアクセス
"Hello, World!"と表示されている。

## 3. ルーティングの登録

1. Routingの必要知識

[Routing](https://echo.labstack.com/guide/routing/#:~:text=guide%20routing-,Routing,-Match%2Dany)

> ルートは、HTTPメソッド、パス、マッチングハンドラを指定することで登録することができます。例えば、以下のコードでは、GETメソッド、パス/hello、そしてHello, World!HTTPレスポンスを送信するハンドラを登録します。
>

```go:
// Handler
func hello(c echo.Context) error {
  	return c.String(http.StatusOK, "Hello, World!")
}

// Route
e.GET("/hello", hello)

```

- Route
    - GETメソッド, パス(/hello), hello(ハンドラ名)
    - 定義済みハンドラを登録する
- Handler
    - 戻り値はerror型で適宜エラーハンドリングを行う
1. Routingの登録

```go:
// e.POST("/users", saveUser)
e.GET("/users/:id", getUser)
// e.PUT("/users/:id", updateUser)
// e.DELETE("/users/:id", deleteUser)

```

※公式手順だと、この時点でのハンドラ定義はGETだけなので、一旦他のルーティングはコメントアウトする。

1. ハンドラ定義(GETメソッド)

```go:
// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
    // User ID from path `users/:id`
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}

```

- 前提知識: プレースホルダ
    - あらかじめ値が入る場所を確保しておき、どこに値を代入する。
    - 値の入る箱を置いておいて、そこに値だけ入れるイメージ
    - SQLインジェクションに対する対策にもなる
- `c.Param("id")`
    - Routingの`/users/:id`に入る値が、変数"id"に入る。

具体的には、以下の流れです。
`http://localhost:1323/users/Joe`にGETメソッドでリクエストを投げる。
`id := c.Param("id")`の変数`id`に"Joe"(文字列)が入る。

1. この時点での`server.go`

```go:
package main
import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/* 2. e.GET("/users/:id", getUser) */
func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func main() {
	e := echo.New() // Echoのインスタンスの作成

	/* Routing */
	// e.POST("/users", saveUser)					// 1.
	e.GET("/users/:id", getUser)				// 2.
	// e.PUT("/users/:id", updateUser)			// 3.
	// e.DELETE("/users/:id", deleteUser)	// 4.

	/* GETリクエスト,  */
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

```

1. `http://localhost:1323/users/Joe`へアクセス

Joeと表示される。

## 4. クエリパラメーターを取得

GETメソッドでリクエストを投げ、クエリを取得する

1. クエリパラメータを取得するハンドラの登録

```go:
// e.GET("/show", show)
func show(c echo.Context) error {
	// teamとmemberというクエリ文字列を取得する
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

```

1. ルーティングの追加

`e.GET("/show", show)`

1. この時点での`server.go`

```go:
package main
import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/* e.GET("/users/:id", getUser) */
func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/show", show)
func show(c echo.Context) error {
	// teamとmemberというクエリ文字列を取得する
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

func main() {
	e := echo.New() // Echoのインスタンスの作成

	/* Routing */
	e.GET("/users/:id", getUser)
	e.GET("/show", show)

	// e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	/* GETリクエスト,  */
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

```

1. `http://localhost:1323/show?team=x-men&member=wolverine`へアクセス

再度リクエストを投げるので、サーバを再起動してからアクセスします。
team:x-men, member:wolverineと表示される。
URLにアクセスした際のクエリ文字列を取得している。

## 5. フォームの送信

1. フォームを取得するハンドラーの登録

```go:
// e.POST("/save", save)

func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

```

- `Context.FormValue(name string)`でフォームデータを名前で取得することができる。
1. ルーティングの登録

フォームはPOSTメソッドで登録します。
`e.POST("/save", save)`

1. この時点での`server.go`

```go:
package main
import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/* e.GET("/users/:id", getUser) */
func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/show", show)
func show(c echo.Context) error {
	// teamとmemberというクエリ文字列を取得する
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func main() {
	e := echo.New() // Echoのインスタンスの作成

	/* Routing */
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)

	// e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	/* GETリクエスト,  */
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

```

1. `curl`でフォームデータを送信する

サーバが起動している状態でフォームデータを送信します。
`$ curl -d "name=Joe Smith" -d "email=joe@labstack.com" <http://localhost:1323/save`>

POSTメソッドでは`-d`オプションの後に送信データを記述する。
name:Joe Smith, email:joe@labstack.comと返ってくる。
送信したのに返ってくるのはハンドラの中でリターンで指定文字列を返す処理を行なっているからです。

## 6. 複数の異なるデータの送信

1. 先ほどのハンドラ(save)を変更(削除orコメントアウトなど)します。
2. 新しいsaveハンドラを登録

```go:
func save(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
  	avatar, err := c.FormFile("avatar")
  	if err != nil {
 		return err
 	}

 	// Source
 	src, err := avatar.Open()
 	if err != nil {
 		return err
 	}
 	defer src.Close()

 	// Destination
 	dst, err := os.Create(avatar.Filename)
 	if err != nil {
 		return err
 	}
 	defer dst.Close()

 	// Copy
 	if _, err = io.Copy(dst, src); err != nil {
  		return err
  	}

	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}

```

1. ローカルのプロジェクトに"avatar.png"を用意する

これは何でも良いです。
フリー画像なりなんなり、用意します。
プロジェクトのディレクトリ構造は以下の通りです。

```shell:
-> tree
.
├── avatar.png
├── go.mod
├── go.sum
└── server.go

0 directories, 4 files

```

1. `$ curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" <http://localhost:1323/save`でアクセス>

`avatar=@/path/to/your/avatar.png`
これは、上のファイル構造で用意出来たら、`avatar=avatar.png`に変えてコマンドを打ちます。

<b>Thank you! Joe Smith</b>と標準出力で返ってきます。

## 7. 構造体へのタグ付け

1. コード

```go:
type User struct {
    Name  string `json:"name" xml:"name" form:"name" query:"name"`
    Email string `json:"email" xml:"email" form:"email" query:"email"`
}

e.POST("/users", func(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    return c.JSON(http.StatusCreated, u)
    // or
    // return c.XML(http.StatusCreated, u)
})

```

1. どういう処理か

いまいち理解しきれてない。
構造体に対してタグ付けを行う。
リクエストする際にそのタグ(キー)に対応した値(バリュー)を指定してレスポンスで受け取ることが出来るということだと思う。

## 8. 静的コンテンツ

`e.Static("/static", "static")`
/static/* 以下のファイルへのリクエストはstatic*以下に保存してあるファイルを提供する。

> 例えば、/static/js/main.js へのリクエストは assets/js/main.js ファイルをフェッチして提供します。
>

## 9. Template Rendering

`Context#Render(code int, name string, data interface{} error)`
テンプレートにデータをレンダリング(情報を整形して表示)し、ステータスコード付きのtext/htmlレスポンスを送信する。

- ミドルウェアとは
    - 今回公式で出てくるのは2つ
        - "Logger", "Recover"
    - ミドルウェアは関数。
    - HTTPのリクエストとレスポンスのサイクルの中で連鎖する
    - Echo#Contextにアクセスし、特定の動作を行うために使用する
        - すべてのリクエストを記録する
        - リクエストの数を制限する
    - **ハンドラは、全てのミドルウェアの実行が終了した後に最後に処理される。**
    - **Echo#Use()を用いて登録されたミドルウェアは、Echo#Use()が呼ばれた後に登録されたパスに対してのみ実行される。**
- レベル(level)とは
    - ルートをレベルで分ける
        - ルータがリクエストを処理する前
        - ルータがリクエストを処理した後

以上の前提知識を踏まえて、コードを区切りながらやってみます。

1. ミドルウェアの登録

```go:
// Root level middleware
e.Use(middleware.Logger())
e.Use(middleware.Recover())

```

- Echo#Use()
    - ミドルウェアの登録をする

[参考](https://ken-aio.github.io/post/2019/02/06/golang-echo-middleware/):

- Logger
    - アクセスログのような機能
    - リクエスト単位のログを出力してくれる
    - フォーマットもいじれる
- Recover
    - アプリケーションのどこかで予期しないpanicを起こしてしまったとしても、サーバは落とさずにエラーレスポンスを返せるようにリカバリーするmiddleware
1. 新しいグループを作成し、そのグループ専用のミドルウェアを登録する

なぜ、グループを作成するのかというと、グループ単位で分離し、セキュリティを確保するためです。
例えば、管理者グループには`BasicAuth`ミドルウェアを登録する。

```go:
// Group level middleware
g := e.Group("/admin") // adminグループの作成
g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
  if username == "joe" && password == "secret" {
    return true, nil
  }
  return false, nil
}))

```

`g.Use`でグループを作成。
※グループを作成した後にミドルウェアを追加することもできる。

- BasicAuth
    - HTTPのBasic認証を提供する
    - 認証情報が有効な場合、次のハンドラを呼び出す。
    - 認証情報がない、または無効な場合は「401 - Unauthorized」レスポンスが送信される。
1. ルート専用のミドルウェアを登録する

```go:
// Route level middleware
track := func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("request to /users")
		return next(c)
	}
}
e.GET("/users", func(c echo.Context) error {
	return c.String(http.StatusOK, "/users")
}, track)

// Handler
/*
func(c echo.Context) error {
	return c.String(http.StatusOK, "/users")
}
*/

// Middleware -> track
/*
func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("request to /users")
		return next(c)
	}
}
*/

```

ルート専用のミドルウェアの登録は、新しいルートを定義する際に登録することが出来る。
`e.GET("/", <Handler>, <Middleware...>)`

# GolangでCRUD処理を実装する
* 使用技術
  + Golang
  + MySQL
  + Docker
    - docker-compose
  + Air

## 1. Dockerの環境構築

今回のCRUDを実装するに当たって、自動でDockerの更新を行ってくれるAirを導入します。この章の目的は以下。

1. Dockerの復習
2. Airを導入した際の挙動の確認
3. 今回のCRUDの実装の環境構築

### 1-1. Dockerの復習

* docker-composeでMySQLを起動する
ProjectName(ルートフォルダ名) -> `docker-crud`

#### 1-1-1. フォルダ構成

```shell:
-> tree
.
└── mysql

    ├── .env_mysql          # 環境変数
    ├── Dockerfile
    ├── docker-compose.yml
    └── init
        └── create_table.sh # コンテナ起動後に実行される

```

#### 1-1-2. Dockerfile

```dockerfile:
FROM mysql:8.0
ENV LANG ja_JP.UTF-8
```

ベースイメージ: mysql8.0

#### 1-1-3. docker-compose.yml

```yml: docker-compose.yml
version: "3.8"

services: # アプリケーションを動かす各要素
  db: # サービスとして定義

    container_name: db  # 任意
    build:
      context: .
      dockerfile: Dockerfile
    platform: linux/amd64
    tty: true
    ports:

      - 3306:3306
    env_file:

      - ./.env_mysql
    volumes:

      - type: volume
        source: mysql-data
        target: /var/lib/mysql

      - type: bind
        source: ./init
        target: /docker-entrypoint-initdb.d

volumes:
  mysql-data:

    name: mysql-volume

```

* env_file
  + 指定したファイルの環境変数をコンテナ内で参照可能
* volumes
  + コンテナのデータを永続化
  + Dockerのボリュームとコンテナを紐づける
* bind
  + コンテナファイルとホストのディレクトリをバインドマウントする
  + `.sh`がコンテナ後に自動実行される

#### 1-1-4. .env_mysql

```env:
MYSQL_DATABASE=test_database
MYSQL_USER=test_user
MYSQL_PASSWORD=pass
MYSQL_ROOT_PASSWORD=root
```

環境変数の記述。

#### 1-1-5. create_table.sh

```shell:
#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table article (
  id int(10) AUTO_INCREMENT NOT NULL primary key,
  title varchar(50) NOT NULL,
  body varchar(1000)
  ); "
$CMD_MYSQL -e "insert into article values (1, '記事1', '記事1です。'); "
$CMD_MYSQL -e "insert into article values (2, '記事2', '記事2です。'); "

```

自動的に実行される。
`${環境変数名}` で、env_fileに記述した環境変数を読み込める。
articleという名前のテーブルを作成し、データを2つ挿入する。

#### 1-1-6. 実行

1. initディレクトリ以下のアクセス権限をchmodコマンドで変更する。
 `chmod a+x ./init/*.sh`

2. フォアグラウンドで実行
 `docker compose up`

3. コンテナに入る
 `docker exec -it db bash`

4. mysqlの使用
 `mysql mysql -utest_user -ppass test_database`

5. table確認

```shell:
mysql> show tables;
+-------------------------+
| Tables_in_test_database |
+-------------------------+
| article                 |
+-------------------------+
1 row in set (0.01 sec)

mysql> select * from article;
+----+---------+------------------+
| id | title   | body             |
+----+---------+------------------+
|  1 | 記事1   | 記事1です。      |
|  2 | 記事2   | 記事2です。      |
+----+---------+------------------+
2 rows in set (0.00 sec)
```

### 1-2. Airの導入と挙動の確認

golangファイルを実行して自動更新の挙動を確認する。

1. 通常のビルドの挙動の確認
2. Airの自動更新の挙動の確認

#### 1-2-1. ディレクトリ構成

`/github.com/docker-crud` 以下

1. `go mod init github.com/docker-crud`
2. 他3つファイルを作成

```shell:
-> tree
.
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── main.go

# └── mysql

#     ├── Dockerfile

#     ├── docker-compose.yml

#     └── init

#         └── create_table.sh

```

1-1で使用したmysqlディレクトリ以下はこの節では使わない。

#### 1-2-2. docker-compose.yml

```yml. docker-compose.yml
version: "3.8"

services:
  reload_test:
    image: reload_test
    container_name: reload_test
    build: .
    ports:
      - 8080:8080
    volumes:
      - .:/app
```

volumesで、ローカルのルートディレクトリとコンテナの/appディレクトリをバインドマウントしているので、変更は即時反映される。

#### 1-2-3. Dockerfile

```dockerfile:
FROM golang:1.17.7-alpine

WORKDIR /app
CMD ["go", "run", "main.go"]

```

#### 1-2-4. main.go

```go:
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Normal</h1>")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
```

#### 1-2-5. 実行

1. `$ docker compose up reload_test`
2. `http://localhost:8080`にアクセス
3. Hello Normalが表示される。

#### 1-2-6. main.goの変更・リロードの挙動の確認

```go: main.go
package main

import (

	"fmt"
	"net/http"

)

func helloHandler(w http. ResponseWriter, r *http. Request) {

	fmt.Fprintf(w, "<h1>Hello Normal Update</h1>")

}

func main() {

	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)

}

```

ブラウザをリロードしてもHello Normalのまま。

`$ docker compose restart`
ブラウザを更新すると、Hello Normal Updateと表示された。

#### 1-2-7. Airの導入 Dockerfile

```dockerfile: Dockerfile
FROM golang:1.17.7-alpine

RUN apk update && apk add git
RUN go get github.com/cosmtrek/air@v1.29.0
WORKDIR /app

# air -c <tomlファイル名>
CMD ["air", "-c", ".air.toml"]
```

gitからAirをインストールし、airコマンドの実行。

#### 1-2-8. .air.toml

[Air](https://github.com/cosmtrek/air)
air_example.tomlから。

```toml: .air.toml

# Config file for [Air](https://github.com/cosmtrek/air) in TOML format

# Working directory

# . or absolute path, please note that the directories following must be under root.

root = "."
tmp_dir = "tmp"

[build]

# Just plain old shell command. You could use `make` as well.

cmd = "go build -o ./tmp/main ."

# Binary file yields from `cmd` .

bin = "tmp/main"

# Customize binary, can setup environment variables when run your app.

full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"

# Watch these filename extensions.

include_ext = ["go", "tpl", "tmpl", "html"]

# Ignore these filename extensions or directories.

exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]

# Watch these directories if you specified.

include_dir = []

# Exclude files.

exclude_file = []

# Exclude specific regular expressions.

exclude_regex = ["_test.go"]

# Exclude unchanged files.

exclude_unchanged = true

# Follow symlink for directories

follow_symlink = true

# This log file places in your tmp_dir.

log = "air.log"

# It's not necessary to trigger build each time file changes if it's too frequent.

delay = 1000 # ms

# Stop running old binary when build errors occur.

stop_on_error = true

# Send Interrupt signal before killing process (windows does not support this feature)

send_interrupt = false

# Delay after sending Interrupt signal

kill_delay = 500 # ms

[log]

# Show log time

time = false

[color]

# Customize each part's color. If no color found, use the raw app log.

main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]

# Delete tmp directory on exit

clean_on_exit = true

```

#### 1-2-9. 実行・Airの挙動の確認

一度Containerは削除する。

```shell:
-> docker compose up reload_test
[+] Running 0/1
 ⠿ reload_test Error     3.6s
[+] Building 11.6s (8/8) FINISHED [internal] load bui  0.0s
 => [internal] load bui  0.0s
 => => transferrin 215B  0.0s
 => [internal] load .do  0.0s
 => => transferring  2B  0.0s
 => [internal] load met  2.0s
 => CACHED [1/4] FROM d  0.0s
 => [2/4] RUN apk updat  5.3s
 => [3/4] RUN go get gi  4.1s
 => [4/4] WORKDIR /app   0.0s
 => exporting to image   0.1s
 => => exporting layers  0.1s
 => => writing image sh  0.0s
 => => naming to docker  0.0s
[+] Running 1/0
 ⠿ Network docker-crud_default[+] Running 2/2
 ⠿ Network docker-crud_default  Created 0.0ss
 ⠿ Container reload_test        Created 0.1s
Attaching to reload_test
reload_test  |
reload_test  |   __    _   ___
reload_test  |  / /\  | | | |_)
reload_test  | /_/--\ |_| |_| \_ , built with Go
reload_test  |
reload_test  | mkdir /app/tmp
reload_test  | watching .
reload_test  | watching mysql
reload_test  | watching mysql/init
reload_test  | !exclude tmp
reload_test  | building...
reload_test  | running...
```

 `http://localhost:8080`

Hello Normal Updateと表示されている。

`main.go` をHello Air!!と変更し、ブラウザでリロード。

```shell:
reload_test  | main.go has changed
reload_test  | building...
reload_test  | running...
```

ログから、変更した時点でホットリロードされていることが分かる。
無事にHello Air‼︎と表示されている。

### 1-3. golangとMySQLをDockerで環境構築

今回のCRUD処理を行う環境構築を行う。

#### 1-3-1. ディレクトリ構成

```shell:
go_blog
├── build
│   ├── app
│   │   └── Dockerfile
│   └── db
│       ├── Dockerfile
│       └── init
│           └── create_table.sh
├── cmd
│   └── go_blog
│       └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── article
│   │   └── article.go
│   └── utility
│       └── database.go
└── web
    └── template
        ├── create.html
        ├── delete.html
        ├── edit.html
        ├── index.html
        └── show.html
```

#### 1-3-2. docker-compose.yml

```yml:
version: "3.8"

services:
  golang_crud:
    container_name: golang_crud
    build:
      context: ./build/app
      dockerfile: Dockerfile
    tty: true
    ports:
      - 8080:8080
    env_file:
      - ./build/db/.env
    expends_on: # サービス起動順: db -> go
      - db
    volumes:
      - type: bind
        source: .
        target: /go/app
    networks:
      - golang_test_network

db:
  container_name: db
  build:
    context: ./build/db
    dockerfile: Dockerfile
  tty: true
  platform: linux/amd64
  ports:
    - 3306:3306
  env_file:
    - ./build/db/.env
  volumes:
    - type: volumes
      source: mysql_test_volume
      target: /var/lib/mysql
    - type: bind
      source: ./build/db/init
      target: /docker-entrypoint-initdb.d
  networks:
    - golang_test_network

volumes:
  mysql_test_volume:
    name: mysql_test_volume

networks:
  golang_test_network:
    external: true
```

#### 1-3-3. Dockerfile

- alpineサーバのgolangを使用
- ホットリロードのAirの導入

```dockerfile: build/db/Dockerfile
FROM mysql:8.0
ENV LANG ja_JP.UTF-8
```

```dockerfile: build/app/Dockerfile
FROM golang:1.17.7-alpine
RUN apk update && apk add git
RUN go get github.com/cosmtrek/air@v1.29.0
RUN mkdir -p /go/app
WORKDIR /go/app

CMD ["air", "-c", ".air.toml"]
```

#### 1-3-4. .air.toml

`build/app/Dockerfile`と同じ階層に`.air.toml`を作成

```toml: build/app/.air.toml
# Config file for [Air](https://github.com/cosmtrek/air) in TOML format

# Working directory
# . or absolute path, please note that the directories following must be under root.
root = "."
tmp_dir = "tmp"

[build]
# Just plain old shell command. You could use `make` as well.
cmd = "go build -o ./tmp/main ./cmd/golang_crud"
# Binary file yields from `cmd`.
bin = "tmp/main"
# Customize binary, can setup environment variables when run your app.
full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
# Watch these filename extensions.
include_ext = ["go", "tpl", "tmpl", "html"]
# Ignore these filename extensions or directories.
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# Watch these directories if you specified.
include_dir = []
# Exclude files.
exclude_file = []
# Exclude specific regular expressions.
exclude_regex = ["_test.go"]
# Exclude unchanged files.
exclude_unchanged = true
# Follow symlink for directories
follow_symlink = true
# This log file places in your tmp_dir.
log = "air.log"
# It's not necessary to trigger build each time file changes if it's too frequent.
delay = 1000 # ms
# Stop running old binary when build errors occur.
stop_on_error = true
# Send Interrupt signal before killing process (windows does not support this feature)
send_interrupt = false
# Delay after sending Interrupt signal
kill_delay = 500 # ms

[log]
# Show log time
time = false

[color]
# Customize each part's color. If no color found, use the raw app log.
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# Delete tmp directory on exit
clean_on_exit = true
```

#### 1-3-5. main.go

`main.go`はhttpのハンドルのみ。
articleパッケージ内にCRUDメソッドを後ほど実装する。

```go: main.go
package main

import (
	"golang_crud/internal/article"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", article.Index)
	http.HandleFunc("/show", article.Show)
	http.HandleFunc("/create", article.Create)
	http.HandleFunc("/edit", article.Edit)
	http.HandleFunc("/delete", article.Delete)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

```

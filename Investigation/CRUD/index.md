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

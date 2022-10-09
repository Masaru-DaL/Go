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

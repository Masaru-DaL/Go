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

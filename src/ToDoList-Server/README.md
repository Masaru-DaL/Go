- [golang TODO Application](#golang-todo-application)
  - [1. 要件定義](#1-要件定義)
  - [2. Web APIの仕様選定](#2-web-apiの仕様選定)
  - [3. Postman](#3-postman)
  - [4. GETメソッドを実装後](#4-getメソッドを実装後)
# golang TODO Application

## 1. 要件定義

* TODOアプリケーションに必要な要件定義
1. タスクは番号で管理する。
2. タスクには名前を付けられる。
3. タスクが終わったかどうかをチェックできる。

| TODOの要件定義 | タスク番号 | タスク名         | 終わったかどうか |
| --------- | ----- | ------------ | -------- |
| テーブルの列名   | id    | name         | finished |
| 列の型       | uuid型(16進数32桁) | 文字列型(string) |  ブール型(boolean)        |

* 完成イメージ

| id(uuid) | name | finished    |
| -------- | ---- | --- |
| 93b7cb6a         |  golangを勉強する    |  [ ]   |
| 93cac93b        |  技術記事を書く    |  [X]   |

## 2. Web APIの仕様選定

* GET `/tasks`
  + **全てのタスク一覧を、サーバのDBから取得する**

* POST `/tasks`
  + **タスクをサーバのDBに追加する**
  + JSONデータで送る。
  + `"name": "golangを勉強する"`
  + idは自動で振られるようにする。

* PUT `/tasks/:id`
  + **idを指定して、サーバのDBのタスクを完了にする**
  + 例えば、`PUT tasks/93b7cb6a`のようにリクエストする。

* DELETE `tasks/:id`
  + **idを指定して、サーバのDBのタスクを削除する**
  + 例えば、`DELETE tasks/93b7cb6a`のようにリクエストする。

## 3. Postman

[Postman](https://www.postman.com/downloads/)でAPIを叩けるようにしておく。

## 4. GETメソッドを実装後

1. `docker compose up -d --build`
2. serverコンテナ内に入る
3. `go run main.go`

```shell:
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.9.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______

                                    O\

⇨ http server started on [::]:8000
```

4. POSTMANでAPIを叩く
GETメソッド: `http://localhost:8000/api/tasks` -> send
ターミナルに表示: `2022-10-19T10:28:29.639749756+09:00 localhost:8000 GET /api/tasks 200`

POSTMANに表示: [] (DBが空なので正常)

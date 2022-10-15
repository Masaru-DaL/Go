- [golang TODO Application](#golang-todo-application)
  - [1. 要件定義](#1-要件定義)
  - [2. Web APIの仕様選定](#2-web-apiの仕様選定)
  - [3. Postman](#3-postman)
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

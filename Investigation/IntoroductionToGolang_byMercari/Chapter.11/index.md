- [メルカリ作のプログラミング言語Go完全入門 読破](#メルカリ作のプログラミング言語go完全入門-読破)
- [11. データベース](#11-データベース)
  - [11-1. データベースへの接続とSQLの実行](#11-1-データベースへの接続とsqlの実行)
      - [11-1-1. データベース](#11-1-1-データベース)
      - [11-1-2. database/sqlパッケージ](#11-1-2-databasesqlパッケージ)
      - [11-1-3. ドライバの登録](#11-1-3-ドライバの登録)
      - [11-1-4. SQLite](#11-1-4-sqlite)
      - [11-1-5. データベースのオープン](#11-1-5-データベースのオープン)
      - [11-1-6. SQLの実行](#11-1-6-sqlの実行)
      - [11-1-7. テーブルの作成](#11-1-7-テーブルの作成)
      - [11-1-8. レコードの挿入](#11-1-8-レコードの挿入)
      - [11-1-9. 複数レコードのスキャン](#11-1-9-複数レコードのスキャン)
      - [11-1-10. レコードの更新](#11-1-10-レコードの更新)
  - [11-2. トランザクション](#11-2-トランザクション)
      - [11-2-1. トランザクション](#11-2-1-トランザクション)
# メルカリ作のプログラミング言語Go完全入門 読破
# 11. データベース
## 11-1. データベースへの接続とSQLの実行
#### 11-1-1. データベース
- データベースとは
  - データを永続化するためのソフトウェア
    - データの保存および検索に特化している
    - データの記録方式によっていくつか種類がある
      - リレーショナルデータモデル
      - ネットワーク型データモデル

- リレーショナルデータベース
  - 表形式でデータを保存しているデータベース
  - 広く使われている
  - MySQL, SQLite, PostgreSQL, Oracle DBなど

- SQL
  - データベース用の問い合わせ言語
    - select, fromなど
  - データを検索・更新・削除など、クエリを記述するための言語

#### 11-1-2. database/sqlパッケージ
RDB(リレーショナルデータベース)にアクセスするためのパッケージ
- 共通機能を提供
  - クエリの発行
  - トランザクション(1setな処理単位)
- データベースの種類ごとにドライバが存在
  - [SQL database drivers](https://github.com/golang/go/wiki/SQLDrivers)

#### 11-1-3. ドライバの登録
**Go公式のデータベースドライバがない！**
- ドライバ
  - 各種RDBに対応したドライバ
    - MySQLやSQLiteなど

- 登録するには
  - インポートするだけ
    - 例: `import _ "modernc.org/sqlite"
  - initで登録される
  - パッケージ自体は直接触らない

#### 11-1-4. SQLite
ファイルベースのデータベース
- 軽量なRDBで、単一のファイルに記録される
- アプリケーションに組み込まれることが多い
  - 他のRDBは、通常はサーバとして動作する

#### 11-1-5. データベースのオープン
[Open関数](https://pkg.go.dev/database/sql#Open:~:text=%E5%88%B6%E5%BE%A1%E3%81%A7%E3%81%8D%E3%81%BE%E3%81%99%E3%80%82-,%E9%96%A2%E6%95%B0%E3%82%92%E9%96%8B%E3%81%8F%20%C2%B6,-func%20Open(driverName)を使用する
データベースハンドルを作成するために使用される。
```go:
/* Open(<driver>, <dataSourceName>) */
db, err := sql.Open("sqlite", "database.db")
/* driver -> データベースドライバーを指定
  dataSourceName -> データベース名や認証資格情報などのデータベース固有の接続情報を指定する */
```

- *sql.DBの特徴
  - 複数のゴールーチンから使用可能
  - コネクションプール機能
  - 一度開いたら使い回す
  - Closeは滅多にしない

[参考](http://go-database-sql.org/overview.html#:~:text=Reading%20and%20Resources-,Overview,-Improve%20this%20page): **sql.DB**とは
> Goでデータベースにアクセスするには、sql.DBを使用する。
> sql.DBはデータベース接続ではないということ。
> また、特定のデータベースソフトウェアの"データベース"や"スキーマ"という概念に対応するものではない。
> これはデータベースのインタフェースと存在を抽象化したもので、ローカルファイル、ネットワーク接続、インメモリ、インプロセスなど様々な携帯がある。

> sql.DBは、ドライバを介して、実際のデータベースへの接続を開いたり閉じたりする。これは必要に応じて接続のプールを管理し、前述のように様々なことを行う。

#### 11-1-6. SQLの実行
*sql.DBのメソッドを使用
https://pkg.go.dev/database/sql#DB.Query
```go:
// INSERTやDELETEなど
/* 行を返すことなく、クエリを実行する */
func (db *DB) Exec(query string, args ...interface{}) (Result, error)

// SELECTなどで複数レコードを取得する場合
/* 行を返す。 */
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)

// SELECTなどで1つのレコードを取得する場合
/* 最大で1行を返すと予想されるクエリを実行する。
  常にnil以外の値を返す。
  エラーは、RowのScanメソッドが呼び出されるまで保留される。 */
func (db *DB) QueryRow(query string, args ...interface{}) *Row
```

#### 11-1-7. テーブルの作成
(*sql.DB).Execを使う
```go:
/* sqlに` ~ `までが代入されていて、それがdb.Execの引数に指定されている */
/* 代入しない場合はdb.Exec(CREATE TABLE...)となる。 */
const sql = `
CREATE TABLE IF NOT EXISTS user (
	id   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	age  INTEGER NOT NULL
);
`
if _, err := db.Exec(sql); err != nil {
	// エラー処理
}
```

#### 11-1-8. レコードの挿入
**SQL文を実行すると、メタデータ(挿入した行のIDや影響を受けた行数)にアクセス可能な`sql.Result`が生成される。
AUTOINCREMENTのIDは、*sql.Resultから取得できる
```go:
type User struct {
	ID   int64
	Name string
	Age  int64
}
users := []*User{{Name: "tenntenn", Age: 32}, {Name: "Gopher", Age: 10}}
for i := range users {
	const sql = "INSERT INTO user(name, age) values (?,?)"

  /* レコードの挿入 */
	r, err := db.Exec(sql, users[i].Name, users[i].Age)
	if err != nil { /* エラー処理 */ }
	id, err := r.LastInsertId()
	if err != nil { /* エラー処理 */ }
	users[i].ID = id
	fmt.Println("INSERT", users[i])
}
```

#### 11-1-9. 複数レコードのスキャン
(*sql.DB).Queryと, *sql.Rowsを使う
rows.Next -> メソッドで取得できることがあることを確認する
rows.Scan -> メソッドでレコードを取得
forで確認->取得。
```go:
rows, err := db.Query("SELECT * FROM user WHERE age = ?", age)
if err != nil { /* エラー処理 */ }
for rows.Next() {
	var u User
	if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
		/* エラー処理 */
	}
	fmt.Println(u)
}
if err := rows.Err(); err != nil { /* エラー処理 */ }
```

#### 11-1-10. レコードの更新
更新したレコード数は*sql.Resultから取得
```go:
r, err := db.Exec("UPDATE user SET age = age + 1 WHERE id = 1")
if err != nil { /* エラー処理 */ }

/* RowsAffected() -> 更新したレコード数を取得している
  仮に存在しないidや同じnameで更新しようとすると0が表示される。 */
cnt, err := r.RowsAffected()
if err != nil { /* エラー処理 */ }
fmt.Println("Affected rows:", cnt)
```

## 11-2. トランザクション
#### 11-2-1. トランザクション


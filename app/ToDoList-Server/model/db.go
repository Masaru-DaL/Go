package model

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

/* DBconnection: DB接続とテーブルを作成する関数 */
func DBconnection() *sql.DB {
	// GetDBConfigを実行し、戻り値をdsnと定義する
	dsn := GetDBConfig()
	// error型のerrを定義する
	var err error
	// dsnを使ってDBに接続する。
	// 戻り値をdb, errに代入する
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}

	// Task型のテーブルを作成する
	CreateTable(db)
	// *gorm.DB型を、*sql.DB型に変換する
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	// sqlDBを返す
	return sqlDB
}

/* DBのdsnを取得する関数 */
func GetDBConfig() string {
	// docker-compose.ymlで設定した各種環境変数を読み込む
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DBNAME")

	// dsn(DBの接続情報つける識別子)を定義する
	dsn := fmt.Sprintf("%s.%s@tcp(%s:%s)/%s", user, password, hostname, port, dbname) + "?charset=utf8mb4&parseTime=True&loc=local"
	// dsnを返す
	return dsn
}

/* Task型のテーブルを作成する関数 */
func CreateTable(db *gorm.DB) {
	// Task型のテーブルを作成する
	db.AutoMigrate(&Task{})
}

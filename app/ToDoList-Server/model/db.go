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

/* DB接続とテーブルを作成する */
func DBconnection() *sql.DB {
	dsn := GetDBConfig()
	var err error
	// Open関数でDB接続を行う
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	// テーブルの作成
	CreateTable(db)
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	return sqlDB
}

/* DBのdsnを取得する */
func GetDBConfig() string {
	/* 接続情報を設定 */
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DBNAME")

	dsn := fmt.Sprintf("%s.%s@tcp(%s:%s)/%s", user, password, hostname, port, dbname) + "?charset=utf8mb4&parseTime=True&loc=local"
	return dsn
}

/* Task型のテーブルを作成する */
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Task{})
}

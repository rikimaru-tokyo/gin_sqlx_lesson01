package myDatabase

import (
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DbInit() データベースを接続する。
func DbInit() *sqlx.DB {
	//dsn形式：username:password@protocol(address)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&loc=Asia%%2fTokyo",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"),
	)

	//MySQLに接続
	conn, err := sqlx.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	return conn
}



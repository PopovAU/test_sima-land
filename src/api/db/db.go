package db

import (
	"api/setting"
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// errors
var (
	ErrNotFound = sql.ErrNoRows
)

// Init db connect
func Init() {
	dataSourceName := fmt.Sprintf("%s:%s@%s/", setting.DBUser, setting.DBPassword, setting.DBHost)
	_db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	db = _db
	initDatabase()
	initTable()
}

func initDatabase() {
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`;", setting.DBName))
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf("USE `%s`;", setting.DBName))
	if err != nil {
		panic(err)
	}
}

func initTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `user` ( " +
		" `id` int(11) NOT NULL  PRIMARY KEY AUTO_INCREMENT, " +
		" `name` varchar(255) NOT NULL " +
		" ) ENGINE=InnoDB DEFAULT CHARSET=utf8;")
	if err != nil {
		panic(err)
	}
}

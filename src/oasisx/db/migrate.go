package main

import (
	"oasisx/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	DBMS := "mysql"
	USER := "root"
	PASS := "1234"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "oasis"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	//db, err := gorm.Open(DBMS, CONNECT)
	db, _ := gorm.Open(DBMS, CONNECT)

	db.CreateTable(&models.User{})
}

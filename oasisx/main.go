package main

import (
	"fmt"
	"log"
	"oasisx/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zenazn/goji"
)

var db *gorm.DB

func main() {

	goji.Get("/user/index", userIndex)
	goji.Get("/user/new", userNew)
	goji.Post("/user/new", userCreate)
	goji.Get("/user/edit/:id", userEdit)
	goji.Post("/user/update/:id", userUpdate)
	goji.Get("/user/delete/:id", userDelete)
	goji.Serve()
}

// データベースの初期化
func init() {
	var err error
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		config.Config.DbUserName,
		config.Config.DbUserPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbName,
	)

	// configから読み込んだ情報を元に、データベースに接続します
	db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connect database..")

	}
}

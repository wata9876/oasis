package main

import (
	"fmt"
	models "gin_sample/model"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/joho/godotenv"
)

//DbEngine DB初期設定
var DbEngine *xorm.Engine

//main DB
func main() {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Println("DB読み込み失敗")
	}

	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("USER")
	PASS := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DBNAME")

	url := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	DbEngine, err = xorm.NewEngine(DBMS, url)

	err = DbEngine.CreateTables(models.Book{})
	err = DbEngine.CreateTables(models.User{})
	if err != nil {
		log.Fatalf("テーブルの生成に失敗しました。: %v", err)
	}

}

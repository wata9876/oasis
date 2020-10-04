package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	error := db.Create(&samples{
		Name: "テスト太郎",
		// age:        18,
		// address:    "東京都千代田区",
		// created_at: getDate(),
		// updated_at: getDate(),
	}).Error
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("データ追加成功")
	}
}

func getDate() string {
	const layout = "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(layout)
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "katsu315"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "oasis"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

// Users ユーザー情報のテーブル情報
type samples struct {
	id   int
	Name string `json:"name"`
	/*age        int    `json:"age"`
	address    string `json:"address"`
	created_at string `json:"created_at" sql:"not null;type:date"`
	updated_at string `json:"updated_at" sql:"not null;type:date"`*/
}

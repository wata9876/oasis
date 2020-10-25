package service

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//init DB初期化
func init() {

	driverName := "mysql"
	DsName := "root:katsu315@tcp(127.0.0.1:3306)/example?parseTime=true&charset=utf8"

	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	//DbEngine.ShowSQL(true)
	//DbEngine.SetMaxOpenConns(2)
	//DbEngine.Sync2(new(model.Book))
	fmt.Println("init data base ok")
}

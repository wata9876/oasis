package service

import (
	"errors"
	"fmt"
	"gin_sample/model"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

// DbEngine DB
var DbEngine *xorm.Engine

//BookService BOOK構造体定義
type BookService struct {
}

//Book bookテーブル
type Book struct {
	ID      int    `xorm:"'id'"`
	title   string `xorm:"'title'"`
	content string `xorm:"'content'"`
}

//AddBook 登録処理
func (BookService) AddBook(book *model.Book) error {

	_, err := DbEngine.Insert(book)
	if err != nil {
		return err
	}
	return nil
}

//GetBookList 一覧表示
func (BookService) GetBookList() []model.Book {
	books := make([]model.Book, 0)
	DbEngine.Find(&books)
	return books
}

//DeleteBook 削除
func (BookService) DeleteBook(id int) error {
	book := new(model.Book)
	_, err := DbEngine.Id(id).Delete(book)
	if err != nil {
		return err
	}
	return nil
}

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

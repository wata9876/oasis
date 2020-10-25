package service

import (
	"fmt"
	"gin_sample/model"

	//_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

// DbEngine DB
var DbEngine *xorm.Engine

//BookService BOOK構造体定義
type BookService struct {
}

//Book bookテーブル
type Book struct {
	ID      int64  `xorm:"'id'"`
	title   string `xorm:"'title'"`
	content string `xorm:"'content'"`
}

//SetBook エラーチェック
func (BookService) SetBook(book *model.Book) error {
	_, err := DbEngine.Insert(book)
	if err != nil {
		return err
	}
	return nil
}

//GetBookList 一覧表示
func (BookService) GetBookList() []model.Book {
	fmt.Println("サービスクラス通っているaa")
	books := make([]model.Book, 0)
	DbEngine.Find(&books)
	fmt.Println(books)
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

package service

import (
	"errors"
	"fmt"
	"strconv"

	"gin_sample/db"
	"gin_sample/model"
	"log"

	"github.com/gin-gonic/gin"
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
	ID      int    `xorm:"'i_d'"`
	Title   string `xorm:"'title'"`
	Author  string `xorm:"'author'"`
	Content string `xorm:"'content'"`
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

// EditBook 更新
func (BookService) EditBook(id int) Book {
	book := Book{ID: id}
	DbEngine.Get(&book)
	fmt.Println(book)
	return book
}

// UpdateBook 書籍更新
func (BookService) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book := new(Book)

	book.Title = c.PostForm("title")
	book.Author = c.PostForm("author")
	book.Content = c.PostForm("content")

	DbEngine.Where("i_d = ?", id).Update(book)
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

	DB := db.Dbinit{}
	driverName, url := DB.GetDbInfo()
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, url)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	fmt.Println("init data base ok")
}

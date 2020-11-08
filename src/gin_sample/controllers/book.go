package controllers

import (
	"errors"
	"fmt"
	"gin_sample/model"
	"gin_sample/service"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// DbEngine DB
var DbEngine *xorm.Engine
var tpl *template.Template

//Book 書籍情報一覧
type Book struct {
	ID      int    `xorm:"pk autoincr int" form:"id" json:"id"`
	Title   string `json:"title" 　　　　　　　xorm:"'title'"`
	Author  string `json:"author" 　　　　　　　xorm:"'author'"`
	Content string `json:"content" 　　　　　　　xorm:"'content'"`
}

//BookList 全件
func BookList(c *gin.Context) {
	bookService := service.BookService{}
	BookLists := bookService.GetBookList()
	c.HTML(200, "index.html", gin.H{"books": BookLists})
}

//BookAdd 新規登録
func BookAdd(c *gin.Context) {
	book := model.Book{}
	book.Title = c.PostForm("title")
	book.Author = c.PostForm("author")
	book.Content = c.PostForm("content")

	err := c.Bind(&book)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	bookService := service.BookService{}
	err = bookService.AddBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.Redirect(302, "/book/list")
}

//BookEdit 更新処理
func BookEdit(c *gin.Context) {

	id, err := PostParamID(c.Param("id"))
	if err != nil {
		log.Println("ID取得失敗")
	}
	BookService := service.BookService{}
	book := BookService.EditBook(int(id))

	c.HTML(200, "edit.html",
		gin.H{
			"id":      book.ID,
			"title":   book.Title,
			"author":  book.Author,
			"content": book.Content,
		})
}

//BookUpdate 書籍更新
func BookUpdate(c *gin.Context) {

	book := new(Book)

	id := c.PostForm("id")
	book.Title = c.PostForm("title")
	book.Author = c.PostForm("author")
	book.Content = c.PostForm("content")

	DbEngine.Where("i_d = ?", id).Update(book)
	c.Redirect(302, "/book/list")
}

//BookDelete 書籍削除
func BookDelete(c *gin.Context) {

	id, err := PostParamID(c.Param("id"))
	if err != nil {
		log.Println("ID取得失敗")
	}

	bookService := service.BookService{}
	bookService.DeleteBook(int(id))

	c.Redirect(302, "/book/list")
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

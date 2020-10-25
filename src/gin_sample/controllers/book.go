package controllers

import (
	"errors"
	"fmt"
	"gin_test/service"
	"html/template"
	"log"
	"strconv"

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
// func BookList(c *gin.Context) {
// 	bookService := service.BookService{}
// 	BookLists := bookService.GetBookList()
// 	c.HTML(200, "index.html", gin.H{"books": BookLists})
// }
//BookList 全件
func BookList(c *gin.Context) {

	books := make([]Book, 0)
	DbEngine.Find(&books)
	fmt.Println(books)

	c.HTML(200, "index.html", gin.H{"books": books})
}

//BookNew 新規作成画面
// func BookNew(c *gin.Context) {
// 	c.HTML(200, "new.html", gin.H{})
// }

//BookAdd 新規登録
func BookAdd(c *gin.Context) {
	book := new(Book)
	book.Title = c.PostForm("title")
	book.Author = c.PostForm("author")
	book.Content = c.PostForm("content")

	DbEngine.Insert(book)

	c.Redirect(302, "/book/list")
}

//BookEdit 更新処理
func BookEdit(c *gin.Context) {

	n := c.Param("id")
	id, _ := strconv.Atoi(n)
	//fmt.Println(id)
	book := Book{ID: id}
	DbEngine.Get(&book)
	//fmt.Println(book.Title)
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
	fmt.Println("update通っている")
	book := new(Book)

	id := c.PostForm("id")
	fmt.Println(id)
	book.Title = c.PostForm("title")
	book.Author = c.PostForm("author")
	book.Content = c.PostForm("content")
	fmt.Println(book.Title)
	DbEngine.Where("i_d = ?", id).Update(book)

	//DbEngine.Where("i_d = ?", id).Update(&book)
	c.Redirect(302, "/book/list")
}

//BookDelete 書籍削除
func BookDelete(c *gin.Context) {

	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	fmt.Println("削除")
	// intID, err := strconv.ParseInt(id, 10, 0)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Bad request")
	// 	return
	// }

	bookService := service.BookService{}
	bookService.DeleteBook(int(id))

	// c.JSON(http.StatusCreated, gin.H{
	// 	"status": "ok",
	// })

	//book := new(Book)
	//DbEngine.Where("i_d = ?", id).Delete(book)
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

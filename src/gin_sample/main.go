package main

import (
	"gin_test/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*.html")
	router.GET("/login", func(c *gin.Context) {
		//router.LoadHTMLGlob("templates/auth/*")
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	router.POST("/login", controllers.PostLogin)

	router.GET("/top", func(c *gin.Context) {
		//router.LoadHTMLGlob("template/menu/*")
		c.HTML(200, "top.html", gin.H{})
	})
	book := router.Group("/book")
	{
		//router.LoadHTMLGlob("template/book/*")
		book.GET("/list", controllers.BookList)
		book.GET("/new", func(c *gin.Context) {
			c.HTML(200, "new.html", gin.H{})
		})

	}
	router.POST("/book/add", controllers.BookAdd)
	router.GET("book/edit/:id", controllers.BookEdit)
	router.POST("book/update/:id", controllers.BookUpdate)
	router.GET("book/delete/:id", controllers.BookDelete)
	//}
	router.Run(":3000")

}

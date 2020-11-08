package router

import (
	"gin_sample/controllers"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//GetRouter ルートを定義
func GetRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("public/template/*.html")

	//bootstrap
	//loginDesign := router.Group("/login")
	router.Static("public/css", "public/css")
	router.Static("public/js", "public/js")

	userDesign := router.Group("/user")
	userDesign.Static("public/css", "public/css")
	userDesign.Static("public/js", "public/js")

	bookDesign := router.Group("/book")
	bookDesign.Static("public/css", "public/css")
	bookDesign.Static("public/js", "public/js")

	// router.GET("/test", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "bootstrap.html", gin.H{})
	// })
	// セッション
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//認証
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	router.POST("/login", controllers.PostLogin)
	router.GET("/logout", controllers.LogOut)

	//トップページ
	menu := router.Group("/menu")
	menu.Use(controllers.SessionCheck())
	{
		menu.GET("/top", controllers.GetMenu)
	}

	//ユーザー情報
	user := router.Group("/user")
	user.Use(controllers.SessionCheck())
	{
		user.GET("/list", controllers.UserList)
		user.GET("/new", func(c *gin.Context) {
			c.HTML(200, "user_new.html", gin.H{})
		})
		user.POST("/add", controllers.UserAdd)
		user.GET("/edit/:id", controllers.UserEdit)
		user.POST("/update/:id", controllers.UserUpdate)
		user.GET("/delete/:id", controllers.UserDelete)
	}

	//書籍情報
	book := router.Group("/book")
	book.Use(controllers.SessionCheck())
	{
		book.GET("/list", controllers.BookList)
		book.GET("/new", func(c *gin.Context) {
			c.HTML(200, "new.html", gin.H{})
		})
		book.POST("/add", controllers.BookAdd)
		book.GET("/edit/:id", controllers.BookEdit)
		book.POST("/update/:id", controllers.BookUpdate)
		book.GET("/delete/:id", controllers.BookDelete)
	}

	return router
}

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
	router.LoadHTMLGlob("template/*.html")

	router.Static("book/assets/css", "assets/css")
	router.Static("user/assets/css", "assets/css")

	//router.Static("assets/css", "assets/css")
	//router.Static("assets/js", "assets/js")

	// セッション
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//認証
	router.GET("/login", func(c *gin.Context) {
		router.LoadHTMLGlob("template/*.html")
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	router.POST("/login", controllers.PostLogin)
	router.GET("/logout", controllers.LogOut)

	//トップページ
	menu := router.Group("/menu")
	menu.Use(controllers.SessionCheck())
	{

		//menu.GET("/top", controllers.GetMenu)
		menu.GET("/top", func(c *gin.Context) {
			router.LoadHTMLGlob("template/menu/*.html")
			controllers.GetMenu(c)
		})
	}

	//ユーザー情報
	user := router.Group("/user")
	user.Use(controllers.SessionCheck())
	{
		//user.GET("/list", controllers.UserList)
		user.GET("/list", func(c *gin.Context) {
			router.LoadHTMLGlob("template/user/*.html")
			controllers.UserList(c)
		})

		user.GET("/new", func(c *gin.Context) {
			router.LoadHTMLGlob("template/user/*.html")
			c.HTML(200, "new.html", gin.H{})
		})
		user.POST("/add", controllers.UserAdd)
		//user.GET("/edit/:id", controllers.UserEdit)
		user.GET("/edit/:id", func(c *gin.Context) {

			router.LoadHTMLGlob("template/user/*.html")
			controllers.UserEdit(c)
		})
		user.POST("/update/:id", controllers.UserUpdate)
		user.GET("/delete/:id", controllers.UserDelete)
	}

	//書籍情報
	book := router.Group("/book")
	//book.Static("/edit/:id", "assets/css")
	book.Use(controllers.SessionCheck())
	{
		//router.LoadHTMLGlob("assets/template/book/*.html")
		// book.Static("assets/css", "assets/css")
		// book.Static("assets/js", "assets/js")
		book.GET("/list", func(c *gin.Context) {
			router.LoadHTMLGlob("template/book/*.html")
			controllers.BookList(c)
		})
		book.GET("/new", func(c *gin.Context) {
			router.LoadHTMLGlob("template/book/*.html")
			c.HTML(200, "new.html", gin.H{})
		})
		book.POST("/add", controllers.BookAdd)
		//book.GET("/edit/:id", controllers.BookEdit)
		book.GET("/edit/:id", func(c *gin.Context) {

			router.LoadHTMLGlob("template/book/*.html")
			controllers.BookEdit(c)

		})
		book.POST("/update/:id", controllers.BookUpdate)
		book.GET("/delete/:id", controllers.BookDelete)
	}

	return router
}

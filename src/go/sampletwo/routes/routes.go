package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	name := c.DefaultQuery("name", "HOGE") // HOGEはデフォルト値?
	//name := c.Query("lastname") // デフォルトがない場合
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"name": name,
	})
}

func HelloParam(c *gin.Context) {
	name := c.Param("name")
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"name": name,
	})
}

func NoRoute(c *gin.Context) {
	// helloに飛ばす
	c.Redirect(http.StatusMovedPermanently, "/hello")
}

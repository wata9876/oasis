package routes

import (
	"github.com/gin-gonic/gin"
)

func HelloJson(c *gin.Context) {
	name := c.DefaultQuery("name", "HOGE") // HOGEはデフォルト値?
	//name := c.Query("lastname") // デフォルトがない場合
	c.JSON(200, gin.H{
		"name": name,
	})
}

func HelloJsonPram(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"name": name,
	})
}

package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

//PostLogin ログイン処理
func PostLogin(c *gin.Context) {
	log.Println("ログイン処理")

	//userID := c.PostForm("userId")
	//authService := service.AuthService{}
	//a := authService.LoginCheck()
	//fmt.Println(userID)
	//fmt.Println(a)
	//UserId := c.PostForm("userId")
	//log.Println(UserId)
	//password := c.PostForm("password")
	//log.Println(password)
	c.Redirect(302, "/top")
}

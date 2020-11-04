package controllers

import (
	"gin_sample/service"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//User ユーザー情報
type User struct {
	ID       int    `xorm:"pk not null autoincr int" form:"id" json:"id"`
	Name     string `json:"name" 　　　　　　　xorm:"'name'"`
	Age      int    `json:"age" 　　　　　　　xorm:"'age'"`
	Address  string `json:"address" 　　　　　　　xorm:"'not null address'"`
	Password string `json:"password" 　　　　　　　xorm:"'password'"`
}

//PostLogin ログイン処理
func PostLogin(c *gin.Context) {
	log.Println("ログイン処理")
	address := c.PostForm("address")
	password := c.PostForm("password")

	if password == "" {
		log.Println("パスワードは必須です")
		return
	}
	UserService := service.UserService{}
	isExist, user := UserService.GetLoginUser(address)

	//認証出来たらセッションに登録
	if isExist {
		AuthService := service.AuthService{}
		AuthService.SetSessionInfo(c, user.ID)
	}
	c.Redirect(302, "/menu/top")
}

//LogOut ログアウト
func LogOut(c *gin.Context) {
	log.Println("ログアウト")
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(302, "/login")
}

//GetMenu メニュー画面
func GetMenu(c *gin.Context) {
	userID, _ := c.Get("userID") // ログインユーザの取得
	c.HTML(http.StatusOK, "top.html", gin.H{"UserId": userID})
}

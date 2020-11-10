package controllers

import (
	"fmt"
	"gin_sample/model"
	"gin_sample/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DbEngine DB
//var DbEngine *xorm.Engine
var err error

//UserList 全件
func UserList(c *gin.Context) {
	userService := service.UserService{}
	userLists := userService.GetUserList()
	c.HTML(200, "index.html", gin.H{"users": userLists})
}

//UserAdd 新規登録
func UserAdd(c *gin.Context) {
	user := model.User{}
	user.Name = c.PostForm("name")
	user.Address = c.PostForm("address")

	UserService := service.UserService{}
	user.Password, err = UserService.GetPasswordHash(c.PostForm("password"))
	n := c.PostForm("age")
	age, _ := strconv.Atoi(n)
	user.Age = age
	fmt.Println(user)
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	userService := service.UserService{}
	err = userService.AddUser(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.Redirect(302, "/user/list")
}

//UserEdit 更新処理
func UserEdit(c *gin.Context) {
	id, err := PostParamID(c.Param("id"))
	if err != nil {
		log.Println("ID取得失敗")
	}
	userService := service.UserService{}
	user := userService.EditUser(int(id))
	c.HTML(200, "edit.html",
		gin.H{
			"id":      user.ID,
			"name":    user.Name,
			"age":     user.Age,
			"address": user.Address,
		})
}

//UserUpdate ユーザー更新
func UserUpdate(c *gin.Context) {
	userService := service.UserService{}
	userService.UpdateUser(c)
	c.Redirect(302, "/user/list")
}

//UserDelete ユーザー削除
func UserDelete(c *gin.Context) {
	id, err := PostParamID(c.Param("id"))
	if err != nil {
		log.Println("ID取得失敗")
	}
	userService := service.UserService{}
	userService.DeleteUser(int(id))
	c.Redirect(302, "/user/list")
}

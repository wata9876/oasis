package controllers

import (
	"fmt"
	"gin_sample/model"
	"gin_sample/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// DbEngine DB
//var DbEngine *xorm.Engine
var err error

//UserList 全件
func UserList(c *gin.Context) {
	userService := service.UserService{}
	userLists := userService.GetUserList()
	c.HTML(200, "user_index.html", gin.H{"users": userLists})
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

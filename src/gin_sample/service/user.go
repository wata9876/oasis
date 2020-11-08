package service

import (
	"errors"
	"fmt"
	"gin_sample/model"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"golang.org/x/crypto/bcrypt"
)

//UserService User構造体定義
type UserService struct {
}

//User userテーブル
type User struct {
	ID      int    `json:"id" xorm:"'i_d'"`
	Name    string `xorm:"name"` // name
	Age     int    `xorm:"age"`  // age
	Address string
}

var isExist bool

//AddUser 登録処理
func (UserService) AddUser(user *model.User) error {

	_, err := DbEngine.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

//GetUserList 一覧表示
func (UserService) GetUserList() []model.User {
	Users := make([]model.User, 0)
	DbEngine.Find(&Users)
	return Users
}

//DeleteUser 削除
func (UserService) DeleteUser(id int) error {
	User := new(model.User)
	_, err := DbEngine.Id(id).Delete(User)
	if err != nil {
		return err
	}
	return nil
}

//GetPasswordHash パスワード暗号化して返却
func (UserService) GetPasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("暗号化失敗")
	}
	return string(hash), err
}

//GetLoginUser ログインユーザー取得
func (UserService) GetLoginUser(address string) (bool, User) {

	var user = User{Address: address}
	DbEngine.Get(&user)
	isExist = false

	if user.ID != 0 {
		log.Println("ログイン成功")
		isExist = true
	}

	return isExist, user
}

// EditUser ユーザー更新
func (UserService) EditUser(id int) User {
	user := User{ID: id}
	DbEngine.Get(&user)
	return user
}

// UpdateUser ユーザー更新
func (UserService) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user := new(User)
	user.Name = c.PostForm("name")
	user.Age, _ = strconv.Atoi(c.PostForm("age"))
	user.Address = c.PostForm("address")

	DbEngine.Where("i_d = ?", id).Update(user)
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
	//DbEngine.Sync2(new(model.User))
	fmt.Println("init data base ok")
}

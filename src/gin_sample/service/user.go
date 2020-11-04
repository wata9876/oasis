package service

import (
	"database/sql"
	"errors"
	"fmt"
	"gin_sample/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"golang.org/x/crypto/bcrypt"
)

//UserService User構造体定義
type UserService struct {
}

//User userテーブル
type User struct {
	ID      int            `json:"id" xorm:"'i_d'"`
	Name    sql.NullString `xorm:"name"` // name
	Age     int            `xorm:"age"`  // age
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

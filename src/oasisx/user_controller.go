package main

import (
	"log"
	"strconv"
	"time"

	"fmt"
	"html/template"
	"net/http"
	"oasisx/models"

	"github.com/wcl48/valval"
	"github.com/zenazn/goji/web"
)

var tpl *template.Template

//FormData モデルクラス定義
type FormData struct {
	User models.User
	Mess string
}

//userIndex ユーザー情報一覧取得
func userIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	Users := []models.User{}
	//db.Find(&Users)
	//fmt.Println(Users)
	db.Select("id, name").Find(&Users)
	//fmt.Println(Users)
	tpl = template.Must(template.ParseFiles("view/user/index.html"))
	if err := tpl.Execute(w, Users); err != nil {
		log.Fatal(err)
	}
}

//userNew 登録画面
func userNew(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("view/user/new.html"))
	tpl.Execute(w, FormData{models.User{}, ""})
}

//userCreate 新規登録
func userCreate(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{Name: r.FormValue("Name")}
	if err := models.UserValidate(User); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("view/user/new.html"))
		tpl.Execute(w, FormData{User, Mess})
	} else {
		newYmd := time.Date(2015, time.December, 31, 0, 0, 0, 0, time.UTC)
		User.CreatedAt = newYmd
		db.Create(&User)
		//db.Create(&User)
		http.Redirect(w, r, "/user/index", 301)
	}
}

//userEdit 更新ページへ遷移
func userEdit(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	User.ID, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	fmt.Println(User.ID)
	db.Find(&User)
	fmt.Println(User.Name)
	tpl = template.Must(template.ParseFiles("view/user/edit.html"))
	tpl.Execute(w, User)
	//tpl.Execute(w, FormData{models.User})
}

//userUpdate 更新処理
func userUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	User := models.User{}
	User.ID, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	fmt.Println(User.ID)
	db.Find(&User)
	fmt.Println(User)
	User.Name = r.FormValue("Name")

	if err := models.UserValidate(User); err != nil {
		var Mess string
		errs := valval.Errors(err)
		for _, errInfo := range errs {
			Mess += fmt.Sprint(errInfo.Error)
		}
		tpl = template.Must(template.ParseFiles("view/user/edit.html"))
		tpl.Execute(w, FormData{User, Mess})
	} else {

		db.Save(&User)
		http.Redirect(w, r, "/user/index", 301)
	}
}

//userDelete
func userDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Println("userdeleteまで来ている")
	User := models.User{}
	User.ID, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Delete(&User)
	http.Redirect(w, r, "/user/index", 301)
}

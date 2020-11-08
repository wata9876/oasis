package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//UserID セッションに持たせる
var UserID interface{}

//PostParamID //IDをint型で返却
func PostParamID(key string) (int, error) {
	i, err := strconv.Atoi(key)
	return int(i), err
}

// JSON 形式で結果を返却
// data interface{} とすると、どのような変数の型でも引数として受け取ることができる
func responseByJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
	return
}

//SessionCheck セッション管理
func SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Println("認証チェック")
		session := sessions.Default(c)
		UserID = session.Get("userID")
		fmt.Println(UserID)
		if UserID == nil {
			log.Println("ログインしていません")
			c.Redirect(302, "/login")
		} else {
			log.Println("ログイン済")
		}
	}
}

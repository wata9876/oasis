package service

import (
	"fmt"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//AuthService 認証構造体定義
type AuthService struct {
}

//SetSessionInfo セッション情報を格納
func (AuthService) SetSessionInfo(c *gin.Context, userID int) {
	log.Println("セッション情報を格納")
	fmt.Println(userID)
	session := sessions.Default(c)
	session.Set("alive", true)
	session.Set("userID", userID)
	session.Save()
}

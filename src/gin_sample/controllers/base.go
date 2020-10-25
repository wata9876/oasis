package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//PostParam //POSTからパラメーター取得
func PostParam(c *gin.Context, key string) (int, error) {
	i, err := strconv.Atoi(c.Param(key))
	return int(i), err
}

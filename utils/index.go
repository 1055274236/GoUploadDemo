/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-09 21:12:53
 */
package utils

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func sendJson(code, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	c.Abort()
}

func SendSuccessJson(msg string, data interface{}, c *gin.Context) {
	sendJson("200", msg, data, c)
}

func SendErrorJson(code, msg string, data interface{}, c *gin.Context) {
	if code == "" {
		code = "500"
	}
	sendJson(code, msg, data, c)
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

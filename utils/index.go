/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-14 00:50:10
 */
package utils

import (
	"math/rand"
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

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

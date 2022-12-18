/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-19 00:56:36
 */
package main

import (
	"UploadApi/config"
	"UploadApi/router"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create(config.GinLogFile)
	gin.DefaultWriter = io.MultiWriter(f)
	gin.SetMode(config.GinRunMode)
	engine := gin.Default()
	router.InitRouter(engine) // 设置路由
	engine.Run(config.PORT)
}

/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-09 03:13:09
 */
package main

import (
	"UploadApi/config"
	"UploadApi/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.Default()
	router.InitRouter(engine) // 设置路由
	engine.Run(config.PORT)
}

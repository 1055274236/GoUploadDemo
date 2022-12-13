/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-14 03:33:08
 */
package router

import (
	"github.com/gin-gonic/gin"

	"UploadApi/controller/download"
	"UploadApi/controller/upload"
)

func InitRouter(r *gin.Engine) {
	r.GET("/download/:folder/:file", download.Download)
	r.POST("/upload", upload.Upload)
}

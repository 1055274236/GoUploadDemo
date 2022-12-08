/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-08 02:33:22
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

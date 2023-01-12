/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 20:19:26
 */
package router

import (
	"github.com/gin-gonic/gin"

	"UploadApi/controller/download"
	"UploadApi/controller/fileinfo"
	"UploadApi/controller/upload"
	"UploadApi/controller/uploadrecord"
)

func InitRouter(r *gin.Engine) {
	r.GET("/download/:folder/:file", download.Download)
	r.POST("/upload", upload.Upload)

	// fileinfo
	r.GET("/infobyid/:id", fileinfo.GetOneById)
	r.GET("/infobyhash/:hash", fileinfo.GetOneByHash)
	r.POST("/infobyid", fileinfo.GetMoreById)
	r.POST("/infobyhash", fileinfo.GetMoreByHash)

	// uploadrecord
	r.GET("/uploadbyid/:id", uploadrecord.GetOneById)
	r.GET("/uploadbyip/:ip", uploadrecord.GetOneByIp)
	r.POST("/uploadbyid", uploadrecord.GetMoreById)
	r.POST("/uploadbyip", uploadrecord.GetMoreByIp)
}

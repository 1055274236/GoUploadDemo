/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-14 03:32:45
 */
package upload

import (
	"UploadApi/config"
	"UploadApi/utils"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// 上传成功返回信息
func UploadSuccess(data interface{}, c *gin.Context) {
	utils.SendSuccessJson("Success", data, c)
}

// 上传失败返回信息
func UploadError(c *gin.Context, errmessage string, err error) {
	if errmessage == "" {
		errmessage = "Error"
	}
	utils.SendErrorJson("500", errmessage, err, c)
}

func Upload(c *gin.Context) {

	// 获取文件长度
	var content_length int64 = c.Request.ContentLength
	if content_length <= 0 || content_length > int64(config.FileContentMaxLength) {
		// log.Printf("content_length error\n")
		UploadError(c, "content_length error", nil)
		return
	}

	// 文件类型
	content_type_, has_key := c.Request.Header["Content-Type"]
	if !has_key {
		log.Printf("Content-Type error\n")
		UploadError(c, "Content-Type error", nil)
		return
	}

	// boundary 作为文件分隔符
	content_type := content_type_[0]
	const BOUNDARY string = "; boundary="
	loc := strings.Index(content_type, BOUNDARY)
	if loc == -1 {
		log.Printf("Content-Type error, no boundary\n")
		UploadError(c, "Content-Type error, no boundary", nil)
		return
	}
	boundary := []byte(content_type[(loc + len(BOUNDARY)):])

	fileInformationArr, err := ParseFileAndSave(c, boundary)

	if err != nil {
		UploadError(c, "", err)
	} else {
		UploadSuccess(fileInformationArr, c)
	}
}

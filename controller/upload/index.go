/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 17:31:30
 */
package upload

import (
	"UploadApi/config"
	"UploadApi/utils"
	"mime"

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
		UploadError(c, "content_length error", nil)
		return
	}

	// 文件类型
	ct := c.Request.Header.Get("Content-Type")
	_, params, err := mime.ParseMediaType(ct)
	_, hasb := params["boundary"]

	if err != nil || !hasb {
		UploadError(c, "请求头Content-Type错误", err)
		return
	}

	// boundary 作为文件分隔符
	boundary := []byte(params["boundary"])

	fileInformationArr, err := ParseFileAndSave(c, boundary)

	returnResult := InsertDB(fileInformationArr, c)

	if err != nil {
		UploadError(c, "", err)
	} else {
		UploadSuccess(returnResult, c)
	}
}

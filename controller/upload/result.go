/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-09 03:22:31
 */
package upload

import (
	"UploadApi/utils"

	"github.com/gin-gonic/gin"
)

// 上传成功返回信息
func UploadSuccess(c *gin.Context) {
	utils.SendSuccessJson("Success", "success", c)
}

// 上传失败返回信息
func UploadError(c *gin.Context, errmessage string) {
	if errmessage == "" {
		errmessage = "Error"
	}
	utils.SendErrorJson("500", errmessage, errmessage, c)
}

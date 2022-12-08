/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-09 02:48:59
 */
package upload

import (
	"UploadApi/utils"

	"github.com/gin-gonic/gin"
)

func UploadSuccess(c *gin.Context) {
	utils.SendSuccessJson("Success", "success", c)
}

func UploadError(c *gin.Context, errmessage string) {
	if errmessage == "" {
		errmessage = "Error"
	}
	utils.SendErrorJson("500", errmessage, errmessage, c)
}

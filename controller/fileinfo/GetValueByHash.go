/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 20:14:29
 */
package fileinfo

import (
	"UploadApi/databases/mysql/fileinfo"
	"UploadApi/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetOneByHash(c *gin.Context) {
	hash, hashFlag := c.Params.Get("hash")
	if !hashFlag {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "参数错误!", "", c)
		return
	}

	result, db := fileinfo.GetValuesByHash([]string{hash})
	if db.RowsAffected > 0 {
		utils.SendSuccessJson("操作成功", result[0], c)
	} else {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "信息不存在!", "", c)
	}
}

func GetMoreByHash(c *gin.Context) {
	hashs := c.PostForm("hashs")
	if hashs == "" {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "参数错误!", "", c)
		return
	}

	result, db := fileinfo.GetValuesByHash(strings.Split(hashs, ","))
	if db.RowsAffected > 0 {
		utils.SendSuccessJson("操作成功", result, c)
	} else {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "信息不存在!", "", c)
	}
}

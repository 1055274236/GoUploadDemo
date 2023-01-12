/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 20:42:30
 */
package uploadrecord

import (
	"UploadApi/databases/mysql/uploadrecord"
	"UploadApi/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetOneById(c *gin.Context) {
	id, idFlag := c.Params.Get("id")
	if !idFlag {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "参数错误!", "", c)
		return
	}

	result, db := uploadrecord.GetValuesById([]string{id})
	if db.RowsAffected > 0 {
		utils.SendSuccessJson("操作成功", result[0], c)
	} else {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "信息不存在", "", c)
	}
}

func GetMoreById(c *gin.Context) {
	ids := c.PostForm("ids")
	if ids == "" {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "参数错误!", "", c)
		return
	}

	result, db := uploadrecord.GetValuesById(strings.Split(ids, ","))
	if db.RowsAffected > 0 {
		utils.SendSuccessJson("操作成功", result, c)
	} else {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "信息不存在", "", c)
	}
}

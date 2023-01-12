/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 20:42:35
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

func GetOneByIp(c *gin.Context) {
	ip, ipFlag := c.Params.Get("ip")
	if !ipFlag {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "参数错误!", "", c)
		return
	}

	result, db := uploadrecord.GetValuesByIp([]string{ip})
	if db.RowsAffected > 0 {
		utils.SendSuccessJson("操作成功", result, c)
	} else {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "信息不存在", "", c)
	}
}

func GetMoreByIp(c *gin.Context) {
	ips := c.PostForm("ips")
	if ips == "" {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "参数错误!", "", c)
		return
	}

	result, db := uploadrecord.GetValuesByIp(strings.Split(ips, ","))
	if db.RowsAffected > 0 {
		utils.SendSuccessJson("操作成功", result, c)
	} else {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "信息不存在", "", c)
	}
}

/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-06 21:25:34
 */
package download

import (
	"UploadApi/config"
	"UploadApi/utils"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	// 获取参数
	folder, folderFlag := c.Params.Get("folder")
	file, fileFlag := c.Params.Get("file")
	if !folderFlag || !fileFlag {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "参数错误!", "", c)
		return
	}

	// 拼接本地文件地址
	p := path.Join(config.StaticFileDir, folder, file)
	if exist, _ := utils.HasDir(p); !exist {
		utils.SendErrorJson(strconv.Itoa(http.StatusInternalServerError), "文件不存在!", "", c)
		return
	}

	c.File(p)
}

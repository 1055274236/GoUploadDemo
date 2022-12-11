/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-12 01:54:14
 */
package download

import (
	"UploadApi/config"
	"UploadApi/utils"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	var (
		folder, _ = c.Params.Get("folder")
		file, _   = c.Params.Get("file")
		p         = path.Join(config.StaticFileDir, folder, file)
		bRange    = c.Request.Header.Get("range")
	)

	// 读取文件
	f, err := os.Open(p)
	if err != nil {
		utils.SendErrorJson("", "读取失败！", err, c)
		return
	}
	defer f.Close()

	// 解析请求头与文件基本信息
	headerParams, err := ParseHeaders(p, bRange)
	if err != nil {
		utils.SendErrorJson("500", "表头解析错误", "", c)
		return
	}

	// 设置响应头
	for key, value := range headerParams.headers {
		c.Writer.Header().Set(key, value[0])
	}

	// 向客户端返回数据
	if bRange != "" {
		readSize := make([]byte, headerParams.end-headerParams.start)
		f.ReadAt(readSize, int64(headerParams.start))
		c.Data(headerParams.statusCode, headerParams.headers["Content-Type"][0], readSize)
	} else if headerParams.size < 1024*1024*20 {
		c.File(p)
	} else {
		c.Done()
	}
}

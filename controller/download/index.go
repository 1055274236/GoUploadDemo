/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-08 22:38:00
 */
package download

import "github.com/gin-gonic/gin"

/// 解析多个文件上传中，每个具体的文件的信息
type FileHeader struct {
	ContentDisposition string
	Name               string
	FileName           string ///< 文件名
	ContentType        string
	ContentLength      int64
}

func Download(c *gin.Context) {

}

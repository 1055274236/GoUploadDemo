/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-09 01:16:55
 */
package upload

// / 解析多个文件上传中，每个具体的文件的信息
type FileHeader struct {
	ContentDisposition string
	Name               string
	FileName           string ///< 文件名
	ContentType        string
	ContentLength      int64
}

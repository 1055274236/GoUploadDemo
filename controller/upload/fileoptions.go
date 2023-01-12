/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-10 21:57:48
 */
package upload

import (
	"mime"
	"strings"
)

// @return FormName, Name, ContentType
// @return 表单名称， 文件名称， 文件类别
func ParseHeader(header []byte) (string, string, string) {
	var (
		FormName    string
		Name        string
		ContentType string
	)
	// headerEncode, err := url.QueryUnescape(string(header))
	// if err != nil {
	// 	log.Println("header 解析错误, ", err)
	// 	panic(err)
	// }
	headerArr := strings.Split(string(header), "\r\n")

	for _, item := range headerArr {
		colonIndex := strings.Index(item, ":")

		// 获取文件信息
		if strings.EqualFold(item[:colonIndex], "Content-Disposition") {
			_, params, err := mime.ParseMediaType(string(item[colonIndex+1:]))
			if err != nil {
				panic("错误输入！文件描述解析错误！")
			}

			name, hasName := params["name"]
			if !hasName {
				name = "files"
			}
			fileName, hasFileName := params["filename"]
			if !hasFileName {
				fileName = ""
			}
			FormName = name
			Name = fileName

		} else if strings.EqualFold(item[:colonIndex], "Content-Type") {
			// 获取文件类别
			mediatype, _, err := mime.ParseMediaType(string(item[colonIndex+1:]))
			if err != nil {
				panic("错误输入！文件类别解析错误！")
			}
			ContentType = mediatype
		} else {
			continue
		}

	}

	return FormName, Name, ContentType
}

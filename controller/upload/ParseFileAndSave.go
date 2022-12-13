/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-14 04:09:30
 */
package upload

import (
	"UploadApi/config"
	"UploadApi/utils"
	"bytes"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type FileHeader struct {
	Name          string
	NewName       string
	FileName      string ///< 文件名
	FilePath      string
	ContentType   string
	ContentLength int
}

// 根据传递的值，建立文本流
// @params suffix 文件后缀
// @return 文件流
// 新文件名称
// 文件位置
func createFileStream(suffix string) (*os.File, string, string) {
	var fileName string
	var filePath string
	fileDir := path.Join(config.StaticFileDir, time.Now().Format("20060102"))

	for {
		fileName = utils.RandStringBytes(16) + "." + suffix
		filePath = path.Join(fileDir, fileName)
		flag, err := utils.HasDir(filePath)
		if err != nil {
			log.Println("文件流创建失败！未知错误", err)
		}
		if !flag {
			break
		}
	}
	utils.CreateNecessaryDir(fileDir)
	f, _ := os.Create(filePath)
	return f, fileName, filePath
}

// 处理文件尺寸
func checkFileSize(fileHeader []FileHeader) {
	// 文件尺寸赋值
	for i := 0; i < len(fileHeader); i++ {
		fileProperty, _ := os.Stat(fileHeader[i].FilePath)
		fileHeader[i].ContentLength = int(fileProperty.Size())
	}
}

func ParseFileAndSave(c *gin.Context, boundary []byte) ([]FileHeader, error) {
	var (
		fileHeader     []FileHeader
		fileHeaderitem FileHeader
	)
	var (
		NAME         = "name=\""
		FILENAME     = "filename=\""
		CONTENT_TYPE = "Content-Type: "
	)
	var (
		f        *os.File
		fileName string
		filePath string
	)

	for {
		buf := make([]byte, 1024*12)

		// 数据接入
		read_len, err := c.Request.Body.Read(buf)
		if err != nil {
			if err != io.EOF {
				// 上传中断，删除不完整文件
				f.Close()
				os.Remove(filePath)
				fileHeader = fileHeader[:len(fileHeader)-1]
			}

			checkFileSize(fileHeader)
			break
		}

		// 文件表头位置查找
		boundary_loc := bytes.Index(buf, boundary)
		start_loc := boundary_loc + len(boundary)
		file_head_loc := bytes.Index(buf, []byte("\r\n\r\n"))

		if boundary_loc > -1 && file_head_loc > -1 {

			header := buf[start_loc:file_head_loc]

			// 数据初始化
			fileHeaderitem = FileHeader{}

			// 解析文件头
			// Name
			nameIndex := bytes.Index(header, []byte(NAME))
			seqIndex := nameIndex + bytes.Index(header[nameIndex:], []byte("\"; "))
			if seqIndex == nameIndex-1 {
				seqIndex = len(header)
			}
			fileHeaderitem.Name = string(header[nameIndex+len(NAME) : seqIndex])
			// FILENAME
			fileNameIndex := bytes.Index(header, []byte(FILENAME))
			seqIndex = fileNameIndex + bytes.Index(header[fileNameIndex:], []byte("\"; "))
			if seqIndex == fileNameIndex-1 {
				seqIndex = len(header)
			}
			fileHeaderitem.FileName = string(header[fileNameIndex+len(FILENAME) : seqIndex])
			// CONTENT_TYPE
			contentTypeIndex := bytes.Index(header, []byte(CONTENT_TYPE))
			seqIndex = contentTypeIndex + bytes.Index(header[contentTypeIndex:], []byte(";"))
			if seqIndex == contentTypeIndex-1 {
				seqIndex = len(header)
			}
			fileHeaderitem.ContentType = string(header[contentTypeIndex+len(CONTENT_TYPE) : seqIndex])

			// 获取后缀
			suffixBad := strings.Split(fileHeaderitem.FileName, ".")
			suffix := ""
			if len(suffixBad) > 1 {
				suffix = suffixBad[1]
			}

			// 创建文件
			f.Close()
			f, fileName, filePath = createFileStream(suffix)
			defer f.Close()

			fileHeaderitem.NewName = fileName
			fileHeaderitem.FilePath = filePath

			// 对成员进行添加
			fileHeader = append(fileHeader, fileHeaderitem)

		} else {
			// 文件数据写入
			f.Write(buf[:read_len])
		}
	}

	return fileHeader, nil
}

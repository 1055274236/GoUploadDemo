/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-21 11:51:03
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
	uuid "github.com/satori/go.uuid"
)

// 根据传递的值，建立文本流
// @params suffix 文件后缀
// @return 文件流
func createFileStream(suffix string, fileHeaderItem *FileHeader) *os.File {
	// 新的随机名称文件
	// 文件名
	fileName := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	// 文件夹名
	fileHeaderItem.Folder = time.Now().Format("20060102")
	// 文件夹地址
	fileDir := path.Join(config.StaticFileDir, fileHeaderItem.Folder)
	// 加后缀之后的文件名
	fileHeaderItem.FileName = fileName + suffix
	// 文件储存地址
	fileHeaderItem.FilePath = path.Join(fileDir, fileHeaderItem.FileName)

	// 创建文件夹
	utils.CreateNecessaryDir(fileDir)
	f, err := os.Create(fileHeaderItem.FilePath)
	if err != nil {
		log.Println("文件创建失败: ", err)
		panic(err)
	}

	return f
}

func ParseFileAndSave(c *gin.Context, boundary []byte) ([]FileHeader, error) {
	var (
		start_buff bytes.Buffer
		end_buff   bytes.Buffer
	)
	start_buff.WriteString("--")
	start_buff.Write(boundary)
	start_buff.WriteString("\r\n")
	end_buff.WriteString("--")
	end_buff.Write(boundary)
	end_buff.WriteString("--")

	var (
		isFirst        = true
		isEnd          = false
		oldChar        []byte
		jointChar      []byte
		fileHeader     []FileHeader
		spaceCharacter = Boundary{
			base:  boundary,
			start: start_buff.Bytes(),
			end:   end_buff.Bytes(),
		}
		fileHeaderitem FileHeader
		f              *os.File
	)

	var (
		header_start_index = -1 // 文件头起始位置
		header_end_index   = -1 // 文件头终止位置
		body_end_index     = -1 // body 传输流终止位置
		file_end_index     = -1 // 文件流终止位置
		header_bytes       []byte
		buf                []byte
	)

	for {
		// 读取片段
		buf = make([]byte, config.FileChunkSize)

		// 数据接入
		read_len, err := c.Request.Body.Read(buf)

		// 文件上传终止
		if err != nil && err != io.EOF {

			if l := len(fileHeader); l > 0 {
				f.Close()

				go utils.RemoveFile(fileHeader[l-1].FilePath)
				fileHeader = fileHeader[:l-1]
			}

			break
		}

		isEnd = bytes.Contains(buf[:read_len], spaceCharacter.end)

		if isFirst {
			isFirst = false
			oldChar = buf[:read_len]
			if isEnd {
				jointChar = oldChar
			} else {
				continue
			}
		} else {
			jointChar = append(oldChar, buf[:read_len]...)
		}

		for {
			// 确定文件终止位置，并将终止位置之前的数据传输到 f 文件流中
			header_start_index = bytes.Index(jointChar, spaceCharacter.start)

			// 查询这次的文件块中的终止符下标
			body_end_index = bytes.Index(jointChar, spaceCharacter.end)
			isEnd = body_end_index != -1
			if header_start_index > -1 {
				file_end_index = header_start_index - 2
			} else {
				file_end_index = body_end_index - 2
			}

			// 写入文件流
			if f != nil && file_end_index > -1 {
				f.Write(jointChar[:file_end_index])
				f.Close()
				jointChar = jointChar[file_end_index+2:]
				continue
			}
			// 当未找到文件头任何可识别部分时，将上一块文件流进行写入
			if header_start_index == -1 && file_end_index == -3 {
				if al := len(jointChar); al > read_len {
					spaceIndex := al - read_len
					f.Write(jointChar[:spaceIndex])
					oldChar = jointChar[spaceIndex:]
				} else {
					oldChar = jointChar
				}
				break
			}

			// 文件头搜索以及解析
			if header_start_index > -1 {
				// 找不到请求头结束的标志则进行下一部分的拼接
				header_end_index = header_start_index + bytes.Index(jointChar[header_start_index:], []byte("\r\n\r\n"))
				if header_end_index == -1 {
					oldChar = jointChar[header_start_index:]
					break
				}

				header_bytes = jointChar[header_start_index+len(spaceCharacter.start) : header_end_index]

				fileHeaderitem = FileHeader{}
				fileHeaderitem.FormName, fileHeaderitem.Name, fileHeaderitem.ContentType = ParseHeader(header_bytes)
				f = createFileStream(path.Ext(fileHeaderitem.Name), &fileHeaderitem)

				fileHeader = append(fileHeader, fileHeaderitem)

				// 更新拼接字符串，删除文件头部分 添加 \r\n\r\n 长度
				jointChar = jointChar[header_end_index+4:]
			} else {
				break
			}
		}

		// 读取完成
		if isEnd {
			break
		}
	}

	return fileHeader, nil
}

/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-10 02:16:18
 */
package upload

import (
	"UploadApi/config"
	"UploadApi/utils"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// / 解析描述文件信息的头部
// / @return FileHeader 文件名等信息的结构体
// / @return bool 解析成功还是失败
func ParseFileHeader(h []byte) (FileHeader, bool) {
	arr := bytes.Split(h, []byte("\r\n"))
	var out_header FileHeader
	out_header.ContentLength = -1
	const (
		CONTENT_DISPOSITION = "Content-Disposition: "
		NAME                = "name=\""
		FILENAME            = "filename=\""
		CONTENT_TYPE        = "Content-Type: "
		CONTENT_LENGTH      = "Content-Length: "
	)
	for _, item := range arr {
		if bytes.HasPrefix(item, []byte(CONTENT_DISPOSITION)) {
			l := len(CONTENT_DISPOSITION)
			arr1 := bytes.Split(item[l:], []byte("; "))
			out_header.ContentDisposition = string(arr1[0])
			if bytes.HasPrefix(arr1[1], []byte(NAME)) {
				out_header.Name = string(arr1[1][len(NAME) : len(arr1[1])-1])
			}
			l = len(arr1[2])
			if bytes.HasPrefix(arr1[2], []byte(FILENAME)) && arr1[2][l-1] == 0x22 {
				out_header.FileName = string(arr1[2][len(FILENAME) : l-1])
			}
		} else if bytes.HasPrefix(item, []byte(CONTENT_TYPE)) {
			l := len(CONTENT_TYPE)
			out_header.ContentType = string(item[l:])
		} else if bytes.HasPrefix(item, []byte(CONTENT_LENGTH)) {
			l := len(CONTENT_LENGTH)
			s := string(item[l:])
			content_length, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Printf("content length error:%s", string(item))
				return out_header, false
			} else {
				out_header.ContentLength = content_length
			}
		} else {
			log.Printf("unknown:%s\n", string(item))
		}
	}
	if len(out_header.FileName) == 0 {
		return out_header, false
	}
	return out_header, true
}

// / 从流中一直读到文件的末位
// / @return []byte 没有写到文件且又属于下一个文件的数据
// / @return bool 是否已经读到流的末位了
// / @return error 是否发生错误
func ReadToBoundary(boundary []byte, stream io.ReadCloser, target io.WriteCloser) ([]byte, bool, error) {
	read_data := make([]byte, 1024*8)
	read_data_len := 0
	buf := make([]byte, 1024*4)
	b_len := len(boundary)
	reach_end := false
	for !reach_end {
		read_len, err := stream.Read(buf)
		if err != nil {
			if err != io.EOF && read_len <= 0 {
				return nil, true, err
			}
			reach_end = true
		}
		//todo: 下面这一句很蠢，值得优化
		copy(read_data[read_data_len:], buf[:read_len]) //追加到另一块buffer，仅仅只是为了搜索方便
		read_data_len += read_len
		if read_data_len < b_len+4 {
			continue
		}
		loc := bytes.Index(read_data[:read_data_len], boundary)
		if loc >= 0 {
			//找到了结束位置
			target.Write(read_data[:loc-4])
			return read_data[loc:read_data_len], reach_end, nil
		}

		target.Write(read_data[:read_data_len-b_len-4])
		copy(read_data[0:], read_data[read_data_len-b_len-4:])
		read_data_len = b_len + 4
	}
	target.Write(read_data[:read_data_len])
	return nil, reach_end, nil
}

// / 解析表单的头部
// / @param read_data 已经从流中读到的数据
// / @param read_total 已经从流中读到的数据长度
// / @param boundary 表单的分割字符串
// / @param stream 输入流
// / @return FileHeader 文件名等信息头
// /         []byte 已经从流中读到的部分
// /         error 是否发生错误
func ParseFromHead(read_data []byte, read_total int, boundary []byte, stream io.ReadCloser) (FileHeader, []byte, error) {
	buf := make([]byte, 1024*4)
	found_boundary := false
	boundary_loc := -1
	var file_header FileHeader
	for {
		read_len, err := stream.Read(buf)
		if err != nil {
			if err != io.EOF {
				return file_header, nil, err
			}
			break
		}
		if read_total+read_len > cap(read_data) {
			return file_header, nil, fmt.Errorf("not found boundary")
		}
		copy(read_data[read_total:], buf[:read_len])
		read_total += read_len
		if !found_boundary {
			boundary_loc = bytes.Index(read_data[:read_total], boundary)
			if boundary_loc == -1 {
				continue
			}
			found_boundary = true
		}
		start_loc := boundary_loc + len(boundary)
		file_head_loc := bytes.Index(read_data[start_loc:read_total], []byte("\r\n\r\n"))
		if file_head_loc == -1 {
			continue
		}
		file_head_loc += start_loc
		ret := false
		file_header, ret = ParseFileHeader(read_data[start_loc:file_head_loc])
		if !ret {
			return file_header, nil, fmt.Errorf("ParseFileHeader fail:%s", string(read_data[start_loc:file_head_loc]))
		}
		return file_header, read_data[file_head_loc+4 : read_total], nil
	}
	return file_header, nil, fmt.Errorf("reach to sream EOF")
}

func Upload(c *gin.Context) {
	// 获取文件长度
	var content_length int64 = c.Request.ContentLength
	if content_length <= 0 || content_length > int64(config.FileContentMaxLength) {
		// log.Printf("content_length error\n")
		UploadError(c, "content_length error")
		return
	}
	content_type_, has_key := c.Request.Header["Content-Type"]
	if !has_key {
		// log.Printf("Content-Type error\n")
		UploadError(c, "Content-Type error")
		return
	}
	if len(content_type_) != 1 {
		// log.Printf("Content-Type count error\n")
		UploadError(c, "Content-Type count error")
		return
	}
	content_type := content_type_[0]
	const BOUNDARY string = "; boundary="
	loc := strings.Index(content_type, BOUNDARY)
	if loc == -1 {
		// log.Printf("Content-Type error, no boundary\n")
		UploadError(c, "Content-Type error, no boundary")
		return
	}
	boundary := []byte(content_type[(loc + len(BOUNDARY)):])
	log.Printf("[%s]\n", boundary)
	read_data := make([]byte, 1024*12)
	var read_total int = 0
	for {
		file_header, _, err := ParseFromHead(read_data, read_total, append(boundary, []byte("\r\n")...), c.Request.Body)
		if err != nil {
			// log.Printf("%v", err)
			UploadError(c, err.Error())
			return
		}
		log.Printf("file :%s\n", file_header.FileName)
		// 创建文件夹以及文件
		fileDir := path.Join(config.StaticFileDir, time.Now().Format("20060102"))
		utils.CreateNecessaryDir(fileDir)
		f, err := os.Create(path.Join(fileDir, file_header.FileName))
		if err != nil {
			// log.Printf("create file fail:%v\n", err)
			UploadError(c, err.Error())
			return
		}

		// 将读取到的数据全部存储到文件中
		// f.Write(file_data)
		// file_data = nil

		//需要反复搜索boundary
		temp_data, reach_end, err := ReadToBoundary(boundary, c.Request.Body, f)
		f.Close()
		if err != nil {
			// log.Printf("%v\n", err)
			UploadError(c, err.Error())
			return
		}
		if reach_end {
			break
		} else {
			copy(read_data[0:], temp_data)
			read_total = len(temp_data)
			continue
		}
	}

	UploadSuccess(c)
}

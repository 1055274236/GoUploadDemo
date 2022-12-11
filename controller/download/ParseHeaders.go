/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-12 01:58:36
 */
package download

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type HeaderParamsType struct {
	statusCode int
	start      int
	end        int
	size       int
	headers    http.Header
}

// 获取文件类型
func GetFileContentType(out *os.File) (string, error) {

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func ParseHeaders(filePath, bRange string) (HeaderParamsType, error) {

	// 检查文件并且初始化
	var headers HeaderParamsType
	headers.headers = http.Header{}
	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return headers, err
	}
	defer f.Close()

	// 初始化局部变量
	fileProperty, _ := f.Stat()
	fileSize := fileProperty.Size()
	contentType, err := GetFileContentType(f)
	if err != nil {
		panic(err)
	}
	headers.start = 0
	headers.end = int(fileSize)
	headers.size = int(fileSize)
	chunkSize := 1024 * 1024 * 8
	if headers.end > chunkSize {
		headers.end = chunkSize
	}

	headers.statusCode = http.StatusOK
	headers.headers.Set("Content-Length", strconv.FormatInt(fileSize, 10))
	headers.headers.Set("Content-Type", contentType)

	if bRange != "" {
		var (
			parts    = strings.Split(strings.Replace(bRange, "bytes=", "", -1), "-")
			start, _ = strconv.Atoi(parts[0])
			end      int
		)
		parts1Value, err := strconv.Atoi(parts[1])
		if err != nil {
			end = int(fileSize)
		} else {
			end = parts1Value
		}
		if end > int(fileSize)-1 {
			end = int(fileSize) - 1
		}
		if end > start+chunkSize {
			end = start + chunkSize
		}

		headers.statusCode = http.StatusPartialContent
		headers.start = start
		headers.end = end
		headers.size = int(fileSize)
		headers.headers.Set("Content-Range", "bytes "+strconv.Itoa(start)+"-"+strconv.Itoa(end)+"/"+strconv.FormatInt(fileSize, 10))
		headers.headers.Set("Accept-Ranges", "bytes")
	}
	return headers, nil
}

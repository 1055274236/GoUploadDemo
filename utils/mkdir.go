/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-11 12:10:56
 */
package utils

import (
	"UploadApi/config"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

// 判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

// 创建文件夹
func CreateDir(p string) {
	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		log.Printf("创建目录异常 -> %v\n", err)
	} else {
		log.Printf("创建成功! -> %v\n", p)
	}
}

func CreateNecessaryDir(p string) {
	if p == "" {
		p = path.Join(config.StaticFileDir, ""+time.Now().Format("20060102"))
	}
	_exist, _ := HasDir(p)
	if !_exist {
		CreateDir(p)
	}
}

func RemoveFile(p string) error {
	var err error = nil
	if hasFile, _ := HasDir(p); hasFile {
		err = os.Remove(p)
		for {
			upDirPath := path.Dir(p)
			// 如果已经是最后一个文件夹，则终止
			if strings.EqualFold(upDirPath, config.StaticFileDir) {
				break
			}
			// 判断该文件夹是否为空文件夹
			if fileArr, _ := os.ReadDir(upDirPath); len(fileArr) == 0 {
				p = upDirPath
				err = os.Remove(upDirPath)
			} else {
				break
			}
		}
	}
	if err != nil {
		log.Println(err)
	}
	return err
}

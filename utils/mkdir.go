/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-09 03:47:28
 */
package utils

import (
	"UploadApi/config"
	"log"
	"os"
	"path"
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

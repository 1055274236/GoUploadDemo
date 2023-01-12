/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 20:48:12
 */
package fileinfo

import (
	"UploadApi/databases/mysql"
	"log"
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

/**
 * @description: 插入新值
 * @author: Ming
 */
func InsertValue(newValue InsertType) (FileInfo, *gorm.DB) {
	var db = mysql.GetContentByTable("file_info")

	// Create New
	var newLine FileInfo = FileInfo{
		ID:     strings.ReplaceAll(uuid.NewV4().String(), "-", ""),
		Folder: newValue.Folder,
		File:   newValue.File,
		Hash:   newValue.Hash,
		Size:   newValue.Size,
	}

	// Insert
	k := db.Create(&newLine)

	return newLine, k
}

/**
 * @description: 更新文件位置
 * @author: Ming
 */
func UpdataValue(newValue FileInfo) FileInfo {
	var db = mysql.GetContentByTable("file_info")

	db.Save(&newValue)

	return newValue
}

/**
 * @description: 根据 ID 数组，返回一个结果数组
 * @author: Ming
 */
func GetValuesById(ids []string) ([]FileInfo, *gorm.DB) {
	var db = mysql.GetContentByTable("file_info")

	// var fileInfo []map[string]interface{}
	var fileInfo []FileInfo
	if len(ids) == 0 {
		log.Println("Data Error! Id 获取失败！")
		return fileInfo, nil
	}
	k := db.Where("id = ?", ids[0])
	for i := 1; i < len(ids); i++ {
		k = k.Or("id = ?", ids[i])
	}
	k.Find(&fileInfo)

	return fileInfo, k
}

/**
 * @description: 根据 hash 数组，返回一个结果数组
 * @author: Ming
 */
func GetValuesByHash(hashs []string) ([]FileInfo, *gorm.DB) {
	var db = mysql.GetContentByTable("file_info")
	var fileInfo []FileInfo

	if len(hashs) == 0 {
		log.Println("Data Error! Hash 获取失败！")
		return fileInfo, nil
	}

	k := db.Where("hash = ?", hashs[0])
	for i := 1; i < len(hashs); i++ {
		k = k.Or("hash = ?", hashs[i])
	}
	k.Find(&fileInfo)

	return fileInfo, k
}

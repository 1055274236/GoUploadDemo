/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 19:55:55
 */
package uploadrecord

import (
	"UploadApi/databases/mysql"
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func InsertValue(newValue InsertType) (UploadRecord, *gorm.DB) {
	var db = mysql.GetContentByTable("upload_record")

	// Create New
	var newLine UploadRecord
	newLine.ID = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	newLine.Ip = newValue.Ip
	newLine.ContentType = newValue.ContentType
	newLine.FileId = newValue.FileId
	newLine.Name = newValue.Name
	// Insert
	k := db.Create(&newLine)

	return newLine, k
}

func GetValuesById(ids []string) ([]UploadRecordAndFileInfo, *gorm.DB) {
	var db = mysql.GetContentByTable("upload_record")
	var result []UploadRecordAndFileInfo

	k := db.Preload("FileInfo").Where("id = ?", ids[0])

	for i := 1; i < len(ids); i++ {
		k.Or("id = ?", ids[i])
	}
	k.Find(&result)

	return result, k
}

func GetValuesByIp(ips []string) ([]UploadRecordAndFileInfo, *gorm.DB) {
	var db = mysql.GetContentByTable("upload_record")
	var result []UploadRecordAndFileInfo

	k := db.Preload("FileInfo").Where("ip = ?", ips[0])

	for i := 1; i < len(ips); i++ {
		k.Or("ip = ?", ips[i])
	}
	k.Find(&result)

	return result, k
}

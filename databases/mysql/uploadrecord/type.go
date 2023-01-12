/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 19:38:37
 */
package uploadrecord

import (
	"time"
)

type UploadRecord struct {
	ID          string    `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	FileId      string    `gorm:"column:file_id"`
	Ip          string    `gorm:"column:ip"`
	ContentType string    `gorm:"column:content_type"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type InsertType struct {
	Name        string `gorm:"column:name"`
	FileId      string `gorm:"column:file_id"`
	Ip          string `gorm:"column:ip"`
	ContentType string `gorm:"column:content_type"`
}

type file_info struct {
	ID        string    `gorm:"column:id"`
	Folder    string    `gorm:"column:folder"`
	File      string    `gorm:"column:file"`
	Hash      string    `gorm:"column:hash"`
	Size      int64     `gorm:"column:size"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type UploadRecordAndFileInfo struct {
	ID          string    `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	FileId      string    `gorm:"column:file_id"`
	Ip          string    `gorm:"column:ip"`
	ContentType string    `gorm:"column:content_type"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	FileInfo    file_info `gorm:"foreignKey:FileId"`
}

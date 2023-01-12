/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-11 23:31:45
 */
package fileinfo

import "time"

type FileInfo struct {
	ID        string    `gorm:"column:id"`
	Folder    string    `gorm:"column:folder"`
	File      string    `gorm:"column:file"`
	Hash      string    `gorm:"column:hash"`
	Size      int64     `gorm:"column:size"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type InsertType struct {
	Folder string `gorm:"column:folder"`
	File   string `gorm:"column:file"`
	Hash   string `gorm:"column:hash"`
	Size   int64  `gorm:"column:size"`
}

type UpdataType struct {
	ID     string `gorm:"column:id"`
	Folder string `gorm:"column:folder"`
	File   string `gorm:"column:file"`
}

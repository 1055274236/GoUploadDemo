/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 20:53:32
 */
package upload

import (
	"UploadApi/config"
	"UploadApi/databases/mysql/fileinfo"
	"UploadApi/databases/mysql/uploadrecord"
	"UploadApi/utils"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func gethash(path string) (hash string) {
	file, _ := os.Open(path)
	defer file.Close()
	h_ob := sha256.New()
	_, err := io.Copy(h_ob, file)
	if err == nil {
		hash := h_ob.Sum(nil)
		hashvalue := hex.EncodeToString(hash)
		return hashvalue
	} else {
		return "哈希错误"
	}
}

func InsertDB(newValues []FileHeader, c *gin.Context) []ReturnInfoType {
	var returnResultArr []ReturnInfoType

	for _, item := range newValues {
		var (
			returnResultItem   ReturnInfoType
			uploadRecordInsert uploadrecord.InsertType   = uploadrecord.InsertType{}
			fileInfoResult     fileinfo.FileInfo         = fileinfo.FileInfo{}
			uploadRecordResult uploadrecord.UploadRecord = uploadrecord.UploadRecord{}
		)

		hash := gethash(item.FilePath)
		dbFileInfoArr, fileInfoDB := fileinfo.GetValuesByHash([]string{hash})

		// 查看文件数据库中是否含有该文件的Hash值
		if fileInfoDB.RowsAffected > 0 {
			fileInfoResult = dbFileInfoArr[0]
			// 当数据库中对应文件未找到时，便更新文件位置 找到便删除新增加文件
			hasDir, _ := utils.HasDir(path.Join(config.StaticFileDir, fileInfoResult.Folder, fileInfoResult.File))
			if !hasDir {
				fileInfoResult.Folder = item.Folder
				fileInfoResult.File = item.FileName
				fileInfoResult = fileinfo.UpdataValue(fileInfoResult)
			} else {
				go utils.RemoveFile(item.FilePath)
			}
		} else {
			// 插入新文件信息
			nowFileInfo, _ := os.Stat(item.FilePath)
			fileInfoInsert := fileinfo.InsertType{
				Folder: item.Folder,
				File:   item.FileName,
				Hash:   hash,
				Size:   nowFileInfo.Size(),
			}
			fileInfoResult, _ = fileinfo.InsertValue(fileInfoInsert)
		}

		// 文件上传记录保存
		uploadRecordInsert = uploadrecord.InsertType{
			Name:        item.Name,
			FileId:      fileInfoResult.ID,
			Ip:          c.ClientIP(),
			ContentType: item.ContentType,
		}
		uploadRecordResult, _ = uploadrecord.InsertValue(uploadRecordInsert)

		returnResultItem = ReturnInfoType{
			Name:        uploadRecordResult.Name,
			Folder:      fileInfoResult.Folder,
			File:        fileInfoResult.File,
			Size:        fileInfoResult.Size,
			Hash:        fileInfoResult.Hash,
			ContentType: uploadRecordResult.ContentType,
			CreatedAt:   uploadRecordResult.CreatedAt,
			Url:         fmt.Sprintf("http://%v/download/%v/%v", c.Request.Host, fileInfoResult.Folder, fileInfoResult.File),
		}
		returnResultArr = append(returnResultArr, returnResultItem)
	}
	return returnResultArr
}

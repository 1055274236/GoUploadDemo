/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 20:53:10
 */
package upload

import "time"

type FileHeader struct {
	Name        string // 文件原始名称
	FormName    string // 上传时表单名称
	Folder      string // 存储文件夹名称
	FileName    string // 包括后缀的新文件名
	FilePath    string // 文件存储路径
	ContentType string // 上传时的ContentType
}

type ReturnInfoType struct {
	Name        string    // 文件名称
	Folder      string    // 本地存储文件夹
	File        string    // 本地存储文件名
	Size        int64     // 本地存储时的尺寸
	Hash        string    // 文件 hash
	Url         string    // 根据本地存储地址拼接的访问连接
	ContentType string    // 返回文件上传时传递的Type
	CreatedAt   time.Time // 该请求完成时间
}

type Boundary struct {
	base  []byte // 基础
	start []byte // 文件头开头
	end   []byte // 文件传输结束
}

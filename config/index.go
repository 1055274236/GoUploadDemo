/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2023-01-12 16:46:35
 */
package config

import (
	"path"
	"runtime"

	"github.com/gin-gonic/gin"
)

var _, rootPath, _, _ = runtime.Caller(0)

var (
	// File
	RootPath             = path.Dir(path.Dir(rootPath))
	StaticFileDir        = path.Join(RootPath, "static")
	GinLogFile           = path.Join(RootPath, "gin.log")
	FileChunkSize        = 12288       // 文件读取块大小 1024 * 12
	FileContentMaxLength = 32212254720 // 30G 1024 * 1024 * 1024 * 30

	// Gin
	PORT       = ":50010"
	GinRunMode = gin.ReleaseMode

	// DB
	DbIp           = "127.0.0.1"
	DbPort         = "3306"
	DbUser         = "file_store"
	DbPassword     = "password"
	DbName         = "file_store"
	DbCharset      = "utf8mb4"
	DbLoc          = "Asia%2FShanghai"
	DbMaxIdleConns = 5 //设置最大连接数
	DbMaxOpenConns = 3 //设置最大的空闲连接数
)

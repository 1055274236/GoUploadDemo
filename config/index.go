/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-19 00:56:20
 */
package config

import (
	"path"
	"runtime"

	"github.com/gin-gonic/gin"
)

var _, rootPath, _, _ = runtime.Caller(0)

var (
	PORT                 = ":50010"
	RootPath             = path.Dir(path.Dir(rootPath))
	StaticFileDir        = path.Join(RootPath, "static")
	GinLogFile           = path.Join(rootPath, "gin.log")
	FileContentMaxLength = 1024 * 1024 * 1024 * 30 // 30G
	GinRunMode           = gin.ReleaseMode
)

/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-09 03:38:27
 */
package config

import (
	"path"
	"runtime"
)

var _, rootPath, _, _ = runtime.Caller(0)

var (
	PORT                 = ":50010"
	RootPath             = path.Dir(path.Dir(rootPath))
	StaticFileDir        = path.Join(RootPath, "static")
	FileContentMaxLength = 1024 * 1024 * 1024 * 30 // 30G
)

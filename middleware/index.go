/*
 * @Description:
 * @Autor: Ming
 * @LastEditors: Ming
 * @LastEditTime: 2022-12-19 01:22:16
 */
package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(c *gin.Engine) {
	c.Use(Cors())
}

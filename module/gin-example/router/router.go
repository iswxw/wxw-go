/*
@Time: 2022/12/24 16:05
@Author: wxw
@File: router
*/
package router

import "github.com/gin-gonic/gin"

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return r
}

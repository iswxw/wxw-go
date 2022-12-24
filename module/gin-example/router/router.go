/*
@Time: 2022/12/24 16:05
@Author: wxw
@File: router
*/
package router

import (
	"github.com/gin-gonic/gin"
	"src/com.wxw/project_actual/module/gin-example/router/debug"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/wxw/gin/debug/ping", debug.Ping)
	return r
}

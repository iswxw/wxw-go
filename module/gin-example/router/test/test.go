/*
@Time: 2022/12/24 17:46
@Author: wxw
@File: hello_world
*/
package test

import (
	"github.com/gin-gonic/gin"
	"src/com.wxw/project_actual/module/gin-example/app/test"
)

// DebugAPI 调试及工具接口
func DebugAPI(r *gin.Engine) {
	r.GET("/wxw/gin/test/ping", test.Ping)
}

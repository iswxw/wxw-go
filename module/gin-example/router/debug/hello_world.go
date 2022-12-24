/*
@Time: 2022/12/24 17:46
@Author: wxw
@File: hello_world
*/
package debug

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping 连通性测试
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

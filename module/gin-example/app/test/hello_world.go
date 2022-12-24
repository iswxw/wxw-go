/*
@Time: 2022/12/24 18:03
@Author: wxw
@File: hello_world
*/
package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping 连通性测试
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

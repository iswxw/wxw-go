/*
@Time: 2022/12/24 17:56
@Author: wxw
@File: resopnse
*/
package param

import (
	"github.com/gin-gonic/gin"
	"src/com.wxw/project_actual/module/gin-example/common/consts"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  consts.GetMsg(errCode),
		Data: data,
	})
	return
}

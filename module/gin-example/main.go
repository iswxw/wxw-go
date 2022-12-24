/*
@Time: 2022/12/24 14:59
@Author: wxw
@File: main
*/
package main

import (
	"src/com.wxw/project_actual/module/gin-example/common/infra/conf"
	"src/com.wxw/project_actual/module/gin-example/common/server"
)

func main() {
	funcInit()
	server.Setup()
}

func funcInit() {
	conf.Setup("")
}

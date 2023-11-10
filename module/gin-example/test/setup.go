/*
@Time: 2022/12/25 11:41
@Author: wxw
@File: setup
*/
package test

import (
	"path/filepath"
	"runtime"
	"src/com.wxw/project_actual/module/gin-example/common/infra/conf"
	"src/com.wxw/project_actual/module/gin-example/common/infra/mysql"
)

// SetupTest 测试初始化配置
func SetupTest() {
	conf.Setup(filepath.Join(curPath(), "../conf/dev/app.toml"))
	mysql.Setup()
}

// curPath 当前根路径
func curPath() string {
	_, fn, _, _ := runtime.Caller(0)
	dir := filepath.Dir(fn)
	return dir
}

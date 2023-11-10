/*
@Time: 2022/12/24 22:50
@Author: wxw
@File: mysql_test
*/
package mysql

import (
	"src/com.wxw/project_actual/module/gin-example/common/infra/conf"
	"testing"
)

func TestSetup(t *testing.T) {
	conf.Setup("")
	Setup()
}

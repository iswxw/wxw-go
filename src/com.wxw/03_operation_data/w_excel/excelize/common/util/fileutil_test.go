/*
@Time : 2022/1/24 00:07
@Author : weixiaowei
@File : fileutil_test
*/
package util

import (
	"fmt"
	"testing"
)

func TestGetPath(t *testing.T) {
	fmt.Println("path=", GetPath("Book1.xlsx"))
}

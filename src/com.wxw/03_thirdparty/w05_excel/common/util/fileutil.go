/*
@Time : 2022/1/23 23:52
@Author : weixiaowei
@File : fileutil
*/
package util

import (
	"fmt"
	"os"
)

func GetPath(fileName string) string {
	rootPath, _ := os.Getwd()
	return fmt.Sprintf("%s/%s", rootPath, fileName)
}

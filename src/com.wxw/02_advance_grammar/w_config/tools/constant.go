/*
@Time: 2022/5/15 11:32
@Author: wxw
@File: constant
*/
package tools

import (
	"fmt"
	"os"
)

// 文件路径
const (
	ConfigFilePath = "/src/com.wxw/04_project_actual/w_config/tools/conf/"
)

// 获取指定文件路径
func GetFilePath(filePath string, fileName string) string {
	basePath, _ := os.Getwd()
	return fmt.Sprintf("%s%s%s", basePath, filePath, fileName)
}

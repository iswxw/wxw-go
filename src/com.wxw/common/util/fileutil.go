/*
@Time: 2022/1/12 23:26
@Author: wxw
@File: file_util
*/
package util

import (
	"fmt"
	"os"
)

// 获取指定文件路径
func GetFilePath(filePath string, fileName string) string {
	basePath, _ := os.Getwd()
	return fmt.Sprintf("%s%s%s", basePath, filePath, fileName)
}

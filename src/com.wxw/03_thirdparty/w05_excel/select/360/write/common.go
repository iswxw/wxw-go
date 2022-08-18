// @Time : 2022/8/18 11:54
// @Author : xiaoweiwei
// @File : common

package write

import (
	"fmt"
	"os"
)

func GetPath(fileName string) string {
	rootPath, _ := os.Getwd()
	return fmt.Sprintf("%s/%s", rootPath, fileName)
}

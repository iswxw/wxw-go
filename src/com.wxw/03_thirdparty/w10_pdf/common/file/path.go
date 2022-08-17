// @Time : 2022/8/17 11:46
// @Author : xiaoweiwei
// @File : path

package file

import (
	"fmt"
	"os"
	"strings"
)

func GetImagePath() string {
	return fmt.Sprintf("%s%s/", GetPdfPath(), "imgs")
}

func GetFontPath() string {
	return fmt.Sprintf("%s%s/", GetPdfPath(), "ttf")
}

func GetSaveFilePath() string {
	return fmt.Sprintf("%s/%s/", strings.Replace(GetPdfPath(), "/common", "", 1), "tmp")
}

func GetPdfPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "file", "", 1)
}

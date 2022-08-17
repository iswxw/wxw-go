// @Time : 2022/8/17 11:49
// @Author : xiaoweiwei
// @File : file_test

package file

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	log.Println(GetSaveFilePath())
}

func TestFilePath(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	replace := strings.Replace(dir, "file", "", 1)
	log.Println(replace)
}

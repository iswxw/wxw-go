/*
@Time: 2022/1/9 14:39
@Author: wxw
@File: read_file_path
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	// 方式一
	dir, _ := os.Getwd()
	fmt.Println("当前路径：", dir)

	// 方式二

	// 方式三
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path) // 执行未见的路径
	fileDir := filepath.Dir(path)
	fmt.Println(fileDir) // for example /home/user

}

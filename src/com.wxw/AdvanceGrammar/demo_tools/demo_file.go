/*
 * @Time : 2021/2/28 21:04
 * @Author : wxw
 * @File : demo_file
 * @Software: GoLand
 * @Link:
 * @Vlog:
 */
package main

import (
	"fmt"
	"os"
)

// 1. 文件对象类型
// 2. 获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("./doc/log/wxw_go.log")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	// 1. 文件对象类型
	fmt.Printf("%T\n", fileObj)
	// 2. 获取文件对象详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return
	}
	fmt.Printf("文件大小是: %d B\n", fileInfo.Size())
	fmt.Printf("文件名称是: %s \n", fileInfo.Name())
}

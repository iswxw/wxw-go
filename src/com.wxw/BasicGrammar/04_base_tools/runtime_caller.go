/*
 * @Time : 2021/2/28 15:31
 * @Author : wxw
 * @File : runtime_caller
 * @Software: GoLand
 * @Link:
 */
package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller() 运行时调用，获取调用者信息
func main() {
	pc, file, line, ok := runtime.Caller(0) // 0 指的是调用的层数，返回 line是行号
	if !ok {
		fmt.Printf("runtime.Caller() failed \n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name() // 获取函数名
	fmt.Println(funcName)
	fmt.Println(file)            // 文件路径
	fmt.Println(path.Base(file)) // 方法所在文件名称
	fmt.Println(line)            // 行号
}

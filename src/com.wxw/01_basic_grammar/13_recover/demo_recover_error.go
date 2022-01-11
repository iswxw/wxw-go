/*
@Time: 2021/10/11 23:57
@Author: wxw
@File: demo1_recover
*/
package main

import "fmt"

func main() {
	recover() // 无任何作用
	panic("停止运行")
	recover() // 不会执行到
	fmt.Println("结束")

	// 输出
	// panic: 停止运行
	//goroutine 1 [running]:exit status 2
}

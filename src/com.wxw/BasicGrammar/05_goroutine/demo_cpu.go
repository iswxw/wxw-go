/*
@Time: 2021/10/24 23:40
@Author: wxw
@File: demo_cpu
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum = ", cpuNum)

	// 可以设置使用多个CPU
	gomaxprocs := runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("gomaxprocs = ", gomaxprocs)
}

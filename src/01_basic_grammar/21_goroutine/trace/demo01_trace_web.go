/*
@Time: 2022/2/20 0:27
@Author: wxw
@File: gpm_web
*/
package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// 可视化追踪 gmp
func main() {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动 trace goroutine
	if err = trace.Start(f); err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	fmt.Println("Hello World")
}

// 当前路径下
// go run demo01_trace_web.go 生成 trace.out文件
// go tool trace trace.out 打开文件分析

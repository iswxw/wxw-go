/*
@Time: 2022/10/15 23:40
@Author: wxw
@File: demo02_trace_debug
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
	}
}

// 当前路径下
// 编译：go build demo02_trace_debug.go
// debug：GODEBUG=schedtrace=1000 ./demo02_trace_debug.exe
//

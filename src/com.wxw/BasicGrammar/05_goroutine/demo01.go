/*
@Time: 2021/10/24 22:44
@Author: wxw
@File: demo01
*/
package main

import (
	"fmt"
	"strconv"
	"time"
)

// 主线程 每隔一秒输出 hello world
// goroutine 每隔一秒输出 hello world
func main() {
	go test() // 开启了一个协程
	for i := 1; i < 10; i++ {
		fmt.Println("hello world! main = " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

// 编写一个函数，每隔一秒，输出一次hello world
func test() {
	for i := 1; i < 10; i++ {
		fmt.Println("hello world! test = " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

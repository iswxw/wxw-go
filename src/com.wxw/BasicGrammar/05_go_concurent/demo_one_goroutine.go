/*
 * @Time : 2021/3/3 21:14
 * @Author : wxw
 * @File : demo_concurent
 * @Software: GoLand
 * @Link: https://www.liwenzhou.com/posts/Go/14_concurrence/
 * @Vlog:
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)
}

func hello() {
	fmt.Println("Hello GoRoutine!")
}

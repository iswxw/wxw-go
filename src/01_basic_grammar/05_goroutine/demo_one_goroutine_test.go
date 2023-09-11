/*
 * @Time : 2021/3/3 21:14
 * @Author : wxw
 * @File : demo_concurent
 * @Software: GoLand
 * @Link: https://www.liwenzhou.com/posts/Go/14_concurrence/
 * @Vlog:
 */
package _goroutine

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	// goroutine 随着main函数的结束而结束，所以我们想要hello 同时执行，需要让main等一等
	// time.Sleep(time.Second)
}

func hello() {
	fmt.Println("Hello GoRoutine!")
}

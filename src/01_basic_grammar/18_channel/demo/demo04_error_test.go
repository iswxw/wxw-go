/*
@Time: 2022/7/25 23:08
@Author: wxw
@File: demo04_error
*/
package demo

import (
	"fmt"
	"testing"
)

// TestError 通道阻塞，内存泄露的场景
func TestError(t *testing.T) {
	var name int
	// 定义一个int类型的通道
	c := make(chan int)
	c <- name
	fmt.Println(<-c)
}

func TestRepeat(t *testing.T) {
	c := make(chan int, 1)
	c <- 1
	close(c)
	close(c)
	fmt.Println("OK")
}

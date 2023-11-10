/*
@Time : 2022/3/22 16:40
@Author : weixiaowei
@File : demo02
*/
package main

import (
	"fmt"
	"testing"
	"time"
)

// go的通道选择器 让你可以同时等待多个通道操作。go协程和通道以及选择器的结合是go的一个强大特性。
func TestDeadLock(t *testing.T) {
	// 定义两个通道
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c2 <- "hello"

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		default:
			fmt.Println("No data received")
		}
	}

}

/*
@Time : 2022/3/22 16:25
@Author : weixiaowei
@File : demo01
*/
package demo

import (
	"fmt"
	"testing"
)

func TestNoCache(t *testing.T) {
	// 定义一个数据
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 定义一个int类型的通道
	c := make(chan int)

	// 调用 sum 方法
	go sum(s[:len(s)/2], c) // 中上数据

	go sum(s[len(s)/2:], c) // 中下数据

	x, y := <-c, <-c // 从通道 c 中接收
	fmt.Println(x, y, x+y)
}

// 实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

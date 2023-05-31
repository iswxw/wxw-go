// @Time : 2023/5/31 20:15
// @Author : xiaoweiwei
// @File : demo05_transfer_single_test

package demo

import (
	"fmt"
	"testing"
	"time"
)

// 基于空结构体，实现channel通道信号传输，详见：https://www.yuque.com/fcant/go/sakaw6#IbdW3
func TestTransferSignle(t *testing.T) {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go doTask1(ch1)
	go doTask2(ch2)

	for {
		select {
		case <-ch1:
			fmt.Println("task1 done")
		case <-ch2:
			fmt.Println("task2 done")
		case <-time.After(time.Second * 3):
			fmt.Println("after 3 seconds")
			return
		}
	}
}

func doTask1(ch chan struct{}) {
	time.Sleep(time.Second)
	fmt.Println("do task1")
	ch <- struct{}{}
}

func doTask2(ch chan struct{}) {
	time.Sleep(time.Second * 2)
	fmt.Println("do task2")
	ch <- struct{}{}
}

/*
@Time : 2022/5/6 15:21
@Author : weixiaowei
@File : demo01_multi_goroutine
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 来源： https://github.com/hyper0x/Golang_Puzzlers/blob/master/src/puzzlers/article25/q0/demo65.go
func main() {

	coordinateWithChan()

	fmt.Println()

	coordinateWithWaitGroup()
}

func coordinateWithWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	num := int32(0)
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	max := int32(10)
	go addNum(&num, 3, max, wg.Done)
	go addNum(&num, 4, max, wg.Done)
	wg.Wait()
}

// 分批地启用执行子任务的 goroutine
//func coordinateWithWaitGroup01() {
//	total := 12
//	stride := 3
//	var num int32
//	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
//	var wg sync.WaitGroup
//	for i := 1; i <= total; i = i + stride {
//		wg.Add(stride)
//		for j := 0; j < stride; j++ {
//			go addNum(&num, i+j, wg.Done)
//		}
//		wg.Wait()
//	}
//	fmt.Println("End.")
//}

func coordinateWithChan() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d [with chan struct{}]\n", num)
	max := int32(10)
	go addNum(&num, 1, max, func() {
		sign <- struct{}{}
	})
	go addNum(&num, 2, max, func() {
		sign <- struct{}{}
	})
	<-sign
	<-sign
}

// addNum 用于原子地增加numP所指的变量的值。
func addNum(numP *int32, id, max int32, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	for i := 0; ; i++ {
		currNum := atomic.LoadInt32(numP)
		if currNum >= max {
			break
		}
		newNum := currNum + 2
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
		} else {
			fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
		}
	}
}
